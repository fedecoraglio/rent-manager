# AI on Kubernetes

The root Kustomize application now includes the AI API, PostgreSQL 17 with
pgvector 0.8.2, and Ollama 0.12.3, in namespace `rent-manager`.
MySQL, the Go backend, and the Angular frontend remain separate workloads.
Argo CD picks up these manifests through its existing `platform/kubernetes` path.

## Components and storage

| Workload | Service | Persistent storage | Memory request / limit |
| --- | --- | --- | --- |
| StatefulSet postgres | postgres:5432 | postgres-data, 5Gi | 256Mi / 1Gi |
| StatefulSet ollama | ollama:11434 | ollama-data, 10Gi | 2Gi / 6Gi |
| Deployment ai | ai:8081 | ai-documents, 5Gi | 512Mi / 1Gi |

PVCs use the cluster's default StorageClass and ReadWriteOnce. Keep one replica
per workload. The API uses Recreate to avoid overlapping filesystem writers;
updates briefly interrupt AI requests. The image runs as UID/GID 10001, and
fsGroup grants access to the document volume. No GPU is required.

Allow at least 8Gi of Docker Desktop memory for this local CPU stack, preferably
more alongside MySQL, the frontend/backend, Argo CD, and other workloads.
Actual capacity and inference latency depend on the host. Requests determine
scheduling; memory limits alone do not reserve that capacity. Ensure enough disk
space for the three claims and container images.

## First deployment

### Validation status (2026-11-13)

Static review and `kubectl kustomize platform/kubernetes` succeeded. Runtime
validation remains blocked in the current session: Docker denies access to
`dockerDesktopLinuxEngine`, and the Kubernetes API at `127.0.0.1:64305` rejects
socket access. The selected context is `kind-rent-manager`. This session cannot
request elevated execution. No cluster resources or existing Secrets were changed.

The local `postgres/secret.yaml` file is absent and its path is Git-ignored.
The cluster Secret's existence is **unknown**, not confirmed missing. Argo CD
ownership, image availability, PVC binding, model downloads, readiness, and
document/chat behavior also remain unverified.

`.ai-deploy-check` contains previous rendered manifests, JSON snapshots, an
Ollama bootstrap shell snippet, and a temporary Go YAML parser/cache. These are
validation artifacts, not deployment inputs; the parser's go.mod contains a
machine-specific module replacement. Do not apply this directory or treat its
snapshots as current cluster state. Its rendered.yaml was refreshed in this review.

Before any apply, inspect live Applications and resource tracking metadata:

```powershell
kubectl --context kind-rent-manager get applications.argoproj.io -A
kubectl --context kind-rent-manager -n argocd get application rent-manager -o yaml
kubectl --context kind-rent-manager -n rent-manager get deployment,statefulset,pvc -o yaml
kubectl --context kind-rent-manager -n rent-manager get secret postgres-secret -o name
```

Only an explicit Secret `NotFound` establishes that it is missing. Authentication,
authorization, and transport failures must stop Secret preparation. Use `create`
for a new Secret rather than applying over an existing one. Generate a random
password locally without printing it; never commit the real secret file. If
PostgreSQL data already exists, recover its matching credentials instead of
generating a replacement.

### Local image without publication

After Docker/cluster access is restored, a local build avoids publishing:

```powershell
docker build -t rent-manager-ai:local-validation ./rent-manager-ai
if ($LASTEXITCODE -ne 0) { throw 'AI image build failed' }
kind load docker-image rent-manager-ai:local-validation --name rent-manager
if ($LASTEXITCODE -ne 0) { throw 'Loading the image into kind failed' }
```

Use that image with `imagePullPolicy: IfNotPresent` in a temporary local
Kustomize patch; keep the tracked GHCR image/policy unchanged. Inspect the
rendered result before applying. If Argo CD manages these resources, first record
and temporarily disable automated reconciliation, wait for any active sync to
finish, and verify no parent Application restores it. Preserve the original sync
policy for restoration. Restoring GitOps before the tested revision/image is
available in Git/GHCR can revert the local image or prune locally added resources.
Do not resume prune until the desired manifests retain all three PVCs.

1. Ensure Docker Desktop, `kind-rent-manager`, a default StorageClass, and the
   nginx ingress controller are available.
2. Build/publish the AI image **before** syncing the new manifests. The new
   `.github/workflows/ai-image.yml` runs on matching pushes to master and supports
   manual execution. It runs the Dockerfile's `mvn verify` and publishes both
   `:master` and a commit SHA tag to GHCR. Make the package readable by the cluster,
   or configure imagePullSecrets if it is private.
3. Prepare the PostgreSQL Secret below before Argo CD sync. It is deliberately
   excluded from Kustomize, as with existing database/backend secrets.

Run from the repository root:

```powershell
kubectl config use-context kind-rent-manager
kubectl get nodes
kubectl get storageclass
kubectl apply -f platform/kubernetes/namespace.yaml
Copy-Item platform/kubernetes/postgres/secret.example.yaml platform/kubernetes/postgres/secret.yaml
# Edit the ignored secret.yaml and replace CHANGE_ME with a private password.
# Do not overwrite an existing secret.yaml or rotate a running database this way.
kubectl apply -f platform/kubernetes/postgres/secret.yaml
kubectl kustomize platform/kubernetes
```

For a manually managed installation, apply the rendered application with
`kubectl apply -k platform/kubernetes`. If Argo CD manages the application, sync
the Git revision instead; avoid competing manual and GitOps changes.
Existing `mysql-secret` and `backend-secret` must still be present.
If automated Argo CD sync observes the first push before CI has published the
image, AI can temporarily show ImagePullBackOff. Verify CI/package availability;
image pulls retry once publication is complete. Missing database secrets must
still be prepared explicitly.

