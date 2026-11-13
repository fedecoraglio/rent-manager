# Source and operation coverage

Reviewed: 2026-11-13. Resolve paths from the repository root. Use this map to open the implementation; excerpts are navigation indexes, not replacements for complete source bodies. Line numbers refer to this review. Do not interpret internal operations as public endpoints.

## Contents

- File inventory
- Operations, data contracts, and rule locations

## File inventory

- `rent-manager-backend/cmd/api/main.go`
- `rent-manager-backend/docker-compose.yml`
- `rent-manager-backend/Dockerfile`
- `rent-manager-backend/go.mod`
- `rent-manager-backend/internal/adapter/config/config.go`
- `rent-manager-backend/internal/adapter/http/auth.go`
- `rent-manager-backend/internal/adapter/http/contract_catalog_response.go`
- `rent-manager-backend/internal/adapter/http/contract_catalog.go`
- `rent-manager-backend/internal/adapter/http/error.go`
- `rent-manager-backend/internal/adapter/http/inflation_index_request.go`
- `rent-manager-backend/internal/adapter/http/inflation_index_response.go`
- `rent-manager-backend/internal/adapter/http/inflation_index.go`
- `rent-manager-backend/internal/adapter/http/location_request.go`
- `rent-manager-backend/internal/adapter/http/location_response.go`
- `rent-manager-backend/internal/adapter/http/location.go`
- `rent-manager-backend/internal/adapter/http/middleware_auth.go`
- `rent-manager-backend/internal/adapter/http/owner_request.go`
- `rent-manager-backend/internal/adapter/http/owner_response.go`
- `rent-manager-backend/internal/adapter/http/owner.go`
- `rent-manager-backend/internal/adapter/http/property_catalog_request.go`
- `rent-manager-backend/internal/adapter/http/property_catalog_response.go`
- `rent-manager-backend/internal/adapter/http/property_catalog.go`
- `rent-manager-backend/internal/adapter/http/property_request.go`
- `rent-manager-backend/internal/adapter/http/property_response.go`
- `rent-manager-backend/internal/adapter/http/property.go`
- `rent-manager-backend/internal/adapter/http/rent_contract_request.go`
- `rent-manager-backend/internal/adapter/http/rent_contract_response.go`
- `rent-manager-backend/internal/adapter/http/rent_contract.go`
- `rent-manager-backend/internal/adapter/http/rent_payment_request.go`
- `rent-manager-backend/internal/adapter/http/rent_payment_response.go`
- `rent-manager-backend/internal/adapter/http/rent_payment_schedule_response.go`
- `rent-manager-backend/internal/adapter/http/rent_payment_suggestion_request.go`
- `rent-manager-backend/internal/adapter/http/rent_payment_suggestion_response.go`
- `rent-manager-backend/internal/adapter/http/rent_payment.go`
- `rent-manager-backend/internal/adapter/http/rental_contract_summary_response.go`
- `rent-manager-backend/internal/adapter/http/request.go`
- `rent-manager-backend/internal/adapter/http/response.go`
- `rent-manager-backend/internal/adapter/http/role_response.go`
- `rent-manager-backend/internal/adapter/http/role.go`
- `rent-manager-backend/internal/adapter/http/router.go`
- `rent-manager-backend/internal/adapter/http/tenant_request.go`
- `rent-manager-backend/internal/adapter/http/tenant_response.go`
- `rent-manager-backend/internal/adapter/http/tenant.go`
- `rent-manager-backend/internal/adapter/http/user_request.go`
- `rent-manager-backend/internal/adapter/http/user_response.go`
- `rent-manager-backend/internal/adapter/http/user.go`
- `rent-manager-backend/internal/adapter/security/bcrypt_password_hasher.go`
- `rent-manager-backend/internal/adapter/security/jwt_token_provider.go`
- `rent-manager-backend/internal/adapter/store/mysql/db.go`
- `rent-manager-backend/internal/adapter/store/mysql/migrations/000001_create_countries_table.down.sql`
- `rent-manager-backend/internal/adapter/store/mysql/migrations/000001_create_countries_table.up.sql`
- `rent-manager-backend/internal/adapter/store/mysql/migrations/000002_create_states_table.down.sql`
- `rent-manager-backend/internal/adapter/store/mysql/migrations/000002_create_states_table.up.sql`
- `rent-manager-backend/internal/adapter/store/mysql/migrations/000003_create_roles_table.down.sql`
- `rent-manager-backend/internal/adapter/store/mysql/migrations/000003_create_roles_table.up.sql`
- `rent-manager-backend/internal/adapter/store/mysql/migrations/000004_create_users_table.down.sql`
- `rent-manager-backend/internal/adapter/store/mysql/migrations/000004_create_users_table.up.sql`
- `rent-manager-backend/internal/adapter/store/mysql/migrations/000005_create_owners_table.down.sql`
- `rent-manager-backend/internal/adapter/store/mysql/migrations/000005_create_owners_table.up.sql`
- `rent-manager-backend/internal/adapter/store/mysql/migrations/000006_create_tenants_table.down.sql`
- `rent-manager-backend/internal/adapter/store/mysql/migrations/000006_create_tenants_table.up.sql`
- `rent-manager-backend/internal/adapter/store/mysql/migrations/000007_create_property_catalogs.down.sql`
- `rent-manager-backend/internal/adapter/store/mysql/migrations/000007_create_property_catalogs.up.sql`
- `rent-manager-backend/internal/adapter/store/mysql/migrations/000008_create_properties_table.down.sql`
- `rent-manager-backend/internal/adapter/store/mysql/migrations/000008_create_properties_table.up.sql`
- `rent-manager-backend/internal/adapter/store/mysql/migrations/000009_create_contract_statuses_table.down.sql`
- `rent-manager-backend/internal/adapter/store/mysql/migrations/000009_create_contract_statuses_table.up.sql`
- `rent-manager-backend/internal/adapter/store/mysql/migrations/000010_create_rental_contracts_table.down.sql`
- `rent-manager-backend/internal/adapter/store/mysql/migrations/000010_create_rental_contracts_table.up.sql`
- `rent-manager-backend/internal/adapter/store/mysql/migrations/000011_create_payment_statuses_table.down.sql`
- `rent-manager-backend/internal/adapter/store/mysql/migrations/000011_create_payment_statuses_table.up.sql`
- `rent-manager-backend/internal/adapter/store/mysql/migrations/000012_create_rent_payments_table.down.sql`
- `rent-manager-backend/internal/adapter/store/mysql/migrations/000012_create_rent_payments_table.up.sql`
- `rent-manager-backend/internal/adapter/store/mysql/migrations/000013_create_inflation_indexes.down.sql`
- `rent-manager-backend/internal/adapter/store/mysql/migrations/000013_create_inflation_indexes.up.sql`
- `rent-manager-backend/internal/adapter/store/mysql/repository/contract_catalog.go`
- `rent-manager-backend/internal/adapter/store/mysql/repository/inflation_index.go`
- `rent-manager-backend/internal/adapter/store/mysql/repository/location.go`
- `rent-manager-backend/internal/adapter/store/mysql/repository/owner.go`
- `rent-manager-backend/internal/adapter/store/mysql/repository/property_catalog.go`
- `rent-manager-backend/internal/adapter/store/mysql/repository/property.go`
- `rent-manager-backend/internal/adapter/store/mysql/repository/rent_contract.go`
- `rent-manager-backend/internal/adapter/store/mysql/repository/rent_payment.go`
- `rent-manager-backend/internal/adapter/store/mysql/repository/role.go`
- `rent-manager-backend/internal/adapter/store/mysql/repository/shared/shared.go`
- `rent-manager-backend/internal/adapter/store/mysql/repository/tenant.go`
- `rent-manager-backend/internal/adapter/store/mysql/repository/user.go`
- `rent-manager-backend/internal/core/domain/adjustment_suggestion.go`
- `rent-manager-backend/internal/core/domain/country.go`
- `rent-manager-backend/internal/core/domain/error_catalog.go`
- `rent-manager-backend/internal/core/domain/error_inflation_index.go`
- `rent-manager-backend/internal/core/domain/error_owner.go`
- `rent-manager-backend/internal/core/domain/error_rent_contract.go`
- `rent-manager-backend/internal/core/domain/error_tenant.go`
- `rent-manager-backend/internal/core/domain/error_user.go`
- `rent-manager-backend/internal/core/domain/error.go`
- `rent-manager-backend/internal/core/domain/errors_property.go`
- `rent-manager-backend/internal/core/domain/errors_rent_payment.go`
- `rent-manager-backend/internal/core/domain/inflation_indexes.go`
- `rent-manager-backend/internal/core/domain/login.go`
- `rent-manager-backend/internal/core/domain/owner.go`
- `rent-manager-backend/internal/core/domain/property_summary.go`
- `rent-manager-backend/internal/core/domain/property.go`
- `rent-manager-backend/internal/core/domain/rent_contract.go`
- `rent-manager-backend/internal/core/domain/rent_payment_calculation.go`
- `rent-manager-backend/internal/core/domain/rent_payment_schedule.go`
- `rent-manager-backend/internal/core/domain/rent_payment_suggestion.go`
- `rent-manager-backend/internal/core/domain/rent_payment.go`
- `rent-manager-backend/internal/core/domain/rental_contract_summary.go`
- `rent-manager-backend/internal/core/domain/role.go`
- `rent-manager-backend/internal/core/domain/state.go`
- `rent-manager-backend/internal/core/domain/tenant.go`
- `rent-manager-backend/internal/core/domain/token.go`
- `rent-manager-backend/internal/core/domain/user.go`
- `rent-manager-backend/internal/core/policy/inflation_index/inflation_index_create_policy.go`
- `rent-manager-backend/internal/core/policy/inflation_index/inflation_index_update_policy.go`
- `rent-manager-backend/internal/core/policy/owner/owner_create_policy.go`
- `rent-manager-backend/internal/core/policy/owner/owner_update_policy.go`
- `rent-manager-backend/internal/core/policy/property/property_create_policy.go`
- `rent-manager-backend/internal/core/policy/property/property_update_policy.go`
- `rent-manager-backend/internal/core/policy/rent_contract/rent_contract_create_policy.go`
- `rent-manager-backend/internal/core/policy/rent_contract/rent_contract_update_policy.go`
- `rent-manager-backend/internal/core/policy/rent_payment/rent_payment_create_policy.go`
- `rent-manager-backend/internal/core/policy/rent_payment/rent_payment_update_policy.go`
- `rent-manager-backend/internal/core/policy/shared/shared.go`
- `rent-manager-backend/internal/core/policy/tenant/tenant_create_policy.go`
- `rent-manager-backend/internal/core/policy/tenant/tenant_update_policy.go`
- `rent-manager-backend/internal/core/policy/user/user_create_policy.go`
- `rent-manager-backend/internal/core/policy/user/user_update_policy.go`
- `rent-manager-backend/internal/core/port/contract_catalog.go`
- `rent-manager-backend/internal/core/port/inflation_index.go`
- `rent-manager-backend/internal/core/port/location.go`
- `rent-manager-backend/internal/core/port/owner.go`
- `rent-manager-backend/internal/core/port/password.go`
- `rent-manager-backend/internal/core/port/property_catalog.go`
- `rent-manager-backend/internal/core/port/property.go`
- `rent-manager-backend/internal/core/port/rent_contract.go`
- `rent-manager-backend/internal/core/port/rent_payment_calculator.go`
- `rent-manager-backend/internal/core/port/rent_payment.go`
- `rent-manager-backend/internal/core/port/role.go`
- `rent-manager-backend/internal/core/port/tenant.go`
- `rent-manager-backend/internal/core/port/token.go`
- `rent-manager-backend/internal/core/port/user.go`
- `rent-manager-backend/internal/core/service/inflation_adjustment_calculator.go`
- `rent-manager-backend/internal/core/service/rent_payment_calculator.go`
- `rent-manager-backend/internal/core/usecase/contract_catalog/list_contract_statuses_use_case.go`
- `rent-manager-backend/internal/core/usecase/contract_catalog/list_interest_calculation_types_use_case.go`
- `rent-manager-backend/internal/core/usecase/contract_catalog/list_rent_adjustment_types_use_case.go`
- `rent-manager-backend/internal/core/usecase/contract_catalog/rent_contract_catalog.go`
- `rent-manager-backend/internal/core/usecase/inflation_index/create_inflation_index_use_case.go`
- `rent-manager-backend/internal/core/usecase/inflation_index/get_inflation_index_by_id_use_case.go`
- `rent-manager-backend/internal/core/usecase/inflation_index/get_inflation_index_by_period_use_case.go`
- `rent-manager-backend/internal/core/usecase/inflation_index/list_inflation_indexes_by_period_range_use_case.go`
- `rent-manager-backend/internal/core/usecase/inflation_index/list_inflation_indexes_use_case.go`
- `rent-manager-backend/internal/core/usecase/inflation_index/update_inflation_index_use_case.go`
- `rent-manager-backend/internal/core/usecase/location/list_countries_use_case.go`
- `rent-manager-backend/internal/core/usecase/location/list_states_by_country_use_case.go`
- `rent-manager-backend/internal/core/usecase/owner/create_owner_use_case.go`
- `rent-manager-backend/internal/core/usecase/owner/get_owner_by_document_number_use_case.go`
- `rent-manager-backend/internal/core/usecase/owner/get_owner_by_email_use_case.go`
- `rent-manager-backend/internal/core/usecase/owner/get_owner_by_id_use_case.go`
- `rent-manager-backend/internal/core/usecase/owner/list_owners_use_case.go`
- `rent-manager-backend/internal/core/usecase/owner/search_owners_use_case.go`
- `rent-manager-backend/internal/core/usecase/owner/update_owner_use_case.go`
- `rent-manager-backend/internal/core/usecase/property_catalog/list_property_statuses_use_case.go`
- `rent-manager-backend/internal/core/usecase/property_catalog/list_property_types_use_case.go`
- `rent-manager-backend/internal/core/usecase/property/create_property_use_case.go`
- `rent-manager-backend/internal/core/usecase/property/get_property_by_code_use_case.go`
- `rent-manager-backend/internal/core/usecase/property/get_property_by_id_use_case.go`
- `rent-manager-backend/internal/core/usecase/property/get_property_summaries_use_case.go`
- `rent-manager-backend/internal/core/usecase/property/list_property_use_case.go`
- `rent-manager-backend/internal/core/usecase/property/search_property_use_case.go`
- `rent-manager-backend/internal/core/usecase/property/update_property_use_case.go`
- `rent-manager-backend/internal/core/usecase/rent_payment/create_rent_payment_use_case.go`
- `rent-manager-backend/internal/core/usecase/rent_payment/get_rent_payment_by_id_use_case.go`
- `rent-manager-backend/internal/core/usecase/rent_payment/get_rent_payment_schedule_use_case.go`
- `rent-manager-backend/internal/core/usecase/rent_payment/get_rent_payment_suggestion_use_case.go`
- `rent-manager-backend/internal/core/usecase/rent_payment/get_rental_contract_summary_use_case.go`
- `rent-manager-backend/internal/core/usecase/rent_payment/list_rent_payments_use_case.go`
- `rent-manager-backend/internal/core/usecase/rent_payment/update_rent_payment_use_case.go`
- `rent-manager-backend/internal/core/usecase/rental_contract/create_rental_contract_use_case.go`
- `rent-manager-backend/internal/core/usecase/rental_contract/get_rental_contract_by_id_use_case.go`
- `rent-manager-backend/internal/core/usecase/rental_contract/list_rental_contracts_use_case.go`
- `rent-manager-backend/internal/core/usecase/rental_contract/update_rental_contract_use_case.go`
- `rent-manager-backend/internal/core/usecase/role/list_roles_use_case.go`
- `rent-manager-backend/internal/core/usecase/tenant/create_tenant_use_case.go`
- `rent-manager-backend/internal/core/usecase/tenant/get_tenant_by_document_number_use_case.go`
- `rent-manager-backend/internal/core/usecase/tenant/get_tenant_by_id_use_case.go`
- `rent-manager-backend/internal/core/usecase/tenant/list_tenant_use_case.go`
- `rent-manager-backend/internal/core/usecase/tenant/search_tenant_use_case.go`
- `rent-manager-backend/internal/core/usecase/tenant/update_tenant_use_case.go`
- `rent-manager-backend/internal/core/usecase/user/create_user_use_case.go`
- `rent-manager-backend/internal/core/usecase/user/delete_user_use_case.go`
- `rent-manager-backend/internal/core/usecase/user/get_user_by_email_use_case.go`
- `rent-manager-backend/internal/core/usecase/user/get_user_by_id_use_case.go`
- `rent-manager-backend/internal/core/usecase/user/list_user_use_case.go`
- `rent-manager-backend/internal/core/usecase/user/login_use_case.go`
- `rent-manager-backend/internal/core/usecase/user/search_user_use_case.go`
- `rent-manager-backend/internal/core/usecase/user/update_user_use_case.go`

