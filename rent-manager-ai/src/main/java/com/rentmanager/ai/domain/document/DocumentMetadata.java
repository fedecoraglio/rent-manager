package com.rentmanager.ai.domain.document;

import java.time.Instant;
import java.util.UUID;

public record DocumentMetadata(UUID id, String propertyId, String contractId, DocumentType type,
        String originalFilename, String storedFilename, String storagePath, String contentType,
        long size, Instant createdAt, Instant updatedAt) {
}
