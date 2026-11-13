---
name: rent-manager-calculations
description: Analyze and modify rent suggestions, compounded IPC adjustments, index selection, daily late interest, calculation bases, due dates, and rounding.
---

# Financial calculations

Read [coverage](references/coverage.md). Main sources: `internal/core/service/inflation_adjustment_calculator.go`, `rent_payment_calculator.go`, and usecase/rent_payment/get_rent_payment_suggestion_use_case.go.

## Suggestion

`GET /v1/rental-contracts/:rentalContractId/payment-suggestions?period=YYYY-MM&payment_date=YYYY-MM-DD` requires JWT. Load the contract, look up the previous month's payment, and use its TotalAmount as MonthlyAmount when nonzero. The repository returns `(nil,nil)` when no previous payment exists, so the contractual amount remains; other errors propagate. Calculate and return base, adjustment, interest, total, dates, and interest code; do not write data.

## IPC adjustment

Apply only with a non-null contract, frequency > 0, type `ipc_argentina`, a period after the initial month, and a month distance divisible by frequency. Normalize to the first day of the month in UTC.

Look up N indexes from period minus frequency through the preceding month. If the count differs from N, fetch the latest N indexes before the period. If still insufficient, return 0 with missing_indexes status; non-applicable cases use not_applicable. The check is count-based; do not assume continuity checks or official publication validation. Indexes are entered manually/via API without querying INDEC.

`percentage = (product(1 + index/100) - 1) * 100`, rounded to two decimals. `adjustment = Round(base * percentage) / 100`. Status/message exist in the internal calculation; the payment-suggestions DTO does not expose them. Do not claim the UI receives missing-index details.

## Late interest and total

Do not apply interest when rate <= 0, type is null, or payment date precedes the day after the full due date. Once that threshold is crossed:

- `daily_from_due_day_next_day`: count from the day after the due date.
- `daily_from_month_first_day`: count from the first day of the month, although activation still requires passing the due date.
- Unknown code: use the day after the due date.

`days = int((paymentDate - start).Hours()/24) + 1`. `interest = (base + adjustment) * dailyRate/100 * days`; this is not compound interest. `total = Round(base + adjustment + interest)` to an integer. Go uses float64, not a monetary decimal type.

CalculateDueDate uses time.Date directly: day 31 in a shorter month normalizes into the following month rather than clamping to month-end. Compare backend UTC dates against the UI's local date.

## Validate changes

Use boundary examples: initial, intermediate, and adjustment months; exact/incomplete/insufficient indexes; payment on the due date and the following day; both interest modes; February/day 31; and a base inherited from the previous month, possibly including interest. Do not describe proposed corrections as current behavior.