## Operations, data contracts, and rule locations

### rent-manager-backend/cmd/api/main.go

```text
21: contractCatalogUseCase "rent-manager-backend/internal/core/usecase/contract_catalog"
22: inflationIndexUseCase "rent-manager-backend/internal/core/usecase/inflation_index"
23: locationUseCase "rent-manager-backend/internal/core/usecase/location"
24: ownerUseCase "rent-manager-backend/internal/core/usecase/owner"
25: propertyUseCase "rent-manager-backend/internal/core/usecase/property"
26: propertyCatalogUseCase "rent-manager-backend/internal/core/usecase/property_catalog"
27: rentPaymentUseCase "rent-manager-backend/internal/core/usecase/rent_payment"
28: rentalContractUseCase "rent-manager-backend/internal/core/usecase/rental_contract"
29: roleUsecase "rent-manager-backend/internal/core/usecase/role"
30: tenantUseCase "rent-manager-backend/internal/core/usecase/tenant"
31: userUseCase "rent-manager-backend/internal/core/usecase/user"
35: func main() {
38: if err != nil {
48: if err != nil {
58: if err != nil {
81: getRentalContractSummaryUseCase := rentPaymentUseCase.NewGetRentalContractSummaryUseCase(
86: getRentPaymentScheduleUseCase := rentPaymentUseCase.NewGetRentPaymentScheduleUseCase(
91: createRentPaymentUseCase := rentPaymentUseCase.NewCreateRentPaymentUseCase(
95: updateRentPaymentUseCase := rentPaymentUseCase.NewUpdateRentPaymentUseCase(
99: getRentPaymentByIDUseCase := rentPaymentUseCase.NewGetRentPaymentByIDUseCase(rentPaymentRepository)
100: listRentPaymentsUseCase := rentPaymentUseCase.NewListRentPaymentsUseCase(rentPaymentRepository)
101: getRentPaymentSuggestionUseCase := rentPaymentUseCase.NewGetRentPaymentSuggestionUseCase(
109: listRoleUseCase := roleUsecase.NewListRolesUseCase(roleRepository)
116: createUserUseCase := userUseCase.NewCreateUserUseCase(userRepository, userCreatePolicy, passwordHasher)
117: getUserByIDUseCase := userUseCase.NewGetUserByIDUseCase(userRepository)
118: getUserByEmailUseCase := userUseCase.NewGetUserByEmailUseCase(userRepository)
119: listUsersUseCase := userUseCase.NewListUsersUseCase(userRepository)
120: searchUsersUseCase := userUseCase.NewSearchUsersUseCase(userRepository)
121: updateUserUseCase := userUseCase.NewUpdateUserUseCase(userRepository, userUpdatePolicy)
122: deleteUserUseCase := userUseCase.NewDeleteUserUseCase(userRepository)
123: loginUseCase := userUseCase.NewLoginUseCase(
143: createOwnerUseCase := ownerUseCase.NewCreateOwnerUseCase(ownerRepository, ownerCreatePolicy)
144: getOwnerByIDUseCase := ownerUseCase.NewGetOwnerByIDUseCase(ownerRepository)
145: getOwnerByDocumentNumberUseCase := ownerUseCase.NewGetOwnerByDocumentNumberUseCase(ownerRepository)
146: listOwnersUseCase := ownerUseCase.NewListOwnersUseCase(ownerRepository)
147: searchOwnersUseCase := ownerUseCase.NewSearchOwnersUseCase(ownerRepository)
148: updateOwnerUseCase := ownerUseCase.NewUpdateOwnerUseCase(ownerRepository, ownerUpdatePolicy)
160: listCountriesUseCase := locationUseCase.NewListCountriesUseCase(locationRepository)
161: listStatesByCountryUseCase := locationUseCase.NewListStatesByCountryUseCase(locationRepository)
168: createTenantUseCase := tenantUseCase.NewCreateTenantUseCase(tenantRepository, tenantCreatePolicy)
169: getTenantByIDUseCase := tenantUseCase.NewGetTenantByIDUseCase(tenantRepository)
170: getTenantByDocumentNumberUseCase := tenantUseCase.NewGetTenantByDocumentNumberUseCase(tenantRepository)
171: listTenantsUseCase := tenantUseCase.NewListTenantsUseCase(tenantRepository)
172: searchTenantsUseCase := tenantUseCase.NewSearchTenantsUseCase(tenantRepository)
173: updateTenantUseCase := tenantUseCase.NewUpdateTenantUseCase(tenantRepository, tenantUpdatePolicy)
186: createPropertyUseCase := propertyUseCase.NewCreatePropertyUseCase(propertyRepository, propertyCreatePolicy)
187: getPropertyByIDUseCase := propertyUseCase.NewGetPropertyByIDUseCase(propertyRepository)
188: getPropertyByCodeUseCase := propertyUseCase.NewGetPropertyByCodeUseCase(propertyRepository)
189: listPropertiesUseCase := propertyUseCase.NewListPropertiesUseCase(propertyRepository)
190: listPropertiesSummariesUseCase := propertyUseCase.NewListPropertiesSummariesUseCase(
194: searchPropertiesUseCase := propertyUseCase.NewSearchPropertiesUseCase(propertyRepository)
195: updatePropertyUseCase := propertyUseCase.NewUpdatePropertyUseCase(propertyRepository, propertyUpdatePolicy)
207: listPropertyTypesUseCase := propertyCatalogUseCase.NewListPropertyTypesUseCase(propertyCatalogRepository)
208: listPropertyStatusesUseCase := propertyCatalogUseCase.NewListPropertyStatusesUseCase(propertyCatalogRepository)
215: listContractStatusesUseCase := contractCatalogUseCase.NewListContractStatusesUseCase(contractCatalogRepository)
216: listInterestCalculationTypesUseCase := contractCatalogUseCase.NewListInterestCalculationTypesUseCase(
219: listRentAdjustmentTypesUseCase := contractCatalogUseCase.NewListRentAdjustmentTypesUseCase(
230: createRentalContractUseCase := rentalContractUseCase.NewCreateRentalContractUseCase(
234: updateRentalContractUseCase := rentalContractUseCase.NewUpdateRentalContractUseCase(
238: getRentalContractByIDUseCase := rentalContractUseCase.NewGetRentalContractByIDUseCase(rentalContractRepository)
239: listRentalContractsUseCase := rentalContractUseCase.NewListRentalContractsUseCase(
259: createInflationIndexUseCase := inflationIndexUseCase.NewCreateInflationIndexUseCase(inflationIndexRepository,
262: updateInflationIndexUseCase := inflationIndexUseCase.NewUpdateInflationIndexUseCase(
266: getInflationIndexByIDUseCase := inflationIndexUseCase.NewGetInflationIndexByIDUseCase(inflationIndexRepository)
267: listInflationIndexesUseCase := inflationIndexUseCase.NewListInflationIndexesUseCase(inflationIndexRepository)
292: if err != nil {
301: if err != nil {
```

