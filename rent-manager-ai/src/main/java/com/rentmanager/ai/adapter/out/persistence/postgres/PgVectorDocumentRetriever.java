package com.rentmanager.ai.adapter.out.persistence.postgres;

import java.util.List;
import java.util.Map;

import com.rentmanager.ai.domain.document.RetrievedDocumentChunk;
import org.springframework.ai.vectorstore.SearchRequest;
import org.springframework.ai.vectorstore.VectorStore;
import org.springframework.ai.vectorstore.filter.FilterExpressionBuilder;
import org.springframework.stereotype.Component;
import com.rentmanager.ai.port.out.DocumentRetriever;

@Component
public final class PgVectorDocumentRetriever implements DocumentRetriever {
    private static final int DEFAULT_TOP_K = 5;
    private final VectorStore vectorStore;

    public PgVectorDocumentRetriever(final VectorStore vectorStore) {
        this.vectorStore = vectorStore;
    }

    @Override
    public List<RetrievedDocumentChunk> retrieve(
            final String propertyId,
            final String contractId,
            final String question) {

        final var filters = new FilterExpressionBuilder();

        final var filter = filters.and(
                filters.eq("propertyId", propertyId),
                filters.eq("contractId", contractId)
        ).build();

        return vectorStore.similaritySearch(
                        SearchRequest.builder()
                                .query(question)
                                .topK(DEFAULT_TOP_K)
                                .filterExpression(filter)
                                .build()
                ).stream()
                .map(this::map)
                .toList();
    }

    private RetrievedDocumentChunk map(final org.springframework.ai.document.Document document) {
        final var metadata = document.getMetadata();

        return new RetrievedDocumentChunk(
                document.getText(),
                value(metadata, "documentId"),
                value(metadata, "propertyId"),
                value(metadata, "contractId"),
                value(metadata, "documentType"),
                value(metadata, "filename"),
                integerValue(metadata, "page")
        );
    }

    private String value(final Map<String, Object> metadata, final String key) {
        final Object value = metadata.get(key);
        return value != null ? value.toString() : null;
    }

    private Integer integerValue(final Map<String, Object> metadata, final String key) {
        final Object value = metadata.get(key);

        if (value == null) {
            return null;
        }

        if (value instanceof Number number) {
            return number.intValue();
        }

        return Integer.valueOf(value.toString());
    }
}
