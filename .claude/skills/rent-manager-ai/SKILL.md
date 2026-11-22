---
name: rent-manager-ai
description: Understand and modify the Rent Manager Java document and RAG service architecture, Spring AI, ports, PDF storage, PostgreSQL, pgvector, and Ollama.
---

# AI service

Read [coverage](references/coverage.md), `rent-manager-ai/pom.xml`, `src/main/resources/application.yml`, and `config/AiConfiguration.java`. The POM declares Java 21, Spring Boot 4.0.8, and Spring AI 2.0.1.

Java root: `src/main/java/com/rentmanager/ai`. `adapter/in/web` implements controllers/DTOs/errors; `application/document` and `application/rag` orchestrate; `domain/document` defines records, scopes, and types; `port/out` defines repository, storage, indexer, and retriever; `adapter/out` implements filesystem, JDBC, embeddings/retrieval, and health. `config` assembles clients and the vector store. Separation is not strict: DocumentService depends on MultipartFile, PdfUploadValidator, and Spring transactions.

## Data and boundaries

Go's MySQL is not shared. PostgreSQL stores the document registry and vectors; the filesystem stores original PDFs. Scope is `(propertyId,contractId)` with string references, without querying the backend to verify property/contract existence. The AI API implements neither JWT nor user authorization. Do not confuse scope filtering with access control.

Flyway creates `documents` and `vector_store`. The vector store shares the registry transaction's JdbcTemplate/DataSource: 768 dimensions, cosine distance, index NONE, initializeSchema false. Flyway must finish before it is constructed.

## Effective configuration

The manually configured chat bean fixes `qwen2.5:3b-instruct`, temperature 0.1; application.yml declares `qwen3:4b`, but does not determine that bean's model. Embeddings fixes `nomic-embed-text` and truncate false. Read the bean before documenting model changes.

Configure `POSTGRES_HOST/PORT/DB/USER/PASSWORD`, `OLLAMA_BASE_URL`, `DOCUMENT_STORAGE_PATH`, `INDEXING_TIMEOUT_SECONDS`, `OLLAMA_EMBEDDING_READ_TIMEOUT_SECONDS`, and `OLLAMA_CHAT_READ_TIMEOUT_SECONDS`. The generic `OLLAMA_READ_TIMEOUT_SECONDS` variable in compose does not match the properties used by these beans.

## Run and verify

Read `../rent-manager-deploy/SKILL.md` for Compose and Kubernetes. Root Kustomize deploys AI, PostgreSQL/pgvector, and Ollama. Read [AI Kubernetes runbook](../../../platform/kubernetes/ai/README.md) and the ai/postgres/ollama manifests for Secret setup, model bootstrap, PVCs, probes, resource sizing, and /ai Ingress routing. Docker UID/GID 10001 matches fsGroup for persistent document writes. The ai-image workflow publishes master and commit SHA tags to GHCR.

From `rent-manager-ai`: `mvn verify` or `mvn spring-boot:run` with the environment prepared. `scripts/smoke-documents.ps1` is the document smoke test; read parameters/effects first. For Kubernetes use `-BaseUrl http://localhost:8085/ai` without Compose-only switches. Actuator readiness includes DB, pgvector, and Ollama; liveness only includes process state. Readiness does not validate every chat response.
