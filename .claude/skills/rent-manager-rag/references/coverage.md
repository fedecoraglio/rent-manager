# Source and operation coverage

Reviewed: 2026-11-13. Resolve paths from the repository root. Use this map to open the implementation; excerpts are navigation indexes, not replacements for complete source bodies. Line numbers refer to this review. Do not interpret internal operations as public endpoints.

## Contents

- File inventory
- Operations, data contracts, and rule locations

## File inventory

- `rent-manager-ai/src/main/java/com/rentmanager/ai/adapter/in/web/controller/DocumentController.java`
- `rent-manager-ai/src/main/java/com/rentmanager/ai/adapter/in/web/controller/DocumentExceptionHandler.java`
- `rent-manager-ai/src/main/java/com/rentmanager/ai/adapter/in/web/controller/RagController.java`
- `rent-manager-ai/src/main/java/com/rentmanager/ai/adapter/in/web/request/RagRequest.java`
- `rent-manager-ai/src/main/java/com/rentmanager/ai/adapter/in/web/response/DocumentResponse.java`
- `rent-manager-ai/src/main/java/com/rentmanager/ai/adapter/in/web/response/RagResponse.java`
- `rent-manager-ai/src/main/java/com/rentmanager/ai/adapter/out/ai/PdfChunker.java`
- `rent-manager-ai/src/main/java/com/rentmanager/ai/adapter/out/ai/PdfUploadValidator.java`
- `rent-manager-ai/src/main/java/com/rentmanager/ai/adapter/out/ai/SpringAiDocumentIndexer.java`
- `rent-manager-ai/src/main/java/com/rentmanager/ai/adapter/out/health/OllamaHealthIndicator.java`
- `rent-manager-ai/src/main/java/com/rentmanager/ai/adapter/out/health/PgVectorHealthIndicator.java`
- `rent-manager-ai/src/main/java/com/rentmanager/ai/adapter/out/persistence/postgres/JdbcDocumentRepository.java`
- `rent-manager-ai/src/main/java/com/rentmanager/ai/adapter/out/persistence/postgres/PgVectorDocumentRetriever.java`
- `rent-manager-ai/src/main/java/com/rentmanager/ai/application/document/DocumentService.java`
- `rent-manager-ai/src/main/java/com/rentmanager/ai/application/rag/RagService.java`
- `rent-manager-ai/src/main/java/com/rentmanager/ai/config/AiConfiguration.java`
- `rent-manager-ai/src/main/java/com/rentmanager/ai/domain/document/DocumentException.java`
- `rent-manager-ai/src/main/java/com/rentmanager/ai/domain/document/DocumentMetadata.java`
- `rent-manager-ai/src/main/java/com/rentmanager/ai/domain/document/DocumentScope.java`
- `rent-manager-ai/src/main/java/com/rentmanager/ai/domain/document/DocumentType.java`
- `rent-manager-ai/src/main/java/com/rentmanager/ai/domain/document/RetrievedDocumentChunk.java`
- `rent-manager-ai/src/main/java/com/rentmanager/ai/port/out/DocumentIndexer.java`
- `rent-manager-ai/src/main/java/com/rentmanager/ai/port/out/DocumentRepository.java`
- `rent-manager-ai/src/main/java/com/rentmanager/ai/port/out/DocumentRetriever.java`
- `rent-manager-ai/src/main/java/com/rentmanager/ai/port/out/DocumentStorage.java`
- `rent-manager-ai/src/main/java/com/rentmanager/ai/RentManagerAiApplication.java`
- `rent-manager-ai/src/main/resources/application.yml`
- `rent-manager-ai/src/main/resources/db/migration/V1__create_documents.sql`
- `rent-manager-ai/src/main/resources/db/migration/V2__create_vector_store.sql`

## Operations, data contracts, and rule locations

### rent-manager-ai/src/main/java/com/rentmanager/ai/adapter/in/web/controller/DocumentController.java

