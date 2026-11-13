# Source and operation coverage

Reviewed: 2026-11-13. Resolve paths from the repository root. Use this map to open the implementation; excerpts are navigation indexes, not replacements for complete source bodies. Line numbers refer to this review. Do not interpret internal operations as public endpoints.

## Contents

- File inventory
- Operations, data contracts, and rule locations

## File inventory

- `rent-manager-backend/internal/adapter/http/rent_payment_schedule_response.go`
- `rent-manager-backend/internal/adapter/http/rental_contract_summary_response.go`
- `rent-manager-backend/internal/core/domain/property_summary.go`
- `rent-manager-backend/internal/core/domain/rent_payment_schedule.go`
- `rent-manager-backend/internal/core/domain/rental_contract_summary.go`
- `rent-manager-backend/internal/core/usecase/property/get_property_summaries_use_case.go`
- `rent-manager-backend/internal/core/usecase/rent_payment/get_rent_payment_schedule_use_case.go`
- `rent-manager-backend/internal/core/usecase/rent_payment/get_rental_contract_summary_use_case.go`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/dashboard/dashboard.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/dashboard/dashboard.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/dashboard/dashboard.mapper.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/dashboard/dashboard.router.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/dashboard/model/dashboard.model.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-contract/overview/rent-contract-overview.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-contract/overview/rent-contract-overview.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/ui/property-summary-card/property-summary-card.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/ui/property-summary-card/property-summary-card.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/ui/property-summary-card/property-summary-card.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/ui/property-summary-card/property-summary-card.model.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/ui/rent-payment-schedule/rent-payment-schedule.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/ui/rent-payment-schedule/rent-payment-schedule.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/ui/rent-payment-schedule/rent-payment-schedule.component.ts`

## Operations, data contracts, and rule locations

### rent-manager-backend/internal/adapter/http/rent_payment_schedule_response.go

```text
8: RentalContractID              int64   `json:"rental_contract_id"`
9: RentPaymentID                 int64   `json:"rent_payment_id"`
10: Period                        string  `json:"period"`
11: DueDate                       string  `json:"due_date"`
12: BaseAmount                    float64 `json:"base_amount"`
13: SuggestedAdjustmentPercentage float64 `json:"suggested_adjustment_percentage"`
14: SuggestedInterestAmount       float64 `json:"suggested_interest_amount"`
15: SuggestedTotalAmount          float64 `json:"suggested_total_amount"`
16: AppliedAdjustmentPercentage   float64 `json:"applied_adjustment_percentage"`
17: AppliedInterestAmount         float64 `json:"applied_interest_amount"`
18: TotalAmount                   float64 `json:"total_amount"`
19: PaidAmount                    float64 `json:"paid_amount"`
20: PaymentDate                   *string `json:"payment_date"`
21: IsPaid                        bool    `json:"is_paid"`
24: func newRentPaymentScheduleItemResponse(
29: if item.PaymentDate != nil {
52: func newRentPaymentScheduleResponse(
```

### rent-manager-backend/internal/adapter/http/rental_contract_summary_response.go

```text
6: RentalContractID     int64   `json:"rental_contract_id"`
7: TotalPayments        int64   `json:"total_payments"`
8: PaidPayments         int64   `json:"paid_payments"`
9: RemainingPayments    int64   `json:"remaining_payments"`
10: CurrentAmount        float64 `json:"current_amount"`
11: NextSuggestedAmount  float64 `json:"next_suggested_amount"`
12: NextPendingPeriod    *string `json:"next_pending_period"`
13: NextAdjustmentPeriod *string `json:"next_adjustment_period"`
16: func newRentalContractSummaryResponse(
22: if summary.NextPendingPeriod != nil {
27: if summary.NextAdjustmentPeriod != nil {
```

### rent-manager-backend/internal/core/usecase/property/get_property_summaries_use_case.go

```text
11: type ListPropertiesSummariesUseCase struct {
13: getRentalContractSummaryUseCase *rent_payment.GetRentalContractSummaryUseCase
16: func NewListPropertiesSummariesUseCase(
18: getRentalContractSummaryUseCase *rent_payment.GetRentalContractSummaryUseCase,
19: ) *ListPropertiesSummariesUseCase {
26: func (uc *ListPropertiesSummariesUseCase) Execute(
32: if err != nil {
40: if err != nil {
49: if contract.Property != nil {
```

### rent-manager-backend/internal/core/usecase/rent_payment/get_rent_payment_schedule_use_case.go

```text
12: type GetRentPaymentScheduleUseCase struct {
18: func NewGetRentPaymentScheduleUseCase(
22: ) *GetRentPaymentScheduleUseCase {
30: func (uc *GetRentPaymentScheduleUseCase) GetRentPaymentSchedule(
35: if err != nil {
41: if err != nil {
55: func (uc *GetRentPaymentScheduleUseCase) buildPaymentSchedule(
72: if rentPayment, exists := paymentsByPeriod[periodKey(currentPeriod)]; exists {
83: if *nextAdjustmentPeriod == currentPeriod {
97: if previousRent != nil {
114: func firstDayOfMonth(value time.Time) time.Time {
118: func periodKey(value time.Time) string {
```

### rent-manager-backend/internal/core/usecase/rent_payment/get_rental_contract_summary_use_case.go

```text
11: type GetRentalContractSummaryUseCase struct {
17: func NewGetRentalContractSummaryUseCase(
21: ) *GetRentalContractSummaryUseCase {
29: func (uc *GetRentalContractSummaryUseCase) GetRentalContractSummary(
34: if err != nil {
40: if err != nil {
48: if rentPayment.IsPaid {
56: if remainingPayments < 0 {
64: if _, exists := paidByPeriod[periodKey(currentPeriod)]; !exists {
75: if exists {
80: if nextPendingPeriod != nil {
87: if err != nil {
109: func calculateNextAdjustmentPeriod(rentalContract *domain.RentalContract) *time.Time {
110: if rentalContract.AdjustmentFrequencyMonths <= 0 {
125: if currentPeriod.After(endPeriod) {
129: if currentPeriod.After(nowPeriod) || currentPeriod.Equal(nowPeriod) {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/dashboard/dashboard.component.ts

```text
25: constructor() {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-contract/overview/rent-contract-overview.component.ts

```text
38: constructor() {
41: switchMap(() => this.router.navigate(['/app/dashboard'])),
47: ngOnInit(): void {
51: void this.router.navigate(['/app/dashboard']);
58: tap(() => {
70: ngOnDestroy(): void {
75: goBack(): void {
79: registerPayment(item: RentPaymentScheduleItem): void {
80: void this.router.navigate(['/app/rent-payments/create'], {
81: queryParams: {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/ui/property-summary-card/property-summary-card.component.ts

```text
21: constructor() {
26: this.router.navigate([
37: gotToViewProperty() {
```
