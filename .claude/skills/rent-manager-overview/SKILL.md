---
name: rent-manager-overview
description: Navigate Rent Manager, locate business and technical flows, and assess changes across the Go backend, Angular frontend, Java AI service, and Kubernetes platform.
---

# Rent Manager map

Resolve source paths from the repository root. This documentation describes local code reviewed on 2026-11-13; check sources before changing behavior. Distinguish implementation, limitations, and proposals. Do not infer capabilities from field names or comments.

## Architecture

| Project | Responsibility | Skill |
| --- | --- | --- |
| rent-manager-backend | Management API and calculations; Go, Gin, and MySQL | `../rent-manager-backend/SKILL.md` |
| rent-manager-frontend | Angular management SPA | `../rent-manager-frontend/SKILL.md` |
| rent-manager-ai | PDFs and RAG queries; Java, PostgreSQL/pgvector, and Ollama | `../rent-manager-ai/SKILL.md` |
| platform | Kubernetes and Argo CD resources | `../rent-manager-platform/SKILL.md` |

The browser calls `/v1` in Go. Java exposes a separate API without `/v1`; Angular routes do not integrate documents/chat, and Go does not proxy to AI. Do not describe that integration as complete.

## Select a flow

| Need | Skill |
| --- | --- |
| Login, JWT, session, logout | `../rent-manager-auth/SKILL.md` |
| Users, roles, passwords | `../rent-manager-users/SKILL.md` |
| Owners | `../rent-manager-owners/SKILL.md` |
| Tenants and addresses | `../rent-manager-tenants/SKILL.md` |
| Properties, addresses, search | `../rent-manager-properties/SKILL.md` |
| Countries, states, types, statuses | `../rent-manager-catalogs/SKILL.md` |
| Contract creation, retrieval, updates | `../rent-manager-contracts/SKILL.md` |
| Payment registration, updates, retrieval | `../rent-manager-payments/SKILL.md` |
| Suggestions, late interest, IPC adjustments | `../rent-manager-calculations/SKILL.md` |
| Schedule, summary, dashboard | `../rent-manager-dashboard/SKILL.md` |
| Inflation index creation, updates, retrieval | `../rent-manager-inflation/SKILL.md` |
| PDF upload, listing, download, deletion | `../rent-manager-documents/SKILL.md` |
| Indexing, embeddings, chat | `../rent-manager-rag/SKILL.md` |
| Configuration, language, loading, UI errors | `../rent-manager-ui-runtime/SKILL.md` |
| Build, CI, containers, GitOps, operations, recovery | `../rent-manager-deploy/SKILL.md` |

## Investigate or modify

1. Identify screen/route, DTO, handler, use case, policy, and repository.
2. Read the relevant coverage reference for source and operation inventories, API fields, and validation conditions extracted from code.
3. Compare rules against migrations; a policy check before INSERT does not ensure uniqueness under concurrency.
4. Preserve date formats and JSON names between Go and Angular. AI has independent HTTP contracts.
5. Verify success, invalid input, missing records, duplicates, dependency failures, and expired sessions as appropriate.
6. Update affected skills when routes, rules, schemas, or deployment change.

## Review scope

Read the [inventory](references/inventory.md) for source files, configuration, and exclusions, and the [operation matrix](references/flows.md) for every Go route and use case. Exclude dependencies, builds, caches, IDE settings, private data, `.env`, real secrets, and database dumps. Cover implemented flows, including API-only paths; do not invent deletion without an endpoint, bank payments, external notifications, scheduled jobs, or AI/UI integration.
