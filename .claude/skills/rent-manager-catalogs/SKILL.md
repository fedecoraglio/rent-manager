---
name: rent-manager-catalogs
description: Work with country, state, role, property type/status, contract status, interest calculation, and rent adjustment catalogs in Rent Manager.
---

# Catalogs and locations

Read [coverage](references/coverage.md) for sources, SQL codes, and DTOs. All these routes require JWT:

| Query | Purpose |
| --- | --- |
| GET /v1/roles | User roles |
| GET /v1/locations/countries | Countries for forms |
| GET /v1/locations/countries/:countryId/states | Country states/provinces |
| GET /v1/property-catalogs/types | Property types |
| GET /v1/property-catalogs/statuses | Property statuses |
| GET /v1/contract-catalogs/statuses | Contract statuses |
| GET /v1/contract-catalogs/interest-calculation-types | Late interest calculation modes |
| GET /v1/contract-catalogs/rent-adjustment-types | Rent adjustment modes |

These are read-only catalogs; no catalog editing UI/endpoint is exposed. Migrations establish initial data/codes. Angular core services call the API, maintain state, and supply selects; inspect each service's caching/loading before adding requests.

Do not rename codes used by business rules without migrating consumers: `ipc_argentina`, `daily_from_due_day_next_day`, and `daily_from_month_first_day` are compared literally. Also verify the active-contract code in repository/rent_contract.go and its seeds rather than inferring it from the UI label.

When adding a value, review the migration, entity/model, DTO, service, translations, form, and relevant calculator. When changing country, avoid retaining an incompatible state. Catalog options do not imply a calculator exists for every option.
