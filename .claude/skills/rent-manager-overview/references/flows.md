# Complete operation matrix

Reviewed on 2026-11-13. Route source: rent-manager-backend/internal/adapter/http/router.go. Each operation inherits /v1 and its group below. Internal use cases do not imply additional HTTP routes.

## Go routes

| Method | Route | Handler |
| --- | --- | --- |
| POST | `/v1/auth/login` | `authHandler.Login` |
| POST | `/v1/users` | `userHandler.CreateUser` |
| GET | `/v1/users` | `userHandler.ListUsers` |
| GET | `/v1/users/search` | `userHandler.SearchUsers` |
| GET | `/v1/users/:id` | `userHandler.GetUserByID` |
| PUT | `/v1/users/:id` | `userHandler.UpdateUser` |
| DELETE | `/v1/users/:id` | `userHandler.DeleteUser` |
| GET | `/v1/roles` | `roleHandler.ListRoles` |
| POST | `/v1/owners` | `ownerHandler.CreateOwner` |
| GET | `/v1/owners` | `ownerHandler.ListOwners` |
| GET | `/v1/owners/search` | `ownerHandler.SearchOwners` |
| GET | `/v1/owners/:id` | `ownerHandler.GetOwnerByID` |
| PUT | `/v1/owners/:id` | `ownerHandler.UpdateOwner` |
| GET | `/v1/locations/countries` | `locationHandler.ListCountries` |
| GET | `/v1/locations/countries/:countryId/states` | `locationHandler.ListStatesByCountry` |
| POST | `/v1/tenants` | `tenantHandler.CreateTenant` |
| GET | `/v1/tenants` | `tenantHandler.ListTenants` |
| GET | `/v1/tenants/search` | `tenantHandler.SearchTenants` |
| GET | `/v1/tenants/:id` | `tenantHandler.GetTenantByID` |
| PUT | `/v1/tenants/:id` | `tenantHandler.UpdateTenant` |
| POST | `/v1/properties` | `propertyHandler.CreateProperty` |
| GET | `/v1/properties` | `propertyHandler.ListProperties` |
| GET | `/v1/properties/search` | `propertyHandler.SearchProperties` |
| GET | `/v1/properties/summaries` | `propertyHandler.ListPropertiesSummary` |
| GET | `/v1/properties/:id` | `propertyHandler.GetPropertyByID` |
| PUT | `/v1/properties/:id` | `propertyHandler.UpdateProperty` |
| GET | `/v1/property-catalogs/types` | `propertyCatalogHandler.ListPropertyTypes` |
| GET | `/v1/property-catalogs/statuses` | `propertyCatalogHandler.ListPropertyStatuses` |
| POST | `/v1/rental-contracts` | `rentalContractHandler.CreateRentalContract` |
| GET | `/v1/rental-contracts` | `rentalContractHandler.ListRentalContracts` |
| GET | `/v1/rental-contracts/:rentalContractId` | `rentalContractHandler.GetRentalContractByID` |
| PUT | `/v1/rental-contracts/:rentalContractId` | `rentalContractHandler.UpdateRentalContract` |
| GET | `/v1/rental-contracts/:rentalContractId/payment-schedules` | `rentPaymentHandler.GetRentPaymentSchedule` |
| GET | `/v1/rental-contracts/:rentalContractId/payment-suggestions` | `rentPaymentHandler.GetRentPaymentSuggestion` |
| GET | `/v1/rental-contracts/:rentalContractId/summary` | `rentPaymentHandler.GetRentalContractSummary` |
| GET | `/v1/contract-catalogs/statuses` | `contractCatalogHandler.ListContractStatuses` |
| GET | `/v1/contract-catalogs/interest-calculation-types` | `contractCatalogHandler.ListInterestCalculationTypes` |
| GET | `/v1/contract-catalogs/rent-adjustment-types` | `contractCatalogHandler.ListRentAdjustmentTypes` |
| POST | `/v1/rent-payments` | `rentPaymentHandler.CreateRentPayment` |
| GET | `/v1/rent-payments` | `rentPaymentHandler.ListRentPayments` |
| GET | `/v1/rent-payments/:id` | `rentPaymentHandler.GetRentPaymentByID` |
| PUT | `/v1/rent-payments/:id` | `rentPaymentHandler.UpdateRentPayment` |
| POST | `/v1/inflation-indexes` | `inflationIndexHandler.Create` |
| PUT | `/v1/inflation-indexes/:id` | `inflationIndexHandler.Update` |
| GET | `/v1/inflation-indexes/:id` | `inflationIndexHandler.GetByID` |
| GET | `/v1/inflation-indexes` | `inflationIndexHandler.List` |

Login and inflation do not use AuthMiddleware. The other routes in this table do.

## AI routes

