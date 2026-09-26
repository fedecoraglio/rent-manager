package com.rentmanager.ai.port.out;

import java.util.UUID;
import com.rentmanager.ai.domain.document.DocumentScope;

public interface DocumentStorage {
    String save(DocumentScope scope, UUID id, byte[] content);
    byte[] read(String path);
    void remove(String path);
    boolean stageDeletion(String path);
    void restoreDeletion(String path);
    void completeDeletion(String path);
}
