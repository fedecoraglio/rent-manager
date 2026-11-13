# Source and operation coverage

Reviewed: 2026-11-13. Resolve paths from the repository root. Use this map to open the implementation; excerpts are navigation indexes, not replacements for complete source bodies. Line numbers refer to this review. Do not interpret internal operations as public endpoints.

## Contents

- File inventory
- Operations, data contracts, and rule locations

## File inventory

- `rent-manager-frontend/angular.json`
- `rent-manager-frontend/Dockerfile`
- `rent-manager-frontend/eslint.config.js`
- `rent-manager-frontend/nginx.conf`
- `rent-manager-frontend/package.json`
- `rent-manager-frontend/postcss.config.js`
- `rent-manager-frontend/projects/rent-manager-frontend/eslint.config.js`
- `rent-manager-frontend/projects/rent-manager-frontend/public/.gitkeep`
- `rent-manager-frontend/projects/rent-manager-frontend/public/assets/config.json`
- `rent-manager-frontend/projects/rent-manager-frontend/public/assets/i18n/en.json`
- `rent-manager-frontend/projects/rent-manager-frontend/public/assets/i18n/es.json`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/app.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/app.config.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/app.routes.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/api/api-url.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/api/api.model.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/auth/auth-api.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/auth/auth.guard.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/auth/auth.interceptor.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/auth/auth.models.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/auth/auth.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/auth/token-storage.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/config/app-config.model.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/config/app-config.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/contract-catalog/contract-catalog-api.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/contract-catalog/contract-catalog.model.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/contract-catalog/contract-catalog.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/error/api-error.interceptor.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/error/api-error.model.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/i18n/language.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/inflation-index/inflation-index-read-api.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/inflation-index/inflation-index-read.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/inflation-index/inflation-index.model.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/loading/loading.interceptor.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/loading/loading.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/location/location-api.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/location/location.model.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/location/location.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/notification/notification.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/owner/owner-read-api.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/owner/owner-read.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/owner/owner.model.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/property-catalog/property-catalog-api.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/property-catalog/property-catalog.model.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/property-catalog/property-catalog.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/property/property-read-api.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/property/property-read.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/property/property.model.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/rent-contract/rent-contract-read-api.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/rent-contract/rent-contract-read.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/rent-contract/rent-contract.model.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/rent-payment/rent-payment-api.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/rent-payment/rent-payment.model.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/rent-payment/rent-payment.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/role/role-api.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/role/role.model.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/role/role.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/tenant/tenant-core.model.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/tenant/tenant-read-api.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/tenant/tenant-read.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/dashboard/dashboard.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/dashboard/dashboard.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/dashboard/dashboard.mapper.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/dashboard/dashboard.router.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/dashboard/model/dashboard.model.ts`
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
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/login/login.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/login/login.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/login/login.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/login/login.routes.ts`
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
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/user/create/user-create.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/user/create/user-create.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/user/create/user-create.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/user/edit/user-edit.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/user/edit/user-edit.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/user/edit/user-edit.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/user/form/field/user-form-field.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/user/form/user-form.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/user/form/user-form.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/user/form/user-form.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/user/list/table/user-list-table.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/user/list/table/user-list-table.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/user/list/table/user-list-table.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/user/list/user-list.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/user/list/user-list.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/user/list/user-list.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/user/model/user.model.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/user/service/user-api.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/user/service/user.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/user/user.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/user/user.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/user/user.routes.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/user/view/user-view.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/user/view/user-view.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/user/view/user-view.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/layout/auth/auth-layout.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/layout/auth/auth-layout.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/layout/auth/auth-layout.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/layout/main/main-layout.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/layout/main/main-layout.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/ui/.gitkeep`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/ui/notification/app-notification.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/ui/notification/app-notification.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/ui/notification/app-notification.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/ui/property-summary-card/property-summary-card.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/ui/property-summary-card/property-summary-card.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/ui/property-summary-card/property-summary-card.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/ui/property-summary-card/property-summary-card.model.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/ui/rent-payment-schedule/rent-payment-schedule.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/ui/rent-payment-schedule/rent-payment-schedule.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/ui/rent-payment-schedule/rent-payment-schedule.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/environments/environment.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/index.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/main.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/styles.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/tsconfig.app.json`
- `rent-manager-frontend/projects/rent-manager-frontend/tsconfig.spec.json`
- `rent-manager-frontend/README.md`
- `rent-manager-frontend/tailwind.config.js`
- `rent-manager-frontend/tsconfig.json`

## Operations, data contracts, and rule locations

### rent-manager-frontend/projects/rent-manager-frontend/src/app/app.component.ts

