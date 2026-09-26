package com.rentmanager.ai.port.out;

import java.util.List;

import com.rentmanager.ai.domain.document.RetrievedDocumentChunk;
import org.springframework.ai.document.Document;

public interface DocumentRetriever {
    List<RetrievedDocumentChunk> retrieve(final String propertyId, final String contractId, final String question);
}
