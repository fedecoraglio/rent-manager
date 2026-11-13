# Source and operation coverage

Reviewed: 2026-11-13. Resolve paths from the repository root. Use this map to open the implementation; excerpts are navigation indexes, not replacements for complete source bodies. Line numbers refer to this review. Do not interpret internal operations as public endpoints.

## Contents

- File inventory
- Operations, data contracts, and rule locations

## File inventory

- `.github/workflows/backend-image.yml`
- `.github/workflows/frontend-image.yml`
- `platform/argocd/rent-manager.yaml`
- `platform/kubernetes/backend/configmap.yaml`
- `platform/kubernetes/backend/deployment.yaml`
- `platform/kubernetes/backend/service.yaml`
- `platform/kubernetes/frontend/configmap.yaml`
- `platform/kubernetes/frontend/deployment.yaml`
- `platform/kubernetes/frontend/service.yaml`
- `platform/kubernetes/ingress.yaml`
- `platform/kubernetes/kustomization.yaml`
- `platform/kubernetes/mysql/pvc.yaml`
- `platform/kubernetes/mysql/secret.example.yaml`
- `platform/kubernetes/mysql/service.yaml`
- `platform/kubernetes/mysql/statefulset.yaml`
- `platform/kubernetes/namespace.yaml`
- `rent-manager-ai/compose.yml`
- `rent-manager-ai/Dockerfile`
- `rent-manager-ai/src/main/java/com/rentmanager/ai/adapter/out/health/OllamaHealthIndicator.java`
- `rent-manager-ai/src/main/java/com/rentmanager/ai/adapter/out/health/PgVectorHealthIndicator.java`
- `rent-manager-ai/src/main/java/com/rentmanager/ai/config/AiConfiguration.java`
- `rent-manager-ai/src/main/resources/application.yml`
- `rent-manager-backend/docker-compose.yml`
- `rent-manager-backend/Dockerfile`
- `rent-manager-backend/internal/adapter/config/config.go`
- `rent-manager-backend/internal/adapter/store/mysql/db.go`
- `rent-manager-frontend/Dockerfile`
- `rent-manager-frontend/nginx.conf`
- `rent-manager-kubernetes-commands.md`

## Operations, data contracts, and rule locations

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

### rent-manager-backend/internal/adapter/config/config.go

```text
43: func New() (*Container, error) {
44: if os.Getenv("APP_ENV") != "production" {
46: if err != nil {
```

### rent-manager-backend/internal/adapter/store/mysql/db.go

```text
31: func New(ctx context.Context, config *config.DB) (*DB, error) {
51: if err != nil {
55: if err := db.PingContext(ctx); err != nil {
69: func (db *DB) Migrate() error {
71: if err != nil {
76: if err != nil {
81: if err != nil && !errors.Is(err, migrate.ErrNoChange) {
88: func (db *DB) ErrorCode(err error) string {
90: if errors.As(err, &mysqlErr) {
97: func (db *DB) Close() {
```
