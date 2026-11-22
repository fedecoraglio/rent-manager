package com.rentmanager.ai.domain.document;

import java.time.Instant;


public record DocumentMetadata(String id, String propertyId, String contractId, DocumentType type,
        String originalFilename, String storedFilename, String storagePath, String contentType,
        long size, Instant createdAt, Instant updatedAt) {
}