### rent-manager-backend/internal/adapter/config/config.go

```text
43: func New() (*Container, error) {
44: if os.Getenv("APP_ENV") != "production" {
46: if err != nil {
```

### rent-manager-backend/internal/adapter/http/auth.go

```text
7: userUseCase "rent-manager-backend/internal/core/usecase/user"
11: loginUseCase *userUseCase.LoginUseCase
14: func NewAuthHandler(loginUseCase *userUseCase.LoginUseCase) *AuthHandler {
21: Email    string `json:"email" binding:"required,email"`
22: Password string `json:"password" binding:"required"`
26: Token string `json:"token"`
29: func (h *AuthHandler) Login(ctx *gin.Context) {
32: if err := ctx.ShouldBindJSON(&req); err != nil {
41: if err != nil {
```

### rent-manager-backend/internal/adapter/http/contract_catalog_response.go

```text
6: ID   int64  `json:"id"`
7: Code string `json:"code"`
8: Name string `json:"name"`
12: ID   int64  `json:"id"`
13: Code string `json:"code"`
14: Name string `json:"name"`
18: ID   int64  `json:"id"`
19: Code string `json:"code"`
20: Name string `json:"name"`
23: func newContractStatusResponse(status domain.ContractStatus) contractStatusResponse {
31: func newContractStatusesResponse(statuses []domain.ContractStatus) []contractStatusResponse {
41: func newInterestCalculationTypeResponse(
51: func newInterestCalculationTypesResponse(
63: func newRentAdjustmentTypeResponse(
73: func newRentAdjustmentTypesResponse(adjustmentTypes []domain.RentAdjustmentType) []rentAdjustmentTypeResponse {
```

