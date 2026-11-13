---
name: rent-manager-tenants
description: Work on tenant creation, updates, search, listing, retrieval, personal details, and addresses in Rent Manager.
---

# Tenants

Read [coverage](references/coverage.md). Follow feature/tenant → core/tenant or write services → TenantHandler → usecase/tenant → policy/tenant → tenant repository and tenants table.

## Business behavior

Register the tenant assigned to a contract, including name, document number, contact details, and address. This module does not perform credit evaluation or execute financial guarantees.

## Implemented flows

`POST /v1/tenants` creates; `PUT /v1/tenants/:id` updates; `GET /v1/tenants` lists; `/search` searches; `/:id` retrieves details. All require JWT. Deletion is not exposed. Document lookup methods are internal and participate in validation.

Policies trim name, phone, document number, and address components, and normalize email. They require name and document number and validate email only when present. They reject an existing document number; updates allow retaining the current record's number. Read the DTO and form for additional address/country/state requirements; do not equate them with policy validation.

For forms, load location catalogs and preserve the country → states relationship. Tenant selection in contracts uses shared read/search services; verify that the submitted ID belongs to the selected tenant.

## Verify

Test creation, search for contract selection, readonly details, updates, duplicate document numbers, empty optional fields, invalid email, missing IDs, and location changes. Inspect SQL/constraints for relationships rather than assuming policies check referenced-record existence.
