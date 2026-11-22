-- Never discard existing contracts to enforce uniqueness. Resolve duplicates explicitly first.
DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM documents WHERE document_type = 'RENT_CONTRACT'
               GROUP BY property_id, contract_id HAVING count(*) > 1) THEN
        RAISE EXCEPTION 'Multiple RENT_CONTRACT documents exist for a scope. Back up and resolve duplicates before migrating.';
    END IF;
END $$;

CREATE TEMP TABLE document_id_migration ON COMMIT DROP AS
SELECT document_id::text AS old_id,
       replace(property_id, '_', '~') || '_' || replace(contract_id, '_', '~') || '_' || document_type
       || CASE WHEN document_type = 'RENT_CONTRACT' THEN '' ELSE '_' || document_id::text END AS new_id
FROM documents;

-- Rebuild the generated FK as text; chunk UUIDs and physical file paths are unchanged.
ALTER TABLE vector_store DROP COLUMN document_id;
ALTER TABLE documents ALTER COLUMN document_id TYPE text USING document_id::text;
UPDATE documents d SET document_id = m.new_id
FROM document_id_migration m WHERE d.document_id = m.old_id;
UPDATE vector_store v SET metadata = jsonb_set(v.metadata, '{documentId}', to_jsonb(m.new_id))
FROM document_id_migration m WHERE v.metadata ->> 'documentId' = m.old_id;
ALTER TABLE vector_store ADD COLUMN document_id text
    GENERATED ALWAYS AS (metadata ->> 'documentId') STORED NOT NULL
    REFERENCES documents(document_id) ON DELETE CASCADE;
CREATE INDEX vector_store_document_idx ON vector_store (document_id);
CREATE UNIQUE INDEX documents_single_rent_contract_idx
    ON documents (property_id, contract_id) WHERE document_type = 'RENT_CONTRACT';
