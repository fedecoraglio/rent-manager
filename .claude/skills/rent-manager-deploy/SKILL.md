---
name: rent-manager-deploy
description: Build, deploy, operate, and troubleshoot Rent Manager with Docker Compose, GitHub Actions, GHCR, kind, Kubernetes, Kustomize, and Argo CD; manage configuration, AI models, and recovery.
---

# Build, deployment, and operations

Read [infrastructure coverage](references/coverage.md), `../rent-manager-platform/SKILL.md`, and current sources before execution. These commands document procedures; creating skills does not imply deploying, restarting, importing data, or publishing code.

## Components and order

Kubernetes manages MySQL, the Go backend, Angular frontend, and the AI stack (PostgreSQL/pgvector, Ollama, and AI API). Argo CD manages repository manifests. AI can also run independently through Compose. Verify context/namespace before mutations; do not recreate an existing cluster as the first troubleshooting step.

## Current builds and CI

- Backend Dockerfile: Go 1.25.4 Alpine, go mod download, go build cmd/api, binary in Alpine 3.22. Embedded migrations.
- Frontend: Node 22 Alpine, npm install and npm run build; browser output to Nginx 1.27 with SPA fallback. Dockerfile does not use npm ci.
- AI: Maven 3.9.11/Java 21 runs mvn verify; Java 21 runtime with app user and writable storage.
- `.github/workflows/backend-image.yml`, frontend-image.yml, and ai-image.yml run on matching pushes to master. They publish GHCR :master images with GITHUB_TOKEN/packages:write. AI also supports manual execution and publishes a commit SHA tag; its Docker build runs mvn verify. Workflows do not update manifests or run kubectl. There is no explicit Go/UI test job.

Publishing a new image under the same tag does not change PodTemplate and does not trigger a Kubernetes rollout by itself. Always pulls when a pod is created, not in an existing pod. Wait for CI to finish before restarting the changed component. For reproducible rollback, prefer an immutable digest/tag and a GitOps change; master does not retain historical image identity in the manifest.

## Prepare local Kubernetes

Requirements: Docker Desktop, kind, kubectl, an available storage class, and an installed nginx ingress controller. The repository does not declare Argo CD installation or kind port configuration. Run `kind create cluster --name rent-manager` only if the intended cluster does not exist.

```powershell
kubectl config use-context kind-rent-manager
kubectl config current-context
kubectl get nodes
kubectl apply -f platform/kubernetes/namespace.yaml
```

Prepare `mysql-secret` and `backend-secret` with keys compatible with StatefulSet/config.go and DB; local secret.yaml files are ignored and excluded from Kustomize. Do not copy their values into documentation. Read secret.example.yaml for the schema and check required APP/JWT/DB settings. MySQL user/database initialization runs on an empty volume; changing a Secret does not automatically rotate an existing database's credentials.

```powershell
kubectl apply -f platform/kubernetes/mysql/secret.yaml
kubectl apply -f platform/kubernetes/backend/secret.yaml
kubectl apply -k platform/kubernetes
kubectl -n rent-manager get pvc,pods,svc,ingress
kubectl -n rent-manager rollout status statefulset/mysql --timeout=120s
kubectl -n rent-manager rollout status deployment/backend --timeout=120s
kubectl -n rent-manager rollout status deployment/frontend --timeout=120s
```

The backend migrates at startup: diagnose connectivity, permissions, and migration state if it restarts. Do not delete the DB/PVC to fix startup failure.

## Argo CD

With Argo CD installed, apply `platform/argocd/rent-manager.yaml`. Observe Application rent-manager in namespace argocd: repository master, path platform/kubernetes, autosync/prune/selfHeal. Manual changes to managed resources may be reverted. Removing a resource from Git can remove it from the cluster through prune; handle PVC/data carefully.

```powershell
kubectl -n argocd get application rent-manager
kubectl -n argocd annotate application rent-manager argocd.argoproj.io/refresh=normal --overwrite
kubectl -n rent-manager rollout restart deployment/backend
kubectl -n rent-manager rollout status deployment/backend --timeout=120s
```

Restart frontend rather than backend when its image/config changes. config.json is mounted with subPath: changing the ConfigMap does not automatically update the mounted file. Recreate affected pods after configuration or environment Secret changes. Do not use manual apply and GitOps as contradictory sources of truth.

## Access and functional verification

```powershell
kubectl -n ingress-nginx port-forward service/ingress-nginx-controller 8085:80
```

Open `http://localhost:8085`: `/` serves Angular, `/v1` reaches Go unchanged, and the separate AI Ingress strips `/ai` before forwarding to ai:8081. Test login and an authenticated Go query. Backend `/` is not a health endpoint. The AI stack has probes; existing frontend/backend/MySQL manifests still lack them. Rollout alone does not validate business functionality.

Alternatives in separate terminals: frontend `kubectl -n rent-manager port-forward svc/frontend 4400:80`; backend `... svc/backend 8082:8080`; MySQL `... svc/mysql 3307:3306`; Argo `kubectl -n argocd port-forward svc/argocd-server 8080:443`. Direct frontend access requires a reachable API base/CORS; it does not provide a /v1 proxy.

