package com.rentmanager.ai.document;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.dao.DataAccessException;
import org.springframework.http.HttpStatus;
import org.springframework.http.ProblemDetail;
import org.springframework.transaction.TransactionException;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.RestControllerAdvice;
import org.springframework.web.servlet.mvc.method.annotation.ResponseEntityExceptionHandler;

@RestControllerAdvice
public final class DocumentExceptionHandler extends ResponseEntityExceptionHandler {
    private static final Logger LOG = LoggerFactory.getLogger(DocumentExceptionHandler.class);

    @ExceptionHandler(DocumentException.class)
    public ProblemDetail documentError(final DocumentException exception) {
        if (exception.status().is5xxServerError()) {
            LOG.error("Document storage operation failed", exception);
        }
        return ProblemDetail.forStatusAndDetail(exception.status(), exception.getMessage());
    }

    @ExceptionHandler({DataAccessException.class, TransactionException.class})
    public ProblemDetail databaseError(final RuntimeException exception) {
        LOG.error("Document registry operation failed", exception);
        return ProblemDetail.forStatusAndDetail(HttpStatus.SERVICE_UNAVAILABLE, "The document registry is unavailable. Please try again later.");
    }
}
