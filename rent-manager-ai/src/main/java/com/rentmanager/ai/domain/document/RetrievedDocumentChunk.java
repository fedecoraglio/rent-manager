package com.rentmanager.ai.domain.document;

public record RetrievedDocumentChunk(
        String text,
        String documentId,
        String propertyId,
        String contractId,
        String documentType,
        String filename,
        Integer page
) {
}
