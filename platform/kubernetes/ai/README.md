# AI through the backend

The browser accesses documents and chat through the authenticated Go backend.
The AI Service remains ClusterIP-only at `http://ai:8081`; there is no AI Ingress.
The existing nginx Ingress routes `/v1` to backend and `/` to frontend.

## HTTP contract

All endpoints below require `Authorization: Bearer <token>`:

| Method | Path | Result |
| --- | --- | --- |
| POST | /v1/rental-contracts/:rentalContractId/documents | Multipart `file` + `type`; 201 with `{data: document}` and backend Location |
| GET | /v1/rental-contracts/:rentalContractId/documents | 200 with `{data: [...]}` |
| GET | /v1/rental-contracts/:rentalContractId/documents/:documentId | PDF inline; `?download=true` for attachment |
| DELETE | /v1/rental-contracts/:rentalContractId/documents/:documentId | 204 |
| POST | /v1/rental-contracts/:rentalContractId/chat | JSON `{question: "..."}`; 200 with `{data: {answer: "..."}}` |

The use case loads the rental contract from MySQL and derives propertyId from it.
The caller cannot supply a different propertyId or an arbitrary upstream URL.
Invalid/missing contracts never reach AI. AI data must use the canonical numeric
property/contract IDs from MySQL; documents previously uploaded under synthetic
scopes such as P1/C1 are not automatically migrated to business contracts.

Access follows the existing application-wide administrative JWT model. Users,
owners and properties have no ownership/access relationship in the current schema;
this change does not claim per-user or per-owner authorization. All authenticated
administrators have the same contract access as existing rental-contract endpoints.

## Architecture

`adapter/http/DocumentHandler` handles binding, multipart limits and response DTOs.
`core/usecase/document` defines separate UploadDocument, ListDocuments, DownloadDocument,
DeleteDocument and AskContractQuestion use cases. Each checks the contract,
validates inputs and orchestrates the
`core/port/DocumentGateway`. `adapter/ai/Client` implements that port using a
bounded HTTP client. `cmd/api/main.go` assembles those dependencies. Core code does
not depend on Gin or net/http. No Go code writes the AI PostgreSQL registry directly.

`core/policy/document/DocumentUploadPolicy.Execute` owns upload size, required
content and allowed document types. `ContractQuestionPolicy.Execute` normalizes
and validates questions. Shared policy helpers validate document/contract IDs.
Policies are injected through constructors from the composition root, matching
other backend write policies. Contract lookup and property scope resolution remain
use-case orchestration. Multipart parsing, MIME headers and HTTP body limits remain
in the HTTP adapter; PDF parsing and text extraction remain in the AI service.

The backend reconstructs multipart while preserving PDF bytes. Maximum file size
is 20 MiB, total multipart body 21 MiB. File headers are encoded safely; the AI
service still validates PDF contents, encryption, filename and readable text.
Requests and downloads are bounded. Request cancellation propagates to AI; there
are no automatic retries of uploads, deletes or chat. Do not assume cancellation
proves an upstream write was rolled back; inspect the document list before retrying.

AI status codes are mapped to backend `{code,message}` errors without forwarding
upstream stack traces or URLs. Supported mappings include 400/404/409/413/415/422,
503 for dependency/protocol failures, and 504 for timeouts. User Authorization and
Cookie headers are never forwarded. Redirects are rejected.

## Configuration

- `AI_BASE_URL`: local default `http://localhost:8081`; Kubernetes `http://ai:8081`.
- `AI_TIMEOUT_SECONDS`: 450 by default, allowed range 1..3600.
- nginx body limit: 21m; read/send timeouts: 480 seconds.

The backend starts even when AI is unavailable; existing Go features continue to
work. Invalid URL/timeout configuration fails startup. Recreate backend pods after
ConfigMap changes because their environment is read only at pod creation.

## Kubernetes rollout

Build/publish the updated backend image before syncing the platform changes.
Never commit Secrets or database backups. Argo CD sync removes the previously
managed `rent-manager-ai` Ingress through prune. With manual apply, removing a
manifest from Kustomize does not remove the old live resource: inspect it and
explicitly remove only that Ingress after the backend routes are available.
Do not delete PVCs or change stored database credentials.

The `ai-from-backend` NetworkPolicy selects `app: ai` and admits TCP/8081 only
from `app: backend` pods in the same namespace. Node-originated health probes are
handled by Kubernetes network-policy semantics. Policy enforcement requires a
supporting CNI and must be tested: merely creating a policy does not prove it works.
The inspected kind cluster uses `kindest/kindnetd`; inter-pod enforcement has not
been verified in this change. No CNI or cluster recreation is performed here.
NetworkPolicies are additive: inspect any other policies selecting AI before
claiming backend-only isolation. Users with pod creation/exec/port-forward RBAC
or node access are outside this application-level boundary.

After sync:

```powershell
kubectl -n rent-manager rollout restart deployment/backend
kubectl -n rent-manager rollout status deployment/backend --timeout=300s
kubectl -n rent-manager get ingress,networkpolicy,pods,pvc
```

Verify an unauthenticated request to a document/chat endpoint receives 401, a
valid token and existing contract can upload/list/download/delete and chat, and
unknown contracts return 404. Confirm `/ai` no longer routes to AI (the frontend
SPA may return HTML there, so HTTP 200 alone does not mean AI is exposed).
Check that AI is reachable from the backend and blocked from unrelated pods.
Keep readiness/health operational checks internal; no public AI health route is added.

For API development, forwarding backend locally is sufficient. The browser never
needs a direct connection or port-forward to AI:

```powershell
kubectl -n rent-manager port-forward service/backend 8082:8080
# Set $env:RENT_MANAGER_TOKEN privately before these calls.
curl.exe -H "Authorization: Bearer $env:RENT_MANAGER_TOKEN" http://localhost:8082/v1/rental-contracts/123/documents
curl.exe -H "Authorization: Bearer $env:RENT_MANAGER_TOKEN" -F "type=RENT_CONTRACT" -F "file=@contract.pdf;type=application/pdf" http://localhost:8082/v1/rental-contracts/123/documents
```

The existing AI smoke-documents script calls the internal AI API directly with
synthetic scopes and expects unwrapped responses. It is an administrator-only AI
test; it cannot be pointed at the backend without adapting authentication, real
contract IDs, routes and response envelopes.

## AI persistence and bootstrap

PostgreSQL 17/pgvector stores registry and vectors on postgres-data (5Gi). Ollama
bootstraps nomic-embed-text and qwen2.5:3b-instruct on ollama-data (10Gi). Original
PDFs live on ai-documents (5Gi), mounted at /app/storage. AI waits for PostgreSQL
credentials and both models. Flyway applies pending migrations at startup.
Keep postgres-secret local/ignored and create it only if explicitly missing; do
not overwrite a live Secret or rotate initialized database credentials this way.
Allow at least 8Gi Docker Desktop memory, preferably more for the full stack.
Backup PostgreSQL and original PDFs together. Compose volumes are not imported.
