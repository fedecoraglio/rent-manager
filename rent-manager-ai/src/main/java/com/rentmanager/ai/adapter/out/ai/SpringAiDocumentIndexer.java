package com.rentmanager.ai.adapter.out.ai;

import com.rentmanager.ai.domain.document.DocumentException;
import com.rentmanager.ai.domain.document.DocumentMetadata;
import com.rentmanager.ai.port.out.DocumentIndexer;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.ai.vectorstore.VectorStore;
import org.springframework.ai.vectorstore.filter.FilterExpressionBuilder;
import org.springframework.http.HttpStatus;
import org.springframework.stereotype.Component;
import org.springframework.transaction.support.TransactionSynchronizationManager;

@Component
public final class SpringAiDocumentIndexer implements DocumentIndexer {
    private static final Logger LOG = LoggerFactory.getLogger(SpringAiDocumentIndexer.class);
    private static final int BATCH_SIZE = 16;
    private final PdfChunker chunker;
    private final VectorStore vectorStore;

    public SpringAiDocumentIndexer(final PdfChunker chunker, final VectorStore vectorStore) {
        this.chunker = chunker;
        this.vectorStore = vectorStore;
    }

    public void index(final DocumentMetadata document, final byte[] pdf) {
        requireTransaction();
        final var chunks = chunker.chunks(document, pdf);
        try {
            for (int start = 0; start < chunks.size(); start += BATCH_SIZE) {
                // add() invokes Ollama's EmbeddingModel, then writes vectors using the shared JDBC transaction.
                vectorStore.add(chunks.subList(start, Math.min(start + BATCH_SIZE, chunks.size())));
            }
        } catch (final RuntimeException exception) {
            // Propagating out of DocumentService rolls back ALL batches and the registry row.
            throw new DocumentException(HttpStatus.SERVICE_UNAVAILABLE,
                    "Document indexing failed. Check PostgreSQL and the local nomic-embed-text model, then retry.", exception);
        }
        LOG.info("Prepared {} vector chunks for document {} in the current transaction", chunks.size(), document.id());
    }

    public void delete(final DocumentMetadata document) {
        requireTransaction();
        final var filters = new FilterExpressionBuilder();
        try {
            vectorStore.delete(filters.and(filters.eq("documentId", document.id().toString()),
                    filters.and(filters.eq("propertyId", document.propertyId()),
                            filters.eq("contractId", document.contractId()))).build());
        } catch (final RuntimeException exception) {
            throw new DocumentException(HttpStatus.SERVICE_UNAVAILABLE, "Document chunks could not be deleted. Please retry.", exception);
        }
    }

    private static void requireTransaction() {
        if (!TransactionSynchronizationManager.isActualTransactionActive()) {
            throw new IllegalStateException("Document indexing must run inside the registry transaction");
        }
    }
}
