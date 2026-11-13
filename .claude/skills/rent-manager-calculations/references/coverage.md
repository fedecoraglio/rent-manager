# Source and operation coverage

Reviewed: 2026-11-13. Resolve paths from the repository root. Use this map to open the implementation; excerpts are navigation indexes, not replacements for complete source bodies. Line numbers refer to this review. Do not interpret internal operations as public endpoints.

## Contents

- File inventory
- Operations, data contracts, and rule locations

## File inventory

- `rent-manager-backend/internal/adapter/http/rent_payment_suggestion_request.go`
- `rent-manager-backend/internal/adapter/http/rent_payment_suggestion_response.go`
- `rent-manager-backend/internal/adapter/store/mysql/repository/shared/shared.go`
- `rent-manager-backend/internal/core/domain/adjustment_suggestion.go`
- `rent-manager-backend/internal/core/domain/rent_payment_suggestion.go`
- `rent-manager-backend/internal/core/policy/inflation_index/inflation_index_create_policy.go`
- `rent-manager-backend/internal/core/policy/inflation_index/inflation_index_update_policy.go`
- `rent-manager-backend/internal/core/policy/shared/shared.go`
- `rent-manager-backend/internal/core/port/rent_payment_calculator.go`
- `rent-manager-backend/internal/core/service/inflation_adjustment_calculator.go`
- `rent-manager-backend/internal/core/service/rent_payment_calculator.go`
- `rent-manager-backend/internal/core/usecase/inflation_index/create_inflation_index_use_case.go`
- `rent-manager-backend/internal/core/usecase/inflation_index/get_inflation_index_by_id_use_case.go`
- `rent-manager-backend/internal/core/usecase/inflation_index/get_inflation_index_by_period_use_case.go`
- `rent-manager-backend/internal/core/usecase/inflation_index/list_inflation_indexes_by_period_range_use_case.go`
- `rent-manager-backend/internal/core/usecase/inflation_index/list_inflation_indexes_use_case.go`
- `rent-manager-backend/internal/core/usecase/inflation_index/update_inflation_index_use_case.go`
- `rent-manager-backend/internal/core/usecase/rent_payment/get_rent_payment_suggestion_use_case.go`

## Operations, data contracts, and rule locations

### rent-manager-backend/internal/adapter/http/rent_payment_suggestion_request.go

```text
4: Period      string `form:"period" binding:"required"`
5: PaymentDate string `form:"payment_date" binding:"required"`
```

### rent-manager-backend/internal/adapter/http/rent_payment_suggestion_response.go

```text
6: InterestCalculationType       string  `json:"interest_calculation_type"`
7: RentalContractID              int64   `json:"rental_contract_id"`
8: Period                        string  `json:"period"`
9: DueDate                       string  `json:"due_date"`
10: PaymentDate                   string  `json:"payment_date"`
11: BaseAmount                    float64 `json:"base_amount"`
12: SuggestedAdjustmentPercentage float64 `json:"suggested_adjustment_percentage"`
13: SuggestedAdjustmentAmount     float64 `json:"suggested_adjustment_amount"`
14: SuggestedInterestAmount       float64 `json:"suggested_interest_amount"`
15: SuggestedTotalAmount          float64 `json:"suggested_total_amount"`
18: func newRentPaymentSuggestionResponse(
```

### rent-manager-backend/internal/adapter/store/mysql/repository/shared/shared.go

```text
3: func NullInt64(value int64) any {
4: if value == 0 {
```

### rent-manager-backend/internal/core/policy/inflation_index/inflation_index_create_policy.go

```text
15: func NewInflationIndexCreatePolicy(
23: func (p *InflationIndexCreatePolicy) Validate(ctx context.Context, inflationIndex *domain.InflationIndex) error {
24: if inflationIndex.Period.IsZero() {
25: return domain.ErrInflationIndexPeriodRequired
28: if inflationIndex.Percentage < 0 {
29: return domain.ErrInflationIndexPercentageInvalid
33: if err == nil && existingInflationIndex != nil {
34: return domain.ErrInflationIndexAlreadyExists
37: if err != nil && !errors.Is(err, domain.ErrInflationIndexNotFound) {
```

### rent-manager-backend/internal/core/policy/inflation_index/inflation_index_update_policy.go

```text
15: func NewInflationIndexUpdatePolicy(
23: func (p *InflationIndexUpdatePolicy) Validate(
27: if inflationIndex.ID <= 0 {
28: return domain.ErrInflationIndexNotFound
31: if inflationIndex.Period.IsZero() {
32: return domain.ErrInflationIndexPeriodRequired
35: if inflationIndex.Percentage < 0 {
36: return domain.ErrInflationIndexPercentageInvalid
40: if err != nil {
45: if err == nil &&
48: return domain.ErrInflationIndexAlreadyExists
51: if err != nil && !errors.Is(err, domain.ErrInflationIndexNotFound) {
```

### rent-manager-backend/internal/core/policy/shared/shared.go

```text
9: func IsValidEmail(email string) bool {
10: if email == "" {
22: func CalculateTotalPayments(startDate *time.Time, endDate *time.Time) int64 {
23: if startDate == nil || endDate == nil || startDate.IsZero() || endDate.IsZero() {
32: if totalMonths < 0 {
39: func IsPeriodWithinContract(period time.Time, rentalContract *domain.RentalContract) bool {
46: func FirstDayOfMonth(value time.Time) time.Time {
```

