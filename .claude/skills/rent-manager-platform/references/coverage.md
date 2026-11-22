# Source and operation coverage

Reviewed: 2026-11-13. Resolve paths from the repository root. Use this map to open the implementation; excerpts are navigation indexes, not replacements for complete source bodies. Line numbers refer to this review. Do not interpret internal operations as public endpoints.

## Contents

- File inventory
- Operations, data contracts, and rule locations

## File inventory

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

## Operations, data contracts, and rule locations

## Kubernetes AI deployment additions

- `platform/kubernetes/ai/configmap.yaml`
- `platform/kubernetes/ai/deployment.yaml`
- `platform/kubernetes/ai/ingress.yaml`
- `platform/kubernetes/ai/pvc.yaml`
- `platform/kubernetes/ai/README.md`
- `platform/kubernetes/ai/service.yaml`
- `platform/kubernetes/postgres/configmap.yaml`
- `platform/kubernetes/postgres/pvc.yaml`
- `platform/kubernetes/postgres/secret.example.yaml`
- `platform/kubernetes/postgres/service.yaml`
- `platform/kubernetes/postgres/statefulset.yaml`
- `platform/kubernetes/ollama/configmap.yaml`
- `platform/kubernetes/ollama/pvc.yaml`
- `platform/kubernetes/ollama/service.yaml`
- `platform/kubernetes/ollama/statefulset.yaml`
- `.github/workflows/ai-image.yml`

These resources are included by the root Kustomization, except the Secret example. Prepare the real ignored Secret before syncing. The runbook describes image publication, bootstrap, persistence, routing, verification, and migration; live deployment still requires target-cluster validation.