| Method | Route | Skill |
| --- | --- | --- |
| POST | `/properties/{propertyId}/contracts/{contractId}/documents` | rent-manager-documents |
| GET | `/properties/{propertyId}/contracts/{contractId}/documents` | rent-manager-documents |
| GET | `/properties/{propertyId}/contracts/{contractId}/documents/{documentId}` | rent-manager-documents |
| DELETE | `/properties/{propertyId}/contracts/{contractId}/documents/{documentId}` | rent-manager-documents |
| POST | `/properties/{propertyId}/contracts/{contractId}/chat` | rent-manager-rag |

Actuator: GET /actuator/health, /actuator/health/readiness, and /actuator/health/liveness. AI does not implement authentication in the reviewed code.

## All Go use cases

| Domain | File |
| --- | --- |
| contract_catalog | `rent-manager-backend/internal/core/usecase/contract_catalog/list_contract_statuses_use_case.go` |
| contract_catalog | `rent-manager-backend/internal/core/usecase/contract_catalog/list_interest_calculation_types_use_case.go` |
| contract_catalog | `rent-manager-backend/internal/core/usecase/contract_catalog/list_rent_adjustment_types_use_case.go` |
| contract_catalog | `rent-manager-backend/internal/core/usecase/contract_catalog/rent_contract_catalog.go` |
| inflation_index | `rent-manager-backend/internal/core/usecase/inflation_index/create_inflation_index_use_case.go` |
| inflation_index | `rent-manager-backend/internal/core/usecase/inflation_index/get_inflation_index_by_id_use_case.go` |
| inflation_index | `rent-manager-backend/internal/core/usecase/inflation_index/get_inflation_index_by_period_use_case.go` |
| inflation_index | `rent-manager-backend/internal/core/usecase/inflation_index/list_inflation_indexes_by_period_range_use_case.go` |
| inflation_index | `rent-manager-backend/internal/core/usecase/inflation_index/list_inflation_indexes_use_case.go` |
| inflation_index | `rent-manager-backend/internal/core/usecase/inflation_index/update_inflation_index_use_case.go` |
| location | `rent-manager-backend/internal/core/usecase/location/list_countries_use_case.go` |
| location | `rent-manager-backend/internal/core/usecase/location/list_states_by_country_use_case.go` |
| owner | `rent-manager-backend/internal/core/usecase/owner/create_owner_use_case.go` |
| owner | `rent-manager-backend/internal/core/usecase/owner/get_owner_by_document_number_use_case.go` |
| owner | `rent-manager-backend/internal/core/usecase/owner/get_owner_by_email_use_case.go` |
| owner | `rent-manager-backend/internal/core/usecase/owner/get_owner_by_id_use_case.go` |
| owner | `rent-manager-backend/internal/core/usecase/owner/list_owners_use_case.go` |
| owner | `rent-manager-backend/internal/core/usecase/owner/search_owners_use_case.go` |
| owner | `rent-manager-backend/internal/core/usecase/owner/update_owner_use_case.go` |
| property_catalog | `rent-manager-backend/internal/core/usecase/property_catalog/list_property_statuses_use_case.go` |
| property_catalog | `rent-manager-backend/internal/core/usecase/property_catalog/list_property_types_use_case.go` |
| property | `rent-manager-backend/internal/core/usecase/property/create_property_use_case.go` |
| property | `rent-manager-backend/internal/core/usecase/property/get_property_by_code_use_case.go` |
| property | `rent-manager-backend/internal/core/usecase/property/get_property_by_id_use_case.go` |
| property | `rent-manager-backend/internal/core/usecase/property/get_property_summaries_use_case.go` |
| property | `rent-manager-backend/internal/core/usecase/property/list_property_use_case.go` |
| property | `rent-manager-backend/internal/core/usecase/property/search_property_use_case.go` |
| property | `rent-manager-backend/internal/core/usecase/property/update_property_use_case.go` |
| rent_payment | `rent-manager-backend/internal/core/usecase/rent_payment/create_rent_payment_use_case.go` |
| rent_payment | `rent-manager-backend/internal/core/usecase/rent_payment/get_rent_payment_by_id_use_case.go` |
| rent_payment | `rent-manager-backend/internal/core/usecase/rent_payment/get_rent_payment_schedule_use_case.go` |
| rent_payment | `rent-manager-backend/internal/core/usecase/rent_payment/get_rent_payment_suggestion_use_case.go` |
| rent_payment | `rent-manager-backend/internal/core/usecase/rent_payment/get_rental_contract_summary_use_case.go` |
| rent_payment | `rent-manager-backend/internal/core/usecase/rent_payment/list_rent_payments_use_case.go` |
| rent_payment | `rent-manager-backend/internal/core/usecase/rent_payment/update_rent_payment_use_case.go` |
| rental_contract | `rent-manager-backend/internal/core/usecase/rental_contract/create_rental_contract_use_case.go` |
| rental_contract | `rent-manager-backend/internal/core/usecase/rental_contract/get_rental_contract_by_id_use_case.go` |
| rental_contract | `rent-manager-backend/internal/core/usecase/rental_contract/list_rental_contracts_use_case.go` |
| rental_contract | `rent-manager-backend/internal/core/usecase/rental_contract/update_rental_contract_use_case.go` |
| role | `rent-manager-backend/internal/core/usecase/role/list_roles_use_case.go` |
| tenant | `rent-manager-backend/internal/core/usecase/tenant/create_tenant_use_case.go` |
| tenant | `rent-manager-backend/internal/core/usecase/tenant/get_tenant_by_document_number_use_case.go` |
| tenant | `rent-manager-backend/internal/core/usecase/tenant/get_tenant_by_id_use_case.go` |
| tenant | `rent-manager-backend/internal/core/usecase/tenant/list_tenant_use_case.go` |
| tenant | `rent-manager-backend/internal/core/usecase/tenant/search_tenant_use_case.go` |
| tenant | `rent-manager-backend/internal/core/usecase/tenant/update_tenant_use_case.go` |
| user | `rent-manager-backend/internal/core/usecase/user/create_user_use_case.go` |
| user | `rent-manager-backend/internal/core/usecase/user/delete_user_use_case.go` |
| user | `rent-manager-backend/internal/core/usecase/user/get_user_by_email_use_case.go` |
| user | `rent-manager-backend/internal/core/usecase/user/get_user_by_id_use_case.go` |
| user | `rent-manager-backend/internal/core/usecase/user/list_user_use_case.go` |
| user | `rent-manager-backend/internal/core/usecase/user/login_use_case.go` |
| user | `rent-manager-backend/internal/core/usecase/user/search_user_use_case.go` |
| user | `rent-manager-backend/internal/core/usecase/user/update_user_use_case.go` |

