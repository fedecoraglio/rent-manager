CREATE TABLE documents (
    document_id UUID PRIMARY KEY,
    property_id VARCHAR(64) NOT NULL,
    contract_id VARCHAR(64) NOT NULL,
    document_type VARCHAR(32) NOT NULL CHECK (document_type IN
        ('RENT_CONTRACT', 'GUARANTEE', 'ANNEX', 'INVENTORY', 'HANDOVER', 'OTHER')),
    original_filename VARCHAR(255) NOT NULL,
    stored_filename VARCHAR(40) NOT NULL,
    storage_path VARCHAR(512) NOT NULL UNIQUE,
    content_type VARCHAR(100) NOT NULL CHECK (content_type = 'application/pdf'),
    file_size BIGINT NOT NULL CHECK (file_size > 0 AND file_size <= 20971520),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX documents_contract_idx ON documents (property_id, contract_id, created_at, document_id);
