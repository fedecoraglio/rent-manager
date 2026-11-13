# Source and operation coverage

Reviewed: 2026-11-13. Resolve paths from the repository root. Use this map to open the implementation; excerpts are navigation indexes, not replacements for complete source bodies. Line numbers refer to this review. Do not interpret internal operations as public endpoints.

## Contents

- File inventory
- Operations, data contracts, and rule locations

## File inventory

- `rent-manager-backend/internal/adapter/http/owner_request.go`
- `rent-manager-backend/internal/adapter/http/owner_response.go`
- `rent-manager-backend/internal/adapter/http/owner.go`
- `rent-manager-backend/internal/adapter/store/mysql/repository/owner.go`
- `rent-manager-backend/internal/core/domain/owner.go`
- `rent-manager-backend/internal/core/policy/owner/owner_create_policy.go`
- `rent-manager-backend/internal/core/policy/owner/owner_update_policy.go`
- `rent-manager-backend/internal/core/port/owner.go`
- `rent-manager-backend/internal/core/usecase/owner/create_owner_use_case.go`
- `rent-manager-backend/internal/core/usecase/owner/get_owner_by_document_number_use_case.go`
- `rent-manager-backend/internal/core/usecase/owner/get_owner_by_email_use_case.go`
- `rent-manager-backend/internal/core/usecase/owner/get_owner_by_id_use_case.go`
- `rent-manager-backend/internal/core/usecase/owner/list_owners_use_case.go`
- `rent-manager-backend/internal/core/usecase/owner/search_owners_use_case.go`
- `rent-manager-backend/internal/core/usecase/owner/update_owner_use_case.go`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/owner/owner-read-api.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/owner/owner-read.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/owner/owner.model.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/owner/create/owner-create.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/owner/create/owner-create.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/owner/create/owner-create.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/owner/edit/owner-edit.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/owner/edit/owner-edit.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/owner/edit/owner-edit.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/owner/form/field/owner-form-field.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/owner/form/owner-form.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/owner/form/owner-form.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/owner/form/owner-form.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/owner/list/owner-list.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/owner/list/owner-list.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/owner/list/owner-list.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/owner/list/table/owner-list-table.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/owner/list/table/owner-list-table.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/owner/list/table/owner-list-table.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/owner/model/owner.model.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/owner/owner.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/owner/owner.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/owner/owner.routes.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/owner/service/owner-write-api.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/owner/service/owner-write.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/owner/view/owner-view.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/owner/view/owner-view.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/owner/view/owner-view.component.ts`

## Operations, data contracts, and rule locations

### rent-manager-backend/internal/adapter/http/owner_request.go

```text
6: Name           string `json:"name" binding:"required"`
7: Email          string `json:"email" binding:"omitempty,email"`
8: Phone          string `json:"phone"`
9: DocumentNumber string `json:"document_number" binding:"required"`
13: Name           string `json:"name" binding:"required"`
14: Email          string `json:"email" binding:"omitempty,email"`
15: Phone          string `json:"phone"`
16: DocumentNumber string `json:"document_number" binding:"required"`
20: Page  uint64 `form:"page" binding:"required,min=1"`
21: Limit uint64 `form:"limit" binding:"required,min=1,max=100"`
25: Value string `form:"value" binding:"required"`
26: Page  uint64 `form:"page" binding:"required,min=1"`
27: Limit uint64 `form:"limit" binding:"required,min=1,max=100"`
30: func newOwnerFromCreateRequest(req createOwnerRequest) *domain.Owner {
39: func newOwnerFromUpdateRequest(id int64, req updateOwnerRequest) *domain.Owner {
```

### rent-manager-backend/internal/adapter/http/owner_response.go

```text
10: ID             int64     `json:"id"`
11: Name           string    `json:"name"`
12: Email          string    `json:"email"`
13: Phone          string    `json:"phone"`
14: DocumentNumber string    `json:"document_number"`
15: CreatedAt      time.Time `json:"created_at"`
16: UpdatedAt      time.Time `json:"updated_at"`
19: func newOwnerResponse(owner *domain.Owner) ownerResponse {
31: func newOwnersResponse(owners []domain.Owner) []ownerResponse {
```

### rent-manager-backend/internal/adapter/http/owner.go

```text
10: ownerUseCase "rent-manager-backend/internal/core/usecase/owner"
14: createOwnerUseCase              *ownerUseCase.CreateOwnerUseCase
15: getOwnerByIDUseCase             *ownerUseCase.GetOwnerByIDUseCase
16: getOwnerByDocumentNumberUseCase *ownerUseCase.GetOwnerByDocumentNumberUseCase
17: listOwnersUseCase               *ownerUseCase.ListOwnersUseCase
18: searchOwnersUseCase             *ownerUseCase.SearchOwnersUseCase
19: updateOwnerUseCase              *ownerUseCase.UpdateOwnerUseCase
22: func NewOwnerHandler(
23: createOwnerUseCase *ownerUseCase.CreateOwnerUseCase,
24: getOwnerByIDUseCase *ownerUseCase.GetOwnerByIDUseCase,
25: getOwnerByDocumentNumberUseCase *ownerUseCase.GetOwnerByDocumentNumberUseCase,
26: listOwnersUseCase *ownerUseCase.ListOwnersUseCase,
27: searchOwnersUseCase *ownerUseCase.SearchOwnersUseCase,
28: updateOwnerUseCase *ownerUseCase.UpdateOwnerUseCase,
40: func (h *OwnerHandler) CreateOwner(ctx *gin.Context) {
43: if err := ctx.ShouldBindJSON(&req); err != nil {
52: if err != nil {
61: func (h *OwnerHandler) GetOwnerByID(ctx *gin.Context) {
63: if err != nil {
69: if err != nil {
78: func (h *OwnerHandler) ListOwners(ctx *gin.Context) {
81: if err := ctx.ShouldBindQuery(&req); err != nil {
87: if err != nil {
96: func (h *OwnerHandler) SearchOwners(ctx *gin.Context) {
99: if err := ctx.ShouldBindQuery(&req); err != nil {
105: if err != nil {
114: func (h *OwnerHandler) UpdateOwner(ctx *gin.Context) {
116: if err != nil {
123: if err := ctx.ShouldBindJSON(&req); err != nil {
132: if err != nil {
```

### rent-manager-backend/internal/adapter/store/mysql/repository/owner.go

```text
19: func NewOwnerRepository(db *mysql.DB) *OwnerRepository {
25: func (or *OwnerRepository) CreateOwner(ctx context.Context, owner *domain.Owner) (*domain.Owner, error) {
42: if err != nil {
48: if err != nil {
51: if errCode := or.db.ErrorCode(err); errCode == "1062" {
59: if err != nil {
67: func (or *OwnerRepository) GetOwnerByID(ctx context.Context, id int64) (*domain.Owner, error) {
75: if err != nil {
89: if err != nil {
90: if errors.Is(err, sql.ErrNoRows) {
101: func (or *OwnerRepository) GetOwnerByDocumentNumber(ctx context.Context, documentNumber string) (*domain.Owner, error) {
109: if err != nil {
123: if err != nil {
124: if errors.Is(err, sql.ErrNoRows) {
135: func (or *OwnerRepository) ListOwners(
140: if page == 0 {
144: if limit == 0 {
158: func (or *OwnerRepository) SearchOwners(
164: if page == 0 {
168: if limit == 0 {
191: func (or *OwnerRepository) UpdateOwner(ctx context.Context, owner *domain.Owner) (*domain.Owner, error) {
192: queryBuilder := or.db.QueryBuilder.Update("owners")
194: if owner.Name != "" {
198: if owner.Email != "" {
202: if owner.Phone != "" {
206: if owner.DocumentNumber != "" {
210: queryBuilder = queryBuilder.Where(sq.Eq{"own_id": owner.ID})
213: if err != nil {
219: if err != nil {
222: if errCode := or.db.ErrorCode(err); errCode == "1062" {
230: if err != nil {
235: if rowsAffected == 0 {
242: func (or *OwnerRepository) baseOwnerQuery() sq.SelectBuilder {
256: func (or *OwnerRepository) findOwners(ctx context.Context, query sq.SelectBuilder) ([]domain.Owner, error) {
258: if err != nil {
264: if err != nil {
284: if err != nil {
292: if err := rows.Err(); err != nil {
```

### rent-manager-backend/internal/core/policy/owner/owner_create_policy.go

```text
16: func NewOwnerCreatePolicy(ownerRepository port.OwnerRepository) *OwnerCreatePolicy {
22: func (ownerPolicy *OwnerCreatePolicy) Execute(ctx context.Context, owner *domain.Owner) error {
23: if owner == nil {
24: return domain.ErrOwnerNil
32: if owner.Name == "" {
33: return domain.ErrOwnerNameEmpty
36: if owner.DocumentNumber == "" {
37: return domain.ErrOwnerDocumentNumberEmpty
40: if owner.Email != "" && !isValidEmail(owner.Email) {
41: return domain.ErrOwnerInvalidEmail
45: if err != nil {
47: if errors.As(err, &appErr) && appErr.Code == domain.ErrCodeDataNotFound {
58: if existingOwner != nil {
59: return domain.ErrOwnerAlreadyExists
65: func isValidEmail(email string) bool {
66: if email == "" {
```

### rent-manager-backend/internal/core/policy/owner/owner_update_policy.go

```text
16: func NewOwnerUpdatePolicy(ownerRepository port.OwnerRepository) *OwnerUpdatePolicy {
22: func (ownerPolicy *OwnerUpdatePolicy) Execute(ctx context.Context, owner *domain.Owner) error {
23: if owner == nil {
24: return domain.ErrOwnerNil
27: if owner.ID <= 0 {
28: return domain.ErrOwnerNotFound
36: if owner.Name == "" {
37: return domain.ErrOwnerNameEmpty
40: if owner.DocumentNumber == "" {
41: return domain.ErrOwnerDocumentNumberEmpty
44: if owner.Email != "" && !isValidEmail(owner.Email) {
45: return domain.ErrOwnerInvalidEmail
52: if err != nil {
55: if errors.As(err, &appErr) && appErr.Code == domain.ErrCodeDataNotFound {
66: if existingOwner != nil && existingOwner.ID != owner.ID {
67: return domain.ErrOwnerAlreadyExists
```

### rent-manager-backend/internal/core/usecase/owner/create_owner_use_case.go

```text
12: type CreateOwnerUseCase struct {
17: func NewCreateOwnerUseCase(
20: ) *CreateOwnerUseCase {
27: func (uc *CreateOwnerUseCase) CreateOwner(ctx context.Context, owner *domain.Owner) (*domain.Owner, error) {
28: if err := uc.ownerCreatePolicy.Execute(ctx, owner); err != nil {
34: if err != nil {
```

### rent-manager-backend/internal/core/usecase/owner/get_owner_by_document_number_use_case.go

```text
11: type GetOwnerByDocumentNumberUseCase struct {
15: func NewGetOwnerByDocumentNumberUseCase(
17: ) *GetOwnerByDocumentNumberUseCase {
23: func (uc *GetOwnerByDocumentNumberUseCase) GetOwnerByDocumentNumber(
28: if err != nil {
```

### rent-manager-backend/internal/core/usecase/owner/get_owner_by_id_use_case.go

```text
11: type GetOwnerByIDUseCase struct {
15: func NewGetOwnerByIDUseCase(ownerRepository port.OwnerRepository) *GetOwnerByIDUseCase {
21: func (uc *GetOwnerByIDUseCase) GetOwnerByID(ctx context.Context, id int64) (*domain.Owner, error) {
23: if err != nil {
```

### rent-manager-backend/internal/core/usecase/owner/list_owners_use_case.go

```text
11: type ListOwnersUseCase struct {
15: func NewListOwnersUseCase(ownerRepository port.OwnerRepository) *ListOwnersUseCase {
21: func (uc *ListOwnersUseCase) ListOwners(
27: if err != nil {
```

### rent-manager-backend/internal/core/usecase/owner/search_owners_use_case.go

```text
11: type SearchOwnersUseCase struct {
15: func NewSearchOwnersUseCase(ownerRepository port.OwnerRepository) *SearchOwnersUseCase {
21: func (uc *SearchOwnersUseCase) SearchOwners(
28: if err != nil {
```

### rent-manager-backend/internal/core/usecase/owner/update_owner_use_case.go

```text
12: type UpdateOwnerUseCase struct {
17: func NewUpdateOwnerUseCase(
20: ) *UpdateOwnerUseCase {
27: func (uc *UpdateOwnerUseCase) UpdateOwner(ctx context.Context, owner *domain.Owner) (*domain.Owner, error) {
28: if err := uc.ownerUpdatePolicy.Execute(ctx, owner); err != nil {
34: if err != nil {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/core/owner/owner-read-api.service.ts

```text
15: return this.apiUrl.build('/owners');
18: search(value: string, page: number, limit: number): Observable<Owner[]> {
21: return this.http
26: list(page: number, limit: number): Observable<Owner[]> {
29: return this.http
34: getById(id: number): Observable<Owner> {
35: return this.http
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/core/owner/owner-read.service.ts

```text
17: get$(id: number): Observable<Owner> {
27: list$(page = 1, limit = 10): Observable<Owner[]> {
37: search$(value: string, page = 1, limit = 10): Observable<Owner[]> {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/owner/create/owner-create.component.ts

```text
31: constructor() {
34: switchMap(() => from(this.router.navigate(['/app/owners']))),
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/owner/edit/owner-edit.component.ts

```text
38: constructor() {
48: switchMap(() => from(this.router.navigate(['/app/owners']))),
64: ngOnInit(): void {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/owner/form/owner-form.component.ts

```text
42: [OwnerFormField.Name]: new FormControl('', Validators.required),
43: [OwnerFormField.Email]: new FormControl('', [Validators.required, Validators.email]),
74: ngAfterViewInit() {
91: ngOnDestroy() {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/owner/list/owner-list.component.ts

```text
50: constructor() {
53: switchMap(() => {
84: ngOnInit(): void {
88: onPageChange(event: PageEvent): void {
94: view(id: number): void {
95: void this.router.navigate(['/app/owners/view', id]);
98: edit(id: number): void {
99: void this.router.navigate(['/app/owners/edit', id]);
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/owner/owner.routes.ts

```text
13: path: '',
18: path: '',
22: path: 'create',
26: path: 'edit/:id',
30: path: 'view/:id',
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/owner/service/owner-write-api.service.ts

```text
16: return this.apiUrl.build('/owners');
19: create(request: CreateOwnerRequest): Observable<Owner> {
20: return this.http
25: update(id: number, request: UpdateOwnerRequest): Observable<Owner> {
26: return this.http
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/owner/service/owner-write.service.ts

```text
19: save$(item: OwnerFormValue): Observable<Owner> {
35: update$(id: number, item: OwnerFormValue): Observable<Owner> {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/owner/view/owner-view.component.ts

```text
34: constructor() {
44: switchMap(() => from(this.router.navigate(['/app/owners']))),
52: switchMap(() => from(this.router.navigate(['/app/owners/edit', this.ownerId]))),
58: ngOnInit(): void {
```
