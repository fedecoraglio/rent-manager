package com.rentmanager.ai.domain.document;

import java.util.UUID;

public final class DocumentId {
    private DocumentId() { }

    public static String create(final DocumentScope scope, final DocumentType type) {
        // References permit underscores but not tildes: escape separators without collisions.
        final String base = scope.propertyId().replace('_', '~') + "_"
                + scope.contractId().replace('_', '~') + "_" + type.name();
        return type == DocumentType.RENT_CONTRACT ? base : base + "_" + UUID.randomUUID();
    }
}
