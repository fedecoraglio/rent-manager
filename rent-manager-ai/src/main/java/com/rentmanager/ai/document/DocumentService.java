package com.rentmanager.ai.document;

import java.time.Instant;
import java.util.List;
import java.util.Objects;
import java.util.UUID;

import com.rentmanager.ai.indexing.DocumentIndexer;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.http.HttpStatus;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;
import org.springframework.transaction.PlatformTransactionManager;
import org.springframework.transaction.support.TransactionSynchronization;
import org.springframework.transaction.support.TransactionSynchronizationManager;
import org.springframework.transaction.support.TransactionTemplate;
import org.springframework.web.multipart.MultipartFile;

@Service
public final class DocumentService {
    private static final Logger LOG = LoggerFactory.getLogger(DocumentService.class);
    private final DocumentRepository repository;
    private final DocumentStorage storage;
    private final PdfUploadValidator validator;
    private final DocumentIndexer indexer;
    private final TransactionTemplate transactions;

    public DocumentService(final DocumentRepository repository, final DocumentStorage storage,
            final PdfUploadValidator validator, final DocumentIndexer indexer,
            final PlatformTransactionManager transactionManager,
            @Value("${app.indexing.transaction-timeout-seconds}") final int transactionTimeoutSeconds) {
        this.repository = repository;
        this.storage = storage;
        this.validator = validator;
        this.indexer = indexer;
        this.transactions = new TransactionTemplate(transactionManager);
        this.transactions.setTimeout(transactionTimeoutSeconds);
    }

    public DocumentMetadata upload(final DocumentScope scope, final DocumentType type, final MultipartFile file) {
        Objects.requireNonNull(type, "Document type is required");
        final var pdf = validator.validate(file);
        final UUID id = UUID.randomUUID();
        final Instant now = Instant.now();
        return Objects.requireNonNull(transactions.execute(status -> {
            final String path = storage.save(scope, id, pdf.content());
            TransactionSynchronizationManager.registerSynchronization(new TransactionSynchronization() {
                @Override
                public void afterCompletion(final int completion) {
                    if (completion == STATUS_ROLLED_BACK) {
                        cleanup(id, () -> storage.remove(path));
                    } else if (completion == STATUS_UNKNOWN) {
                        // Never delete the PDF if the database might have committed its registry row.
                        LOG.error("Upload transaction outcome unknown for document {}; retain file and reconcile registry", id);
                    }
                }
            });
            final var document = new DocumentMetadata(id, scope.propertyId(), scope.contractId(), type,
                    pdf.filename(), id + ".pdf", path, "application/pdf", pdf.content().length, now, now);
            repository.insert(document);
            indexer.index(document, storage.read(path));
            return document;
        }));
    }

    public List<DocumentMetadata> list(final DocumentScope scope) {
        return repository.list(scope);
    }

    public Download download(final DocumentScope scope, final UUID id) {
        final var document = repository.find(scope, id, false).orElseThrow(DocumentException::notFound);
        return new Download(document, storage.read(document.storagePath()));
    }

    public void delete(final DocumentScope scope, final UUID id) {
        transactions.executeWithoutResult(status -> {
            // Serializes competing deletes of the same document before touching its file.
            final var document = repository.find(scope, id, true).orElseThrow(DocumentException::notFound);
            indexer.delete(document);
            final boolean staged = storage.stageDeletion(document.storagePath());
            if (staged) {
                TransactionSynchronizationManager.registerSynchronization(new TransactionSynchronization() {
                    @Override
                    public void afterCommit() {
                        // A failure here is reported to the caller, but cannot undo the committed DELETE.
                        try {
                            storage.completeDeletion(document.storagePath());
                        } catch (final RuntimeException exception) {
                            LOG.error("Committed deletion requires file cleanup for document {}", id, exception);
                            throw new DocumentException(HttpStatus.SERVICE_UNAVAILABLE,
                                    "The registry entry was deleted, but physical file cleanup requires reconciliation.", exception);
                        }
                    }

                    @Override
                    public void afterCompletion(final int completion) {
                        if (completion == STATUS_ROLLED_BACK) {
                            cleanup(id, () -> storage.restoreDeletion(document.storagePath()));
                        } else if (completion == STATUS_UNKNOWN) {
                            LOG.error("Delete transaction outcome unknown for document {}; retain .deleting file and reconcile registry", id);
                        }
                    }
                });
            }
            repository.delete(scope, id);
        });
    }

    private static void cleanup(final UUID id, final Runnable action) {
        try {
            action.run();
        } catch (final RuntimeException exception) {
            // The database transaction is already complete; never pretend we can roll it back here.
            LOG.error("Filesystem cleanup requires reconciliation for document {}", id, exception);
        }
    }

    public record Download(DocumentMetadata document, byte[] content) {
    }
}
