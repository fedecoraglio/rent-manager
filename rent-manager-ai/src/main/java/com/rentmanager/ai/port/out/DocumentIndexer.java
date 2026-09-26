package com.rentmanager.ai.port.out;

import com.rentmanager.ai.domain.document.DocumentMetadata;

public interface DocumentIndexer {
    void index(DocumentMetadata document, byte[] pdf);
    void delete(DocumentMetadata document);
}
