package com.rentmanager.ai.application.document;

import java.time.Instant;
import java.util.Optional;
import java.util.UUID;
import java.io.ByteArrayOutputStream;
import org.apache.pdfbox.pdmodel.PDDocument;
import org.apache.pdfbox.pdmodel.PDPage;

import com.rentmanager.ai.adapter.out.ai.PdfUploadValidator;
import com.rentmanager.ai.domain.document.*;
import com.rentmanager.ai.port.out.*;
import org.junit.jupiter.api.Test;
import org.springframework.mock.web.MockMultipartFile;
import org.springframework.transaction.TransactionDefinition;
import org.springframework.transaction.support.AbstractPlatformTransactionManager;
import org.springframework.transaction.support.DefaultTransactionStatus;

import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.ArgumentMatchers.*;
import static org.mockito.Mockito.*;

class DocumentReplacementTest {
    private final DocumentScope scope = new DocumentScope("P1", "C1");
    private final DocumentRepository repository = mock(DocumentRepository.class);
    private final DocumentStorage storage = mock(DocumentStorage.class);
    private final DocumentIndexer indexer = mock(DocumentIndexer.class);
    private final PdfUploadValidator validator = new PdfUploadValidator();
    private final MockMultipartFile file = pdfFile();
    private final DocumentService service = new DocumentService(repository, storage, validator, indexer,
            new AbstractPlatformTransactionManager() {
                protected Object doGetTransaction() { return new Object(); }
                protected void doBegin(Object transaction, TransactionDefinition definition) { }
                protected void doCommit(DefaultTransactionStatus status) { }
                protected void doRollback(DefaultTransactionStatus status) { }
            }, 300);

    private DocumentMetadata prepareReplacement() {
        var old = new DocumentMetadata("P1_C1_RENT_CONTRACT", "P1", "C1", DocumentType.RENT_CONTRACT,
                "old.pdf", "old.pdf", "old-path", "application/pdf", 1, Instant.EPOCH, Instant.EPOCH);
        when(repository.find(scope, old.id(), true)).thenReturn(Optional.of(old));
        when(storage.save(eq(scope), any(UUID.class), any(byte[].class))).thenReturn("new-path");
        when(storage.read("new-path")).thenReturn(new byte[]{1});
        return old;
    }

    private static MockMultipartFile pdfFile() {
        try (var pdf = new PDDocument(); var output = new ByteArrayOutputStream()) {
            pdf.addPage(new PDPage());
            pdf.save(output);
            return new MockMultipartFile("file", "contract.pdf", "application/pdf", output.toByteArray());
        } catch (java.io.IOException exception) {
            throw new IllegalStateException(exception);
        }
    }

    @Test void replacementKeepsIdentityAndRemovesOldFileOnlyAfterCommit() {
        var old = prepareReplacement();
        doAnswer(invocation -> {
            verify(storage, never()).remove("old-path");
            return null;
        }).when(indexer).index(any(), any());
        var result = service.upload(scope, DocumentType.RENT_CONTRACT, file);
        assertEquals(old.id(), result.id());
        assertEquals(old.createdAt(), result.createdAt());
        assertEquals("new-path", result.storagePath());
        verify(repository).lockScope(scope);
        verify(indexer).delete(old);
        verify(repository).delete(scope, old.id());
        verify(storage).remove("old-path");
        verify(storage, never()).remove("new-path");
    }

    @Test void indexingFailureRetainsOldFileAndCleansNewVersion() {
        var old = prepareReplacement();
        doThrow(new IllegalStateException("Ollama unavailable")).when(indexer).index(any(), any());
        assertThrows(IllegalStateException.class, () -> service.upload(scope, DocumentType.RENT_CONTRACT, file));
        verify(storage).remove("new-path");
        verify(storage, never()).remove(old.storagePath());
    }

    @Test void otherTypesAllowMultipleIdsAndScopeSeparatorsCannotCollide() {
        assertNotEquals(DocumentId.create(scope, DocumentType.ANNEX), DocumentId.create(scope, DocumentType.ANNEX));
        assertNotEquals(DocumentId.create(new DocumentScope("A_B", "C"), DocumentType.RENT_CONTRACT),
                DocumentId.create(new DocumentScope("A", "B_C"), DocumentType.RENT_CONTRACT));
    }
}
