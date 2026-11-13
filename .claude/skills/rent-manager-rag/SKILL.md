---
name: rent-manager-rag
description: Work on PDF extraction and chunking, embeddings, scoped vector retrieval, and RAG chat responses about Rent Manager contracts.
---

# Indexing and RAG chat

Read [coverage](references/coverage.md) and `../rent-manager-ai/SKILL.md` for effective model configuration.

## Indexing

Document upload calls PdfChunker: read one page per document, skip empty pages, normalize whitespace, split uppercase clause headings, then use TokenTextSplitter (500 tokens, minChunkSizeChars 100, minChunkLengthToEmbed 5, keepSeparator true). Preserve documentId/propertyId/contractId/documentType/filename/page on each chunk. No usable text: 422, without OCR.

SpringAiDocumentIndexer requires an active transaction and adds batches of 16 through VectorStore. Each add invokes nomic-embed-text and stores 768-dimensional vectors; all batches share the registry transaction. Add failures are wrapped as 503 and roll back the set. Deletion requires the same transaction and a three-ID filter.

## Business query

POST `/properties/{propertyId}/contracts/{contractId}/chat` with `{ "question": "..." }` asks about documents in the scope. RagController passes strings directly to RagService without constructing DocumentScope or applying Bean Validation to the question; do not claim it applies the document controller's validations.

PgVectorDocumentRetriever expands the question with contractual terms in English/Spanish and performs topK=10 similarity search filtered by propertyId AND contractId. It has no document type/documentId filter or explicit threshold configured in SearchRequest. Convert results to RetrievedDocumentChunk.

With no chunks, return fixed missing-information text without calling the LLM. With chunks, build context containing Source, Page, and Content; ChatClient receives a system prompt to answer only from context, avoid invention, and remain concise. HTTP response is `{answer: ...}`, without structured citations or streaming. There is no conversation history or memory between calls.

## Limitations and verification

The model instruction does not guarantee accuracy; preserve scope, metadata, and the no-results path. RagService prints chunk text to stdout: consider sensitivity when changing observability. Do not confuse vector filtering with authorization. Test different contracts with similar text, out-of-context questions, missing documents, embedding errors, missing models, and timeouts. Changing embeddings/dimensions requires a reindexing plan; do not mix incompatible model vectors.
