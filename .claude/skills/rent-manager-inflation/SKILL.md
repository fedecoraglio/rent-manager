---
name: rent-manager-inflation
description: Work on inflation index creation, listing, retrieval, updates, and consumption by IPC rental adjustments in Rent Manager.
---

# Inflation indexes

Read [coverage](references/coverage.md). Record monthly percentages used by IPC suggestions. There is no automatic INDEC integration or import scheduler.

## Flows

UI feature/inflation-index uses core/inflation-index for reads and feature services for writes. Create → period/percentage form → POST `/v1/inflation-indexes`; edit → load ID → PUT `/:id`; list/details use root GET and GET `/:id`. There is no DELETE. These routes do NOT use AuthMiddleware in router.go, even though the screen is under `/app`.

Policy requires a nonzero period and percentage >= 0, rejecting another record for the same period. Update checks existence by ID and excludes the current record. Read DTO/use case for period parsing/normalization and pagination. Period/range/latest-before queries exist as internal calculation operations, not dedicated endpoints.

## Effects

Changing an index affects future recalculated suggestions; it does not rewrite stored payment values. `IPCInflationAdjustmentCalculator` compounds rates and uses an index-count fallback; read `../rent-manager-calculations/SKILL.md` before changing selection/rounding.

## Verify

Check duplicate months, zero/negative percentages, missing IDs, updates retaining the month, incomplete ranges, and latest available indexes. Inspect HTTP mapping for inflation errors: not all appear in statusFromErrorCode and some may become 500. Do not describe the Angular guard as API security.
