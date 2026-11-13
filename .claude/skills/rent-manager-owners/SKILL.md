---
name: rent-manager-owners
description: Work on owner creation, updates, search, listing, and retrieval in Rent Manager, including document identity, contact details, and property associations.
---

# Owners

Read [coverage](references/coverage.md). Use Angular `feature/owner`, `core/owner` for reads, and Go `internal/adapter/http/owner*.go`, `core/usecase/owner`, `core/policy/owner`, and repository/owner.go.

## Business behavior

Register the property owner so they can be selected when creating/updating properties. Store name, document number, and contact details. The document number identifies duplicates; email is optional but must be valid when supplied.

## Flows

- Create: form → write service → `POST /v1/owners` → binding → entity → create policy → INSERT → response → success navigation.
- Update: load by ID, complete the form, and `PUT /v1/owners/:id`; exclude the current ID when checking duplicate document numbers.
- List: `GET /v1/owners` with DTO pagination; transform data into UI models.
- Search: `GET /v1/owners/search`; follow the DTO/API service query and limits; reuse this flow to select an owner.
- Detail: `GET /v1/owners/:id`; the view uses a readonly form.

Trim name, phone, and document number, and trim/lowercase email. Require name/document number; do not confuse internal email lookup with a dedicated endpoint. Policy checks do not replace SQL constraints. There is no owner DELETE.

## Verify

Check duplicate document numbers, updates without changing the document number, empty/malformed email, missing records, pagination, and selection from a property form. Inspect joins and foreign keys in migrations before changing identity or relationships.
