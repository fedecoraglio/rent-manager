# Source and operation coverage

Reviewed: 2026-11-13. Resolve paths from the repository root. Use this map to open the implementation; excerpts are navigation indexes, not replacements for complete source bodies. Line numbers refer to this review. Do not interpret internal operations as public endpoints.

## Contents

- File inventory
- Operations, data contracts, and rule locations

## File inventory

- `rent-manager-backend/internal/adapter/http/inflation_index_request.go`
- `rent-manager-backend/internal/adapter/http/inflation_index_response.go`
- `rent-manager-backend/internal/adapter/http/inflation_index.go`
- `rent-manager-backend/internal/adapter/store/mysql/repository/inflation_index.go`
- `rent-manager-backend/internal/core/domain/inflation_indexes.go`
- `rent-manager-backend/internal/core/policy/inflation_index/inflation_index_create_policy.go`
- `rent-manager-backend/internal/core/policy/inflation_index/inflation_index_update_policy.go`
- `rent-manager-backend/internal/core/port/inflation_index.go`
- `rent-manager-backend/internal/core/service/inflation_adjustment_calculator.go`
- `rent-manager-backend/internal/core/usecase/inflation_index/create_inflation_index_use_case.go`
- `rent-manager-backend/internal/core/usecase/inflation_index/get_inflation_index_by_id_use_case.go`
- `rent-manager-backend/internal/core/usecase/inflation_index/get_inflation_index_by_period_use_case.go`
- `rent-manager-backend/internal/core/usecase/inflation_index/list_inflation_indexes_by_period_range_use_case.go`
- `rent-manager-backend/internal/core/usecase/inflation_index/list_inflation_indexes_use_case.go`
- `rent-manager-backend/internal/core/usecase/inflation_index/update_inflation_index_use_case.go`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/inflation-index/inflation-index-read-api.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/inflation-index/inflation-index-read.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/inflation-index/inflation-index.model.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/inflation-index/create/inflation-index-create.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/inflation-index/create/inflation-index-create.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/inflation-index/create/inflation-index-create.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/inflation-index/edit/inflation-index-edit.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/inflation-index/edit/inflation-index-edit.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/inflation-index/edit/inflation-index-edit.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/inflation-index/form/field/inflation-index-form-field.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/inflation-index/form/inflation-index-form.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/inflation-index/form/inflation-index-form.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/inflation-index/form/inflation-index-form.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/inflation-index/inflation-index.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/inflation-index/inflation-index.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/inflation-index/inflation-index.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/inflation-index/inflation-index.routes.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/inflation-index/list/inflation-index-list.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/inflation-index/list/inflation-index-list.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/inflation-index/list/inflation-index-list.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/inflation-index/list/table/inflation-index-list-table.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/inflation-index/list/table/inflation-index-list-table.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/inflation-index/list/table/inflation-index-list-table.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/inflation-index/model/inflation-index.model.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/inflation-index/service/inflation-index-write-api.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/inflation-index/service/inflation-index-write.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/inflation-index/view/inflation-index-view.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/inflation-index/view/inflation-index-view.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/inflation-index/view/inflation-index-view.component.ts`

## Operations, data contracts, and rule locations

### rent-manager-backend/internal/adapter/http/inflation_index_request.go

```text
4: Period     string  `json:"period" binding:"required"` // yyyy-mm
5: Percentage float64 `json:"percentage" binding:"required"`
6: Source     string  `json:"source"`
7: Notes      string  `json:"notes"`
10: Period     string  `json:"period" binding:"required"` // yyyy-mm
11: Percentage float64 `json:"percentage" binding:"required"`
12: Source     string  `json:"source"`
13: Notes      string  `json:"notes"`
```

### rent-manager-backend/internal/adapter/http/inflation_index_response.go

```text
9: ID         int64   `json:"id"`
10: Period     string  `json:"period"`
11: Percentage float64 `json:"percentage"`
12: Source     string  `json:"source"`
13: Notes      string  `json:"notes"`
14: CreatedAt  string  `json:"created_at"`
15: UpdatedAt  string  `json:"updated_at"`
18: func newInflationIndexResponse(inflationIndex *domain.InflationIndex) InflationIndexResponse {
30: func newInflationIndexesResponse(inflationIndexes []domain.InflationIndex) []InflationIndexResponse {
```

### rent-manager-backend/internal/adapter/http/inflation_index.go

```text
16: createInflationIndexUseCase  *inflation_index.CreateInflationIndexUseCase
17: updateInflationIndexUseCase  *inflation_index.UpdateInflationIndexUseCase
18: getInflationIndexByIDUseCase *inflation_index.GetInflationIndexByIDUseCase
19: listInflationIndexesUseCase  *inflation_index.ListInflationIndexesUseCase
22: func NewInflationIndexHandler(
23: createInflationIndexUseCase *inflation_index.CreateInflationIndexUseCase,
24: updateInflationIndexUseCase *inflation_index.UpdateInflationIndexUseCase,
25: getInflationIndexByIDUseCase *inflation_index.GetInflationIndexByIDUseCase,
26: listInflationIndexesUseCase *inflation_index.ListInflationIndexesUseCase,
36: func (h *InflationIndexHandler) Create(ctx *gin.Context) {
39: if err := ctx.ShouldBindJSON(&request); err != nil {
46: if err != nil {
60: if err != nil {
70: func (h *InflationIndexHandler) Update(ctx *gin.Context) {
72: if err != nil || id <= 0 {
80: if err := ctx.ShouldBindJSON(&request); err != nil {
87: if err != nil {
102: if err != nil {
111: func (h *InflationIndexHandler) GetByID(ctx *gin.Context) {
113: if err != nil || id <= 0 {
120: if err != nil {
129: func (h *InflationIndexHandler) List(ctx *gin.Context) {
133: if value := ctx.Query("page"); value != "" {
135: if err == nil && parsed > 0 {
140: if value := ctx.Query("limit"); value != "" {
142: if err == nil && parsed > 0 {
152: if err != nil {
```

### rent-manager-backend/internal/adapter/store/mysql/repository/inflation_index.go

```text
20: func NewInflationIndexRepository(db *mysql.DB) *InflationIndexRepository {
26: func (r *InflationIndexRepository) Create(
46: if err != nil {
52: if err != nil {
55: if errCode := r.db.ErrorCode(err); errCode == "1062" {
56: return domain.ErrInflationIndexAlreadyExists
63: if err != nil {
69: if err != nil {
78: func (r *InflationIndexRepository) Update(
91: if err != nil {
97: if err != nil {
100: if errCode := r.db.ErrorCode(err); errCode == "1062" {
101: return domain.ErrInflationIndexAlreadyExists
108: if err != nil {
113: if rowsAffected == 0 {
114: return domain.ErrInflationIndexNotFound
118: if err != nil {
127: func (r *InflationIndexRepository) GetByID(
138: func (r *InflationIndexRepository) GetByPeriod(
149: func (r *InflationIndexRepository) List(
154: if page == 0 {
158: if limit == 0 {
172: func (r *InflationIndexRepository) ListByPeriodRange(
185: func (r *InflationIndexRepository) ListLatestBeforePeriod(
196: if err != nil {
208: func (r *InflationIndexRepository) baseInflationIndexQuery() sq.SelectBuilder {
222: func (r *InflationIndexRepository) findOne(
227: if err != nil {
231: if len(inflationIndexes) == 0 {
238: func (r *InflationIndexRepository) findMany(
243: if err != nil {
249: if err != nil {
252: if errors.Is(err, sql.ErrNoRows) {
265: if err := rows.Scan(
281: if err := rows.Err(); err != nil {
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

### rent-manager-frontend/projects/rent-manager-frontend/src/app/core/inflation-index/inflation-index-read-api.service.ts

```text
15: return this.apiUrl.build('/inflation-indexes');
18: list$(page: number, limit: number): Observable<InflationIndex[]> {
21: return this.http
26: getById$(id: number): Observable<InflationIndex> {
27: return this.http
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/core/inflation-index/inflation-index-read.service.ts

```text
18: get$(id: number): Observable<InflationIndex> {
28: list$(page = 1, limit = 20): Observable<InflationIndex[]> {
33: tap((items) => {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/inflation-index/create/inflation-index-create.component.ts

```text
32: constructor() {
35: switchMap(() => from(this.router.navigate(['/app/inflation-indexes']))),
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/inflation-index/edit/inflation-index-edit.component.ts

```text
39: constructor() {
44: switchMap(() => from(this.router.navigate(['/app/inflation-indexes']))),
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/inflation-index/form/inflation-index-form.component.ts

```text
27: [InflationIndexFormField.Period]: new FormControl('', Validators.required),
29: Validators.required,
30: Validators.min(0),
71: ngAfterViewInit() {
90: ngOnDestroy() {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/inflation-index/inflation-index.routes.ts

```text
5: path: '',
10: path: 'create',
17: path: 'edit/:id',
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/inflation-index/list/inflation-index-list.component.ts

```text
37: constructor() {
42: switchMap(() => from(this.router.navigate(['/app/inflation-indexes/create']))),
48: view(id: number): void {
49: void this.router.navigate(['/app/inflation-indexes/view', id]);
52: edit(id: number): void {
53: void this.router.navigate(['/app/inflation-indexes/edit', id]);
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/inflation-index/service/inflation-index-write-api.service.ts

```text
19: return this.apiUrl.build('/inflation-indexes');
22: create(request: InflationIndexCreateRequest): Observable<InflationIndex> {
23: return this.http
28: update(id: number, request: InflationIndexUpdateRequest): Observable<InflationIndex> {
29: return this.http
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/inflation-index/service/inflation-index-write.service.ts

```text
18: save$(request: InflationIndexCreateRequest): Observable<InflationIndex> {
27: update$(id: number, request: InflationIndexUpdateRequest): Observable<InflationIndex> {
```
