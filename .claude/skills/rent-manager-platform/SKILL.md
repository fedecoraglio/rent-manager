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

Frontend/backend deployments use GHCR `:master` images with Always. There are no Kubernetes resources for AI, PostgreSQL, or Ollama. This composition also lacks probes, resource limits, TLS, and ingress controller installation.

`kustomization.yaml` includes namespace, ingress, MySQL PVC/Service/StatefulSet, and non-secret backend/frontend resources. Secrets are prepared separately and do not appear in resources. Keep `backend-secret`, `mysql-secret`, key names, and database credentials consistent without reproducing values in skills.

Use `../rent-manager-deploy/SKILL.md` for setup, CI, updates, troubleshooting, and rollback procedures. Do not confuse declared configuration with the state of a running cluster.