### rent-manager-backend/internal/core/service/inflation_adjustment_calculator.go

```text
18: func NewIPCInflationAdjustmentCalculator(
26: func (calculator *IPCInflationAdjustmentCalculator) CalculateSuggestedAdjustment(
36: if rentalContract == nil {
40: if rentalContract.AdjustmentFrequencyMonths <= 0 {
44: if rentalContract.AdjustmentType == nil {
48: if rentalContract.AdjustmentType.Code != rentAdjustmentTypeIPCArgentina {
55: if normalizedPeriod.Before(normalizedStartDate) {
61: if monthsFromStart == 0 {
65: if monthsFromStart%int(rentalContract.AdjustmentFrequencyMonths) != 0 {
82: if err != nil {
88: // Fall back to the latest available indexes if the expected range is incomplete.
89: if len(indexes) != expectedIndexes {
95: if err != nil {
99: if len(indexes) != expectedIndexes {
124: func monthDiff(from time.Time, to time.Time) int {
```

### rent-manager-backend/internal/core/service/rent_payment_calculator.go

```text
28: func NewRentPaymentCalculator(
36: func (calculator *RentPaymentCalculator) Calculate(
49: if calculator.inflationAdjustmentCalculator != nil {
52: if err != nil {
56: if calculatedSuggestion != nil {
83: func calculateInterestAmount(
89: if rentalContract.DailyInterestPercentage <= 0 {
93: if rentalContract.InterestCalculationType == nil {
102: if paymentDate.Before(firstOverdueDate) {
111: if paymentDate.Before(startDate) {
116: if lateDays <= 0 {
123: func getInterestStartDate(rentalContract *domain.RentalContract, period time.Time) time.Time {
127: case interestTypeDailyFromDueDayNextDay:
139: case interestTypeDailyFromMonthFirstDay:
156: func firstDayOfMonth(value time.Time) time.Time {
```

### rent-manager-backend/internal/core/usecase/inflation_index/create_inflation_index_use_case.go

```text
11: type CreateInflationIndexUseCase struct {
16: func NewCreateInflationIndexUseCase(
19: ) *CreateInflationIndexUseCase {
26: func (uc *CreateInflationIndexUseCase) Execute(
30: if err := uc.inflationIndexCreatePolicy.Validate(ctx, inflationIndex); err != nil {
34: if err := uc.inflationIndexRepository.Create(ctx, inflationIndex); err != nil {
```

### rent-manager-backend/internal/core/usecase/inflation_index/get_inflation_index_by_id_use_case.go

```text
10: type GetInflationIndexByIDUseCase struct {
14: func NewGetInflationIndexByIDUseCase(
16: ) *GetInflationIndexByIDUseCase {
22: func (uc *GetInflationIndexByIDUseCase) Execute(ctx context.Context, id int64) (*domain.InflationIndex, error) {
```

### rent-manager-backend/internal/core/usecase/inflation_index/get_inflation_index_by_period_use_case.go

```text
11: type GetInflationIndexByPeriodUseCase struct {
15: func NewGetInflationIndexByPeriodUseCase(
17: ) *GetInflationIndexByPeriodUseCase {
23: func (uc *GetInflationIndexByPeriodUseCase) Execute(
```

### rent-manager-backend/internal/core/usecase/inflation_index/list_inflation_indexes_by_period_range_use_case.go

```text
11: type ListInflationIndexesByPeriodRangeUseCase struct {
15: func NewListInflationIndexesByPeriodRangeUseCase(
17: ) *ListInflationIndexesByPeriodRangeUseCase {
23: func (uc *ListInflationIndexesByPeriodRangeUseCase) Execute(
```

### rent-manager-backend/internal/core/usecase/inflation_index/list_inflation_indexes_use_case.go

```text
10: type ListInflationIndexesUseCase struct {
14: func NewListInflationIndexesUseCase(
16: ) *ListInflationIndexesUseCase {
22: func (uc *ListInflationIndexesUseCase) Execute(
```

### rent-manager-backend/internal/core/usecase/inflation_index/update_inflation_index_use_case.go

```text
11: type UpdateInflationIndexUseCase struct {
16: func NewUpdateInflationIndexUseCase(
19: ) *UpdateInflationIndexUseCase {
26: func (uc *UpdateInflationIndexUseCase) Execute(
30: if err := uc.inflationIndexUpdatePolicy.Validate(ctx, inflationIndex); err != nil {
34: if err := uc.inflationIndexRepository.Update(ctx, inflationIndex); err != nil {
```

### rent-manager-backend/internal/core/usecase/rent_payment/get_rent_payment_suggestion_use_case.go

```text
12: type GetRentPaymentSuggestionUseCase struct {
18: func NewGetRentPaymentSuggestionUseCase(
22: ) *GetRentPaymentSuggestionUseCase {
30: func (uc *GetRentPaymentSuggestionUseCase) GetRentPaymentSuggestion(
37: if err != nil {
44: if errPeriod != nil {
48: if previousRent != nil && previousRent.TotalAmount != 0 {
58: if err != nil {
66: if rentalContract.InterestCalculationType != nil {
84: func CalculateDueDate(period time.Time, dueDay int64) time.Time {
```
