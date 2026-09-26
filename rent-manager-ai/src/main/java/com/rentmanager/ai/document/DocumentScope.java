package com.rentmanager.ai.document;

import java.util.regex.Pattern;

public record DocumentScope(String propertyId, String contractId) {
    private static final Pattern REFERENCE = Pattern.compile("[A-Za-z0-9_-]{1,64}");

    public DocumentScope {
        if (propertyId == null || contractId == null
                || !REFERENCE.matcher(propertyId).matches() || !REFERENCE.matcher(contractId).matches()) {
            throw DocumentException.invalid("Property and contract references must contain 1–64 letters, digits, underscores or hyphens.");
        }
    }
}
