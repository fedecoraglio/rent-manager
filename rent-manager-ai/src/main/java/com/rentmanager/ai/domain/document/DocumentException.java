package com.rentmanager.ai.domain.document;

import org.springframework.http.HttpStatus;

public final class DocumentException extends RuntimeException {
    private final HttpStatus status;

    public DocumentException(final HttpStatus status, final String message) {
        super(message);
        this.status = status;
    }

    public DocumentException(final HttpStatus status, final String message, final Throwable cause) {
        super(message, cause);
        this.status = status;
    }

    public HttpStatus status() {
        return status;
    }

    public static DocumentException invalid(final String message) {
        return new DocumentException(HttpStatus.BAD_REQUEST, message);
    }

    public static DocumentException notFound() {
        return new DocumentException(HttpStatus.NOT_FOUND, "Document not found in this property and contract.");
    }
}
