---
name: rent-manager-properties
description: Work on property creation, search, listing, retrieval, and updates, including owners, types, statuses, and addresses in Rent Manager.
---

# Properties

Read [coverage](references/coverage.md). UI: feature/property and core/property; Go: property handlers, use cases, policies, and repository.

## Business behavior

Maintain a property identified by code, title, and address, associated with an owner, type, and status, so contracts can reference it. Dashboard card retrieval uses active contracts and is not a complete inventory of all properties.

## Operations

JWT is required for `POST /v1/properties`, `PUT /v1/properties/:id`, `GET /v1/properties`, `/search`, and `/:id`. The summary route is registered as `properties.GET("summaries", ...)`; verify its effective URL and Angular consumer before changing it. There is no DELETE.

Creation and updates normalize code, title, description, and address. They require positive owner, type, status, country, and state IDs; code, title, street, and city must not be empty. They reject duplicate codes, excluding the current ID during updates. Read DTOs for binding and repository/migrations for actual referential validation.

## UI flow

List/search → open create or load ID → form and catalogs → owner selection → write service → handler → use case → policy → SQL → data → update/navigation. Selects obtain shared catalogs; on country changes, reload states according to the form implementation.

## Verify

Check unique codes, unrelated/missing IDs, required fields, addresses/catalogs, paginated lists, and readonly details. Read `../rent-manager-dashboard/SKILL.md` for aggregation and contract navigation; do not duplicate financial calculations inside property logic.
