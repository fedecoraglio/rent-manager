package com.rentmanager.ai.adapter.out.persistence.postgres;

import java.util.List;
import java.util.Map;

import com.rentmanager.ai.domain.document.RetrievedDocumentChunk;
import com.rentmanager.ai.port.out.DocumentRetriever;
import org.springframework.ai.document.Document;
import org.springframework.ai.vectorstore.SearchRequest;
import org.springframework.ai.vectorstore.VectorStore;
import org.springframework.ai.vectorstore.filter.FilterExpressionBuilder;
import org.springframework.stereotype.Component;

@Component
public final class PgVectorDocumentRetriever implements DocumentRetriever {

    private static final int DEFAULT_TOP_K = 10;

    private final VectorStore vectorStore;

    public PgVectorDocumentRetriever(final VectorStore vectorStore) {
        this.vectorStore = vectorStore;
    }

    @Override
    public List<RetrievedDocumentChunk> retrieve(
            final String propertyId,
            final String contractId,
            final String question) {

        final FilterExpressionBuilder filters = new FilterExpressionBuilder();

        final var filter = filters.and(
                filters.eq("propertyId", propertyId),
                filters.eq("contractId", contractId)
        ).build();

        final String retrievalQuery = buildRetrievalQuery(question);

        return vectorStore.similaritySearch(
                        SearchRequest.builder()
                                .query(retrievalQuery)
                                .topK(DEFAULT_TOP_K)
                                .filterExpression(filter)
                                .build()
                ).stream()
                .map(this::map)
                .toList();
    }

    private String buildRetrievalQuery(final String question) {
        return """
                Rental contract information.
                Contract terms and conditions.
                Contract term, duration, validity, start date, end date and expiration date.
                Fecha de inicio, término, duración, vigencia, finalización y vencimiento del contrato de locación.
                
                User question:
                %s
                """.formatted(question);
    }

    private RetrievedDocumentChunk map(final Document document) {
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