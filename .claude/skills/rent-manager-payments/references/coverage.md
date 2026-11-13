# Source and operation coverage

Reviewed: 2026-11-13. Resolve paths from the repository root. Use this map to open the implementation; excerpts are navigation indexes, not replacements for complete source bodies. Line numbers refer to this review. Do not interpret internal operations as public endpoints.

## Contents

- File inventory
- Operations, data contracts, and rule locations

## File inventory

- `rent-manager-backend/internal/adapter/http/rent_payment_request.go`
- `rent-manager-backend/internal/adapter/http/rent_payment_response.go`
- `rent-manager-backend/internal/adapter/http/rent_payment_schedule_response.go`
- `rent-manager-backend/internal/adapter/http/rent_payment_suggestion_request.go`
- `rent-manager-backend/internal/adapter/http/rent_payment_suggestion_response.go`
- `rent-manager-backend/internal/adapter/http/rent_payment.go`
- `rent-manager-backend/internal/adapter/store/mysql/repository/rent_payment.go`
- `rent-manager-backend/internal/core/domain/rent_payment_calculation.go`
- `rent-manager-backend/internal/core/domain/rent_payment_schedule.go`
- `rent-manager-backend/internal/core/domain/rent_payment_suggestion.go`
- `rent-manager-backend/internal/core/domain/rent_payment.go`
- `rent-manager-backend/internal/core/policy/rent_payment/rent_payment_create_policy.go`
- `rent-manager-backend/internal/core/policy/rent_payment/rent_payment_update_policy.go`
- `rent-manager-backend/internal/core/port/rent_payment_calculator.go`
- `rent-manager-backend/internal/core/port/rent_payment.go`
- `rent-manager-backend/internal/core/service/rent_payment_calculator.go`
- `rent-manager-backend/internal/core/usecase/rent_payment/create_rent_payment_use_case.go`
- `rent-manager-backend/internal/core/usecase/rent_payment/get_rent_payment_by_id_use_case.go`
- `rent-manager-backend/internal/core/usecase/rent_payment/get_rent_payment_schedule_use_case.go`
- `rent-manager-backend/internal/core/usecase/rent_payment/get_rent_payment_suggestion_use_case.go`
- `rent-manager-backend/internal/core/usecase/rent_payment/get_rental_contract_summary_use_case.go`
- `rent-manager-backend/internal/core/usecase/rent_payment/list_rent_payments_use_case.go`
- `rent-manager-backend/internal/core/usecase/rent_payment/update_rent_payment_use_case.go`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/rent-payment/rent-payment-api.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/rent-payment/rent-payment.model.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/rent-payment/rent-payment.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-payment/create/rent-payment-create.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-payment/create/rent-payment-create.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-payment/create/rent-payment-create.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-payment/edit/rent-payment-edit.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-payment/edit/rent-payment-edit.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-payment/edit/rent-payment-edit.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-payment/form/field/rent-payment-form-field.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-payment/form/rent-payment-form.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-payment/form/rent-payment-form.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-payment/form/rent-payment-form.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-payment/model/rent-payment-model.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-payment/rent-payment.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-payment/rent-payment.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-payment/rent-payment.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-payment/rent-payment.routes.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-payment/service/rent-payment-write-api.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-payment/service/rent-payment-write.service.ts`

## Operations, data contracts, and rule locations

### rent-manager-backend/internal/adapter/http/rent_payment_request.go

```text
10: RentalContractID              int64   `json:"rental_contract_id" binding:"required,min=1"`
11: Period                        string  `json:"period" binding:"required"`
12: DueDate                       string  `json:"due_date" binding:"required"`
13: PaymentDate                   string  `json:"payment_date"`
14: BaseAmount                    float64 `json:"base_amount" binding:"required,gt=0"`
15: SuggestedAdjustmentPercentage float64 `json:"suggested_adjustment_percentage"`
16: AppliedAdjustmentPercentage   float64 `json:"applied_adjustment_percentage"`
17: SuggestedInterestAmount       float64 `json:"suggested_interest_amount"`
18: AppliedInterestAmount         float64 `json:"applied_interest_amount"`
19: TotalAmount                   float64 `json:"total_amount" binding:"required,gt=0"`
20: PaidAmount                    float64 `json:"paid_amount"`
21: IsPaid                        bool    `json:"is_paid"`
22: Notes                         string  `json:"notes"`
26: RentalContractID              int64   `json:"rental_contract_id" binding:"required,min=1"`
27: Period                        string  `json:"period" binding:"required"`
28: DueDate                       string  `json:"due_date" binding:"required"`
29: PaymentDate                   string  `json:"payment_date"`
30: BaseAmount                    float64 `json:"base_amount" binding:"required,gt=0"`
31: SuggestedAdjustmentPercentage float64 `json:"suggested_adjustment_percentage"`
32: AppliedAdjustmentPercentage   float64 `json:"applied_adjustment_percentage"`
33: SuggestedInterestAmount       float64 `json:"suggested_interest_amount"`
34: AppliedInterestAmount         float64 `json:"applied_interest_amount"`
35: TotalAmount                   float64 `json:"total_amount" binding:"required,gt=0"`
36: PaidAmount                    float64 `json:"paid_amount"`
37: IsPaid                        bool    `json:"is_paid"`
38: Notes                         string  `json:"notes"`
42: RentalContractID int64  `form:"rental_contract_id"`
43: Page             uint64 `form:"page" binding:"required,min=1"`
44: Limit            uint64 `form:"limit" binding:"required,min=1,max=100"`
47: func newRentPaymentFromCreateRequest(
51: if err != nil {
56: if err != nil {
61: if err != nil {
65: if req.PaidAmount == 0 {
85: func newRentPaymentFromUpdateRequest(
90: if err != nil {
99: func parsePeriod(value string) (time.Time, error) {
103: func parseOptionalDate(value string) (*time.Time, error) {
104: if value == "" {
109: if err != nil {
```

### rent-manager-backend/internal/adapter/http/rent_payment_response.go

```text
10: ID                            int64     `json:"id"`
11: RentalContractID              int64     `json:"rental_contract_id"`
12: Period                        string    `json:"period"`
13: DueDate                       string    `json:"due_date"`
14: PaymentDate                   *string   `json:"payment_date"`
15: BaseAmount                    float64   `json:"base_amount"`
16: SuggestedAdjustmentPercentage float64   `json:"suggested_adjustment_percentage"`
17: AppliedAdjustmentPercentage   float64   `json:"applied_adjustment_percentage"`
18: SuggestedInterestAmount       float64   `json:"suggested_interest_amount"`
19: AppliedInterestAmount         float64   `json:"applied_interest_amount"`
20: TotalAmount                   float64   `json:"total_amount"`
21: PaidAmount                    float64   `json:"paid_amount"`
22: IsPaid                        bool      `json:"is_paid"`
23: Notes                         string    `json:"notes"`
24: CreatedAt                     time.Time `json:"created_at"`
25: UpdatedAt                     time.Time `json:"updated_at"`
28: func newRentPaymentResponse(
33: if rentPayment.PaymentDate != nil {
58: func newRentPaymentsResponse(
```

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

### rent-manager-backend/internal/adapter/http/rent_payment.go

```text
10: rentPaymentUseCase "rent-manager-backend/internal/core/usecase/rent_payment"
14: getRentPaymentScheduleUseCase   *rentPaymentUseCase.GetRentPaymentScheduleUseCase
15: createRentPaymentUseCase        *rentPaymentUseCase.CreateRentPaymentUseCase
16: updateRentPaymentUseCase        *rentPaymentUseCase.UpdateRentPaymentUseCase
17: getRentPaymentByIDUseCase       *rentPaymentUseCase.GetRentPaymentByIDUseCase
18: listRentPaymentsUseCase         *rentPaymentUseCase.ListRentPaymentsUseCase
19: getRentPaymentSuggestionUseCase *rentPaymentUseCase.GetRentPaymentSuggestionUseCase
20: getRentalContractSummaryUseCase *rentPaymentUseCase.GetRentalContractSummaryUseCase
23: func NewRentPaymentHandler(
24: getRentPaymentScheduleUseCase *rentPaymentUseCase.GetRentPaymentScheduleUseCase,
25: createRentPaymentUseCase *rentPaymentUseCase.CreateRentPaymentUseCase,
26: updateRentPaymentUseCase *rentPaymentUseCase.UpdateRentPaymentUseCase,
27: getRentPaymentByIDUseCase *rentPaymentUseCase.GetRentPaymentByIDUseCase,
28: listRentPaymentsUseCase *rentPaymentUseCase.ListRentPaymentsUseCase,
29: getRentPaymentSuggestionUseCase *rentPaymentUseCase.GetRentPaymentSuggestionUseCase,
30: getRentalContractSummaryUseCase *rentPaymentUseCase.GetRentalContractSummaryUseCase,
43: func (h *RentPaymentHandler) GetRentPaymentSchedule(ctx *gin.Context) {
45: if err != nil {
51: if err != nil {
60: func (h *RentPaymentHandler) CreateRentPayment(ctx *gin.Context) {
63: if err := ctx.ShouldBindJSON(&req); err != nil {
69: if err != nil {
75: if err != nil {
84: func (h *RentPaymentHandler) UpdateRentPayment(ctx *gin.Context) {
86: if err != nil {
93: if err := ctx.ShouldBindJSON(&req); err != nil {
99: if err != nil {
105: if err != nil {
114: func (h *RentPaymentHandler) GetRentPaymentByID(ctx *gin.Context) {
116: if err != nil {
122: if err != nil {
131: func (h *RentPaymentHandler) ListRentPayments(ctx *gin.Context) {
134: if err := ctx.ShouldBindQuery(&req); err != nil {
145: if err != nil {
154: func (h *RentPaymentHandler) GetRentPaymentSuggestion(ctx *gin.Context) {
156: if err != nil {
163: if err := ctx.ShouldBindQuery(&req); err != nil {
169: if err != nil {
175: if err != nil {
186: if err != nil {
195: func (h *RentPaymentHandler) GetRentalContractSummary(ctx *gin.Context) {
197: if err != nil {
206: if err != nil {
```

### rent-manager-backend/internal/adapter/store/mysql/repository/rent_payment.go

```text
18: func (rpr *RentPaymentRepository) CreateRentPayment(
56: if err != nil {
62: if err != nil {
65: if errCode := rpr.db.ErrorCode(err); errCode == "1062" {
73: if err != nil {
81: func (rpr *RentPaymentRepository) UpdateRentPayment(
85: queryBuilder := rpr.db.QueryBuilder.Update("rent_payments")
87: if rentPayment.RentalContractID != 0 {
91: if !rentPayment.Period.IsZero() {
95: if !rentPayment.DueDate.IsZero() {
99: if rentPayment.PaymentDate != nil {
103: if rentPayment.BaseAmount != 0 {
127: if rentPayment.TotalAmount != 0 {
135: queryBuilder = queryBuilder.Where(sq.Eq{"rpa_id": rentPayment.ID})
138: if err != nil {
144: if err != nil {
147: if errCode := rpr.db.ErrorCode(err); errCode == "1062" {
155: if err != nil {
160: if rowsAffected == 0 {
167: func (rpr *RentPaymentRepository) GetRentPaymentByID(
179: if err != nil {
202: if err != nil {
203: if errors.Is(err, sql.ErrNoRows) {
211: if paymentDate.Valid {
218: func (rpr *RentPaymentRepository) GetRentPaymentByContractIDAndPeriod(
232: if err != nil {
255: if err != nil {
256: if errors.Is(err, sql.ErrNoRows) {
264: if paymentDate.Valid {
270: func (rpr *RentPaymentRepository) ListRentPayments(
276: if page == 0 {
280: if limit == 0 {
291: if rentalContractID > 0 {
292: query = query.Where(sq.Eq{"rpa.rco_id": rentalContractID})
298: func NewRentPaymentRepository(db *mysql.DB) *RentPaymentRepository {
304: func (rpr *RentPaymentRepository) ListRentPaymentsByContractID(
315: func (rpr *RentPaymentRepository) baseRentPaymentQuery() sq.SelectBuilder {
338: func (rpr *RentPaymentRepository) findRentPayments(
343: if err != nil {
349: if err != nil {
379: if err != nil {
384: if paymentDate.Valid {
391: if err := rows.Err(); err != nil {
```

### rent-manager-backend/internal/core/policy/rent_payment/rent_payment_create_policy.go

```text
17: func NewRentPaymentCreatePolicy(
27: func (policy *RentPaymentCreatePolicy) Execute(ctx context.Context, rentPayment *domain.RentPayment) error {
28: if err := policy.validate(rentPayment); err != nil {
37: if err != nil {
40: if errors.As(err, &appErr) && appErr.Code == domain.ErrCodeDataNotFound {
47: if existingPayment != nil {
48: return domain.ErrRentPaymentAlreadyExists
54: func (policy *RentPaymentCreatePolicy) validate(rentPayment *domain.RentPayment) error {
55: if rentPayment == nil {
56: return domain.ErrRentPaymentNil
61: if rentPayment.RentalContractID <= 0 {
62: return domain.ErrRentPaymentContractIDEmpty
65: if rentPayment.Period.IsZero() {
66: return domain.ErrRentPaymentPeriodEmpty
69: if rentPayment.DueDate.IsZero() {
70: return domain.ErrRentPaymentDueDateEmpty
73: if rentPayment.BaseAmount <= 0 {
74: return domain.ErrRentPaymentBaseAmountInvalid
77: if rentPayment.TotalAmount <= 0 {
78: return domain.ErrRentPaymentTotalAmountInvalid
81: if rentPayment.PaidAmount < 0 {
82: return domain.ErrRentPaymentPaidAmountInvalid
```

### rent-manager-backend/internal/core/policy/rent_payment/rent_payment_update_policy.go

```text
18: func NewRentPaymentUpdatePolicy(
28: func (policy *RentPaymentUpdatePolicy) Execute(ctx context.Context, rentPayment *domain.RentPayment) error {
29: if rentPayment == nil {
30: return domain.ErrRentPaymentNil
33: if rentPayment.ID <= 0 {
34: return domain.ErrRentPaymentNotFound
37: if err := policy.validate(rentPayment); err != nil {
45: if err != nil {
53: if !shared.IsPeriodWithinContract(rentPayment.Period, rentalContract) {
54: return domain.ErrRentPaymentPeriodOutsideContract
62: if err != nil {
65: if errors.As(err, &appErr) && appErr.Code == domain.ErrCodeDataNotFound {
72: if existingPayment != nil && existingPayment.ID != rentPayment.ID {
73: return domain.ErrRentPaymentAlreadyExists
79: func (policy *RentPaymentUpdatePolicy) validate(rentPayment *domain.RentPayment) error {
82: if rentPayment.RentalContractID <= 0 {
83: return domain.ErrRentPaymentContractIDEmpty
86: if rentPayment.Period.IsZero() {
87: return domain.ErrRentPaymentPeriodEmpty
90: if rentPayment.DueDate.IsZero() {
91: return domain.ErrRentPaymentDueDateEmpty
94: if rentPayment.BaseAmount <= 0 {
95: return domain.ErrRentPaymentBaseAmountInvalid
98: if rentPayment.TotalAmount <= 0 {
99: return domain.ErrRentPaymentTotalAmountInvalid
102: if rentPayment.PaidAmount < 0 {
103: return domain.ErrRentPaymentPaidAmountInvalid
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

### rent-manager-backend/internal/core/usecase/rent_payment/create_rent_payment_use_case.go

```text
12: type CreateRentPaymentUseCase struct {
17: func NewCreateRentPaymentUseCase(
20: ) *CreateRentPaymentUseCase {
27: func (uc *CreateRentPaymentUseCase) CreateRentPayment(
31: if err := uc.rentPaymentCreatePolicy.Execute(ctx, rentPayment); err != nil {
37: if err != nil {
```

### rent-manager-backend/internal/core/usecase/rent_payment/get_rent_payment_by_id_use_case.go

```text
11: type GetRentPaymentByIDUseCase struct {
15: func NewGetRentPaymentByIDUseCase(
17: ) *GetRentPaymentByIDUseCase {
23: func (uc *GetRentPaymentByIDUseCase) GetRentPaymentByID(
28: if err != nil {
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

### rent-manager-backend/internal/core/usecase/rent_payment/list_rent_payments_use_case.go

```text
11: type ListRentPaymentsUseCase struct {
15: func NewListRentPaymentsUseCase(
17: ) *ListRentPaymentsUseCase {
23: func (uc *ListRentPaymentsUseCase) ListRentPayments(
35: if err != nil {
```

### rent-manager-backend/internal/core/usecase/rent_payment/update_rent_payment_use_case.go

```text
12: type UpdateRentPaymentUseCase struct {
17: func NewUpdateRentPaymentUseCase(
20: ) *UpdateRentPaymentUseCase {
27: func (uc *UpdateRentPaymentUseCase) UpdateRentPayment(
31: if err := uc.rentPaymentUpdatePolicy.Execute(ctx, rentPayment); err != nil {
37: if err != nil {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/core/rent-payment/rent-payment-api.service.ts

```text
21: return this.apiUrl.build('/rent-payments');
25: return this.apiUrl.build('/rental-contracts');
39: return this.http
44: getById$(id: number): Observable<ApiResponse<RentPayment>> {
45: return this.http
50: getSchedule$(rentalContractId: number): Observable<ApiResponse<RentPaymentScheduleItem[]>> {
51: return this.http
65: return this.http
72: getSummary$(rentalContractId: number): Observable<ApiResponse<RentalContractSummary>> {
73: return this.http
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/core/rent-payment/rent-payment.service.ts

```text
29: list$(rentalContractId = 0, page = 1, limit = 10): Observable<RentPayment[]> {
34: catchError((e: HttpErrorResponse) => {
41: get$(id: number): Observable<RentPayment> {
46: catchError((e: HttpErrorResponse) => {
53: getSchedule$(rentalContractId: number): Observable<RentPaymentScheduleItem[]> {
58: catchError((e: HttpErrorResponse) => {
74: catchError((e: HttpErrorResponse) => {
81: getSummary$(rentalContractId: number): Observable<RentalContractSummary> {
86: catchError((e: HttpErrorResponse) => {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-payment/create/rent-payment-create.component.ts

```text
41: constructor() {
53: switchMap((date) => {
59: tap((response) => {
63: catchError(() => {
76: switchMap(() => from(this.router.navigate(['/app/rent-contracts/view', rentalContractId]))),
82: switchMap(() => from(this.router.navigate(['/app/rent-contracts/view', rentalContractId]))),
88: refreshSuggestion(paymentDate: string): void {
92: save(): void {
103: cancel(): void {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-payment/edit/rent-payment-edit.component.ts

```text
33: ngOnInit(): void {
44: save(): void {
54: cancel(): void {
55: void this.router.navigate(['/app/rental-contracts']);
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-payment/form/rent-payment-form.component.ts

```text
41: [RentPaymentFormField.RentalContractId]: new FormControl(0, Validators.required),
42: [RentPaymentFormField.Period]: new FormControl('', Validators.required),
43: [RentPaymentFormField.DueDate]: new FormControl('', Validators.required),
44: [RentPaymentFormField.PaymentDate]: new FormControl('', Validators.required),
45: [RentPaymentFormField.BaseAmount]: new FormControl(0, [Validators.required, Validators.min(1)]),
51: Validators.required,
52: Validators.min(1),
96: ngAfterViewInit(): void {
143: ngOnDestroy(): void {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-payment/rent-payment.routes.ts

```text
11: path: '',
16: path: 'create',
20: path: 'edit/:id',
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-payment/service/rent-payment-write-api.service.ts

```text
17: return this.apiUrl.build('/rent-payments');
20: create$(request: CreateRentPaymentRequest): Observable<RentPayment> {
21: return this.http
26: update$(id: number, request: UpdateRentPaymentRequest): Observable<RentPayment> {
27: return this.http
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-payment/service/rent-payment-write.service.ts

```text
23: save$(rentPayment: RentPaymentFormValue): Observable<RentPayment> {
34: update$(id: number, rentPayment: RentPaymentFormValue): Observable<RentPayment> {
```
