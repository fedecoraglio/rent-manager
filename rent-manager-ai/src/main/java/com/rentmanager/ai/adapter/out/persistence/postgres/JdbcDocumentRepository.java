package com.rentmanager.ai.adapter.out.persistence.postgres;

import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Timestamp;
import java.util.List;
import java.util.Optional;
import java.util.UUID;

import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Component;
import com.rentmanager.ai.domain.document.DocumentException;
import com.rentmanager.ai.domain.document.DocumentMetadata;
import com.rentmanager.ai.domain.document.DocumentScope;
import com.rentmanager.ai.domain.document.DocumentType;
import com.rentmanager.ai.port.out.DocumentRepository;

// JdbcTemplate already translates SQL exceptions; no repository proxy is needed.
@Component
public final class JdbcDocumentRepository implements DocumentRepository {
    private final JdbcTemplate jdbc;

    public JdbcDocumentRepository(final JdbcTemplate jdbc) {
        this.jdbc = jdbc;
    }

    public void insert(final DocumentMetadata document) {
        jdbc.update("""
                        INSERT INTO documents (document_id, property_id, contract_id, document_type,
                            original_filename, stored_filename, storage_path, content_type, file_size, created_at, updated_at)
                        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
                        """, document.id(), document.propertyId(), document.contractId(), document.type().name(),
                document.originalFilename(), document.storedFilename(), document.storagePath(),
                document.contentType(), document.size(), Timestamp.from(document.createdAt()), Timestamp.from(document.updatedAt()));
    }

    public List<DocumentMetadata> list(final DocumentScope scope) {
        return jdbc.query("""
                SELECT * FROM documents WHERE property_id = ? AND contract_id = ?
                ORDER BY created_at DESC, document_id
                """, JdbcDocumentRepository::map, scope.propertyId(), scope.contractId());
    }

    public Optional<DocumentMetadata> find(final DocumentScope scope, final UUID id, final boolean lock) {
        final String sql = "SELECT * FROM documents WHERE property_id = ? AND contract_id = ? AND document_id = ?"
                + (lock ? " FOR UPDATE" : "");
        return jdbc.query(sql, JdbcDocumentRepository::map, scope.propertyId(), scope.contractId(), id).stream().findFirst();
    }

    public void delete(final DocumentScope scope, final UUID id) {
        final int deleted = jdbc.update("DELETE FROM documents WHERE property_id = ? AND contract_id = ? AND document_id = ?",
                scope.propertyId(), scope.contractId(), id);
        if (deleted != 1) {
            throw DocumentException.notFound();
        }
    }

    private static DocumentMetadata map(final ResultSet row, final int index) throws SQLException {
        return new DocumentMetadata(row.getObject("document_id", UUID.class), row.getString("property_id"),
                row.getString("contract_id"), DocumentType.valueOf(row.getString("document_type")),
                row.getString("original_filename"), row.getString("stored_filename"), row.getString("storage_path"),
                row.getString("content_type"), row.getLong("file_size"), row.getTimestamp("created_at").toInstant(),
                row.getTimestamp("updated_at").toInstant());
    }
}
