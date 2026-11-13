---
name: rent-manager-dashboard
description: Work on the dashboard, property cards, contract summaries, monthly payment schedules, next installments, and next adjustments in Rent Manager.
---

# Dashboard, summary, and schedule

Read [coverage](references/coverage.md). These are three related but distinct calculations; do not assume their amounts are equivalent.

## Business dashboard

DashboardComponent → PropertyReadService.listSummary$ → GET `/v1/properties/summaries` → ListPropertiesSummariesUseCase → paginated active contracts → one GetRentalContractSummary per contract → PropertySummary → DashboardMapper → PropertySummaryCard. Include only properties with active contracts, not all properties. Aggregation queries each contract and fails when any summary fails.

## Contract summary

`GET /v1/rental-contracts/:id/summary`: load contract/payments; map only IsPaid records by period; count unique paid periods, remaining installments with a zero floor, and the first pending month within TotalPayments. If the previous month is paid, use its TotalAmount as the current amount. Calculate the next suggestion using that amount and the current UTC date. Next adjustment: advance by frequency from contract start and choose the first month >= current month within contract end; this can be null.

## Payment schedule

`GET /v1/rental-contracts/:id/payment-schedules`: generate TotalPayments rows from the starting month. If a record exists for a period, use its amounts/dates/payment flag. Otherwise, calculate the due date and use the previous record's TotalAmount or contractual MonthlyAmount, marking it unpaid. In the next adjustment month calculated relative to today, populate suggestions through the calculator.

Observed limitations: nextAdjustmentPeriod is dereferenced even though it can be nil; calculator and previous lookup errors are ignored; PaidAmount is initialized with the amount even in an unpaid row. Do not interpret PaidAmount alone as a collected payment. The schedule is neither a persisted ledger nor a full projection of future inflation.

## Navigation

Overview loads the contract, schedule, and property. Registering payment from a row navigates to `/app/rent-payments/create` with contract/period. Also inspect the contract view and ui/rent-payment-schedule component for actual actions and presentation.

## Verify

Check contracts without payments, fully paid contracts, gaps between months, existing unpaid rows, contracts ending before today, null adjustment, missing indexes, dependency errors, and navigation to register/edit. Compare summary and schedule against stored data; do not infer bank balances or general accounting debt from these structures.