## AI with Kubernetes

Read [AI Kubernetes runbook](../../../platform/kubernetes/ai/README.md). Publish the AI image and prepare ignored postgres/secret.yaml from its example before syncing. Kustomize includes the non-secret ai/postgres/ollama manifests; Argo CD already watches that path. PostgreSQL and AI share postgres-secret credentials. Separate PVCs persist database (5Gi), models (10Gi), and PDFs (5Gi); Compose data is not imported.

Ollama's init container runs a temporary loopback server and pulls missing nomic-embed-text/qwen2.5:3b-instruct models into persistent storage. AI init containers wait for authenticated database access and both models; Flyway creates the vector extension and tables. AI uses Recreate and UID/GID/fsGroup 10001 for single-writer persistent storage. This CPU stack needs at least 8Gi Docker Desktop memory, preferably more alongside other workloads.

Use `kubectl -n rent-manager logs ollama-0 -c pull-models -f`, `rollout status statefulset/postgres --timeout=300s`, `rollout status statefulset/ollama --timeout=1800s`, and `rollout status deployment/ai --timeout=1800s`. Direct access: `kubectl -n rent-manager port-forward service/ai 8081:8081`, without /ai. Through the existing ingress forward, use http://localhost:8085/ai/actuator/health/readiness and the document smoke script with `-BaseUrl http://localhost:8085/ai`, without Compose-only switches.

The AI Ingress allows 21m bodies and 330-second read/send timeouts. AI has no authentication; keep this ingress local to development. Angular integration is separate. Wait for CI before restarting deployment/ai for a changed master image; prefer commit SHA tags for reproducible releases. Updating the Secret does not rotate initialized PostgreSQL passwords.

## AI with Compose

From rent-manager-ai, prepare .env with POSTGRES_USER/PASSWORD and required variables; use .env.example without copying credentials. `docker compose up -d --build` starts postgres pgvector 17, Ollama, and Java on 8081. By default, Postgres exposes loopback port 5433, Ollama 11434, and AI 8081; inside the network use postgres:5432 and ollama:11434. Persist postgres-data, ollama-data, and documents volumes.

```powershell
docker compose exec ollama ollama pull nomic-embed-text
docker compose exec ollama ollama pull qwen2.5:3b-instruct
docker compose logs rent-manager-ai
curl.exe http://localhost:8081/actuator/health/readiness
curl.exe http://localhost:8081/actuator/health/liveness
```

Healthy Ollama does not imply models are downloaded. The chat bean uses qwen2.5:3b-instruct even though YAML mentions qwen3:4b. Effective timeout variables: OLLAMA_EMBEDDING_READ_TIMEOUT_SECONDS and OLLAMA_CHAT_READ_TIMEOUT_SECONDS; deliberately resolve the mismatch with compose's OLLAMA_READ_TIMEOUT_SECONDS when changing timeouts.

Backend Compose in rent-manager-backend/docker-compose.yml starts only local MySQL and uses a different database name from Kubernetes: align configuration with the selected environment. The old platform/docker/docker-compose.yml reference in the root document points to a missing file.

## Troubleshooting

| Symptom | Inspect |
| --- | --- |
| Kubernetes API connection refused | Running Docker/kind, context, kubeconfig endpoint, and port; do not assume an application problem |
| MySQL Pending | PVC/StorageClass, events, node resources |
| Backend CrashLoop | Startup logs, DB_HOST=mysql, Secret, permissions, migrations, database readiness |
| ImagePullBackOff | Image name/tag, GHCR access, pull credentials for private images |
| Changes not visible | Completed CI, published image, recreated pod, actual imageID, ConfigMap subPath |
| UI opens but API fails | config.json, /v1 in Ingress, Service/endpoints, token, CORS |
| AI 503 or slow uploads | Available models, Ollama/Java logs, DB, pgvector, filesystem, effective timeout |

Use `kubectl -n rent-manager describe pod <pod>`, `logs deployment/backend --previous` for the previous crash, `get events --sort-by=.metadata.creationTimestamp`, and `get endpoints backend frontend mysql`. The live cluster may differ from Git: inspect before concluding.

## Data and recovery

The root document swaps backup/import descriptions: mysqldump exports; mysql with SQL input imports. `backup.sql` is inside rent-manager-backend, not the repository root. Do not copy commands containing literal passwords or run DROP DATABASE for diagnosis. Export with managed credentials, verify contents/restoration in an isolated target, and only then consider restoring existing data. With PowerShell, account for binary redirection/encoding and use a compatible tool.

Do not run `docker compose down -v` or delete PVCs as rollback. Back up AI PostgreSQL and PDFs together; retain models or the ability to download them. Reconcile .deleting/orphaned files against registry/vectors after unknown transaction outcomes. An image rollback does not automatically revert migrations or data. No live cluster was modified or validated when these skills were generated.
