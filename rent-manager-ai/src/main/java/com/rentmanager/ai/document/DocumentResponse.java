package com.rentmanager.ai.document;

import java.time.Instant;
import java.util.UUID;

public record DocumentResponse(UUID id, DocumentType type, String filename, String contentType,
        long size, Instant createdAt) {
    public static DocumentResponse from(final DocumentMetadata document) {
        return new DocumentResponse(document.id(), document.type(), document.originalFilename(),
                document.contentType(), document.size(), document.createdAt());
    }
}