### rent-manager-backend/internal/adapter/http/contract_catalog.go

```text
8: contractCatalogUseCase "rent-manager-backend/internal/core/usecase/contract_catalog"
12: listContractStatusesUseCase         *contractCatalogUseCase.ListContractStatusesUseCase
13: listInterestCalculationTypesUseCase *contractCatalogUseCase.ListInterestCalculationTypesUseCase
14: listRentAdjustmentTypesUseCase      *contractCatalogUseCase.ListRentAdjustmentTypesUseCase
17: func NewContractCatalogHandler(
18: listContractStatusesUseCase *contractCatalogUseCase.ListContractStatusesUseCase,
19: listInterestCalculationTypesUseCase *contractCatalogUseCase.ListInterestCalculationTypesUseCase,
20: listRentAdjustmentTypesUseCase *contractCatalogUseCase.ListRentAdjustmentTypesUseCase,
29: func (h *ContractCatalogHandler) ListContractStatuses(ctx *gin.Context) {
31: if err != nil {
40: func (h *ContractCatalogHandler) ListInterestCalculationTypes(ctx *gin.Context) {
42: if err != nil {
51: func (h *ContractCatalogHandler) ListRentAdjustmentTypes(ctx *gin.Context) {
53: if err != nil {
```

### rent-manager-backend/internal/adapter/http/error.go

```text
13: Code    string `json:"code"`
14: Message string `json:"message"`
17: func validationError(ctx *gin.Context, err error) {
24: func handleError(ctx *gin.Context, err error) {
27: if errors.As(err, &appErr) {
41: func statusFromErrorCode(code domain.ErrorCode) int {
43: case domain.ErrCodeDataNotFound, domain.ErrCodeUserNotFound, domain.ErrCodePropertyNotFound:
45: case domain.ErrCodeConflictingData, domain.ErrCodeUserAlreadyExists, domain.ErrCodePropertyAlreadyExists:
47: case domain.ErrCodeInvalidData,
54: case domain.ErrCodeOwnerNotFound:
56: case domain.ErrCodeOwnerAlreadyExists:
58: case domain.ErrCodeOwnerNil,
72: case domain.ErrCodeUnauthorized:
74: case domain.ErrCodeForbidden:
76: case domain.ErrCodeRentalContractNotFound:
78: case domain.ErrCodeRentalContractAlreadyExists:
80: case domain.ErrCodeRentalContractNil,
96: case domain.ErrCodeRentPaymentNotFound:
98: case domain.ErrCodeRentPaymentAlreadyExists:
100: case domain.ErrCodeRentPaymentNil,
```

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

### rent-manager-backend/internal/adapter/http/location_response.go

```text
10: ID        int64     `json:"id"`
11: Code      string    `json:"code"`
12: Name      string    `json:"name"`
13: CreatedAt time.Time `json:"created_at"`
14: UpdatedAt time.Time `json:"updated_at"`
18: ID        int64     `json:"id"`
19: CountryID int64     `json:"country_id"`
20: Code      string    `json:"code"`
21: Name      string    `json:"name"`
22: CreatedAt time.Time `json:"created_at"`
23: UpdatedAt time.Time `json:"updated_at"`
26: func newCountryResponse(country domain.Country) countryResponse {
36: func newCountriesResponse(countries []domain.Country) []countryResponse {
46: func newStateResponse(state domain.State) stateResponse {
57: func newStatesResponse(states []domain.State) []stateResponse {
```

### rent-manager-backend/internal/adapter/http/location.go

```text
8: locationUseCase "rent-manager-backend/internal/core/usecase/location"
12: listCountriesUseCase       *locationUseCase.ListCountriesUseCase
13: listStatesByCountryUseCase *locationUseCase.ListStatesByCountryUseCase
16: func NewLocationHandler(
17: listCountriesUseCase *locationUseCase.ListCountriesUseCase,
18: listStatesByCountryUseCase *locationUseCase.ListStatesByCountryUseCase,
26: func (h *LocationHandler) ListCountries(ctx *gin.Context) {
28: if err != nil {
37: func (h *LocationHandler) ListStatesByCountry(ctx *gin.Context) {
40: if err := ctx.ShouldBindUri(&req); err != nil {
49: if err != nil {
```

### rent-manager-backend/internal/adapter/http/middleware_auth.go

```text
19: func AuthMiddleware(tokenProvider port.TokenProvider) gin.HandlerFunc {
22: if header == "" {
28: if !strings.HasPrefix(header, bearerPrefix) {
35: if tokenValue == "" {
42: if err != nil {
```

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

### rent-manager-backend/internal/adapter/http/request.go

```text
4: Page  uint64 `form:"page" binding:"required,min=1"`
5: Limit uint64 `form:"limit" binding:"required,min=1,max=100"`
```

### rent-manager-backend/internal/adapter/http/response.go

```text
10: Data any `json:"data"`
13: func handleSuccess(ctx *gin.Context, data any) {
```

### rent-manager-backend/internal/adapter/http/role_response.go

```text
8: ID   int64  `json:"id"`
9: Name string `json:"name"`
12: func newRoleResponse(role *domain.Role) roleResponse {
19: func newRolesResponse(roles []domain.Role) []roleResponse {
```

### rent-manager-backend/internal/adapter/http/role.go

```text
9: listRolesUseCase *role.ListRolesUseCase
12: func NewRoleHandler(listRolesUseCase *role.ListRolesUseCase) *RoleHandler {
16: func (rh *RoleHandler) ListRoles(ctx *gin.Context) {
19: if err := ctx.ShouldBindQuery(&req); err != nil {
25: if err != nil {
```

### rent-manager-backend/internal/adapter/http/router.go

