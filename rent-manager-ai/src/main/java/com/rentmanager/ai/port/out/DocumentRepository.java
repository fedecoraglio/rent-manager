package com.rentmanager.ai.port.out;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

import com.rentmanager.ai.domain.document.DocumentMetadata;
import com.rentmanager.ai.domain.document.DocumentScope;

public interface DocumentRepository {
    void insert(DocumentMetadata document);
    List<DocumentMetadata> list(DocumentScope scope);
    Optional<DocumentMetadata> find(DocumentScope scope, UUID id, boolean lock);
    void delete(DocumentScope scope, UUID id);
}
