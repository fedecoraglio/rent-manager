package com.rentmanager.ai.port.out;

import com.rentmanager.ai.domain.document.RetrievedDocumentChunk;

import java.util.List;

public interface DocumentRetriever {
    List<RetrievedDocumentChunk> retrieve(final String propertyId, final String contractId, final String question);
}