```text
26: @RequestMapping("/properties/{propertyId}/contracts/{contractId}/documents")
30: public DocumentController(final DocumentService service) {
34: @PostMapping(consumes = MediaType.MULTIPART_FORM_DATA_VALUE)
35: public ResponseEntity<DocumentResponse> upload(@PathVariable final String propertyId,
43: @GetMapping
44: public List<DocumentResponse> list(@PathVariable final String propertyId, @PathVariable final String contractId) {
48: @GetMapping("/{documentId}")
49: public ResponseEntity<byte[]> download(@PathVariable final String propertyId, @PathVariable final String contractId,
61: @DeleteMapping("/{documentId}")
62: public ResponseEntity<Void> delete(@PathVariable final String propertyId, @PathVariable final String contractId,
```

### rent-manager-ai/src/main/java/com/rentmanager/ai/adapter/in/web/controller/DocumentExceptionHandler.java

```text
16: private static final Logger LOG = LoggerFactory.getLogger(DocumentExceptionHandler.class);
19: public ProblemDetail documentError(final DocumentException exception) {
20: if (exception.status().is5xxServerError()) {
27: public ProblemDetail databaseError(final RuntimeException exception) {
29: return ProblemDetail.forStatusAndDetail(HttpStatus.SERVICE_UNAVAILABLE, "The document registry is unavailable. Please try again later.");
```

### rent-manager-ai/src/main/java/com/rentmanager/ai/adapter/in/web/controller/RagController.java

```text
10: @RequestMapping("/properties/{propertyId}/contracts/{contractId}/chat")
15: public RagController(final RagService ragService) {
19: @PostMapping
20: public ResponseEntity<RagResponse> ask(
```

### rent-manager-ai/src/main/java/com/rentmanager/ai/adapter/in/web/request/RagRequest.java

```text
3: public record RagRequest(String question) {
```

### rent-manager-ai/src/main/java/com/rentmanager/ai/adapter/in/web/response/DocumentResponse.java

```text
9: public record DocumentResponse(UUID id, DocumentType type, String filename, String contentType,
11: public static DocumentResponse from(final DocumentMetadata document) {
```

### rent-manager-ai/src/main/java/com/rentmanager/ai/adapter/in/web/response/RagResponse.java

```text
3: public record RagResponse(String answer) {
```

### rent-manager-ai/src/main/java/com/rentmanager/ai/adapter/out/ai/PdfChunker.java

```text
22: private static final Pattern CLAUSE_START = Pattern.compile(
26: public List<Document> chunks(final DocumentMetadata document, final byte[] pdf) {
32: throw new DocumentException(
33: HttpStatus.UNPROCESSABLE_CONTENT,
42: if (page.getText() == null || page.getText().isBlank()) {
60: if (!clause.isBlank()) {
79: if (chunks.isEmpty()) {
80: throw new DocumentException(
81: HttpStatus.UNPROCESSABLE_CONTENT,
89: private static List<String> splitClauses(final String text) {
93: private static String normalize(final String text) {
102: private CloseablePageReader(final byte[] pdf) {
107: public void close() throws IOException {
```

### rent-manager-ai/src/main/java/com/rentmanager/ai/adapter/out/ai/PdfUploadValidator.java

```text
17: public ValidatedPdf validate(final MultipartFile file) {
18: if (file.isEmpty()) {
19: throw DocumentException.invalid("The PDF file is empty.");
21: if (file.getSize() > MAX_BYTES) {
22: throw new DocumentException(HttpStatus.CONTENT_TOO_LARGE, "PDF files must not exceed 20 MiB.");
27: if (filename.isBlank() || filename.length() > 255 || filename.codePoints().anyMatch(Character::isISOControl)) {
28: throw DocumentException.invalid("Provide a filename of at most 255 characters without control characters.");
30: if (!filename.toLowerCase(Locale.ROOT).endsWith(".pdf")
32: throw new DocumentException(HttpStatus.UNSUPPORTED_MEDIA_TYPE, "Only application/pdf files with a .pdf filename are supported.");
36: if (content.length > MAX_BYTES) {
37: throw new DocumentException(HttpStatus.CONTENT_TOO_LARGE, "PDF files must not exceed 20 MiB.");
39: if (content.length < 5 || !new String(content, 0, 5, StandardCharsets.US_ASCII).equals("%PDF-")) {
40: throw DocumentException.invalid("The uploaded file is not a valid PDF.");
43: if (pdf.isEncrypted() || pdf.getNumberOfPages() == 0) {
44: throw DocumentException.invalid("The PDF must have at least one page and must not be encrypted.");
49: throw new DocumentException(HttpStatus.BAD_REQUEST, "The PDF could not be read; upload a valid, unencrypted PDF.", exception);
53: public record ValidatedPdf(String filename, byte[] content) {
```

