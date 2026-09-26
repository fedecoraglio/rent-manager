package com.rentmanager.ai.document;

import java.util.UUID;

public interface DocumentStorage {
    String save(DocumentScope scope, UUID id, byte[] content);
    byte[] read(String path);
    void remove(String path);
    boolean stageDeletion(String path);
    void restoreDeletion(String path);
    void completeDeletion(String path);
}