```powershell
kubectl -n argocd annotate application rent-manager argocd.argoproj.io/refresh=normal --overwrite
kubectl -n argocd get application rent-manager
kubectl -n rent-manager get pvc,pods,svc,ingress
kubectl -n rent-manager rollout status statefulset/postgres --timeout=300s
kubectl -n rent-manager logs ollama-0 -c pull-models -f
kubectl -n rent-manager rollout status statefulset/ollama --timeout=1800s
kubectl -n rent-manager rollout status deployment/ai --timeout=1800s
```

The first Ollama init container starts a temporary local server and downloads
`nomic-embed-text` and `qwen2.5:3b-instruct` into its PVC. Downloads need outbound
internet and can take several minutes. Cached models are checked and reused on
subsequent starts. Failures retry through init-container restart behavior; inspect
`pull-models` logs for network or disk problems. No inference runs in probes.

The AI init containers wait for authenticated PostgreSQL access and both models.
Flyway then applies V1/V2, including `CREATE EXTENSION IF NOT EXISTS vector` and
the registry/vector tables. The PostgreSQL image already contains the extension
binaries; pgvector is an extension, not a separate server. The initialized
POSTGRES_USER has administrative privileges, matching the local Compose setup;
separating migration and runtime users is outside this local deployment change.

The AI startup/liveness probes use the process-only endpoint; readiness includes
database, pgvector, and Ollama. A dependency outage removes the API from ready
Service endpoints without making liveness depend on that dependency.

## Access and smoke checks

Through the existing ingress-controller port-forward:

```powershell
kubectl -n ingress-nginx port-forward service/ingress-nginx-controller 8085:80
```

The separate AI Ingress strips `/ai` before forwarding, while `/v1` continues to
reach Go and `/` reaches Angular. It allows 21MiB request bodies for multipart
PDF uploads and 330-second proxy read/send timeouts for indexing/chat.

```powershell
curl.exe http://localhost:8085/ai/actuator/health/readiness
./rent-manager-ai/scripts/smoke-documents.ps1 -BaseUrl http://localhost:8085/ai
```

The smoke script creates temporary scoped PDFs and deletes them. Its
RestartService/CheckIndexing/CheckEmbeddingFailure switches target Compose,
so do not use those switches for Kubernetes.
The default script does not test chat or directly inspect vector rows. Test chat
separately while a test document still exists, confirm a grounded response with
citations, then verify another contract cannot retrieve its content. Do not use
a successful document smoke run as evidence that chat was validated.

Alternatively, use direct local access (no `/ai` prefix):

```powershell
kubectl -n rent-manager port-forward service/ai 8081:8081
curl.exe http://localhost:8081/actuator/health/readiness
```

Document/chat paths are otherwise unchanged, for example
`/ai/properties/P1/contracts/C1/documents` through Ingress. The API's Location
header remains an application-relative path without `/ai`; clients using the
prefix must resolve it against the configured AI base URL. This change deploys
the API; it does not add Angular integration or Go authentication to AI. The AI
API currently has no authentication; keep this ingress local to the development
environment unless access control is added.

## Updates and configuration

`ai-config` defines server, storage, dependency URLs, JVM sizing, and effective
timeout variable names. `postgres-config` supplies the shared database name,
and `postgres-secret` supplies the same credentials to PostgreSQL and AI.
`ollama-config` defines model names and limits inference to one loaded model
and one parallel request to reduce local memory pressure.

Keep model names aligned with `AiConfiguration.java`, whose explicit chat bean
currently uses `qwen2.5:3b-instruct` rather than the YAML's `qwen3:4b`.
Changing an embedding model/dimension requires a reindexing plan.

After CI finishes publishing a replacement `:master` image:

```powershell
kubectl -n rent-manager rollout restart deployment/ai
kubectl -n rent-manager rollout status deployment/ai --timeout=1800s
```

An unchanged mutable tag does not trigger rollout automatically. For reproducible
releases/rollback, change the manifest to the published commit SHA tag and let
Argo CD reconcile it. Restart affected workloads after ConfigMap/Secret changes;
environment variables are read when pods are created. Changing initialized
PostgreSQL credentials in a Secret does not change the stored database password.

## Troubleshooting and data

```powershell
kubectl -n rent-manager describe pod ai-<pod-suffix>
kubectl -n rent-manager logs deployment/ai -c wait-for-postgres
kubectl -n rent-manager logs deployment/ai -c wait-for-models
kubectl -n rent-manager logs deployment/ai -c ai
kubectl -n rent-manager logs ollama-0 -c pull-models
kubectl -n rent-manager exec ollama-0 -c ollama -- ollama list
kubectl -n rent-manager logs postgres-0
kubectl -n rent-manager get events --sort-by=.metadata.creationTimestamp
```

Missing Secret: CreateContainerConfigError. Missing registry image/access:
ImagePullBackOff. Unbound PVC or insufficient CPU/memory: Pending. Model pull
failure: Init:CrashLoopBackOff. Inspect startup logs and Flyway errors if the API
restarts after its init containers complete. The init loops intentionally keep
waiting for dependencies; deployment progress timeout does not repair them.

This creates new Kubernetes volumes; it does not copy existing Compose volumes,
host PDFs, PostgreSQL rows, or model caches. Back up and restore PostgreSQL plus
document files consistently if existing data must be migrated; do not copy only
vectors or only PDFs. Keep both stacks stopped from writing during the final
migration. Downloading cached models again is safe but may be slow.

Do not delete PVCs, reset databases, or run `compose down -v` as a rollout fix.
Argo CD prune can delete resources removed from Git; preserve PVC manifests and
backups. Image rollback does not roll back schema or data changes.
