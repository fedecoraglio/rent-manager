---
name: rent-manager-documents
description: Work on PDF upload, validation, listing, download, viewing, and transactional deletion by property and contract in rent-manager-ai.
---

# PDF documents

Read [coverage](references/coverage.md). HTTP base independent of Go: `/properties/{propertyId}/contracts/{contractId}/documents`. Do not prefix it with `/v1` or assume an Angular screen. Types: RENT_CONTRACT, GUARANTEE, ANNEX, INVENTORY, HANDOVER, OTHER.

## Business and technical upload flow

1. Accept multipart POST with `type` and `file`; validate scope references containing 1–64 alphanumeric characters, `_`, or `-`.
2. PdfUploadValidator requires a nonempty file <=20 MiB, basename <=255 without control characters, .pdf extension, application/pdf MIME, %PDF- header, PDFBox readability, at least one page, and no encryption.
3. Create UUID and open TransactionTemplate with configurable timeout. Save the original in the filesystem using scope/UUID, not an arbitrary user filename.
4. Insert DocumentMetadata into PostgreSQL and read the saved file for indexing.
5. Extract text/chunks and insert embeddings within the same JDBC transaction. Success: 201, Location, and DocumentResponse.
6. Rollback: revert registry/vectors and compensate by removing the file; if transaction outcome is UNKNOWN, retain it and reconcile. Cleanup failures are logged; this is not guaranteed distributed rollback.

## Listing and download

Base GET lists metadata within the scope. GET `/{documentId}` requires a UUID and scoped lookup, returning application/pdf inline by default; `?download=true` returns an attachment. Use the original name in UTF-8 Content-Disposition, cache-control no-store, and nosniff. A document in another scope is treated as missing. Go is not queried to verify contract existence, and AI does not authenticate the caller.

## Deletion

DELETE by UUID: transaction → locked SELECT → delete chunks filtered by documentId/propertyId/contractId → stage physical deletion through stageDeletion → delete registry record. After commit, complete deletion; on rollback, restore the file. Unknown outcome: retain `.deleting`. If cleanup fails after commit, return a service error and reconcile; do not attempt to restore a committed registry record. Normal success: 204.

## Errors and recovery

Follow DocumentExceptionHandler for 400/413/415/422/404/503 and framework exceptions; do not assume Go's JSON format. PDFs without text require OCR, which is unsupported. Indexing failures roll back the entire upload; no asynchronous pending state is retained.

Smoke-check upload/list/view/attachment/delete, another scope, fake/encrypted/empty/oversized/textless PDFs, unavailable Ollama, and filesystem failures. Reconcile DB, vectors, and files before removing orphans; do not delete storage wholesale.
