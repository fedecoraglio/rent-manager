package com.rentmanager.ai.indexing;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.Arrays;
import java.util.List;
import java.util.UUID;
import java.util.concurrent.atomic.AtomicInteger;
import java.util.stream.IntStream;

import com.rentmanager.ai.document.DocumentException;
import com.rentmanager.ai.document.DocumentRepository;
import com.rentmanager.ai.document.DocumentScope;
import com.rentmanager.ai.document.DocumentService;
import com.rentmanager.ai.document.DocumentType;
import com.rentmanager.ai.document.FileSystemDocumentStorage;
import com.rentmanager.ai.document.PdfFixtures;
import com.rentmanager.ai.document.PdfUploadValidator;
import org.flywaydb.core.Flyway;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.condition.EnabledIfEnvironmentVariable;
import org.junit.jupiter.api.io.TempDir;
import org.springframework.ai.document.Document;
import org.springframework.ai.embedding.Embedding;
import org.springframework.ai.embedding.EmbeddingModel;
import org.springframework.ai.embedding.EmbeddingRequest;
import org.springframework.ai.embedding.EmbeddingResponse;
import org.springframework.ai.vectorstore.pgvector.PgVectorStore;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.jdbc.datasource.DataSourceTransactionManager;
import org.springframework.jdbc.datasource.DriverManagerDataSource;
import org.springframework.mock.web.MockMultipartFile;

import static org.assertj.core.api.Assertions.assertThat;
import static org.assertj.core.api.Assertions.assertThatThrownBy;

// Opt-in real PostgreSQL tests: mvn -Dtest=DocumentIndexingIT test (see README).
// A deterministic embedding model lets us fail AFTER the first vector batch was written.
@EnabledIfEnvironmentVariable(named = "TEST_POSTGRES_URL", matches = ".+")
final class DocumentIndexingIT {
    private final String schema = "indexing_test_" + UUID.randomUUID().toString().replace("-", "");
    private final DocumentScope scope = new DocumentScope("integration-property", "contract-A");
    private final AtomicInteger calls = new AtomicInteger();
    private final AtomicInteger failOnCall = new AtomicInteger(-1);
    private final AtomicInteger rowsBeforeFailure = new AtomicInteger();
    @TempDir
    private Path storageRoot;
    private JdbcTemplate jdbc;
    private DocumentService service;

    @BeforeEach
    void prepareIsolatedSchema() throws IOException {
        final String baseUrl = System.getenv("TEST_POSTGRES_URL");
        final String url = baseUrl + (baseUrl.contains("?") ? "&" : "?") + "currentSchema=" + schema + ",public";
        final var dataSource = new DriverManagerDataSource(url, System.getenv("POSTGRES_USER"), System.getenv("POSTGRES_PASSWORD"));
        jdbc = new JdbcTemplate(dataSource);
        Flyway.configure().dataSource(dataSource).schemas(schema).defaultSchema(schema).load().migrate();
        final EmbeddingModel embeddingModel = new EmbeddingModel() {
            @Override
            public EmbeddingResponse call(final EmbeddingRequest request) {
                if (calls.incrementAndGet() == failOnCall.get()) {
                    rowsBeforeFailure.set(jdbc.queryForObject("SELECT count(*) FROM vector_store", Integer.class));
                    throw new IllegalStateException("Simulated Ollama failure on later batch");
                }
                final List<Embedding> embeddings = IntStream.range(0, request.getInstructions().size())
                        .mapToObj(index -> {
                            final float[] vector = new float[768];
                            vector[0] = 1;
                            return new Embedding(vector, index);
                        }).toList();
                return new EmbeddingResponse(embeddings);
            }

            @Override
            public float[] embed(final Document document) {
                return embed(document.getText());
            }
        };
        final var vectorStore = PgVectorStore.builder(jdbc, embeddingModel).schemaName(schema)
                .dimensions(768).initializeSchema(false).build();
        service = new DocumentService(new DocumentRepository(jdbc), new FileSystemDocumentStorage(storageRoot.toString()),
                new PdfUploadValidator(), new DocumentIndexer(new PdfChunker(), vectorStore),
                new DataSourceTransactionManager(dataSource), 300);
    }

