package com.rentmanager.ai.document;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.time.Instant;
import java.util.Optional;
import java.util.UUID;

import com.rentmanager.ai.indexing.DocumentIndexer;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.io.TempDir;
import org.springframework.dao.DataAccessResourceFailureException;
import org.springframework.http.HttpStatus;
import org.springframework.transaction.TransactionDefinition;
import org.springframework.transaction.TransactionSystemException;
import org.springframework.transaction.support.AbstractPlatformTransactionManager;
import org.springframework.transaction.support.DefaultTransactionStatus;

import static org.assertj.core.api.Assertions.assertThat;
import static org.assertj.core.api.Assertions.assertThatThrownBy;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.doThrow;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.never;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.verifyNoInteractions;
import static org.mockito.Mockito.when;

final class DocumentServiceTest {
    @TempDir
    private Path root;
    private final DocumentRepository repository = mock(DocumentRepository.class);
    private final DocumentIndexer indexer = mock(DocumentIndexer.class);
    private final DocumentScope scope = new DocumentScope("123", "456");

    @Test
    void failedInsertRollsBackAndRemovesPhysicalFile() throws IOException {
        final var service = service(new TestTransactions(false));
        doThrow(new DataAccessResourceFailureException("database write failed")).when(repository).insert(any());
        assertThatThrownBy(() -> service.upload(scope, DocumentType.RENT_CONTRACT, PdfFixtures.upload()))
                .isInstanceOf(DataAccessResourceFailureException.class);
        try (final var files = Files.walk(root)) {
            assertThat(files.filter(Files::isRegularFile).count()).isZero();
        }
    }

    @Test
    void failedIndexingRemovesTheSavedPdf() throws IOException {
        doThrow(new DocumentException(HttpStatus.SERVICE_UNAVAILABLE, "Embedding failed")).when(indexer).index(any(), any());
        final var service = service(new TestTransactions(false));
        assertThatThrownBy(() -> service.upload(scope, DocumentType.RENT_CONTRACT, PdfFixtures.upload()))
                .isInstanceOf(DocumentException.class);
        verify(repository).insert(any());
        try (final var files = Files.walk(root)) {
            assertThat(files.filter(Files::isRegularFile).count()).isZero();
        }
    }

    @Test
    void ambiguousCommitKeepsPdfForReconciliation() throws IOException {
        final var service = service(new TestTransactions(true));
        assertThatThrownBy(() -> service.upload(scope, DocumentType.RENT_CONTRACT, PdfFixtures.upload()))
                .isInstanceOf(TransactionSystemException.class);
        try (final var files = Files.walk(root)) {
            assertThat(files.filter(Files::isRegularFile).count()).isEqualTo(1);
        }
    }

    @Test
    void failedDeleteRestoresFileAndSuccessfulDeleteRemovesIt() throws IOException {
        final var storage = new FileSystemDocumentStorage(root.toString());
        final UUID id = UUID.randomUUID();
        final byte[] content = PdfFixtures.pdf(false);
        final String path = storage.save(scope, id, content);
        final var metadata = new DocumentMetadata(id, scope.propertyId(), scope.contractId(), DocumentType.OTHER,
                "contract.pdf", id + ".pdf", path, "application/pdf", content.length, Instant.now(), Instant.now());
        when(repository.find(scope, id, true)).thenReturn(Optional.of(metadata));
        doThrow(new DataAccessResourceFailureException("delete failed")).doNothing().when(repository).delete(scope, id);
        final var service = service(new TestTransactions(false));
        assertThatThrownBy(() -> service.delete(scope, id)).isInstanceOf(DataAccessResourceFailureException.class);
        assertThat(storage.read(path)).isEqualTo(content);
        assertThat(root.resolve(path + ".deleting")).doesNotExist();
        service.delete(scope, id);
        assertThat(root.resolve(path)).doesNotExist();
        assertThat(root.resolve(path + ".deleting")).doesNotExist();
    }

    @Test
    void failedFileSaveNeverInsertsARegistryRow() throws IOException {
        final var storage = mock(DocumentStorage.class);
        when(storage.save(any(), any(), any())).thenThrow(new DocumentException(HttpStatus.SERVICE_UNAVAILABLE, "Disk is full"));
        final var service = new DocumentService(repository, storage, new PdfUploadValidator(), indexer, new TestTransactions(false), 300);
        assertThatThrownBy(() -> service.upload(scope, DocumentType.OTHER, PdfFixtures.upload())).isInstanceOf(DocumentException.class);
        verifyNoInteractions(repository);
    }

    @Test
    void reportsPostCommitCleanupFailureWithoutPretendingToRestoreDeletedRow() {
        final var storage = mock(DocumentStorage.class);
        final UUID id = UUID.randomUUID();
        final String path = "properties/123/contracts/456/" + id + ".pdf";
        final var metadata = new DocumentMetadata(id, scope.propertyId(), scope.contractId(), DocumentType.OTHER,
                "contract.pdf", id + ".pdf", path, "application/pdf", 123, Instant.now(), Instant.now());
        when(repository.find(scope, id, true)).thenReturn(Optional.of(metadata));
        when(storage.stageDeletion(path)).thenReturn(true);
        doThrow(new DocumentException(HttpStatus.SERVICE_UNAVAILABLE, "Disk unavailable")).when(storage).completeDeletion(path);
        final var service = new DocumentService(repository, storage, new PdfUploadValidator(), indexer, new TestTransactions(false), 300);
        assertThatThrownBy(() -> service.delete(scope, id)).isInstanceOfSatisfying(DocumentException.class,
                error -> assertThat(error.getMessage()).contains("reconciliation"));
        verify(repository).delete(scope, id);
        verify(storage, never()).restoreDeletion(any());
    }

    @Test
    void foreignContractCannotReadOrDeleteFile() {
        final var storage = mock(DocumentStorage.class);
        final UUID id = UUID.randomUUID();
        final var service = new DocumentService(repository, storage, new PdfUploadValidator(), indexer, new TestTransactions(false), 300);
        when(repository.find(scope, id, false)).thenReturn(Optional.empty());
        when(repository.find(scope, id, true)).thenReturn(Optional.empty());
        assertThatThrownBy(() -> service.download(scope, id)).isInstanceOf(DocumentException.class);
        assertThatThrownBy(() -> service.delete(scope, id)).isInstanceOf(DocumentException.class);
        verifyNoInteractions(storage);
    }

    private DocumentService service(final TestTransactions transactions) throws IOException {
        return new DocumentService(repository, new FileSystemDocumentStorage(root.toString()), new PdfUploadValidator(), indexer, transactions, 300);
    }

    // Exercises real Spring transaction synchronization, with deterministic commit/rollback outcomes.
    private static final class TestTransactions extends AbstractPlatformTransactionManager {
        private final boolean failCommit;

        private TestTransactions(final boolean failCommit) {
            this.failCommit = failCommit;
        }

        @Override
        protected Object doGetTransaction() {
            return new Object();
        }

        @Override
        protected void doBegin(final Object transaction, final TransactionDefinition definition) {
        }

        @Override
        protected void doCommit(final DefaultTransactionStatus status) {
            if (failCommit) {
                throw new TransactionSystemException("Commit response lost");
            }
        }

        @Override
        protected void doRollback(final DefaultTransactionStatus status) {
        }
    }
}