```text
15: constructor() {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/app.config.ts

```text
32: provideAppInitializer(() => {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/app.routes.ts

```text
9: path: '',
13: path: 'login',
17: path: '',
19: redirectTo: 'login',
24: path: 'app',
26: canActivate: [authGuard],
29: path: '',
31: redirectTo: 'dashboard',
34: path: 'dashboard',
39: path: 'users',
43: path: 'owners',
47: path: 'tenants',
51: path: 'properties',
55: path: 'rent-contracts',
59: path: 'rent-payments',
63: path: 'rent-contracts/overview/:id',
70: path: 'inflation-indexes',
79: path: '**',
80: redirectTo: 'login',
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/core/api/api-url.service.ts

```text
9: build(path: string): string {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/core/auth/auth-api.service.ts

```text
15: return this.apiUrl.build('/auth/login');
18: login(request: LoginRequest): Observable<LoginResponse> {
19: return this.http
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/core/auth/auth.service.ts

```text
17: constructor() {
21: login(request: LoginRequest): Observable<LoginResponse> {
28: logout(): void {
32: isAuthenticated(): boolean {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/core/auth/token-storage.service.ts

```text
7: getToken(): string | null {
8: return localStorage.getItem(TOKEN_KEY);
11: setToken(token: string): void {
12: localStorage.setItem(TOKEN_KEY, token);
15: clear(): void {
16: localStorage.removeItem(TOKEN_KEY);
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/core/config/app-config.service.ts

```text
15: const config = await firstValueFrom(this.http.get<AppConfig>('/assets/config.json'));
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

### rent-manager-frontend/projects/rent-manager-frontend/src/app/core/error/api-error.interceptor.ts

```text
12: catchError((error: HttpErrorResponse) => {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/core/i18n/language.service.ts

```text
13: constructor() {
17: initialize(): void {
23: changeLanguage(language: Language): void {
24: localStorage.setItem(LANGUAGE_KEY, language);
29: const language = localStorage.getItem(LANGUAGE_KEY);
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

### rent-manager-frontend/projects/rent-manager-frontend/src/app/core/loading/loading.service.ts

```text
9: show(): void {
13: hide(): void {
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

### rent-manager-frontend/projects/rent-manager-frontend/src/app/core/notification/notification.service.ts

```text
18: showError(message: string): void {
25: showSuccess(message: string): void {
32: showInfo(message: string): void {
39: clear(): void {
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

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/dashboard/dashboard.component.ts

```text
25: constructor() {
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

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/login/login.component.ts

```text
44: email: ['', [Validators.required, Validators.email]],
45: password: ['', [Validators.required]],
48: constructor() {
57: void this.router.navigate(['/app']);
66: submit(): void {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/login/login.routes.ts

```text
7: path: '',
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

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/user/create/user-create.component.ts

```text
29: constructor() {
32: switchMap(() => from(this.router.navigate(['/app/users']))),
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/user/edit/user-edit.component.ts

```text
36: constructor() {
46: switchMap(() => from(this.router.navigate(['/app/users']))),
62: ngOnInit(): void {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/user/form/user-form.component.ts

```text
41: [UserFormField.Name]: new FormControl('', Validators.required),
42: [UserFormField.Email]: new FormControl('', [Validators.required, Validators.email]),
44: [UserFormField.RoleId]: new FormControl(2, Validators.required),
83: ngAfterViewInit() {
102: ngOnDestroy() {
111: this.passwordField?.setValidators([Validators.required]);
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/user/list/user-list.component.ts

```text
68: constructor() {
71: switchMap(() => {
111: ngOnInit(): void {
115: onPageChange(event: PageEvent): void {
121: viewUser(id: number): void {
122: void this.router.navigate(['/app/users/view', id]);
125: editUser(id: number): void {
126: void this.router.navigate(['/app/users/edit', id]);
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/user/service/user-api.service.ts

```text
15: return this.apiUrl.build('/users');
18: create(request: CreateUserRequest): Observable<User> {
19: return this.http
24: list(page: number, limit: number): Observable<User[]> {
27: return this.http
32: search(value: string, page: number, limit: number): Observable<User[]> {
35: return this.http
40: getById(id: number): Observable<User> {
41: return this.http
46: update(id: number, request: UpdateUserRequest): Observable<User> {
47: return this.http
52: delete(id: number): Observable<void> {
53: return this.http.delete<void>(`${this.baseUrl}/${id}`);
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/user/service/user.service.ts

```text
17: save$(user: UserFormValue): Observable<User> {
31: update$(id: number, user: UserFormValue): Observable<User> {
48: get$(id: number): Observable<User> {
58: list$(page = 1, limit = 10): Observable<User[]> {
68: search$(value: string, page = 1, limit = 10): Observable<User[]> {
78: delete$(id: number): Observable<void> {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/user/user.routes.ts

```text
13: path: '',
18: path: '',
22: path: 'create',
26: path: 'edit/:id',
30: path: 'view/:id',
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/user/view/user-view.component.ts

```text
32: constructor() {
42: switchMap(() => from(this.router.navigate(['/app/users']))),
50: switchMap(() => from(this.router.navigate(['/app/users/edit', this.userId]))),
56: ngOnInit(): void {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/layout/auth/auth-layout.component.ts

```text
17: changeLanguage(language: 'es' | 'en'): void {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/layout/main/main-layout.component.ts

```text
90: constructor() {
112: logout(): void {
114: void this.router.navigate(['/login']);
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/ui/property-summary-card/property-summary-card.component.ts

```text
21: constructor() {
26: this.router.navigate([
37: gotToViewProperty() {
```