```text
21: func NewRouter(
37: if config.Env == "production" {
61: router.Use(
69: v1 := router.Group("/v1")
71: auth := v1.Group("/auth")
73: auth.POST("/login", authHandler.Login)
76: users := v1.Group("/users")
77: users.Use(authMiddleware)
79: users.POST("", userHandler.CreateUser)
80: users.GET("", userHandler.ListUsers)
81: users.GET("/search", userHandler.SearchUsers)
82: users.GET("/:id", userHandler.GetUserByID)
83: users.PUT("/:id", userHandler.UpdateUser)
84: users.DELETE("/:id", userHandler.DeleteUser)
86: roles := v1.Group("/roles")
87: roles.Use(authMiddleware)
89: roles.GET("", roleHandler.ListRoles)
91: owners := v1.Group("/owners")
92: owners.Use(authMiddleware)
94: owners.POST("", ownerHandler.CreateOwner)
95: owners.GET("", ownerHandler.ListOwners)
96: owners.GET("/search", ownerHandler.SearchOwners)
97: owners.GET("/:id", ownerHandler.GetOwnerByID)
98: owners.PUT("/:id", ownerHandler.UpdateOwner)
100: locations := v1.Group("/locations")
101: locations.Use(authMiddleware)
103: locations.GET("/countries", locationHandler.ListCountries)
104: locations.GET("/countries/:countryId/states", locationHandler.ListStatesByCountry)
106: tenants := v1.Group("/tenants")
107: tenants.Use(authMiddleware)
109: tenants.POST("", tenantHandler.CreateTenant)
110: tenants.GET("", tenantHandler.ListTenants)
111: tenants.GET("/search", tenantHandler.SearchTenants)
112: tenants.GET("/:id", tenantHandler.GetTenantByID)
113: tenants.PUT("/:id", tenantHandler.UpdateTenant)
115: properties := v1.Group("/properties")
116: properties.Use(authMiddleware)
118: properties.POST("", propertyHandler.CreateProperty)
119: properties.GET("", propertyHandler.ListProperties)
120: properties.GET("/search", propertyHandler.SearchProperties)
121: properties.GET("summaries", propertyHandler.ListPropertiesSummary)
122: properties.GET("/:id", propertyHandler.GetPropertyByID)
123: properties.PUT("/:id", propertyHandler.UpdateProperty)
125: propertyCatalogs := v1.Group("/property-catalogs")
126: propertyCatalogs.Use(authMiddleware)
128: propertyCatalogs.GET("/types", propertyCatalogHandler.ListPropertyTypes)
129: propertyCatalogs.GET("/statuses", propertyCatalogHandler.ListPropertyStatuses)
131: rentalContracts := v1.Group("/rental-contracts")
132: rentalContracts.Use(authMiddleware)
134: rentalContracts.POST("", rentalContractHandler.CreateRentalContract)
135: rentalContracts.GET("", rentalContractHandler.ListRentalContracts)
136: rentalContracts.GET("/:rentalContractId", rentalContractHandler.GetRentalContractByID)
137: rentalContracts.PUT("/:rentalContractId", rentalContractHandler.UpdateRentalContract)
138: rentalContracts.GET("/:rentalContractId/payment-schedules", rentPaymentHandler.GetRentPaymentSchedule)
139: rentalContracts.GET("/:rentalContractId/payment-suggestions", rentPaymentHandler.GetRentPaymentSuggestion)
140: rentalContracts.GET("/:rentalContractId/summary", rentPaymentHandler.GetRentalContractSummary)
142: contractCatalogs := v1.Group("/contract-catalogs")
143: contractCatalogs.Use(authMiddleware)
145: contractCatalogs.GET("/statuses", contractCatalogHandler.ListContractStatuses)
146: contractCatalogs.GET("/interest-calculation-types", contractCatalogHandler.ListInterestCalculationTypes)
147: contractCatalogs.GET("/rent-adjustment-types", contractCatalogHandler.ListRentAdjustmentTypes)
149: rentPayments := v1.Group("/rent-payments")
150: rentPayments.Use(authMiddleware)
152: rentPayments.POST("", rentPaymentHandler.CreateRentPayment)
153: rentPayments.GET("", rentPaymentHandler.ListRentPayments)
154: rentPayments.GET("/:id", rentPaymentHandler.GetRentPaymentByID)
155: rentPayments.PUT("/:id", rentPaymentHandler.UpdateRentPayment)
157: inflationIndexes := v1.Group("/inflation-indexes")
159: inflationIndexes.POST("", inflationIndexHandler.Create)
160: inflationIndexes.PUT("/:id", inflationIndexHandler.Update)
161: inflationIndexes.GET("/:id", inflationIndexHandler.GetByID)
162: inflationIndexes.GET("", inflationIndexHandler.List)
170: func (r *Router) Serve(listenAddr string) error {
```

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

### rent-manager-backend/internal/adapter/http/user_request.go

```text
6: Name     string `json:"name" binding:"required"`
7: Email    string `json:"email" binding:"required,email"`
8: Password string `json:"password" binding:"required"`
9: RoleID   int64  `json:"role_id" binding:"required,min=1"`
13: Name         string `json:"name" binding:"required"`
14: Email        string `json:"email" binding:"required,email"`
16: RoleID       int64 `json:"role_id" binding:"required,min=1"`
20: Value string `form:"value" binding:"required"`
21: Page  uint64 `form:"page" binding:"required,min=1"`
22: Limit uint64 `form:"limit" binding:"required,min=1,max=100"`
25: func newUserFromCreateRequest(req createUserRequest) *domain.User {
34: func newUserFromUpdateRequest(id int64, req updateUserRequest) *domain.User {
```

### rent-manager-backend/internal/adapter/http/user_response.go

```text
9: ID        int64     `json:"id"`
10: Name      string    `json:"name"`
11: Email     string    `json:"email"`
12: RoleID    int64     `json:"role_id"`
13: CreatedAt time.Time `json:"created_at"`
14: UpdatedAt time.Time `json:"updated_at"`
17: func newUserResponse(user *domain.User) userResponse {
28: func newUsersResponse(users []domain.User) []userResponse {
```

### rent-manager-backend/internal/adapter/http/user.go

```text
5: userUseCase "rent-manager-backend/internal/core/usecase/user"
12: createUserUseCase     *userUseCase.CreateUserUseCase
13: getUserByIDUseCase    *userUseCase.GetUserByIDUseCase
14: getUserByEmailUseCase *userUseCase.GetUserByEmailUseCase
15: listUsersUseCase      *userUseCase.ListUserUseCase
16: searchUsersUseCase    *userUseCase.SearchUserUseCase
17: updateUserUseCase     *userUseCase.UpdateUserUseCase
18: deleteUserUseCase     *userUseCase.DeleteUserUseCase
21: func NewUserHandler(
22: createUserUseCase *userUseCase.CreateUserUseCase,
23: getUserByIDUseCase *userUseCase.GetUserByIDUseCase,
24: getUserByEmailUseCase *userUseCase.GetUserByEmailUseCase,
25: listUsersUseCase *userUseCase.ListUserUseCase,
26: searchUsersUseCase *userUseCase.SearchUserUseCase,
27: updateUserUseCase *userUseCase.UpdateUserUseCase,
28: deleteUserUseCase *userUseCase.DeleteUserUseCase,
41: func (h *UserHandler) CreateUser(ctx *gin.Context) {
44: if err := ctx.ShouldBindJSON(&req); err != nil {
52: if err != nil {
60: func (h *UserHandler) GetUserByID(ctx *gin.Context) {
62: if err != nil {
68: if err != nil {
76: func (h *UserHandler) ListUsers(ctx *gin.Context) {
79: if err := ctx.ShouldBindQuery(&req); err != nil {
85: if err != nil {
93: func (h *UserHandler) SearchUsers(ctx *gin.Context) {
96: if err := ctx.ShouldBindQuery(&req); err != nil {
102: if err != nil {
110: func (h *UserHandler) UpdateUser(ctx *gin.Context) {
112: if err != nil {
119: if err := ctx.ShouldBindJSON(&req); err != nil {
127: if err != nil {
135: func (h *UserHandler) DeleteUser(ctx *gin.Context) {
137: if err != nil {
142: if err := h.deleteUserUseCase.DeleteUser(ctx, id); err != nil {
```

### rent-manager-backend/internal/adapter/security/bcrypt_password_hasher.go

```text
9: func NewBcryptPasswordHasher(cost int) *BcryptPasswordHasher {
15: func (h *BcryptPasswordHasher) Hash(password string) (string, error) {
17: if err != nil {
24: func (h *BcryptPasswordHasher) Compare(hash string, password string) error {
```

### rent-manager-backend/internal/adapter/security/jwt_token_provider.go

```text
17: func NewJWTTokenProvider(secretKey string, ttl time.Duration) *JWTTokenProvider {
24: func (p *JWTTokenProvider) GenerateToken(user *domain.User) (string, error) {
39: func (p *JWTTokenProvider) ValidateToken(tokenValue string) (*domain.TokenClaims, error) {
41: if token.Method != jwt.SigningMethodHS256 {
47: if err != nil {
51: if !token.Valid {
56: if !ok {
61: if !ok {
66: if !ok {
71: if !ok {
76: if !ok {
```

### rent-manager-backend/internal/adapter/store/mysql/db.go

```text
31: func New(ctx context.Context, config *config.DB) (*DB, error) {
51: if err != nil {
55: if err := db.PingContext(ctx); err != nil {
69: func (db *DB) Migrate() error {
71: if err != nil {
76: if err != nil {
81: if err != nil && !errors.Is(err, migrate.ErrNoChange) {
88: func (db *DB) ErrorCode(err error) string {
90: if errors.As(err, &mysqlErr) {
97: func (db *DB) Close() {
```

### rent-manager-backend/internal/adapter/store/mysql/migrations/000001_create_countries_table.up.sql

