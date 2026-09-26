package com.rentmanager.ai.adapter.out.persistence.postgres;

import java.util.List;

import org.springframework.ai.document.Document;
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
    public List<Document> retrieve(final String propertyId, final String contractId, final String question) {
        final var filters = new FilterExpressionBuilder();
        final var filter = filters.and(filters.eq("propertyId", propertyId), filters.eq("contractId", contractId)).build();
        return vectorStore.similaritySearch(SearchRequest.builder().query(question).topK(DEFAULT_TOP_K).filterExpression(filter).build());
    }
}
