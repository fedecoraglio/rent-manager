package com.rentmanager.ai.adapter.in.web;

import java.time.Instant;
import java.util.UUID;

import com.rentmanager.ai.domain.document.DocumentMetadata;
import com.rentmanager.ai.domain.document.DocumentType;

public record DocumentResponse(UUID id, DocumentType type, String filename, String contentType,
                               long size, Instant createdAt) {
    public static DocumentResponse from(final DocumentMetadata document) {
        return new DocumentResponse(document.id(), document.type(), document.originalFilename(),
                document.contentType(), document.size(), document.createdAt());
    }
}
