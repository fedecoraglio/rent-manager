# Source and operation coverage

Reviewed: 2026-11-13. Resolve paths from the repository root. Use this map to open the implementation; excerpts are navigation indexes, not replacements for complete source bodies. Line numbers refer to this review. Do not interpret internal operations as public endpoints.

## Contents

- File inventory
- Operations, data contracts, and rule locations

## File inventory

- `rent-manager-backend/internal/adapter/http/contract_catalog_response.go`
- `rent-manager-backend/internal/adapter/http/contract_catalog.go`
- `rent-manager-backend/internal/adapter/http/location_request.go`
- `rent-manager-backend/internal/adapter/http/location_response.go`
- `rent-manager-backend/internal/adapter/http/location.go`
- `rent-manager-backend/internal/adapter/http/property_catalog_request.go`
- `rent-manager-backend/internal/adapter/http/property_catalog_response.go`
- `rent-manager-backend/internal/adapter/http/property_catalog.go`
- `rent-manager-backend/internal/adapter/http/role_response.go`
- `rent-manager-backend/internal/adapter/http/role.go`
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
- `rent-manager-backend/internal/adapter/store/mysql/repository/location.go`
- `rent-manager-backend/internal/adapter/store/mysql/repository/property_catalog.go`
- `rent-manager-backend/internal/adapter/store/mysql/repository/role.go`
- `rent-manager-backend/internal/core/domain/error_catalog.go`
- `rent-manager-backend/internal/core/domain/role.go`
- `rent-manager-backend/internal/core/port/contract_catalog.go`
- `rent-manager-backend/internal/core/port/location.go`
- `rent-manager-backend/internal/core/port/property_catalog.go`
- `rent-manager-backend/internal/core/port/role.go`
- `rent-manager-backend/internal/core/usecase/contract_catalog/list_contract_statuses_use_case.go`
- `rent-manager-backend/internal/core/usecase/contract_catalog/list_interest_calculation_types_use_case.go`
- `rent-manager-backend/internal/core/usecase/contract_catalog/list_rent_adjustment_types_use_case.go`
- `rent-manager-backend/internal/core/usecase/contract_catalog/rent_contract_catalog.go`
- `rent-manager-backend/internal/core/usecase/location/list_countries_use_case.go`
- `rent-manager-backend/internal/core/usecase/location/list_states_by_country_use_case.go`
- `rent-manager-backend/internal/core/usecase/property_catalog/list_property_statuses_use_case.go`
- `rent-manager-backend/internal/core/usecase/property_catalog/list_property_types_use_case.go`
- `rent-manager-backend/internal/core/usecase/role/list_roles_use_case.go`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/contract-catalog/contract-catalog-api.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/contract-catalog/contract-catalog.model.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/contract-catalog/contract-catalog.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/location/location-api.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/location/location.model.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/location/location.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/property-catalog/property-catalog-api.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/property-catalog/property-catalog.model.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/property-catalog/property-catalog.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/role/role-api.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/role/role.model.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/role/role.service.ts`

## Operations, data contracts, and rule locations

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

### rent-manager-backend/internal/core/usecase/role/list_roles_use_case.go

```text
9: type ListRolesUseCase struct {
13: func NewListRolesUseCase(roleRepository port.RoleRepository) *ListRolesUseCase {
17: func (lr *ListRolesUseCase) ListRoles(ctx context.Context, page uint64, limit uint64) ([]domain.Role, error) {
19: if err != nil {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/core/contract-catalog/contract-catalog-api.service.ts

```text
20: return this.apiUrl.build('/contract-catalogs');
23: listContractStatuses$(): Observable<ApiResponse<ContractStatus[]>> {
24: return this.http
29: listInterestCalculationTypes$(): Observable<ApiResponse<InterestCalculationType[]>> {
30: return this.http
35: listRentAdjustmentTypes$(): Observable<ApiResponse<RentAdjustmentType[]>> {
36: return this.http
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/core/contract-catalog/contract-catalog.service.ts

```text
24: listContractStatuses$(): Observable<ContractStatus[]> {
29: catchError((e: HttpErrorResponse) => {
36: listInterestCalculationTypes$(): Observable<InterestCalculationType[]> {
41: catchError((e: HttpErrorResponse) => {
48: listRentAdjustmentTypes$(): Observable<RentAdjustmentType[]> {
53: catchError((e: HttpErrorResponse) => {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/core/location/location-api.service.ts

```text
15: return this.apiUrl.build('/locations');
18: listCountries$(): Observable<ApiResponse<Country[]>> {
19: return this.http
24: listStates$(countryId: number): Observable<ApiResponse<State[]>> {
25: return this.http
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/core/location/location.service.ts

```text
18: listCountries$(): Observable<Country[]> {
21: tap((apiResponse) => {
25: catchError((e: HttpErrorResponse) => {
32: listStates$(countryId: number): Observable<State[]> {
35: tap((apiResponse) => {
39: catchError((e: HttpErrorResponse) => {
46: getCountryById(id: number): Country | null {
50: getStateById(id: number): State | null {
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

### rent-manager-frontend/projects/rent-manager-frontend/src/app/core/role/role-api.service.ts

```text
15: return this.apiUrl.build('/roles');
18: listRoles$(page = 1, limit = 10): Observable<ApiResponse<Role[]>> {
20: return this.http
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/core/role/role.service.ts

```text
19: listRoles$(page = 1, limit = 10): Observable<Role[]> {
22: tap((apiResponse) => {
26: catchError((e: HttpErrorResponse) => {
33: getRoleById(id: number): Role | null {
37: getRoleNameById(id: number): string {
```