## All screen routes

Read parent routes in app.routes.ts for the /app prefix. Each feature mounts its child paths.

### rent-manager-frontend/projects/rent-manager-frontend/src/app/app.routes.ts

```text
path: '',
component: AuthLayoutComponent,
path: 'login',
path: '',
redirectTo: 'login',
path: 'app',
component: MainLayoutComponent,
canActivate: [authGuard],
path: '',
redirectTo: 'dashboard',
path: 'dashboard',
loadComponent: () =>
path: 'users',
path: 'owners',
path: 'tenants',
path: 'properties',
path: 'rent-contracts',
path: 'rent-payments',
path: 'rent-contracts/overview/:id',
loadComponent: () =>
path: 'inflation-indexes',
path: '**',
redirectTo: 'login',
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/dashboard/dashboard.router.ts

```text
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/inflation-index/inflation-index.routes.ts

```text
path: '',
loadComponent: () =>
path: 'create',
loadComponent: () =>
path: 'edit/:id',
loadComponent: () =>
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/login/login.routes.ts

```text
path: '',
component: LoginComponent,
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/owner/owner.routes.ts

```text
path: '',
component: OwnerComponent,
path: '',
component: OwnerListComponent,
path: 'create',
component: OwnerCreateComponent,
path: 'edit/:id',
component: OwnerEditComponent,
path: 'view/:id',
component: OwnerViewComponent,
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/property/property.routes.ts

```text
path: '',
component: PropertyComponent,
path: '',
component: PropertyListComponent,
path: 'create',
component: PropertyCreateComponent,
path: 'edit/:id',
component: PropertyEditComponent,
path: 'view/:id',
component: PropertyViewComponent,
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-contract/rent-contract.routes.ts

```text
path: '',
component: RentContractComponent,
path: '',
component: RentContractListComponent,
path: 'create',
component: RentContractCreateComponent,
path: 'edit/:id',
component: RentContractEditComponent,
path: 'view/:id',
component: RentContractViewComponent,
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/rent-payment/rent-payment.routes.ts

```text
path: '',
component: RentPaymentComponent,
path: 'create',
component: RentPaymentCreateComponent,
path: 'edit/:id',
component: RentPaymentEditComponent,
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/tenant/tenant.routes.ts

```text
path: '',
component: TenantComponent,
path: '',
component: TenantListComponent,
path: 'create',
component: TenantCreateComponent,
path: 'edit/:id',
component: TenantEditComponent,
path: 'view/:id',
component: TenantViewComponent,
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/user/user.routes.ts

```text
path: '',
component: UserComponent,
path: '',
component: UserListComponent,
path: 'create',
component: UserCreateComponent,
path: 'edit/:id',
component: UserEditComponent,
path: 'view/:id',
component: UserViewComponent,
```