### rent-manager-ai/src/main/java/com/rentmanager/ai/adapter/out/ai/SpringAiDocumentIndexer.java

```text
9: import org.springframework.ai.vectorstore.filter.FilterExpressionBuilder;
16: private static final Logger LOG = LoggerFactory.getLogger(SpringAiDocumentIndexer.class);
21: public SpringAiDocumentIndexer(final PdfChunker chunker, final VectorStore vectorStore) {
26: public void index(final DocumentMetadata document, final byte[] pdf) {
36: throw new DocumentException(HttpStatus.SERVICE_UNAVAILABLE,
42: public void delete(final DocumentMetadata document) {
50: throw new DocumentException(HttpStatus.SERVICE_UNAVAILABLE, "Document chunks could not be deleted. Please retry.", exception);
54: private static void requireTransaction() {
55: if (!TransactionSynchronizationManager.isActualTransactionActive()) {
56: throw new IllegalStateException("Document indexing must run inside the registry transaction");
```

### rent-manager-ai/src/main/java/com/rentmanager/ai/adapter/out/health/OllamaHealthIndicator.java

```text
16: public OllamaHealthIndicator(@Value("${spring.ai.ollama.base-url}") final String baseUrl) {
24: public Health health() {
```

### rent-manager-ai/src/main/java/com/rentmanager/ai/adapter/out/health/PgVectorHealthIndicator.java

```text
12: public PgVectorHealthIndicator(final JdbcTemplate jdbc) {
17: public Health health() {
```

### rent-manager-ai/src/main/java/com/rentmanager/ai/adapter/out/persistence/postgres/JdbcDocumentRepository.java

```text
23: public JdbcDocumentRepository(final JdbcTemplate jdbc) {
27: public void insert(final DocumentMetadata document) {
37: public List<DocumentMetadata> list(final DocumentScope scope) {
44: public Optional<DocumentMetadata> find(final DocumentScope scope, final UUID id, final boolean lock) {
50: public void delete(final DocumentScope scope, final UUID id) {
53: if (deleted != 1) {
54: throw DocumentException.notFound();
58: private static DocumentMetadata map(final ResultSet row, final int index) throws SQLException {
```

### rent-manager-ai/src/main/java/com/rentmanager/ai/adapter/out/persistence/postgres/PgVectorDocumentRetriever.java

```text
11: import org.springframework.ai.vectorstore.filter.FilterExpressionBuilder;
21: public PgVectorDocumentRetriever(final VectorStore vectorStore) {
26: public List<RetrievedDocumentChunk> retrieve(
43: .topK(DEFAULT_TOP_K)
44: .filterExpression(filter)
51: private String buildRetrievalQuery(final String question) {
63: private RetrievedDocumentChunk map(final Document document) {
77: private String value(final Map<String, Object> metadata, final String key) {
82: private Integer integerValue(final Map<String, Object> metadata, final String key) {
85: if (value == null) {
89: if (value instanceof Number number) {
```

### rent-manager-ai/src/main/java/com/rentmanager/ai/application/document/DocumentService.java

```text
29: private static final Logger LOG = LoggerFactory.getLogger(DocumentService.class);
36: public DocumentService(final DocumentRepository repository, final DocumentStorage storage,
48: public DocumentMetadata upload(final DocumentScope scope, final DocumentType type, final MultipartFile file) {
57: public void afterCompletion(final int completion) {
58: if (completion == STATUS_ROLLED_BACK) {
60: } else if (completion == STATUS_UNKNOWN) {
74: public List<DocumentMetadata> list(final DocumentScope scope) {
78: public Download download(final DocumentScope scope, final UUID id) {
83: public void delete(final DocumentScope scope, final UUID id) {
89: if (staged) {
92: public void afterCommit() {
98: throw new DocumentException(HttpStatus.SERVICE_UNAVAILABLE,
104: public void afterCompletion(final int completion) {
105: if (completion == STATUS_ROLLED_BACK) {
107: } else if (completion == STATUS_UNKNOWN) {
117: private static void cleanup(final UUID id, final Runnable action) {
126: public record Download(DocumentMetadata document, byte[] content) {
```