    @AfterEach
    void removeOnlyThisTestsSchema() {
        if (jdbc != null) {
            // schema is exclusively generated here, never supplied by an environment variable or user.
            jdbc.execute("DROP SCHEMA " + schema + " CASCADE");
        }
    }

    @Test
    void rollsBackEarlierVectorBatchesRegistryAndFileWhenLaterEmbeddingsFail() throws IOException {
        final String[] pages = new String[17];
        Arrays.fill(pages, "This rental contract ends on 31 December 2027. Rent is adjusted quarterly.");
        failOnCall.set(2);
        assertThatThrownBy(() -> service.upload(scope, DocumentType.RENT_CONTRACT, upload(pages)))
                .isInstanceOf(DocumentException.class).hasMessageContaining("indexing failed");
        assertThat(rowsBeforeFailure.get()).isEqualTo(16);
        assertThat(jdbc.queryForObject("SELECT count(*) FROM vector_store", Integer.class)).isZero();
        assertThat(service.list(scope)).isEmpty();
        try (final var files = Files.walk(storageRoot)) {
            assertThat(files.filter(Files::isRegularFile).count()).isZero();
        }
    }

    @Test
    void persistsScopedPageMetadataAndDeletionCannotTouchAnotherDocumentsChunks() throws IOException {
        final var first = service.upload(scope, DocumentType.RENT_CONTRACT, upload("Rent is adjusted quarterly.", "The guarantor is Juan Perez."));
        final var otherScope = new DocumentScope(scope.propertyId(), "contract-B");
        final var second = service.upload(otherScope, DocumentType.GUARANTEE, upload("The guarantee lasts for two years."));
        assertThat(jdbc.queryForList("SELECT metadata->>'page' FROM vector_store WHERE document_id = ? ORDER BY metadata->>'page'",
                String.class, first.id())).containsExactly("1", "2");
        assertThat(jdbc.queryForObject("""
                SELECT count(*) FROM vector_store WHERE document_id = ?
                AND metadata->>'propertyId' = ? AND metadata->>'contractId' = ?
                AND metadata->>'documentType' = 'RENT_CONTRACT' AND metadata->>'filename' = 'contract.pdf'
                AND vector_dims(embedding) = 768
                """, Integer.class, first.id(), scope.propertyId(), scope.contractId())).isEqualTo(2);
        assertThatThrownBy(() -> service.delete(otherScope, first.id())).isInstanceOf(DocumentException.class);
        assertThat(jdbc.queryForObject("SELECT count(*) FROM vector_store", Integer.class)).isEqualTo(3);
        final int embeddingCalls = calls.get();
        service.delete(scope, first.id());
        assertThat(jdbc.queryForObject("SELECT count(*) FROM vector_store WHERE document_id = ?", Integer.class, first.id())).isZero();
        assertThat(jdbc.queryForObject("SELECT count(*) FROM vector_store WHERE document_id = ?", Integer.class, second.id())).isEqualTo(1);
        assertThat(service.download(otherScope, second.id()).content()).isNotEmpty();
        assertThat(calls.get()).isEqualTo(embeddingCalls); // Deletion needs PostgreSQL, not Ollama.
    }

    @Test
    void restoresChunksAndFileIfRegistryDeletionFails() throws IOException {
        final var document = service.upload(scope, DocumentType.OTHER, upload("Pets are allowed with written permission."));
        // A referencing test table forces the final registry DELETE to fail after the chunk DELETE.
        jdbc.execute("CREATE TABLE deletion_blocker (document_id UUID REFERENCES documents(document_id))");
        jdbc.update("INSERT INTO deletion_blocker VALUES (?)", document.id());
        assertThatThrownBy(() -> service.delete(scope, document.id())).isInstanceOf(RuntimeException.class);
        assertThat(jdbc.queryForObject("SELECT count(*) FROM vector_store WHERE document_id = ?", Integer.class, document.id())).isEqualTo(1);
        assertThat(service.download(scope, document.id()).content()).isNotEmpty();
        assertThat(storageRoot.resolve(document.storagePath() + ".deleting")).doesNotExist();
    }

    private static MockMultipartFile upload(final String... pages) throws IOException {
        return new MockMultipartFile("file", "contract.pdf", "application/pdf", PdfFixtures.textPdf(pages));
    }
}
