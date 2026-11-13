# Source and operation coverage

Reviewed: 2026-11-13. Resolve paths from the repository root. Use this map to open the implementation; excerpts are navigation indexes, not replacements for complete source bodies. Line numbers refer to this review. Do not interpret internal operations as public endpoints.

## Contents

- File inventory
- Operations, data contracts, and rule locations

## File inventory

- `rent-manager-backend/internal/adapter/http/tenant_request.go`
- `rent-manager-backend/internal/adapter/http/tenant_response.go`
- `rent-manager-backend/internal/adapter/http/tenant.go`
- `rent-manager-backend/internal/adapter/store/mysql/repository/tenant.go`
- `rent-manager-backend/internal/core/domain/tenant.go`
- `rent-manager-backend/internal/core/policy/tenant/tenant_create_policy.go`
- `rent-manager-backend/internal/core/policy/tenant/tenant_update_policy.go`
- `rent-manager-backend/internal/core/port/tenant.go`
- `rent-manager-backend/internal/core/usecase/tenant/create_tenant_use_case.go`
- `rent-manager-backend/internal/core/usecase/tenant/get_tenant_by_document_number_use_case.go`
- `rent-manager-backend/internal/core/usecase/tenant/get_tenant_by_id_use_case.go`
- `rent-manager-backend/internal/core/usecase/tenant/list_tenant_use_case.go`
- `rent-manager-backend/internal/core/usecase/tenant/search_tenant_use_case.go`
- `rent-manager-backend/internal/core/usecase/tenant/update_tenant_use_case.go`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/tenant/tenant-core.model.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/tenant/tenant-read-api.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/tenant/tenant-read.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/tenant/create/tenant-create.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/tenant/create/tenant-create.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/tenant/create/tenant-create.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/tenant/edit/tenant-edit.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/tenant/edit/tenant-edit.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/tenant/edit/tenant-edit.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/tenant/form/field/tenant-form-field.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/tenant/form/tenant-form.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/tenant/form/tenant-form.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/tenant/form/tenant-form.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/tenant/list/table/tenant-list-table.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/tenant/list/table/tenant-list-table.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/tenant/list/table/tenant-list-table.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/tenant/list/tenant-list.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/tenant/list/tenant-list.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/tenant/list/tenant-list.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/tenant/model/tenant.model.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/tenant/service/tenant-write-api.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/tenant/service/tenant-write.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/tenant/tenant.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/tenant/tenant.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/tenant/tenant.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/tenant/tenant.routes.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/tenant/view/tenant-view.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/tenant/view/tenant-view.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/tenant/view/tenant-view.component.ts`

## Operations, data contracts, and rule locations

### rent-manager-backend/internal/adapter/http/tenant_request.go

```text
6: CountryID      int64  `json:"country_id"`
7: StateID        int64  `json:"state_id"`
8: Name           string `json:"name" binding:"required"`
9: Email          string `json:"email" binding:"omitempty,email"`
10: Phone          string `json:"phone"`
11: DocumentNumber string `json:"document_number" binding:"required"`
13: City         string `json:"city"`
14: Street       string `json:"street"`
15: StreetNumber string `json:"street_number"`
16: Floor        string `json:"floor"`
17: Apartment    string `json:"apartment"`
18: PostalCode   string `json:"postal_code"`
22: CountryID      int64  `json:"country_id"`
23: StateID        int64  `json:"state_id"`
24: Name           string `json:"name" binding:"required"`
25: Email          string `json:"email" binding:"omitempty,email"`
26: Phone          string `json:"phone"`
27: DocumentNumber string `json:"document_number" binding:"required"`
29: City         string `json:"city"`
30: Street       string `json:"street"`
31: StreetNumber string `json:"street_number"`
32: Floor        string `json:"floor"`
33: Apartment    string `json:"apartment"`
34: PostalCode   string `json:"postal_code"`
38: Page  uint64 `form:"page" binding:"required,min=1"`
39: Limit uint64 `form:"limit" binding:"required,min=1,max=100"`
43: Value string `form:"value" binding:"required"`
44: Page  uint64 `form:"page" binding:"required,min=1"`
45: Limit uint64 `form:"limit" binding:"required,min=1,max=100"`
48: func newTenantFromCreateRequest(req createTenantRequest) *domain.Tenant {
66: func newTenantFromUpdateRequest(id int64, req updateTenantRequest) *domain.Tenant {
```

### rent-manager-backend/internal/adapter/http/tenant_response.go

```text
10: ID             int64  `json:"id"`
11: CountryID      int64  `json:"country_id"`
12: StateID        int64  `json:"state_id"`
13: Name           string `json:"name"`
14: Email          string `json:"email"`
15: Phone          string `json:"phone"`
16: DocumentNumber string `json:"document_number"`
18: City         string `json:"city"`
19: Street       string `json:"street"`
20: StreetNumber string `json:"street_number"`
21: Floor        string `json:"floor"`
22: Apartment    string `json:"apartment"`
23: PostalCode   string `json:"postal_code"`
25: CreatedAt time.Time `json:"created_at"`
26: UpdatedAt time.Time `json:"updated_at"`
29: func newTenantResponse(tenant *domain.Tenant) tenantResponse {
49: func newTenantsResponse(tenants []domain.Tenant) []tenantResponse {
```

### rent-manager-backend/internal/adapter/http/tenant.go

```text
10: tenantUseCase "rent-manager-backend/internal/core/usecase/tenant"
14: createTenantUseCase              *tenantUseCase.CreateTenantUseCase
15: getTenantByIDUseCase             *tenantUseCase.GetTenantByIDUseCase
16: getTenantByDocumentNumberUseCase *tenantUseCase.GetTenantByDocumentNumberUseCase
17: listTenantUseCase                *tenantUseCase.ListTenantUseCase
18: searchTenantUseCase              *tenantUseCase.SearchTenantsUseCase
19: updateTenantUseCase              *tenantUseCase.UpdateTenantUseCase
22: func NewTenantHandler(
23: createTenantUseCase *tenantUseCase.CreateTenantUseCase,
24: getTenantByIDUseCase *tenantUseCase.GetTenantByIDUseCase,
25: getTenantByDocumentNumberUseCase *tenantUseCase.GetTenantByDocumentNumberUseCase,
26: listTenantUseCase *tenantUseCase.ListTenantUseCase,
27: searchTenantUseCase *tenantUseCase.SearchTenantsUseCase,
28: updateTenantUseCase *tenantUseCase.UpdateTenantUseCase,
40: func (h *TenantHandler) CreateTenant(ctx *gin.Context) {
43: if err := ctx.ShouldBindJSON(&req); err != nil {
49: if err != nil {
58: func (h *TenantHandler) GetTenantByID(ctx *gin.Context) {
60: if err != nil {
66: if err != nil {
75: func (h *TenantHandler) ListTenants(ctx *gin.Context) {
78: if err := ctx.ShouldBindQuery(&req); err != nil {
84: if err != nil {
93: func (h *TenantHandler) SearchTenants(ctx *gin.Context) {
96: if err := ctx.ShouldBindQuery(&req); err != nil {
102: if err != nil {
111: func (h *TenantHandler) UpdateTenant(ctx *gin.Context) {
113: if err != nil {
120: if err := ctx.ShouldBindJSON(&req); err != nil {
126: if err != nil {
```

### rent-manager-backend/internal/adapter/store/mysql/repository/tenant.go

```text
20: func NewTenantRepository(db *mysql.DB) *TenantRepository {
26: func (tr *TenantRepository) CreateTenant(ctx context.Context, tenant *domain.Tenant) (*domain.Tenant, error) {
59: if err != nil {
65: if err != nil {
68: if errCode := tr.db.ErrorCode(err); errCode == "1062" {
76: if err != nil {
84: func (tr *TenantRepository) GetTenantByID(ctx context.Context, id int64) (*domain.Tenant, error) {
92: if err != nil {
114: if err != nil {
115: if errors.Is(err, sql.ErrNoRows) {
126: func (tr *TenantRepository) GetTenantByDocumentNumber(ctx context.Context, docNumber string) (*domain.Tenant, error) {
134: if err != nil {
156: if err != nil {
157: if errors.Is(err, sql.ErrNoRows) {
168: func (tr *TenantRepository) ListTenants(ctx context.Context, page uint64, limit uint64) ([]domain.Tenant, error) {
169: if page == 0 {
173: if limit == 0 {
187: func (tr *TenantRepository) SearchTenants(
193: if page == 0 {
197: if limit == 0 {
221: func (tr *TenantRepository) UpdateTenant(
225: queryBuilder := tr.db.QueryBuilder.Update("tenants")
227: if tenant.CountryID != 0 {
231: if tenant.StateID != 0 {
235: if tenant.Name != "" {
239: if tenant.Email != "" {
243: if tenant.Phone != "" {
247: if tenant.DocumentNumber != "" {
251: if tenant.City != "" {
255: if tenant.Street != "" {
259: if tenant.StreetNumber != "" {
263: if tenant.Floor != "" {
267: if tenant.Apartment != "" {
271: if tenant.PostalCode != "" {
275: queryBuilder = queryBuilder.Where(sq.Eq{"ten_id": tenant.ID})
278: if err != nil {
284: if err != nil {
287: if errCode := tr.db.ErrorCode(err); errCode == "1062" {
295: if err != nil {
300: if rowsAffected == 0 {
307: func (tr *TenantRepository) findTenants(ctx context.Context, query sq.SelectBuilder) ([]domain.Tenant, error) {
309: if err != nil {
315: if err != nil {
343: if err != nil {
351: if err := rows.Err(); err != nil {
359: func (tr *TenantRepository) baseTenantQuery() sq.SelectBuilder {
```

### rent-manager-backend/internal/core/policy/tenant/tenant_create_policy.go

```text
17: func NewTenantCreatePolicy(tenantRepository port.TenantRepository) *TenantCreatePolicy {
23: func (tenantPolicy *TenantCreatePolicy) Execute(ctx context.Context, tenant *domain.Tenant) error {
24: if tenant == nil {
25: return domain.ErrTenantNil
39: if tenant.Name == "" {
40: return domain.ErrTenantNameEmpty
43: if tenant.DocumentNumber == "" {
44: return domain.ErrTenantDocumentNumberEmpty
47: if tenant.Email != "" && !shared.IsValidEmail(tenant.Email) {
48: return domain.ErrTenantInvalidEmail
52: if err != nil {
54: if errors.As(err, &appErr) && appErr.Code == domain.ErrCodeDataNotFound {
65: if existingTenant != nil {
66: return domain.ErrTenantAlreadyExists
```

### rent-manager-backend/internal/core/policy/tenant/tenant_update_policy.go

```text
17: func NewTenantUpdatePolicy(tenantRepository port.TenantRepository) *TenantUpdatePolicy {
23: func (tenantPolicy *TenantUpdatePolicy) Execute(ctx context.Context, tenant *domain.Tenant) error {
24: if tenant == nil {
25: return domain.ErrTenantNil
28: if tenant.ID <= 0 {
29: return domain.ErrTenantNotFound
43: if tenant.Name == "" {
44: return domain.ErrTenantNameEmpty
47: if tenant.DocumentNumber == "" {
48: return domain.ErrTenantDocumentNumberEmpty
51: if tenant.Email != "" && !shared.IsValidEmail(tenant.Email) {
52: return domain.ErrTenantInvalidEmail
59: if err != nil {
62: if errors.As(err, &appErr) && appErr.Code == domain.ErrCodeDataNotFound {
73: if existingTenant != nil && existingTenant.ID != tenant.ID {
74: return domain.ErrTenantAlreadyExists
```

### rent-manager-backend/internal/core/usecase/tenant/create_tenant_use_case.go

```text
12: type CreateTenantUseCase struct {
17: func NewCreateTenantUseCase(
20: ) *CreateTenantUseCase {
27: func (uc *CreateTenantUseCase) CreateTenant(
31: if err := uc.tenantCreatePolicy.Execute(ctx, tenant); err != nil {
37: if err != nil {
```

### rent-manager-backend/internal/core/usecase/tenant/get_tenant_by_document_number_use_case.go

```text
11: type GetTenantByDocumentNumberUseCase struct {
15: func NewGetTenantByDocumentNumberUseCase(tenantRepository port.TenantRepository) *GetTenantByDocumentNumberUseCase {
21: func (uc *GetTenantByDocumentNumberUseCase) GetTenantByDocumentNumber(
26: if err != nil {
```

### rent-manager-backend/internal/core/usecase/tenant/get_tenant_by_id_use_case.go

```text
11: type GetTenantByIDUseCase struct {
15: func NewGetTenantByIDUseCase(tenantRepository port.TenantRepository) *GetTenantByIDUseCase {
21: func (uc *GetTenantByIDUseCase) GetTenantByID(ctx context.Context, id int64) (*domain.Tenant, error) {
23: if err != nil {
```

### rent-manager-backend/internal/core/usecase/tenant/list_tenant_use_case.go

```text
11: type ListTenantUseCase struct {
15: func NewListTenantsUseCase(tenantRepository port.TenantRepository) *ListTenantUseCase {
21: func (uc *ListTenantUseCase) ListTenants(
27: if err != nil {
```

### rent-manager-backend/internal/core/usecase/tenant/search_tenant_use_case.go

```text
11: type SearchTenantsUseCase struct {
15: func NewSearchTenantsUseCase(tenantRepository port.TenantRepository) *SearchTenantsUseCase {
21: func (uc *SearchTenantsUseCase) SearchTenants(
28: if err != nil {
```

### rent-manager-backend/internal/core/usecase/tenant/update_tenant_use_case.go

```text
12: type UpdateTenantUseCase struct {
17: func NewUpdateTenantUseCase(
20: ) *UpdateTenantUseCase {
27: func (uc *UpdateTenantUseCase) UpdateTenant(
31: if err := uc.tenantUpdatePolicy.Execute(ctx, tenant); err != nil {
37: if err != nil {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/core/tenant/tenant-read-api.service.ts

```text
15: return this.apiUrl.build('/tenants');
18: list$(page: number, limit: number): Observable<Tenant[]> {
21: return this.http
26: search$(value: string, page: number, limit: number): Observable<Tenant[]> {
29: return this.http
34: getById$(id: number): Observable<Tenant> {
35: return this.http
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/core/tenant/tenant-read.service.ts

```text
21: get$(id: number): Observable<Tenant> {
31: list$(page = 1, limit = 10): Observable<Tenant[]> {
41: search$(value: string, page = 1, limit = 10): Observable<Tenant[]> {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/tenant/create/tenant-create.component.ts

```text
29: constructor() {
32: switchMap(() => from(this.router.navigate(['/app/tenants']))),
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/tenant/edit/tenant-edit.component.ts

```text
39: constructor() {
42: switchMap(() => from(this.router.navigate(['/app/tenants']))),
58: ngOnInit(): void {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/tenant/form/tenant-form.component.ts

```text
44: [TenantFormField.CountryId]: new FormControl(1, Validators.required),
45: [TenantFormField.StateId]: new FormControl(0, Validators.required),
46: [TenantFormField.Name]: new FormControl('', Validators.required),
47: [TenantFormField.Email]: new FormControl('', [Validators.email]),
48: [TenantFormField.DocumentNumber]: new FormControl('', Validators.required),
97: ngAfterViewInit(): void {
122: ngOnDestroy(): void {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/tenant/list/tenant-list.component.ts

```text
52: constructor() {
55: switchMap(() => {
86: ngOnInit(): void {
90: onPageChange(event: PageEvent): void {
96: viewTenant(id: number): void {
97: void this.router.navigate(['/app/tenants/view', id]);
100: editTenant(id: number): void {
101: void this.router.navigate(['/app/tenants/edit', id]);
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/tenant/service/tenant-write-api.service.ts

```text
16: return this.apiUrl.build('/tenants');
19: create$(request: CreateTenantRequest): Observable<Tenant> {
20: return this.http
25: update$(id: number, request: UpdateTenantRequest): Observable<Tenant> {
26: return this.http
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/tenant/service/tenant-write.service.ts

```text
21: save$(tenant: TenantFormValue): Observable<Tenant> {
30: update$(id: number, tenant: TenantFormValue): Observable<Tenant> {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/tenant/tenant.routes.ts

```text
13: path: '',
18: path: '',
22: path: 'create',
26: path: 'edit/:id',
30: path: 'view/:id',
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/tenant/view/tenant-view.component.ts

```text
33: constructor() {
36: switchMap(() => from(this.router.navigate(['/app/tenants']))),
44: switchMap(() => from(this.router.navigate(['/app/tenants/edit', this.tenantId]))),
50: ngOnInit(): void {
```
