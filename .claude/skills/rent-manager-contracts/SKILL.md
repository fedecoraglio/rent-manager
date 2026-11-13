---
name: rent-manager-contracts
description: Work on rental contract creation, listing, retrieval, and updates, financial terms, validity, installments, and the active-contract-per-property restriction.
---

# Rental contracts

Read [coverage](references/coverage.md). Use UI feature/rent-contract and core/rent-contract; Go rent_contract*.go adapters/domain, usecase/rental_contract, and policy/rent_contract.

## Business behavior

Link a property and tenant and record validity, currency, monthly rent, deposit, due day, daily late interest percentage, interest/adjustment types, adjustment frequency, total installments, status, and notes. The application records terms and payments; it does not sign contracts or execute transfers.

## Operations and flow

`POST /v1/rental-contracts`, `GET /v1/rental-contracts`, `GET /v1/rental-contracts/:rentalContractId`, and PUT at the same ID require JWT. List → create/edit/detail → form with property/tenant search and catalogs → write service → handler/DTO → use case → policy → SQL. There is no DELETE or separate renewal/cancellation endpoint; status changes go through update.

Normalize currency to uppercase and trim notes. Require positive references/types/status, dates present with end after start, positive monthly amount, nonnegative deposit and late interest, due day 1–31, and positive frequency. Creation requires positive TotalPayments; the update policy does not repeat that check. Binding may impose additional restrictions.

Both policies query the property's active contract before saving; creation rejects an existing one, while update excludes its own ID. This is not general date-range overlap detection and does not guarantee uniqueness under concurrency. Inspect the active SQL condition and constraints before changing behavior.

Dates are converted in the adapter; do not send timestamps where the DTO expects a date. Distinguish contractual monthly_amount, stored installment amounts, and next payment suggestions. Installment count controls schedule length; it is not automatically derived from the date range.

## Verify

Test active-contract conflicts, property/status changes, invalid dates, frequency/installments, relationship selection, readonly details, and navigation to payments/overview. Read the dedicated skills for summaries, schedules, and suggestions; do not automatically apply a suggested change to the contract.
