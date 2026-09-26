package com.rentmanager.ai.port.out;

import java.util.List;

import org.springframework.ai.document.Document;

public interface DocumentRetriever {
    List<Document> retrieve(String propertyId, String contractId, String question);
}