```text
1: CREATE TABLE countries (
2: cou_id BIGINT NOT NULL AUTO_INCREMENT,
3: cou_code VARCHAR(10) NOT NULL,
4: cou_name VARCHAR(150) NOT NULL,
5: cou_created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
6: cou_updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
8: PRIMARY KEY (cou_id),
9: UNIQUE KEY uk_countries_code (cou_code),
10: UNIQUE KEY uk_countries_name (cou_name)
```

### rent-manager-backend/internal/adapter/store/mysql/migrations/000002_create_states_table.up.sql

```text
1: CREATE TABLE states (
2: sta_id BIGINT NOT NULL AUTO_INCREMENT,
3: cou_id BIGINT NOT NULL,
4: sta_code VARCHAR(20) NULL,
5: sta_name VARCHAR(150) NOT NULL,
6: sta_created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
7: sta_updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
9: PRIMARY KEY (sta_id),
10: UNIQUE KEY uk_states_country_name (cou_id, sta_name),
11: UNIQUE KEY uk_states_country_code (cou_id, sta_code),
14: CONSTRAINT fk_states_country
15: FOREIGN KEY (cou_id) REFERENCES countries(cou_id)
18: INSERT INTO countries (cou_code, cou_name) VALUES ('AR','Argentina');
19: INSERT INTO states (cou_id,sta_code,sta_name)
```

### rent-manager-backend/internal/adapter/store/mysql/migrations/000003_create_roles_table.up.sql

```text
1: CREATE TABLE roles (
2: rol_id BIGINT NOT NULL AUTO_INCREMENT,
3: rol_code VARCHAR(50) NOT NULL,
4: rol_name VARCHAR(100) NOT NULL,
5: rol_created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
7: PRIMARY KEY (rol_id),
8: UNIQUE KEY uk_roles_code (rol_code)
11: INSERT INTO roles (rol_code, rol_name)
```

### rent-manager-backend/internal/adapter/store/mysql/migrations/000004_create_users_table.up.sql

```text
1: CREATE TABLE users (
2: usr_id BIGINT NOT NULL AUTO_INCREMENT,
3: rol_id BIGINT NOT NULL,
5: usr_name VARCHAR(255) NOT NULL,
6: usr_email VARCHAR(255) NOT NULL,
7: usr_password_hash VARCHAR(255) NOT NULL,
9: usr_created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
10: usr_updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
12: PRIMARY KEY (usr_id),
13: UNIQUE KEY uk_users_email (usr_email),
16: CONSTRAINT fk_users_role
17: FOREIGN KEY (rol_id) REFERENCES roles(rol_id)
20: INSERT INTO users (usr_id, rol_id, usr_name, usr_email, usr_password_hash, usr_created_at, usr_updated_at)
```

### rent-manager-backend/internal/adapter/store/mysql/migrations/000005_create_owners_table.up.sql

```text
1: CREATE TABLE owners (
2: own_id BIGINT NOT NULL AUTO_INCREMENT,
3: own_name VARCHAR(255) NOT NULL,
4: own_email VARCHAR(255) NULL,
5: own_phone VARCHAR(50) NULL,
6: own_document_number VARCHAR(100) NULL,
7: own_created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
8: own_updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
10: PRIMARY KEY (own_id),
11: UNIQUE KEY uk_owners_document_number (own_document_number)
```

### rent-manager-backend/internal/adapter/store/mysql/migrations/000006_create_tenants_table.up.sql

```text
1: CREATE TABLE tenants (
2: ten_id BIGINT NOT NULL AUTO_INCREMENT,
4: cou_id BIGINT NULL,
5: sta_id BIGINT NULL,
7: ten_name VARCHAR(255) NOT NULL,
8: ten_email VARCHAR(255) NULL,
9: ten_phone VARCHAR(50) NULL,
10: ten_document_number VARCHAR(100) NULL,
12: ten_city VARCHAR(150) NULL,
13: ten_street VARCHAR(255) NULL,
14: ten_street_number VARCHAR(50) NULL,
15: ten_floor VARCHAR(50) NULL,
16: ten_apartment VARCHAR(50) NULL,
17: ten_postal_code VARCHAR(20) NULL,
19: ten_created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
20: ten_updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
22: PRIMARY KEY (ten_id),
24: UNIQUE KEY uk_tenants_document_number (ten_document_number),
29: CONSTRAINT fk_tenants_country
30: FOREIGN KEY (cou_id) REFERENCES countries(cou_id),
32: CONSTRAINT fk_tenants_state
33: FOREIGN KEY (sta_id) REFERENCES states(sta_id)
```

### rent-manager-backend/internal/adapter/store/mysql/migrations/000007_create_property_catalogs.up.sql

```text
1: CREATE TABLE property_types (
2: pty_id BIGINT NOT NULL AUTO_INCREMENT,
3: pty_code VARCHAR(50) NOT NULL,
4: pty_name VARCHAR(100) NOT NULL,
6: PRIMARY KEY (pty_id),
7: UNIQUE KEY uk_property_types_code (pty_code)
10: INSERT INTO property_types (pty_code, pty_name)
17: CREATE TABLE property_statuses (
18: pst_id BIGINT NOT NULL AUTO_INCREMENT,
19: pst_code VARCHAR(50) NOT NULL,
20: pst_name VARCHAR(100) NOT NULL,
22: PRIMARY KEY (pst_id),
23: UNIQUE KEY uk_property_statuses_code (pst_code)
26: INSERT INTO property_statuses (pst_code, pst_name)
```

### rent-manager-backend/internal/adapter/store/mysql/migrations/000008_create_properties_table.up.sql

```text
1: CREATE TABLE properties (
2: pro_id BIGINT NOT NULL AUTO_INCREMENT,
3: own_id BIGINT NOT NULL,
4: pty_id BIGINT NOT NULL,
5: pst_id BIGINT NOT NULL,
6: cou_id BIGINT NOT NULL,
7: sta_id BIGINT NOT NULL,
9: pro_code VARCHAR(100) NOT NULL,
10: pro_title VARCHAR(255) NOT NULL,
11: pro_description TEXT NULL,
13: pro_street VARCHAR(255) NOT NULL,
14: pro_street_number VARCHAR(50) NULL,
15: pro_floor VARCHAR(50) NULL,
16: pro_apartment VARCHAR(50) NULL,
17: pro_city VARCHAR(100) NOT NULL,
18: pro_postal_code VARCHAR(50) NULL,
20: pro_created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
21: pro_updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
23: PRIMARY KEY (pro_id),
24: UNIQUE KEY uk_properties_code (pro_code),
32: CONSTRAINT fk_properties_owner
33: FOREIGN KEY (own_id) REFERENCES owners(own_id),
35: CONSTRAINT fk_properties_type
36: FOREIGN KEY (pty_id) REFERENCES property_types(pty_id),
38: CONSTRAINT fk_properties_status
39: FOREIGN KEY (pst_id) REFERENCES property_statuses(pst_id),
41: CONSTRAINT fk_properties_country
42: FOREIGN KEY (cou_id) REFERENCES countries(cou_id),
44: CONSTRAINT fk_properties_state
45: FOREIGN KEY (sta_id) REFERENCES states(sta_id)
```

### rent-manager-backend/internal/adapter/store/mysql/migrations/000009_create_contract_statuses_table.up.sql

```text
1: CREATE TABLE contract_statuses (
2: cst_id BIGINT NOT NULL AUTO_INCREMENT,
3: cst_code VARCHAR(50) NOT NULL,
4: cst_name VARCHAR(100) NOT NULL,
6: PRIMARY KEY (cst_id),
7: UNIQUE KEY uk_contract_statuses_code (cst_code)
10: INSERT INTO contract_statuses (cst_code, cst_name)
```

### rent-manager-backend/internal/adapter/store/mysql/migrations/000010_create_rental_contracts_table.up.sql

