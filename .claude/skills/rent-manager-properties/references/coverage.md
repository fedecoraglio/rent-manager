# Source and operation coverage

Reviewed: 2026-11-13. Resolve paths from the repository root. Use this map to open the implementation; excerpts are navigation indexes, not replacements for complete source bodies. Line numbers refer to this review. Do not interpret internal operations as public endpoints.

## Contents

- File inventory
- Operations, data contracts, and rule locations

## File inventory

- `rent-manager-backend/internal/adapter/http/property_catalog_request.go`
- `rent-manager-backend/internal/adapter/http/property_catalog_response.go`
- `rent-manager-backend/internal/adapter/http/property_catalog.go`
- `rent-manager-backend/internal/adapter/http/property_request.go`
- `rent-manager-backend/internal/adapter/http/property_response.go`
- `rent-manager-backend/internal/adapter/http/property.go`
- `rent-manager-backend/internal/adapter/store/mysql/repository/property_catalog.go`
- `rent-manager-backend/internal/adapter/store/mysql/repository/property.go`
- `rent-manager-backend/internal/core/domain/property_summary.go`
- `rent-manager-backend/internal/core/domain/property.go`
- `rent-manager-backend/internal/core/policy/property/property_create_policy.go`
- `rent-manager-backend/internal/core/policy/property/property_update_policy.go`
- `rent-manager-backend/internal/core/port/property_catalog.go`
- `rent-manager-backend/internal/core/port/property.go`
- `rent-manager-backend/internal/core/usecase/property_catalog/list_property_statuses_use_case.go`
- `rent-manager-backend/internal/core/usecase/property_catalog/list_property_types_use_case.go`
- `rent-manager-backend/internal/core/usecase/property/create_property_use_case.go`
- `rent-manager-backend/internal/core/usecase/property/get_property_by_code_use_case.go`
- `rent-manager-backend/internal/core/usecase/property/get_property_by_id_use_case.go`
- `rent-manager-backend/internal/core/usecase/property/get_property_summaries_use_case.go`
- `rent-manager-backend/internal/core/usecase/property/list_property_use_case.go`
- `rent-manager-backend/internal/core/usecase/property/search_property_use_case.go`
- `rent-manager-backend/internal/core/usecase/property/update_property_use_case.go`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/property-catalog/property-catalog-api.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/property-catalog/property-catalog.model.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/property-catalog/property-catalog.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/property/property-read-api.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/property/property-read.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/property/property.model.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/property/create/property-create.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/property/create/property-create.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/property/create/property-create.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/property/edit/property-edit.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/property/edit/property-edit.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/property/edit/property-edit.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/property/form/field/property-form-field.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/property/form/property-form.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/property/form/property-form.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/property/form/property-form.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/property/list/property-list.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/property/list/property-list.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/property/list/property-list.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/property/list/table/property-list-table.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/property/list/table/property-list-table.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/property/list/table/property-list-table.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/property/model/property.model.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/property/property.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/property/property.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/property/property.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/property/property.routes.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/property/service/property-write-api.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/property/service/property-write.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/property/view/property-view.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/property/view/property-view.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/property/view/property-view.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/ui/property-summary-card/property-summary-card.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/ui/property-summary-card/property-summary-card.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/ui/property-summary-card/property-summary-card.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/ui/property-summary-card/property-summary-card.model.ts`

## Operations, data contracts, and rule locations

### rent-manager-backend/internal/adapter/http/property_catalog_response.go

```text
6: ID   int64  `json:"id"`
7: Code string `json:"code"`
8: Name string `json:"name"`
12: ID   int64  `json:"id"`
13: Code string `json:"code"`
14: Name string `json:"name"`
17: func newPropertyTypeResponse(propertyType domain.PropertyType) propertyTypeResponse {
25: func newPropertyTypesResponse(propertyTypes []domain.PropertyType) []propertyTypeResponse {
35: func newPropertyStatusResponse(propertyStatus domain.PropertyStatus) propertyStatusResponse {
43: func newPropertyStatusesResponse(propertyStatuses []domain.PropertyStatus) []propertyStatusResponse {
```

### rent-manager-backend/internal/adapter/http/property_catalog.go

```text
8: propertyCatalogUseCase "rent-manager-backend/internal/core/usecase/property_catalog"
12: listPropertyTypesUseCase    *propertyCatalogUseCase.ListPropertyTypesUseCase
13: listPropertyStatusesUseCase *propertyCatalogUseCase.ListPropertyStatusesUseCase
16: func NewPropertyCatalogHandler(
17: listPropertyTypesUseCase *propertyCatalogUseCase.ListPropertyTypesUseCase,
18: listPropertyStatusesUseCase *propertyCatalogUseCase.ListPropertyStatusesUseCase,
26: func (h *PropertyCatalogHandler) ListPropertyTypes(ctx *gin.Context) {
28: if err != nil {
37: func (h *PropertyCatalogHandler) ListPropertyStatuses(ctx *gin.Context) {
39: if err != nil {
```

### rent-manager-backend/internal/adapter/http/property_request.go

```text
6: OwnerID  int64 `json:"owner_id" binding:"required,min=1"`
7: TypeID   int64 `json:"type_id" binding:"required,min=1"`
8: StatusID int64 `json:"status_id" binding:"required,min=1"`
10: CountryID int64 `json:"country_id" binding:"required,min=1"`
11: StateID   int64 `json:"state_id" binding:"required,min=1"`
13: Code        string `json:"code" binding:"required"`
14: Title       string `json:"title" binding:"required"`
15: Description string `json:"description"`
17: Street       string `json:"street" binding:"required"`
18: StreetNumber string `json:"street_number"`
19: Floor        string `json:"floor"`
20: Apartment    string `json:"apartment"`
21: City         string `json:"city" binding:"required"`
22: PostalCode   string `json:"postal_code"`
26: OwnerID  int64 `json:"owner_id" binding:"required,min=1"`
27: TypeID   int64 `json:"type_id" binding:"required,min=1"`
28: StatusID int64 `json:"status_id" binding:"required,min=1"`
30: CountryID int64 `json:"country_id" binding:"required,min=1"`
31: StateID   int64 `json:"state_id" binding:"required,min=1"`
33: Code        string `json:"code" binding:"required"`
34: Title       string `json:"title" binding:"required"`
35: Description string `json:"description"`
37: Street       string `json:"street" binding:"required"`
38: StreetNumber string `json:"street_number"`
39: Floor        string `json:"floor"`
40: Apartment    string `json:"apartment"`
41: City         string `json:"city" binding:"required"`
42: PostalCode   string `json:"postal_code"`
46: Page  uint64 `form:"page" binding:"required,min=1"`
47: Limit uint64 `form:"limit" binding:"required,min=1,max=100"`
51: Value string `form:"value" binding:"required"`
52: Page  uint64 `form:"page" binding:"required,min=1"`
53: Limit uint64 `form:"limit" binding:"required,min=1,max=100"`
56: func newPropertyFromCreateRequest(req createPropertyRequest) *domain.Property {
75: func newPropertyFromUpdateRequest(
```

### rent-manager-backend/internal/adapter/http/property_response.go

```text
10: ID           int64     `json:"id"`
11: OwnerID      int64     `json:"owner_id"`
12: TypeID       int64     `json:"type_id"`
13: StatusID     int64     `json:"status_id"`
14: CountryID    int64     `json:"country_id"`
15: StateID      int64     `json:"state_id"`
16: Code         string    `json:"code"`
17: Title        string    `json:"title"`
18: Description  string    `json:"description"`
19: Street       string    `json:"street"`
20: StreetNumber string    `json:"street_number"`
21: Floor        string    `json:"floor"`
22: Apartment    string    `json:"apartment"`
23: City         string    `json:"city"`
24: PostalCode   string    `json:"postal_code"`
25: CreatedAt    time.Time `json:"created_at"`
26: UpdatedAt    time.Time `json:"updated_at"`
30: PropertyID            int64                         `json:"id"`
31: PropertyTitle         string                        `json:"title"`
32: RentalContractSummary rentalContractSummaryResponse `json:"summary"`
35: func newPropertyResponse(property *domain.Property) propertyResponse {
57: func newPropertySummaryResponse(property *domain.PropertySummary) propertySummaryResponse {
65: func newPropertiesResponse(properties []domain.Property) []propertyResponse {
76: func newPropertiesSummaryResponse(properties []domain.PropertySummary) []propertySummaryResponse {
```

### rent-manager-backend/internal/adapter/http/property.go

```text
10: propertyUseCase "rent-manager-backend/internal/core/usecase/property"
14: createPropertyUseCase          *propertyUseCase.CreatePropertyUseCase
15: getPropertyByIDUseCase         *propertyUseCase.GetPropertyByIDUseCase
16: getPropertyByCodeUseCase       *propertyUseCase.GetPropertyByCodeUseCase
17: listPropertiesUseCase          *propertyUseCase.ListPropertiesUseCase
18: listPropertiesSummariesUseCase *propertyUseCase.ListPropertiesSummariesUseCase
19: searchPropertiesUseCase        *propertyUseCase.SearchPropertiesUseCase
20: updatePropertyUseCase          *propertyUseCase.UpdatePropertyUseCase
23: func NewPropertyHandler(
24: createPropertyUseCase *propertyUseCase.CreatePropertyUseCase,
25: getPropertyByIDUseCase *propertyUseCase.GetPropertyByIDUseCase,
26: getPropertyByCodeUseCase *propertyUseCase.GetPropertyByCodeUseCase,
27: listPropertiesUseCase *propertyUseCase.ListPropertiesUseCase,
28: listPropertiesSummariesUseCase *propertyUseCase.ListPropertiesSummariesUseCase,
29: searchPropertiesUseCase *propertyUseCase.SearchPropertiesUseCase,
30: updatePropertyUseCase *propertyUseCase.UpdatePropertyUseCase,
43: func (h *PropertyHandler) CreateProperty(ctx *gin.Context) {
46: if err := ctx.ShouldBindJSON(&req); err != nil {
55: if err != nil {
64: func (h *PropertyHandler) GetPropertyByID(ctx *gin.Context) {
66: if err != nil {
72: if err != nil {
81: func (h *PropertyHandler) ListProperties(ctx *gin.Context) {
84: if err := ctx.ShouldBindQuery(&req); err != nil {
90: if err != nil {
99: func (h *PropertyHandler) ListPropertiesSummary(ctx *gin.Context) {
102: if err := ctx.ShouldBindQuery(&req); err != nil {
108: if err != nil {
117: func (h *PropertyHandler) SearchProperties(ctx *gin.Context) {
120: if err := ctx.ShouldBindQuery(&req); err != nil {
126: if err != nil {
135: func (h *PropertyHandler) UpdateProperty(ctx *gin.Context) {
137: if err != nil {
144: if err := ctx.ShouldBindJSON(&req); err != nil {
153: if err != nil {
```

### rent-manager-backend/internal/adapter/store/mysql/repository/property_catalog.go

```text
15: func NewPropertyCatalogRepository(db *mysql.DB) *PropertyCatalogRepository {
21: func (pr *PropertyCatalogRepository) ListPropertyTypes(ctx context.Context) ([]domain.PropertyType, error) {
32: if err != nil {
38: if err != nil {
49: if err := rows.Scan(
61: if err := rows.Err(); err != nil {
69: func (pr *PropertyCatalogRepository) ListPropertyStatuses(ctx context.Context) ([]domain.PropertyStatus, error) {
80: if err != nil {
86: if err != nil {
97: if err := rows.Scan(
109: if err := rows.Err(); err != nil {
```

### rent-manager-backend/internal/adapter/store/mysql/repository/property.go

```text
19: func NewPropertyRepository(db *mysql.DB) *PropertyRepository {
25: func (pr *PropertyRepository) CreateProperty(ctx context.Context, property *domain.Property) (*domain.Property, error) {
62: if err != nil {
68: if err != nil {
71: if errCode := pr.db.ErrorCode(err); errCode == "1062" {
79: if err != nil {
87: func (pr *PropertyRepository) GetPropertyByID(ctx context.Context, id int64) (*domain.Property, error) {
95: if err != nil {
119: if err != nil {
120: if errors.Is(err, sql.ErrNoRows) {
131: func (pr *PropertyRepository) GetPropertyByCode(ctx context.Context, code string) (*domain.Property, error) {
139: if err != nil {
163: if err != nil {
164: if errors.Is(err, sql.ErrNoRows) {
175: func (pr *PropertyRepository) ListProperties(
180: if page == 0 {
184: if limit == 0 {
198: func (pr *PropertyRepository) SearchProperties(
204: if page == 0 {
208: if limit == 0 {
233: func (pr *PropertyRepository) UpdateProperty(
237: queryBuilder := pr.db.QueryBuilder.Update("properties")
239: if property.OwnerID != 0 {
243: if property.TypeID != 0 {
247: if property.StatusID != 0 {
251: if property.CountryID != 0 {
255: if property.StateID != 0 {
259: if property.Code != "" {
263: if property.Title != "" {
267: if property.Description != "" {
271: if property.Street != "" {
275: if property.StreetNumber != "" {
279: if property.Floor != "" {
283: if property.Apartment != "" {
287: if property.City != "" {
291: if property.PostalCode != "" {
295: queryBuilder = queryBuilder.Where(sq.Eq{"pro_id": property.ID})
298: if err != nil {
304: if err != nil {
307: if errCode := pr.db.ErrorCode(err); errCode == "1062" {
315: if err != nil {
320: if rowsAffected == 0 {
327: func (pr *PropertyRepository) basePropertyQuery() sq.SelectBuilder {
351: func (pr *PropertyRepository) findProperties(ctx context.Context, query sq.SelectBuilder) ([]domain.Property, error) {
353: if err != nil {
359: if err != nil {
389: if err != nil {
397: if err := rows.Err(); err != nil {
```

### rent-manager-backend/internal/core/policy/property/property_create_policy.go

```text
16: func NewPropertyCreatePolicy(propertyRepository port.PropertyRepository) *PropertyCreatePolicy {
22: func (propertyPolicy *PropertyCreatePolicy) Execute(
26: if property == nil {
27: return domain.ErrPropertyNil
40: if property.OwnerID <= 0 {
41: return domain.ErrPropertyOwnerIDEmpty
44: if property.TypeID <= 0 {
45: return domain.ErrPropertyTypeIDEmpty
48: if property.StatusID <= 0 {
49: return domain.ErrPropertyStatusIDEmpty
52: if property.CountryID <= 0 {
53: return domain.ErrPropertyCountryIDEmpty
56: if property.StateID <= 0 {
57: return domain.ErrPropertyStateIDEmpty
60: if property.Code == "" {
61: return domain.ErrPropertyCodeEmpty
64: if property.Title == "" {
65: return domain.ErrPropertyTitleEmpty
68: if property.Street == "" {
69: return domain.ErrPropertyStreetEmpty
72: if property.City == "" {
73: return domain.ErrPropertyCityEmpty
77: if err != nil {
80: if errors.As(err, &appErr) && appErr.Code == domain.ErrCodeDataNotFound {
91: if existingProperty != nil {
92: return domain.ErrPropertyAlreadyExists
```

### rent-manager-backend/internal/core/policy/property/property_update_policy.go

```text
16: func NewPropertyUpdatePolicy(propertyRepository port.PropertyRepository) *PropertyUpdatePolicy {
22: func (propertyPolicy *PropertyUpdatePolicy) Execute(ctx context.Context, property *domain.Property) error {
23: if property == nil {
24: return domain.ErrPropertyNil
27: if property.ID <= 0 {
28: return domain.ErrPropertyNotFound
41: if property.OwnerID <= 0 {
42: return domain.ErrPropertyOwnerIDEmpty
45: if property.TypeID <= 0 {
46: return domain.ErrPropertyTypeIDEmpty
49: if property.StatusID <= 0 {
50: return domain.ErrPropertyStatusIDEmpty
53: if property.CountryID <= 0 {
54: return domain.ErrPropertyCountryIDEmpty
57: if property.StateID <= 0 {
58: return domain.ErrPropertyStateIDEmpty
61: if property.Code == "" {
62: return domain.ErrPropertyCodeEmpty
65: if property.Title == "" {
66: return domain.ErrPropertyTitleEmpty
69: if property.Street == "" {
70: return domain.ErrPropertyStreetEmpty
73: if property.City == "" {
74: return domain.ErrPropertyCityEmpty
78: if err != nil {
81: if errors.As(err, &appErr) && appErr.Code == domain.ErrCodeDataNotFound {
92: if existingProperty != nil && existingProperty.ID != property.ID {
93: return domain.ErrPropertyAlreadyExists
```

### rent-manager-backend/internal/core/usecase/property_catalog/list_property_statuses_use_case.go

```text
11: type ListPropertyStatusesUseCase struct {
15: func NewListPropertyStatusesUseCase(
17: ) *ListPropertyStatusesUseCase {
23: func (uc *ListPropertyStatusesUseCase) ListPropertyStatuses(
27: if err != nil {
```

### rent-manager-backend/internal/core/usecase/property_catalog/list_property_types_use_case.go

```text
11: type ListPropertyTypesUseCase struct {
15: func NewListPropertyTypesUseCase(
17: ) *ListPropertyTypesUseCase {
23: func (uc *ListPropertyTypesUseCase) ListPropertyTypes(
27: if err != nil {
```

### rent-manager-backend/internal/core/usecase/property/create_property_use_case.go

```text
12: type CreatePropertyUseCase struct {
17: func NewCreatePropertyUseCase(
20: ) *CreatePropertyUseCase {
27: func (uc *CreatePropertyUseCase) CreateProperty(
31: if err := uc.propertyCreatePolicy.Execute(ctx, property); err != nil {
37: if err != nil {
```

### rent-manager-backend/internal/core/usecase/property/get_property_by_code_use_case.go

```text
11: type GetPropertyByCodeUseCase struct {
15: func NewGetPropertyByCodeUseCase(propertyRepository port.PropertyRepository) *GetPropertyByCodeUseCase {
21: func (uc *GetPropertyByCodeUseCase) GetPropertyByCode(ctx context.Context, code string) (*domain.Property, error) {
23: if err != nil {
```

### rent-manager-backend/internal/core/usecase/property/get_property_by_id_use_case.go

```text
11: type GetPropertyByIDUseCase struct {
15: func NewGetPropertyByIDUseCase(propertyRepository port.PropertyRepository) *GetPropertyByIDUseCase {
21: func (uc *GetPropertyByIDUseCase) GetPropertyByID(ctx context.Context, id int64) (*domain.Property, error) {
23: if err != nil {
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

### rent-manager-backend/internal/core/usecase/property/list_property_use_case.go

```text
11: type ListPropertiesUseCase struct {
15: func NewListPropertiesUseCase(propertyRepository port.PropertyRepository) *ListPropertiesUseCase {
21: func (uc *ListPropertiesUseCase) ListProperties(
27: if err != nil {
```

### rent-manager-backend/internal/core/usecase/property/search_property_use_case.go

```text
11: type SearchPropertiesUseCase struct {
15: func NewSearchPropertiesUseCase(propertyRepository port.PropertyRepository) *SearchPropertiesUseCase {
21: func (uc *SearchPropertiesUseCase) SearchProperties(
28: if err != nil {
```

### rent-manager-backend/internal/core/usecase/property/update_property_use_case.go

```text
12: type UpdatePropertyUseCase struct {
17: func NewUpdatePropertyUseCase(
20: ) *UpdatePropertyUseCase {
27: func (uc *UpdatePropertyUseCase) UpdateProperty(
31: if err := uc.propertyUpdatePolicy.Execute(ctx, property); err != nil {
37: if err != nil {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/core/property-catalog/property-catalog-api.service.ts

```text
15: return this.apiUrl.build('/property-catalogs');
18: listPropertyTypes$(): Observable<ApiResponse<PropertyType[]>> {
19: return this.http
24: listPropertyStatuses$(): Observable<ApiResponse<PropertyStatus[]>> {
25: return this.http
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/core/property-catalog/property-catalog.service.ts

```text
18: listPropertyTypes$(): Observable<PropertyType[]> {
23: catchError((e: HttpErrorResponse) => {
30: listPropertyStatuses$(): Observable<PropertyStatus[]> {
35: catchError((e: HttpErrorResponse) => {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/core/property/property-read-api.service.ts

```text
15: return this.apiUrl.build('/properties');
18: list$(page: number, limit: number): Observable<Property[]> {
21: return this.http
26: listSummaries$(page: number, limit: number): Observable<ApiResponse<PropertySummary[]>> {
29: return this.http.get<ApiResponse<PropertySummary[]>>(`${this.baseUrl}/summaries`, { params });
32: search$(value: string, page: number, limit: number): Observable<Property[]> {
35: return this.http
40: getById$(id: number): Observable<Property> {
41: return this.http
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/core/property/property-read.service.ts

```text
23: get$(id: number): Observable<Property> {
33: list$(page = 1, limit = 10): Observable<Property[]> {
43: listSummary$(page = 1, limit = 10): Observable<PropertySummary[]> {
54: search$(value: string, page = 1, limit = 10): Observable<Property[]> {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/property/create/property-create.component.ts

```text
31: constructor() {
34: switchMap(() => from(this.router.navigate(['/app/properties']))),
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/property/edit/property-edit.component.ts

```text
38: constructor() {
41: switchMap(() => from(this.router.navigate(['/app/properties']))),
59: ngOnInit(): void {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/property/form/property-form.component.ts

```text
52: [PropertyFormField.OwnerId]: new FormControl(0, Validators.required),
53: [PropertyFormField.TypeId]: new FormControl(0, Validators.required),
54: [PropertyFormField.StatusId]: new FormControl(0, Validators.required),
55: [PropertyFormField.CountryId]: new FormControl(1, Validators.required),
56: [PropertyFormField.StateId]: new FormControl(0, Validators.required),
57: [PropertyFormField.Code]: new FormControl('', Validators.required),
58: [PropertyFormField.Title]: new FormControl('', Validators.required),
60: [PropertyFormField.Street]: new FormControl('', Validators.required),
64: [PropertyFormField.City]: new FormControl('', Validators.required),
76: constructor() {
127: ngAfterViewInit(): void {
156: ngOnDestroy(): void {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/property/list/property-list.component.ts

```text
52: constructor() {
55: switchMap(() => {
86: ngOnInit(): void {
90: onPageChange(event: PageEvent): void {
96: viewProperty(id: number): void {
97: void this.router.navigate(['/app/properties/view', id]);
100: editProperty(id: number): void {
101: void this.router.navigate(['/app/properties/edit', id]);
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/property/property.routes.ts

```text
13: path: '',
18: path: '',
22: path: 'create',
26: path: 'edit/:id',
30: path: 'view/:id',
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/property/service/property-write-api.service.ts

```text
16: return this.apiUrl.build('/properties');
19: create$(request: CreatePropertyRequest): Observable<Property> {
20: return this.http
25: update$(id: number, request: UpdatePropertyRequest): Observable<Property> {
26: return this.http
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/property/service/property-write.service.ts

```text
22: save$(property: PropertyFormValue): Observable<Property> {
35: update$(id: number, property: PropertyFormValue): Observable<Property> {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/property/view/property-view.component.ts

```text
33: constructor() {
36: switchMap(() => from(this.router.navigate(['/app/properties']))),
44: switchMap(() => from(this.router.navigate(['/app/properties/edit', this.propertyId]))),
50: ngOnInit(): void {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/ui/property-summary-card/property-summary-card.component.ts

```text
21: constructor() {
26: this.router.navigate([
37: gotToViewProperty() {
```
