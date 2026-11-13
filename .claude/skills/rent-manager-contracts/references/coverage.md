# Source and operation coverage

Reviewed: 2026-11-13. Resolve paths from the repository root. Use this map to open the implementation; excerpts are navigation indexes, not replacements for complete source bodies. Line numbers refer to this review. Do not interpret internal operations as public endpoints.

## Contents

- File inventory
- Operations, data contracts, and rule locations

## File inventory

- `rent-manager-backend/internal/adapter/http/rent_contract_request.go`
- `rent-manager-backend/internal/adapter/http/rent_contract_response.go`
- `rent-manager-backend/internal/adapter/http/rent_contract.go`
- `rent-manager-backend/internal/adapter/http/rental_contract_summary_response.go`
- `rent-manager-backend/internal/adapter/store/mysql/repository/rent_contract.go`
- `rent-manager-backend/internal/core/domain/rent_contract.go`
- `rent-manager-backend/internal/core/domain/rental_contract_summary.go`
- `rent-manager-backend/internal/core/policy/rent_contract/rent_contract_create_policy.go`
- `rent-manager-backend/internal/core/policy/rent_contract/rent_contract_update_policy.go`
- `rent-manager-backend/internal/core/port/rent_contract.go`
- `rent-manager-backend/internal/core/usecase/contract_catalog/rent_contract_catalog.go`
- `rent-manager-backend/internal/core/usecase/rental_contract/create_rental_contract_use_case.go`
- `rent-manager-backend/internal/core/usecase/rental_contract/get_rental_contract_by_id_use_case.go`
- `rent-manager-backend/internal/core/usecase/rental_contract/list_rental_contracts_use_case.go`
- `rent-manager-backend/internal/core/usecase/rental_contract/update_rental_contract_use_case.go`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/rent-contract/rent-contract-read-api.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/rent-contract/rent-contract-read.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/rent-contract/rent-contract.model.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-contract/create/rent-contract-create.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-contract/create/rent-contract-create.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-contract/create/rent-contract-create.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-contract/edit/rent-contract-edit.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-contract/edit/rent-contract-edit.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-contract/edit/rent-contract-edit.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-contract/form/field/rent-contract-form-field.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-contract/form/rent-contract-form.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-contract/form/rent-contract-form.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-contract/form/rent-contract-form.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-contract/list/rent-contract-list.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-contract/list/rent-contract-list.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-contract/list/rent-contract-list.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-contract/list/table/rent-contract-list-table.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-contract/list/table/rent-contract-list-table.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-contract/list/table/rent-contract-list-table.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-contract/model/rent-contract.model.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-contract/overview/rent-contract-overview.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-contract/overview/rent-contract-overview.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-contract/rent-contract.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-contract/rent-contract.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-contract/rent-contract.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-contract/rent-contract.routes.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-contract/service/rent-contract-write-api.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-contract/service/rent-contract-write.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-contract/view/rent-contract-view.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-contract/view/rent-contract-view.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-contract/view/rent-contract-view.component.ts`

## Operations, data contracts, and rule locations

### rent-manager-backend/internal/adapter/http/rent_contract_request.go

```text
10: PropertyID int64 `json:"property_id" binding:"required,min=1"`
11: TenantID   int64 `json:"tenant_id" binding:"required,min=1"`
12: StatusID   int64 `json:"status_id" binding:"required,min=1"`
14: InterestCalculationTypeID int64 `json:"interest_calculation_type_id" binding:"required,min=1"`
15: AdjustmentTypeID          int64 `json:"adjustment_type_id" binding:"required,min=1"`
17: StartDate string `json:"start_date" binding:"required"`
18: EndDate   string `json:"end_date" binding:"required"`
20: MonthlyAmount float64 `json:"monthly_amount" binding:"required,gt=0"`
21: DepositAmount float64 `json:"deposit_amount"`
22: Currency      string  `json:"currency" binding:"required"`
24: DueDay                    int64   `json:"due_day" binding:"required,min=1,max=31"`
25: DailyInterestPercentage   float64 `json:"daily_interest_percentage"`
26: AdjustmentFrequencyMonths int64   `json:"adjustment_frequency_months" binding:"required,min=1"`
28: Notes string `json:"notes"`
32: PropertyID int64 `json:"property_id" binding:"required,min=1"`
33: TenantID   int64 `json:"tenant_id" binding:"required,min=1"`
34: StatusID   int64 `json:"status_id" binding:"required,min=1"`
36: InterestCalculationTypeID int64 `json:"interest_calculation_type_id" binding:"required,min=1"`
37: AdjustmentTypeID          int64 `json:"adjustment_type_id" binding:"required,min=1"`
39: StartDate string `json:"start_date" binding:"required"`
40: EndDate   string `json:"end_date" binding:"required"`
42: MonthlyAmount float64 `json:"monthly_amount" binding:"required,gt=0"`
43: DepositAmount float64 `json:"deposit_amount"`
44: Currency      string  `json:"currency" binding:"required"`
46: DueDay                    int64   `json:"due_day" binding:"required,min=1,max=31"`
47: DailyInterestPercentage   float64 `json:"daily_interest_percentage"`
48: AdjustmentFrequencyMonths int64   `json:"adjustment_frequency_months" binding:"required,min=1"`
50: Notes string `json:"notes"`
54: PropertyID int64  `form:"property_id"`
55: Page       uint64 `form:"page" binding:"required,min=1"`
56: Limit      uint64 `form:"limit" binding:"required,min=1,max=100"`
59: func newRentalContractFromCreateRequest(
63: if err != nil {
68: if err != nil {
90: func newRentalContractFromUpdateRequest(
97: if err != nil {
106: func parseDate(value string) (time.Time, error) {
```

### rent-manager-backend/internal/adapter/http/rent_contract_response.go

```text
10: ID                        int64            `json:"id"`
11: PropertyID                int64            `json:"property_id"`
12: TenantID                  int64            `json:"tenant_id"`
13: StatusID                  int64            `json:"status_id"`
14: InterestCalculationTypeID int64            `json:"interest_calculation_type_id"`
15: AdjustmentTypeID          int64            `json:"adjustment_type_id"`
16: StartDate                 string           `json:"start_date"`
17: EndDate                   string           `json:"end_date"`
18: TotalPayments             int64            `json:"total_payments" binding:"required,min=0"`
19: MonthlyAmount             float64          `json:"monthly_amount"`
20: DepositAmount             float64          `json:"deposit_amount"`
21: Currency                  string           `json:"currency"`
22: DueDay                    int64            `json:"due_day"`
23: DailyInterestPercentage   float64          `json:"daily_interest_percentage"`
24: AdjustmentFrequencyMonths int64            `json:"adjustment_frequency_months"`
25: Property                  propertyResponse `json:"property"`
26: Tenant                    tenantResponse   `json:"tenant"`
27: Notes                     string           `json:"notes"`
28: CreatedAt                 time.Time        `json:"created_at"`
29: UpdatedAt                 time.Time        `json:"updated_at"`
32: func newRentalContractResponse(
36: if rentalContract.Property != nil {
41: if rentalContract.Tenant != nil {
68: func newRentalContractsResponse(
81: func formatDate(value time.Time) string {
82: if value.IsZero() {
```

### rent-manager-backend/internal/adapter/http/rent_contract.go

```text
10: rentalContractUseCase "rent-manager-backend/internal/core/usecase/rental_contract"
14: createRentalContractUseCase  *rentalContractUseCase.CreateRentalContractUseCase
15: updateRentalContractUseCase  *rentalContractUseCase.UpdateRentalContractUseCase
16: getRentalContractByIDUseCase *rentalContractUseCase.GetRentalContractByIDUseCase
17: listRentalContractsUseCase   *rentalContractUseCase.ListRentalContractsUseCase
20: func NewRentalContractHandler(
21: createRentalContractUseCase *rentalContractUseCase.CreateRentalContractUseCase,
22: updateRentalContractUseCase *rentalContractUseCase.UpdateRentalContractUseCase,
23: getRentalContractByIDUseCase *rentalContractUseCase.GetRentalContractByIDUseCase,
24: listRentalContractsUseCase *rentalContractUseCase.ListRentalContractsUseCase,
34: func (h *RentalContractHandler) CreateRentalContract(ctx *gin.Context) {
37: if err := ctx.ShouldBindJSON(&req); err != nil {
43: if err != nil {
49: if err != nil {
58: func (h *RentalContractHandler) UpdateRentalContract(ctx *gin.Context) {
60: if err != nil {
67: if err := ctx.ShouldBindJSON(&req); err != nil {
73: if err != nil {
79: if err != nil {
88: func (h *RentalContractHandler) GetRentalContractByID(ctx *gin.Context) {
90: if err != nil {
96: if err != nil {
105: func (h *RentalContractHandler) ListRentalContracts(ctx *gin.Context) {
108: if err := ctx.ShouldBindQuery(&req); err != nil {
119: if err != nil {
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

### rent-manager-backend/internal/adapter/store/mysql/repository/rent_contract.go

```text
17: func NewRentalContractRepository(db *mysql.DB) *RentalContractRepository {
21: func (rcr *RentalContractRepository) CreateRentalContract(
63: if err != nil {
69: if err != nil {
75: if err != nil {
83: func (rcr *RentalContractRepository) GetRentalContractByID(
94: if err != nil {
102: if err != nil {
103: if errors.Is(err, sql.ErrNoRows) {
114: func (rcr *RentalContractRepository) UpdateRentalContract(
118: queryBuilder := rcr.db.QueryBuilder.Update("rental_contracts")
120: if rentalContract.PropertyID != 0 {
124: if rentalContract.TenantID != 0 {
128: if rentalContract.StatusID != 0 {
132: if rentalContract.InterestCalculationTypeID != 0 {
136: if rentalContract.AdjustmentTypeID != 0 {
140: if !rentalContract.StartDate.IsZero() {
144: if !rentalContract.EndDate.IsZero() {
148: if rentalContract.MonthlyAmount != 0 {
152: if rentalContract.DepositAmount >= 0 {
156: if rentalContract.Currency != "" {
160: if rentalContract.DueDay != 0 {
164: if rentalContract.TotalPayments > 0 {
168: if rentalContract.DailyInterestPercentage >= 0 {
175: if rentalContract.AdjustmentFrequencyMonths != 0 {
183: queryBuilder = queryBuilder.Where(sq.Eq{"rco_id": rentalContract.ID})
186: if err != nil {
192: if err != nil {
198: if err != nil {
203: if rowsAffected == 0 {
210: func (rcr *RentalContractRepository) ListRentalContracts(
216: if page == 0 {
220: if limit == 0 {
231: if propertyID > 0 {
232: query = query.Where(sq.Eq{"rc.pro_id": propertyID})
238: func (rcr *RentalContractRepository) GetActiveRentalContractByPropertyID(
251: if err != nil {
259: if err != nil {
260: if errors.Is(err, sql.ErrNoRows) {
271: func (rcr *RentalContractRepository) ListActiveRentalContracts(
276: if page == 0 {
280: if limit == 0 {
292: if err != nil {
299: func (rcr *RentalContractRepository) findRentalContracts(
304: if err != nil {
310: if err != nil {
320: if err != nil {
328: if err := rows.Err(); err != nil {
336: func (rcr *RentalContractRepository) baseRentalContractQuery() sq.SelectBuilder {
383: func scanRentalContract(scanner rentalContractScanner) (domain.RentalContract, error) {
433: if err != nil {
437: if interestCalculationTypeID.Valid {
445: if adjustmentTypeID.Valid {
453: if statusID.Valid {
461: if propertyID.Valid {
465: if tenantID.Valid {
```

### rent-manager-backend/internal/core/policy/rent_contract/rent_contract_create_policy.go

```text
17: func NewRentalContractCreatePolicy(
25: func (policy *RentalContractCreatePolicy) Execute(
29: if err := policy.validate(rentalContract); err != nil {
37: if err != nil {
40: if errors.As(err, &appErr) && appErr.Code == domain.ErrCodeDataNotFound {
51: if activeContract != nil {
52: return domain.ErrRentalContractAlreadyExists
58: func (policy *RentalContractCreatePolicy) validate(
61: if rentalContract == nil {
62: return domain.ErrRentalContractNil
68: if rentalContract.PropertyID <= 0 {
69: return domain.ErrRentalContractPropertyIDEmpty
72: if rentalContract.TenantID <= 0 {
73: return domain.ErrRentalContractTenantIDEmpty
76: if rentalContract.StatusID <= 0 {
77: return domain.ErrRentalContractStatusIDEmpty
80: if rentalContract.InterestCalculationTypeID <= 0 {
81: return domain.ErrRentalContractInterestCalculationTypeIDEmpty
84: if rentalContract.AdjustmentTypeID <= 0 {
85: return domain.ErrRentalContractAdjustmentTypeIDEmpty
88: if rentalContract.StartDate.IsZero() {
89: return domain.ErrRentalContractStartDateEmpty
92: if rentalContract.EndDate.IsZero() {
93: return domain.ErrRentalContractEndDateEmpty
96: if !rentalContract.EndDate.After(rentalContract.StartDate) {
97: return domain.ErrRentalContractInvalidDateRange
100: if rentalContract.MonthlyAmount <= 0 {
101: return domain.ErrRentalContractMonthlyAmountInvalid
104: if rentalContract.DepositAmount < 0 {
105: return domain.ErrRentalContractDepositAmountInvalid
108: if rentalContract.Currency == "" {
109: return domain.ErrRentalContractCurrencyEmpty
112: if rentalContract.DueDay < 1 || rentalContract.DueDay > 31 {
113: return domain.ErrRentalContractDueDayInvalid
116: if rentalContract.DailyInterestPercentage < 0 {
117: return domain.ErrRentalContractDailyInterestPercentageInvalid
120: if rentalContract.AdjustmentFrequencyMonths <= 0 {
121: return domain.ErrRentalContractAdjustmentFrequencyInvalid
124: if rentalContract.TotalPayments <= 0 {
125: return domain.ErrRentalContractTotalPaymentsInvalid
131: func normalizeDate(date time.Time) time.Time {
```

### rent-manager-backend/internal/core/policy/rent_contract/rent_contract_update_policy.go

```text
16: func NewRentalContractUpdatePolicy(
24: func (policy *RentalContractUpdatePolicy) Execute(
28: if rentalContract == nil {
29: return domain.ErrRentalContractNil
32: if rentalContract.ID <= 0 {
33: return domain.ErrRentalContractNotFound
39: if rentalContract.PropertyID <= 0 {
40: return domain.ErrRentalContractPropertyIDEmpty
43: if rentalContract.TenantID <= 0 {
44: return domain.ErrRentalContractTenantIDEmpty
47: if rentalContract.StatusID <= 0 {
48: return domain.ErrRentalContractStatusIDEmpty
51: if rentalContract.InterestCalculationTypeID <= 0 {
52: return domain.ErrRentalContractInterestCalculationTypeIDEmpty
55: if rentalContract.AdjustmentTypeID <= 0 {
56: return domain.ErrRentalContractAdjustmentTypeIDEmpty
59: if rentalContract.StartDate.IsZero() {
60: return domain.ErrRentalContractStartDateEmpty
63: if rentalContract.EndDate.IsZero() {
64: return domain.ErrRentalContractEndDateEmpty
67: if !rentalContract.EndDate.After(rentalContract.StartDate) {
68: return domain.ErrRentalContractInvalidDateRange
71: if rentalContract.MonthlyAmount <= 0 {
72: return domain.ErrRentalContractMonthlyAmountInvalid
75: if rentalContract.DepositAmount < 0 {
76: return domain.ErrRentalContractDepositAmountInvalid
79: if rentalContract.Currency == "" {
80: return domain.ErrRentalContractCurrencyEmpty
83: if rentalContract.DueDay < 1 || rentalContract.DueDay > 31 {
84: return domain.ErrRentalContractDueDayInvalid
87: if rentalContract.DailyInterestPercentage < 0 {
88: return domain.ErrRentalContractDailyInterestPercentageInvalid
91: if rentalContract.AdjustmentFrequencyMonths <= 0 {
92: return domain.ErrRentalContractAdjustmentFrequencyInvalid
99: if err != nil {
102: if errors.As(err, &appErr) && appErr.Code == domain.ErrCodeDataNotFound {
113: if activeContract != nil && activeContract.ID != rentalContract.ID {
114: return domain.ErrRentalContractAlreadyExists
```

### rent-manager-backend/internal/core/usecase/rental_contract/create_rental_contract_use_case.go

```text
13: type CreateRentalContractUseCase struct {
18: func NewCreateRentalContractUseCase(
21: ) *CreateRentalContractUseCase {
28: func (uc *CreateRentalContractUseCase) CreateRentalContract(
34: if err := uc.rentalContractCreatePolicy.Execute(ctx, rentalContract); err != nil {
43: if err != nil {
```

### rent-manager-backend/internal/core/usecase/rental_contract/get_rental_contract_by_id_use_case.go

```text
11: type GetRentalContractByIDUseCase struct {
15: func NewGetRentalContractByIDUseCase(
17: ) *GetRentalContractByIDUseCase {
23: func (uc *GetRentalContractByIDUseCase) GetRentalContractByID(
28: if err != nil {
```

### rent-manager-backend/internal/core/usecase/rental_contract/list_rental_contracts_use_case.go

```text
11: type ListRentalContractsUseCase struct {
15: func NewListRentalContractsUseCase(
17: ) *ListRentalContractsUseCase {
23: func (uc *ListRentalContractsUseCase) ListRentalContracts(
35: if err != nil {
```

### rent-manager-backend/internal/core/usecase/rental_contract/update_rental_contract_use_case.go

```text
13: type UpdateRentalContractUseCase struct {
18: func NewUpdateRentalContractUseCase(
21: ) *UpdateRentalContractUseCase {
28: func (uc *UpdateRentalContractUseCase) UpdateRentalContract(
33: if err := uc.rentalContractUpdatePolicy.Execute(ctx, rentalContract); err != nil {
42: if err != nil {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/core/rent-contract/rent-contract-read-api.service.ts

```text
15: return this.apiUrl.build('/rental-contracts');
18: list$(propertyId: number, page: number, limit: number): Observable<ApiResponse<RentContract[]>> {
25: return this.http
30: getById$(id: number): Observable<ApiResponse<RentContract>> {
31: return this.http
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/core/rent-contract/rent-contract-read.service.ts

```text
18: list$(propertyId = 0, page = 1, limit = 10): Observable<RentContract[]> {
23: catchError((e: HttpErrorResponse) => {
30: get$(id: number): Observable<RentContract> {
35: catchError((e: HttpErrorResponse) => {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-contract/create/rent-contract-create.component.ts

```text
30: constructor() {
33: switchMap(() => from(this.router.navigate(['/app/rent-contracts']))),
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-contract/edit/rent-contract-edit.component.ts

```text
39: constructor() {
42: switchMap(() => from(this.router.navigate(['/app/rent-contracts']))),
60: ngOnInit(): void {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-contract/form/rent-contract-form.component.ts

```text
52: [RentContractFormField.PropertyId]: new FormControl(0, Validators.required),
53: [RentContractFormField.TenantId]: new FormControl(0, Validators.required),
54: [RentContractFormField.StatusId]: new FormControl(0, Validators.required),
55: [RentContractFormField.InterestCalculationTypeId]: new FormControl(0, Validators.required),
56: [RentContractFormField.AdjustmentTypeId]: new FormControl(0, Validators.required),
57: [RentContractFormField.StartDate]: new FormControl('', Validators.required),
58: [RentContractFormField.EndDate]: new FormControl('', Validators.required),
60: Validators.required,
61: Validators.min(1),
63: [RentContractFormField.DepositAmount]: new FormControl(0, Validators.min(0)),
64: [RentContractFormField.Currency]: new FormControl('ARS', Validators.required),
66: Validators.required,
67: Validators.min(1),
68: Validators.max(31),
70: [RentContractFormField.DailyInterestPercentage]: new FormControl(0, Validators.min(0)),
72: Validators.required,
73: Validators.min(1),
140: ngAfterViewInit(): void {
169: ngOnDestroy(): void {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-contract/list/rent-contract-list.component.ts

```text
50: constructor() {
74: ngOnInit(): void {
78: onPropertyChange(propertyId: number): void {
84: onPageChange(event: PageEvent): void {
90: viewRentContract(id: number): void {
91: void this.router.navigate(['/app/rent-contracts/view', id]);
94: editRentContract(id: number): void {
95: void this.router.navigate(['/app/rent-contracts/edit', id]);
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

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-contract/rent-contract.component.ts

```text
19: constructor() {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-contract/rent-contract.routes.ts

```text
13: path: '',
18: path: '',
22: path: 'create',
26: path: 'edit/:id',
30: path: 'view/:id',
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-contract/service/rent-contract-write-api.service.ts

```text
16: return this.apiUrl.build('/rental-contracts');
19: create$(request: CreateRentContractRequest): Observable<RentContract> {
20: return this.http
25: update$(id: number, request: UpdateRentContractRequest): Observable<RentContract> {
26: return this.http
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-contract/service/rent-contract-write.service.ts

```text
22: save$(rentContract: RentContractFormValue): Observable<RentContract> {
33: update$(id: number, rentContract: RentContractFormValue): Observable<RentContract> {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-contract/view/rent-contract-view.component.ts

```text
61: constructor() {
64: switchMap(() => from(this.router.navigate(['/app/rent-contracts']))),
73: from(this.router.navigate(['/app/rent-contracts/edit', this.rentContractId])),
80: ngOnDestroy(): void {
85: ngOnInit(): void {
96: tap((rentContract) => {
145: registerPayment(item: RentPaymentScheduleItem): void {
146: void this.router.navigate(['/app/rent-payments/create'], {
147: queryParams: {
154: editPayment(): void {
```