```text
1: CREATE TABLE interest_calculation_types (
2: ict_id BIGINT NOT NULL AUTO_INCREMENT,
3: ict_code VARCHAR(50) NOT NULL,
4: ict_name VARCHAR(100) NOT NULL,
6: PRIMARY KEY (ict_id),
7: UNIQUE KEY uk_interest_calculation_types_code (ict_code)
10: INSERT INTO interest_calculation_types (ict_code, ict_name)
16: CREATE TABLE rent_adjustment_types (
17: rat_id BIGINT NOT NULL AUTO_INCREMENT,
18: rat_code VARCHAR(50) NOT NULL,
19: rat_name VARCHAR(100) NOT NULL,
21: PRIMARY KEY (rat_id),
22: UNIQUE KEY uk_rent_adjustment_types_code (rat_code)
25: INSERT INTO rent_adjustment_types (rat_code, rat_name)
29: CREATE TABLE rental_contracts (
30: rco_id BIGINT NOT NULL AUTO_INCREMENT,
31: pro_id BIGINT NOT NULL,
32: ten_id BIGINT NOT NULL,
33: cst_id BIGINT NOT NULL,
34: ict_id BIGINT NOT NULL,
36: rco_start_date DATE NOT NULL,
37: rco_end_date DATE NOT NULL,
38: rco_total_payments INT NOT NULL,
39: rco_monthly_amount DECIMAL(12,2) NOT NULL,
40: rco_deposit_amount DECIMAL(12,2) NOT NULL DEFAULT 0,
41: rco_currency VARCHAR(10) NOT NULL DEFAULT 'ARS',
43: rco_due_day INT NOT NULL DEFAULT 10,
44: rco_daily_interest_percentage DECIMAL(5,2) NOT NULL DEFAULT 0,
46: rco_notes TEXT NULL,
48: rat_id BIGINT NOT NULL,
49: rco_adjustment_frequency_months INT NOT NULL DEFAULT 4,
51: rco_created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
52: rco_updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
54: PRIMARY KEY (rco_id),
62: CONSTRAINT fk_rental_contracts_adjustment_type
63: FOREIGN KEY (rat_id) REFERENCES rent_adjustment_types(rat_id),
65: CONSTRAINT fk_rental_contracts_property
66: FOREIGN KEY (pro_id) REFERENCES properties(pro_id),
68: CONSTRAINT fk_rental_contracts_tenant
69: FOREIGN KEY (ten_id) REFERENCES tenants(ten_id),
71: CONSTRAINT fk_rental_contracts_status
72: FOREIGN KEY (cst_id) REFERENCES contract_statuses(cst_id),
74: CONSTRAINT fk_rental_contracts_interest_type
75: FOREIGN KEY (ict_id) REFERENCES interest_calculation_types(ict_id)
```

### rent-manager-backend/internal/adapter/store/mysql/migrations/000011_create_payment_statuses_table.up.sql

```text
1: CREATE TABLE payment_statuses (
2: pas_id BIGINT NOT NULL AUTO_INCREMENT,
3: pas_code VARCHAR(50) NOT NULL,
4: pas_name VARCHAR(100) NOT NULL,
6: PRIMARY KEY (pas_id),
7: UNIQUE KEY uk_payment_statuses_code (pas_code)
10: INSERT INTO payment_statuses (pas_code, pas_name)
```

### rent-manager-backend/internal/adapter/store/mysql/migrations/000012_create_rent_payments_table.up.sql

```text
1: CREATE TABLE rent_payments (
2: rpa_id BIGINT NOT NULL AUTO_INCREMENT,
3: rco_id BIGINT NOT NULL,
5: rpa_period DATE NOT NULL,
6: rpa_due_date DATE NOT NULL,
7: rpa_payment_date DATE NULL,
9: rpa_base_amount DECIMAL(12,2) NOT NULL,
10: rpa_suggested_adjustment_percentage DECIMAL(8,4) NULL,
11: rpa_applied_adjustment_percentage DECIMAL(8,4) NULL,
13: rpa_suggested_interest_amount DECIMAL(12,2) NULL,
14: rpa_applied_interest_amount DECIMAL(12,2) NULL,
16: rpa_total_amount DECIMAL(12,2) NOT NULL,
17: rpa_paid_amount DECIMAL(12,2) NULL,
19: rpa_is_paid BOOLEAN NOT NULL DEFAULT FALSE,
20: rpa_notes TEXT NULL,
22: rpa_created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
23: rpa_updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
25: PRIMARY KEY (rpa_id),
27: UNIQUE KEY uk_rent_payments_contract_period (rco_id, rpa_period),
34: CONSTRAINT fk_rent_payments_contract
35: FOREIGN KEY (rco_id) REFERENCES rental_contracts(rco_id)
```

### rent-manager-backend/internal/adapter/store/mysql/migrations/000013_create_inflation_indexes.up.sql

```text
1: CREATE TABLE inflation_indexes (
2: ixi_id BIGINT AUTO_INCREMENT PRIMARY KEY,
3: ixi_period DATE NOT NULL,
4: ixi_percentage DECIMAL(10,4) NOT NULL,
5: ixi_source VARCHAR(100) NULL,
6: ixi_notes TEXT NULL,
7: ixi_created_at DATETIME NOT NULL,
8: ixi_updated_at DATETIME NOT NULL,
10: UNIQUE KEY uk_inflation_indexes_period (ixi_period)
```

### rent-manager-backend/internal/adapter/store/mysql/repository/contract_catalog.go

