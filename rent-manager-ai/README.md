# Rent Manager AI

Phases 1–2: an independent Java 21 service with PostgreSQL/pgvector and Ollama connectivity,
plus a PDF document registry and filesystem storage. No indexing or chat exists yet.
The Go backend, MySQL, Angular and platform are unchanged.

## Versions and scope

- Java 21; Spring Boot 4.0.8; Maven 3.9.11 in the Docker build.
- Spring AI BOM 2.0.1 pins future AI modules. A BOM manages dependency versions; it does not
  enable models, download models, create vectors or execute RAG.
- PostgreSQL 17 with pgvector 0.8.2; Ollama 0.11.11, running on CPU by default.
- Flyway manages the document schema; Apache PDFBox 3.0.8 validates PDF structure without
  extracting text or generating embeddings. Plain JDBC handles registry queries.
- Model configuration reserves `qwen3:4b` for generation and `nomic-embed-text` for embeddings.
  These settings take effect when the Spring AI model starters are added in later phases.

[Spring AI compatibility](https://docs.spring.io/spring-ai/reference/getting-started.html)
lists Boot 4.0.x and 4.1.x for AI 2.0.x.
[Boot requirements](https://docs.spring.io/spring-boot/4.0/system-requirements.html)
support Java 21. No preview dependencies or paid model providers are used.

## Start locally with Docker

Run these commands from `rent-manager-ai`. Docker Desktop must be running in Linux container mode.
The image build runs `mvn verify`, including tests, so local Maven and Java are optional.

1. Copy `.env.example` to `.env` (`Copy-Item .env.example .env` in PowerShell).
2. Set `POSTGRES_PASSWORD` to a local password. `.env` is ignored by Git and Docker builds.
3. Run:

```sh
docker compose up -d --build
curl http://localhost:8081/actuator/health
curl http://localhost:8081/actuator/health/readiness
curl http://localhost:8081/actuator/health/liveness
```

In PowerShell, use `Invoke-RestMethod http://localhost:8081/actuator/health/readiness`.
The application may take several seconds to start after Compose returns.
Readiness must return HTTP 200 and `UP` for `db`, `pgvector`, `ollama` and `readinessState`.
`db` checks JDBC connectivity; `pgvector` checks the installed extension; `ollama` calls
`GET /api/tags` with bounded connection/read timeouts. Readiness returns HTTP 503 on dependency failure.
Liveness checks only the application so a dependency outage does not mean the JVM must restart.
Only health endpoints are exposed through Actuator, with component status but no connection/error details.

Additional checks:

```sh
docker compose exec ollama ollama list
docker compose logs rent-manager-ai
```

For a simpler SQL check in any shell, run `docker compose exec postgres psql -U rent_manager_ai -d rent_manager_ai`
(substitute your configured user/database), then execute:

```sql
SELECT extname, extversion FROM pg_extension WHERE extname = 'vector';
```

All published ports bind to localhost. Default host ports are 8081 (AI), 5433 (PostgreSQL),
and 11434 (Ollama); change `AI_PORT`, `POSTGRES_PORT`, or `OLLAMA_PORT` in `.env` if occupied.
The containers use their internal service names/ports regardless of host port mappings.
This is a local learning environment; authentication and authorization are not implemented.

## Download the models

```sh
docker compose exec ollama ollama pull qwen3:4b
docker compose exec ollama ollama pull nomic-embed-text
docker compose exec ollama ollama list
```

Model downloads require internet access and several GB of disk space. Once downloaded,
inference can run locally without cloud APIs. Phase 1 health checks verify the Ollama server,
not model availability or inference quality; model pulls are deliberately explicit.
Allow several GB of available RAM for Qwen and additional memory for the other containers.
The initial CPU setup requires no GPU; inference performance depends on the host.

## Run Java on the host instead

Install Java 21 and Maven 3.9.x, then start dependencies:

```sh
docker compose up -d postgres ollama
```

Set environment variables in the shell/IDE running Java. Compose reads `.env`; Spring Boot
does **not** automatically read that file. Export its database values yourself.

| Variable | Host default / requirement |
| --- | --- |
| `POSTGRES_HOST` | `localhost` |
| `POSTGRES_PORT` | `5433` |
| `POSTGRES_DB` | `rent_manager_ai` |
| `POSTGRES_USER` | Required; match `.env` |
| `POSTGRES_PASSWORD` | Required; match `.env` |
| `OLLAMA_BASE_URL` | `http://localhost:11434` |
| `DOCUMENT_STORAGE_PATH` | `./storage` |
| `SERVER_PORT` | `8081` |

```sh
mvn verify
mvn spring-boot:run
```

Stop the Compose application container first if it occupies the same port:
`docker compose stop rent-manager-ai`.
No JPA, MySQL access or business-domain replication is introduced. JDBC is sufficient for
connectivity and the document registry.

## Persistence and shutdown

Named volumes preserve PostgreSQL data, downloaded models and `/app/storage` across container
restarts/recreation. The document volume is mounted and writable by the non-root application user;
PDFs are stored there. Set `DOCUMENT_STORAGE_PATH` for host execution. Compose
intentionally fixes it to the mounted `/app/storage` directory.

```sh
docker compose down
```

This keeps volumes. `docker compose down -v` permanently deletes this stack's database, models
and documents. PostgreSQL initialization SQL runs only for a new database volume; changing
`.env` credentials does not change users/passwords already stored in that volume.
For an existing database missing pgvector, connect with an administrative user and execute
`CREATE EXTENSION IF NOT EXISTS vector;`. Flyway applies versioned registry migrations from
`src/main/resources/db/migration` at startup, including to an existing Phase 1 volume.
Do not edit an already-applied migration; add a new version for later schema changes.

## Phase 2 document API

All routes use `/properties/{propertyId}/contracts/{contractId}/documents`.
References are opaque external IDs (1–64 ASCII letters, digits, underscores or hyphens).
The AI service does not check whether they exist in Go/MySQL. Every registry query, including
download and deletion, is scoped by both references. This is data scoping, not user authorization.

| Method | Suffix | Result |
| --- | --- | --- |
| POST | none | Multipart `file` and `type`; HTTP 201, metadata and Location header |
| GET | none | HTTP 200, array of registered documents |
| GET | `/{documentId}` | HTTP 200, inline PDF; append `?download=true` for attachment |
| DELETE | `/{documentId}` | HTTP 204 after registry and file deletion |

Types: `RENT_CONTRACT`, `GUARANTEE`, `ANNEX`, `INVENTORY`, `HANDOVER`, `OTHER`.
Uploads require `application/pdf`, a `.pdf` filename, valid PDF structure, at least one page,
no encryption, and at most 20 MiB. Scanned PDFs can be stored; there is no OCR or text indexing.
The total multipart request limit is 21 MiB. Filename directories are stripped; filenames
with control characters or more than 255 characters are rejected.

```sh
curl -F "file=@contract.pdf;type=application/pdf" -F "type=RENT_CONTRACT" http://localhost:8081/properties/123/contracts/456/documents
curl http://localhost:8081/properties/123/contracts/456/documents
curl -o downloaded.pdf http://localhost:8081/properties/123/contracts/456/documents/DOCUMENT_UUID
curl -X DELETE http://localhost:8081/properties/123/contracts/456/documents/DOCUMENT_UUID
```

Use `curl.exe` in Windows PowerShell if `curl` is aliased. Replace `DOCUMENT_UUID` with the upload response ID.
Upload and list metadata use this shape:

```json
{
  "id": "ee0807ab-4571-4a19-b89f-3258c6d885c2",
  "type": "RENT_CONTRACT",
  "filename": "contract.pdf",
  "contentType": "application/pdf",
  "size": 245812,
  "createdAt": "2026-09-26T12:00:00Z"
}
```

List responses wrap these records in an array. Internal paths and stored filenames are not
returned. Downloads preserve the original filename and use `no-store` and `nosniff` headers.
Invalid PDFs, invalid IDs or unknown types return 400; unsupported file types return 415;
oversized uploads return 413; missing or out-of-scope documents return 404. Registry/storage
outages return 503. Errors use Spring's `application/problem+json` format without SQL or
physical paths in the response.

PDFs live at `properties/{propertyId}/contracts/{contractId}/{documentId}.pdf` beneath the
storage root. Duplicate original filenames never overwrite one another. PostgreSQL's
`documents` table records UUID, references, type, original/stored filenames, relative path,
MIME type, byte size and UTC creation/update timestamps; it never stores PDF bytes. Relative
paths allow moving the storage volume without rewriting registry paths.

### Transaction and filesystem cleanup

The service uses an explicit `TransactionTemplate`, so the orchestration remains visible
and service classes can be `final` without transaction proxies.

- Upload validates first, begins a database transaction, writes a temporary file and renames
  it to its UUID filename, then inserts metadata. A confirmed rollback removes the PDF.
  A write failure cleans up the partial temporary file before any registry insert.
- Delete locks the scoped registry row, renames the PDF to `<uuid>.pdf.deleting`, deletes the
  row and commits. After commit it unlinks the staged file; rollback restores the original.
  A missing physical file does not prevent deleting a stale registry row.
- If the commit outcome is unknown, retain the physical file and log the document UUID.
  Removing it could destroy a PDF whose registry insert actually committed. A post-commit
  unlink failure returns 503 and logs the UUID; the registry row is already deleted.
- A process crash or a second filesystem failure during compensation can require manual
  reconciliation. There is no distributed transaction, background cleanup job or automatic
  crash recovery in Phase 2. No vector rows exist yet; vector cleanup starts with Phase 3.

For reconciliation, stop the AI application while leaving PostgreSQL available. Use the
logged UUID and registry `storage_path` to inspect the exact file. A `.deleting` file with a
remaining registry row should be restored to its original name (provided no original exists);
one without a row can be removed. An upload PDF without a registry row can be removed after
confirming no upload is active. Remove stale `.upload-*.tmp` files only while the service is
stopped. Preserve anything ambiguous and investigate before deleting. A missing registered
file requires restoring a backup or explicitly deleting its stale registry entry via the API.

### Verification

`mvn verify` (also run during the Docker build) covers PDF validation, filesystem safety,
duplicate filenames, scoped access and transaction compensation/uncertain outcomes.
The symlink test runs on Linux/macOS; Docker runs it on Windows hosts too.

With the stack running, this PowerShell 7 script exercises real HTTP APIs and PostgreSQL:

```powershell
./scripts/smoke-documents.ps1
# Also restart the local Compose service between upload and download to verify persistence:
./scripts/smoke-documents.ps1 -RestartService
```

It generates its own valid PDF, uses unique external references and deletes its test documents.
It checks duplicate uploads, byte-identical download, isolated listing, foreign-scope reads
and deletes, invalid PDF/type, multipart size limits and successful deletion.

Phase 2 validation: all 17 tests passed in the Java 21 Docker build. The live smoke test
passed against PostgreSQL, including document persistence across a service restart.
Readiness returned `UP` after verification.

## Next increments

Phase 1 validation: the Docker build completed `mvn verify` with five passing tests.
The full Compose stack started successfully, readiness reported all four components `UP`,
SQL confirmed pgvector 0.8.2, and the non-root service user could write to the document volume.
Stopping Ollama produced readiness HTTP 503 while liveness remained HTTP 200; restarting
Ollama restored readiness. Model downloads and inference are not part of this bootstrap check.

1. **Phase 2 complete:** document registry migration, filesystem storage and scoped
   upload/list/download/delete APIs. No PDF content is extracted or indexed yet.
2. **Phase 3:** explicitly show `PagePdfDocumentReader -> TokenTextSplitter -> embedding model ->
   PgVectorStore.add()`. Embeddings represent text numerically for similarity search; Qwen is
   not needed to index documents. Keep a dedicated registry and metadata on every vector chunk.
3. **Phase 4:** test retrieval scoped by both property and contract, including cross-contract
   isolation, before adding generation.
4. **Phase 5:** retrieve context and ask Qwen to answer only from that context, including an
   explicit not-found response. This is the generation step of RAG.
5. **Phases 6–9:** Angular documents, stateless assistant, error cleanup and source citations.
6. **Phase 10:** evaluate Kubernetes separately after local acceptance.

Frontend inspection on `develop` found standalone Angular 21 Material components, signals,
Tailwind styling, translations, `core` API/read services and feature write services. Future UI
belongs under `feature/rent-contract/view` with document/assistant components and service-owned
HTTP calls. Extend `core/config/AppConfig` and `/assets/config.json` with an independent
`aiApiBaseUrl`, following the existing `apiBaseUrl` naming. No Angular files change in Phases 1–2.
