---
name: rent-manager-payments
description: Work on rental payment registration, retrieval, listing, and updates, periods, suggested and applied amounts, and payment flags in Rent Manager.
---

# Rental payments

Read [coverage](references/coverage.md) and `../rent-manager-calculations/SKILL.md`. A recorded payment represents a contract installment/period, not a bank instruction.

## Creation flow

1. Open from a contract/schedule with `rentalContractId` and `period` query parameters.
2. RentPaymentCreateComponent takes today's local date and requests a suggestion; changing the date uses switchMap to replace the previous request.
3. The form receives suggested base, adjustment, interest, and total. Keep suggested and applied values in separate fields.
4. Saving requires a ready suggestion, no refresh in progress, and a valid form.
5. RentPaymentWriteService sends `POST /v1/rent-payments`; the handler converts the DTO to an entity, and the use case applies policy and persists submitted values. The create use case does not automatically recalculate those amounts.
6. Return to contract details after saving or canceling.

## Other operations

`GET /v1/rent-payments` lists with rental_contract_id/page/limit; `GET /:id` retrieves and `PUT /:id` updates. All require JWT. The UI exposes create/edit in feature/rent-payment, without a standalone list equivalent to other entities. Use the schedule for display. There is no DELETE.

During editing, RentPaymentEditComponent calls cancel() after a successful save or cancellation. That method navigates to `/app/rental-contracts`, which does not exist (the actual module is `/app/rent-contracts`), so navigation reaches the router fallback. Document this as a current limitation; do not confuse it with failed persistence or claim editing correctly returns to the contract.

## Rules and limitations

Require a positive contract ID, period/due date present, positive base and total, and nonnegative paid_amount; trim notes. Uniqueness is by contract/period: creation rejects duplicates, updates exclude the same ID. Update also fetches the contract and checks the period is within validity through shared.IsPeriodWithinContract; creation does not make that same query even though it receives the contract repository.

HTTP period: `YYYY-MM`; dates: `YYYY-MM-DD`. PaymentDate is optional in the backend DTO but required in the UI. IsPaid and PaidAmount are independent fields; do not infer payment from the amount. The current form has no paid_amount control despite a getter: check the actual payload before assuming UI support for partial payments.

## Verify

Check create/update for the same period, duplicates, periods outside validity, empty dates, suggestion errors, and rapid date changes. Verify the final payload and server reload, not only displayed values. Preserve suggested/applied distinctions and rounding when changing calculations.
