package com.rentmanager.ai.adapter.in.web.response;

import java.time.Instant;


import com.rentmanager.ai.domain.document.DocumentMetadata;
import com.rentmanager.ai.domain.document.DocumentType;

public record DocumentResponse(String id, DocumentType type, String filename, String contentType,
                               long size, Instant createdAt) {
    public static DocumentResponse from(final DocumentMetadata document) {
        return new DocumentResponse(document.id(), document.type(), document.originalFilename(),
                document.contentType(), document.size(), document.createdAt());
    }
}
