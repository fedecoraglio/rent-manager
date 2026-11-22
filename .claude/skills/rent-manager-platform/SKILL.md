---
name: rent-manager-platform
description: Understand Rent Manager platform resources, Kustomize composition, Kubernetes topology, configuration, persistence, and the Argo CD application.
---

# Platform architecture

Read [resources](references/coverage.md). `platform/argocd/rent-manager.yaml` declares an Application in the argocd namespace watching repository `master` and rendering `platform/kubernetes`. Destination: internal cluster and rent-manager namespace; automated sync with prune and selfHeal.

| Component | Resources and behavior |
| --- | --- |
| Entry point | Ingress with nginx class: `/v1` to backend:8080 and `/` to frontend:80 |
| Frontend | Single-replica Deployment, ClusterIP Service, ConfigMap containing config.json mounted through subPath |
| Backend | Single-replica Deployment, ClusterIP Service, envFrom backend-config and backend-secret |
| MySQL | Single-replica StatefulSet, mysql:8.0, headless Service, 2Gi RWO mysql-data PVC |
| AI API | Single-replica Recreate Deployment, ai:8081 ClusterIP Service, ai-config, 5Gi ai-documents PVC, startup/liveness/readiness probes |
| PostgreSQL/pgvector | Single-replica StatefulSet, pgvector/pgvector:0.8.2-pg17, postgres:5432 headless Service, postgres-config/postgres-secret, 5Gi postgres-data PVC |
| Ollama | Single-replica StatefulSet, ollama/ollama:0.12.3, ollama:11434 headless Service, 10Gi model PVC, model-pull init container |
| AI entry point | Separate nginx Ingress strips /ai, allows 21m multipart bodies and 330-second proxy timeouts |

Frontend/backend/AI deployments use GHCR `:master` images with Always. The new AI stack has probes and resource requests/limits; existing frontend/backend/MySQL resources remain unchanged. TLS and ingress controller installation are outside this composition. The AI stack runs on CPU and uses separate persistent volumes. Keep one replica; filesystem storage is not shared object storage.

`kustomization.yaml` includes namespace, both ingresses, MySQL resources, backend/frontend resources, and the non-secret AI/PostgreSQL/Ollama stack. Secrets are prepared separately and excluded from resources. Keep `backend-secret`, `mysql-secret`, `postgres-secret`, key names, and credentials consistent without reproducing values in skills. Prepare the ignored postgres/secret.yaml from postgres/secret.example.yaml before sync.

Read [AI Kubernetes runbook](../../../platform/kubernetes/ai/README.md). Ollama initializes cached models before serving; AI init containers wait for authenticated PostgreSQL access and both models. Flyway creates the vector extension/schema. Document storage uses fsGroup 10001, matching the Docker image UID/GID. Existing Compose data is not migrated automatically.

Use `../rent-manager-deploy/SKILL.md` for setup, CI, updates, troubleshooting, and rollback procedures. Do not confuse declared configuration with the state of a running cluster.