```text
15: func NewContractCatalogRepository(db *mysql.DB) *ContractCatalogRepository {
21: func (ccr *ContractCatalogRepository) ListContractStatuses(
34: if err != nil {
40: if err != nil {
51: if err := rows.Scan(
63: if err := rows.Err(); err != nil {
71: func (ccr *ContractCatalogRepository) ListInterestCalculationTypes(
84: if err != nil {
90: if err != nil {
101: if err := rows.Scan(
113: if err := rows.Err(); err != nil {
121: func (ccr *ContractCatalogRepository) ListRentAdjustmentTypes(
134: if err != nil {
140: if err != nil {
151: if err := rows.Scan(
163: if err := rows.Err(); err != nil {
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

### rent-manager-backend/internal/adapter/store/mysql/repository/location.go

```text
15: func NewLocationRepository(db *mysql.DB) *LocationRepository {
21: func (lr *LocationRepository) ListCountries(ctx context.Context) ([]domain.Country, error) {
34: if err != nil {
40: if err != nil {
51: if err := rows.Scan(
65: if err := rows.Err(); err != nil {
73: func (lr *LocationRepository) ListStatesByCountry(ctx context.Context, countryID int64) ([]domain.State, error) {
88: if err != nil {
94: if err != nil {
105: if err := rows.Scan(
120: if err := rows.Err(); err != nil {
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

### rent-manager-backend/internal/adapter/store/mysql/repository/role.go

```text
13: func NewRoleRepository(db *mysql.DB) *RoleRepository {
17: func (rr *RoleRepository) ListRoles(ctx context.Context, page uint64, limit uint64) ([]domain.Role, error) {
18: if page == 0 {
22: if limit == 0 {
41: if err != nil {
46: if err != nil {
62: if err != nil {
69: if err := rows.Err(); err != nil {
```

### rent-manager-backend/internal/adapter/store/mysql/repository/shared/shared.go

```text
3: func NullInt64(value int64) any {
4: if value == 0 {
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

### rent-manager-backend/internal/adapter/store/mysql/repository/user.go

```text
18: func NewUserRepository(db *mysql.DB) *UserRepository {
22: func (ur *UserRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
39: if err != nil {
45: if errExe != nil {
47: if errCode := ur.db.ErrorCode(errExe); errCode == "1062" {
54: if errLastInsert != nil {
62: func (ur *UserRepository) GetUserByID(ctx context.Context, id int64) (*domain.User, error) {
70: if err != nil {
84: if err != nil {
86: if errors.Is(err, sql.ErrNoRows) {
95: func (ur *UserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
103: if errSql != nil {
117: if errQuery != nil {
119: if errors.Is(errQuery, sql.ErrNoRows) {
128: func (ur *UserRepository) ListUsers(ctx context.Context, page uint64, limit uint64) ([]domain.User, error) {
129: if page == 0 {
133: if limit == 0 {
147: func (ur *UserRepository) SearchUsersByNameOrEmail(
153: if page == 0 {
157: if limit == 0 {
178: func (ur *UserRepository) PathUser(ctx context.Context, user *domain.User) (*domain.User, error) {
179: queryBuilder := ur.db.QueryBuilder.Update("users")
181: if user.RoleID != 0 {
185: if user.Name != "" {
189: if user.Email != "" {
193: if user.PasswordHash != "" {
196: queryBuilder = queryBuilder.Where(sq.Eq{"usr_id": user.ID})
199: if err != nil {
204: if errExe != nil {
206: if errCode := ur.db.ErrorCode(errExe); errCode == "1062" {
213: if errRow != nil {
218: if rowsAffected == 0 {
226: func (ur *UserRepository) DeleteUser(ctx context.Context, id int64) error {
232: if errSql != nil {
238: if err != nil {
244: if errRows != nil {
249: if rowsAffected == 0 {
250: return domain.ErrDataNotFound
256: func (ur *UserRepository) findUsers(ctx context.Context, query sq.SelectBuilder) ([]domain.User, error) {
260: if err != nil {
266: if errQuery != nil {
284: if err != nil {
292: if err := rows.Err(); err != nil {
300: func (ur *UserRepository) baseUserQuery() sq.SelectBuilder {
309: ).From("users")
```

### rent-manager-backend/internal/core/domain/error.go

```text
13: func (e *AppError) Error() string {
14: if e == nil {
18: if e.Cause != nil {
25: func (e *AppError) Unwrap() error {
29: func NewAppError(code ErrorCode, description string) *AppError {
36: func WrapAppError(code ErrorCode, description string, cause error) *AppError {
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

### rent-manager-backend/internal/core/policy/user/user_create_policy.go

```text
16: func NewUserCreatePolicy(userRepository port.UserRepository) *UserCreatePolicy {
22: func (userPolicy *UserCreatePolicy) Execute(ctx context.Context, user *domain.User) error {
23: if user == nil {
24: return domain.ErrUserNil
30: if user.Name == "" {
31: return domain.ErrUserNameEmpty
34: if user.Email == "" {
35: return domain.ErrUserEmailEmpty
38: if !shared.IsValidEmail(user.Email) {
39: return domain.ErrUserEmailInvalid
42: if user.PasswordHash == "" {
43: return domain.ErrUserPasswordEmpty
47: if err != nil {
49: if errors.As(err, &appErr) && appErr.Code == domain.ErrCodeDataNotFound {
60: if existingUser != nil {
61: return domain.ErrUserAlreadyExists
```

### rent-manager-backend/internal/core/policy/user/user_update_policy.go

```text
16: func NewUserUpdatePolicy(userRepository port.UserRepository) *UserUpdatePolicy {
22: func (userPolicy *UserUpdatePolicy) Execute(ctx context.Context, user *domain.User) error {
23: if user == nil {
24: return domain.ErrUserNil
27: if user.ID == 0 {
28: return domain.ErrUserNotFound
31: if user.Name == "" {
33: return domain.ErrUserNameEmpty
36: if user.Email != "" {
40: if !shared.IsValidEmail(user.Email) {
41: return domain.ErrUserEmailInvalid
44: if user.PasswordHash == "" {
49: if err != nil {
51: if errors.As(err, &appErr) && appErr.Code == domain.ErrCodeDataNotFound {
62: if existingUser != nil && existingUser.ID != user.ID {
63: return domain.ErrUserAlreadyExists
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

### rent-manager-backend/internal/core/usecase/contract_catalog/list_contract_statuses_use_case.go

```text
11: type ListContractStatusesUseCase struct {
15: func NewListContractStatusesUseCase(
17: ) *ListContractStatusesUseCase {
23: func (uc *ListContractStatusesUseCase) ListContractStatuses(
27: if err != nil {
```

### rent-manager-backend/internal/core/usecase/contract_catalog/list_interest_calculation_types_use_case.go

```text
11: type ListInterestCalculationTypesUseCase struct {
15: func NewListInterestCalculationTypesUseCase(
17: ) *ListInterestCalculationTypesUseCase {
23: func (uc *ListInterestCalculationTypesUseCase) ListInterestCalculationTypes(
27: if err != nil {
```

### rent-manager-backend/internal/core/usecase/contract_catalog/list_rent_adjustment_types_use_case.go

```text
11: type ListRentAdjustmentTypesUseCase struct {
15: func NewListRentAdjustmentTypesUseCase(
17: ) *ListRentAdjustmentTypesUseCase {
23: func (uc *ListRentAdjustmentTypesUseCase) ListRentAdjustmentTypes(
27: if err != nil {
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

### rent-manager-backend/internal/core/usecase/location/list_countries_use_case.go

```text
11: type ListCountriesUseCase struct {
15: func NewListCountriesUseCase(
17: ) *ListCountriesUseCase {
23: func (uc *ListCountriesUseCase) ListCountries(
27: if err != nil {
```

### rent-manager-backend/internal/core/usecase/location/list_states_by_country_use_case.go

```text
11: type ListStatesByCountryUseCase struct {
15: func NewListStatesByCountryUseCase(
17: ) *ListStatesByCountryUseCase {
23: func (uc *ListStatesByCountryUseCase) ListStatesByCountry(
31: if err != nil {
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

### rent-manager-backend/internal/core/usecase/role/list_roles_use_case.go

```text
9: type ListRolesUseCase struct {
13: func NewListRolesUseCase(roleRepository port.RoleRepository) *ListRolesUseCase {
17: func (lr *ListRolesUseCase) ListRoles(ctx context.Context, page uint64, limit uint64) ([]domain.Role, error) {
19: if err != nil {
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

### rent-manager-backend/internal/core/usecase/user/create_user_use_case.go

```text
10: type CreateUserUseCase struct {
16: func NewCreateUserUseCase(
20: ) *CreateUserUseCase {
28: func (uc *CreateUserUseCase) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
29: if err := uc.userCreatePolicy.Execute(ctx, user); err != nil {
34: if err != nil {
41: if err != nil {
```

### rent-manager-backend/internal/core/usecase/user/delete_user_use_case.go

```text
8: type DeleteUserUseCase struct {
12: func NewDeleteUserUseCase(userRepository port.UserRepository) *DeleteUserUseCase {
18: func (uc *DeleteUserUseCase) DeleteUser(ctx context.Context, id int64) error {
```

### rent-manager-backend/internal/core/usecase/user/get_user_by_email_use_case.go

```text
9: type GetUserByEmailUseCase struct {
13: func NewGetUserByEmailUseCase(
15: ) *GetUserByEmailUseCase {
21: func (uc *GetUserByEmailUseCase) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
23: if err != nil {
```

### rent-manager-backend/internal/core/usecase/user/get_user_by_id_use_case.go

```text
9: type GetUserByIDUseCase struct {
13: func NewGetUserByIDUseCase(
15: ) *GetUserByIDUseCase {
21: func (uc *GetUserByIDUseCase) GetUserByID(ctx context.Context, id int64) (*domain.User, error) {
23: if err != nil {
```

### rent-manager-backend/internal/core/usecase/user/list_user_use_case.go

```text
9: type ListUserUseCase struct {
13: func NewListUsersUseCase(
15: ) *ListUserUseCase {
21: func (uc *ListUserUseCase) ListUsers(ctx context.Context, page uint64, limit uint64) ([]domain.User, error) {
23: if err != nil {
```

### rent-manager-backend/internal/core/usecase/user/login_use_case.go

```text
10: type LoginUseCase struct {
16: func NewLoginUseCase(
20: ) *LoginUseCase {
28: func (uc *LoginUseCase) Login(ctx context.Context, login *domain.Login) (string, error) {
29: if login == nil {
34: if err != nil {
38: if err := uc.passwordHasher.Compare(user.PasswordHash, login.Password); err != nil {
43: if err != nil {
```

### rent-manager-backend/internal/core/usecase/user/search_user_use_case.go

```text
9: type SearchUserUseCase struct {
13: func NewSearchUsersUseCase(
15: ) *SearchUserUseCase {
21: func (uc *SearchUserUseCase) SearchUsersByNameOrEmail(
28: if err != nil {
```

### rent-manager-backend/internal/core/usecase/user/update_user_use_case.go

```text
10: type UpdateUserUseCase struct {
15: func NewUpdateUserUseCase(userRepository port.UserRepository, policy *userPolicy.UserUpdatePolicy) *UpdateUserUseCase {
19: func (uc *UpdateUserUseCase) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
21: if err != nil {
25: if err != nil {
```