### rent-manager-ai/src/main/java/com/rentmanager/ai/application/rag/RagService.java

```text
17: public RagService(final DocumentRetriever documentRetriever, final ChatClient chatClient) {
22: public String ask(final String propertyId, final String contractId, final String question) {
26: if (chunks.isEmpty()) {
```

### rent-manager-ai/src/main/java/com/rentmanager/ai/config/AiConfiguration.java

```text
27: public EmbeddingModel embeddingModel(
44: .model("nomic-embed-text")
53: public VectorStore vectorStore(final JdbcTemplate jdbc, final EmbeddingModel embeddingModel) {
56: .dimensions(768)
57: .distanceType(PgVectorStore.PgDistanceType.COSINE_DISTANCE)
58: .indexType(PgVectorStore.PgIndexType.NONE)
65: public ChatModel chatModel(
82: .model("qwen2.5:3b-instruct")
90: public ChatClient chatClient(final ChatModel chatModel) {
```

### rent-manager-ai/src/main/java/com/rentmanager/ai/domain/document/DocumentException.java

```text
8: public DocumentException(final HttpStatus status, final String message) {
13: public DocumentException(final HttpStatus status, final String message, final Throwable cause) {
18: public HttpStatus status() {
22: public static DocumentException invalid(final String message) {
23: return new DocumentException(HttpStatus.BAD_REQUEST, message);
26: public static DocumentException notFound() {
27: return new DocumentException(HttpStatus.NOT_FOUND, "Document not found in this property and contract.");
```

### rent-manager-ai/src/main/java/com/rentmanager/ai/domain/document/DocumentMetadata.java

```text
6: public record DocumentMetadata(UUID id, String propertyId, String contractId, DocumentType type,
```

### rent-manager-ai/src/main/java/com/rentmanager/ai/domain/document/DocumentScope.java

```text
5: public record DocumentScope(String propertyId, String contractId) {
6: private static final Pattern REFERENCE = Pattern.compile("[A-Za-z0-9_-]{1,64}");
9: if (propertyId == null || contractId == null
11: throw DocumentException.invalid("Property and contract references must contain 1–64 letters, digits, underscores or hyphens.");
```

### rent-manager-ai/src/main/java/com/rentmanager/ai/domain/document/RetrievedDocumentChunk.java

```text
3: public record RetrievedDocumentChunk(
```

### rent-manager-ai/src/main/java/com/rentmanager/ai/RentManagerAiApplication.java

```text
8: public static void main(final String[] args) {
```

### rent-manager-ai/src/main/resources/db/migration/V1__create_documents.sql

```text
1: CREATE TABLE documents (
2: document_id UUID PRIMARY KEY,
3: property_id VARCHAR(64) NOT NULL,
4: contract_id VARCHAR(64) NOT NULL,
5: document_type VARCHAR(32) NOT NULL CHECK (document_type IN
7: original_filename VARCHAR(255) NOT NULL,
8: stored_filename VARCHAR(40) NOT NULL,
9: storage_path VARCHAR(512) NOT NULL UNIQUE,
10: content_type VARCHAR(100) NOT NULL CHECK (content_type = 'application/pdf'),
11: file_size BIGINT NOT NULL CHECK (file_size > 0 AND file_size <= 20971520),
12: created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
13: updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
16: CREATE INDEX documents_contract_idx ON documents (property_id, contract_id, created_at, document_id);
```

### rent-manager-ai/src/main/resources/db/migration/V2__create_vector_store.sql

```text
1: CREATE EXTENSION IF NOT EXISTS vector;
5: CREATE TABLE vector_store (
6: id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
7: content TEXT NOT NULL CHECK (length(trim(content)) > 0),
8: metadata JSONB NOT NULL CHECK (
11: embedding VECTOR(768) NOT NULL,
13: document_id UUID GENERATED ALWAYS AS ((metadata ->> 'documentId')::uuid) STORED NOT NULL
14: REFERENCES documents(document_id) ON DELETE CASCADE
17: CREATE INDEX vector_store_document_idx ON vector_store (document_id);
18: CREATE INDEX vector_store_metadata_idx ON vector_store USING GIN (metadata jsonb_path_ops);
```
