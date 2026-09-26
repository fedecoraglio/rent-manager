CREATE EXTENSION IF NOT EXISTS vector;

-- Spring AI owns content/metadata/embedding mapping. Flyway owns the schema.
-- nomic-embed-text produces 768-dimensional vectors. Changing models requires reindexing.
CREATE TABLE vector_store (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    content TEXT NOT NULL CHECK (length(trim(content)) > 0),
    metadata JSONB NOT NULL CHECK (
        metadata ?& ARRAY['documentId', 'propertyId', 'contractId', 'documentType', 'filename', 'page']
    ),
    embedding VECTOR(768) NOT NULL,
    -- Derive the FK from metadata instead of maintaining a second document identifier in Java.
    document_id UUID GENERATED ALWAYS AS ((metadata ->> 'documentId')::uuid) STORED NOT NULL
        REFERENCES documents(document_id) ON DELETE CASCADE
);

CREATE INDEX vector_store_document_idx ON vector_store (document_id);
CREATE INDEX vector_store_metadata_idx ON vector_store USING GIN (metadata jsonb_path_ops);
-- Begin with exact similarity search; approximate indexes can be evaluated with retrieval in Phase 4.
