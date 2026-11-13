---
name: rent-manager-frontend
description: Understand and modify the Rent Manager Angular architecture, standalone routes, read and write services, signals, reactive forms, and shared components.
---

# Angular frontend

Use `rent-manager-frontend/projects/rent-manager-frontend/src/app` as the application root. Read [coverage](references/coverage.md), `app.routes.ts`, and `app.config.ts` before cross-cutting changes.

Angular 21, TypeScript 5.9, RxJS 7.8, Material/CDK, Tailwind, and ngx-translate are declared in `package.json`. The application uses standalone components, frequently OnPush, signals for state, and observables for requests/events.

| Directory | Responsibility |
| --- | --- |
| core/api | URL construction and response envelope |
| core/auth | Session, token, guard, interceptor |
| core/config | Configuration loaded at startup |
| core/<domain> | Shared models, read services, catalogs |
| feature/<domain> | Routes, create, edit, view, lists, forms, writes |
| layout | Public login layout and authenticated layout |
| ui | Reusable notifications, property cards, payment schedule |
| assets | JSON configuration and languages |

Follow each domain's pattern: service names are not uniform. Users uses `user-api.service.ts`/`user.service.ts`; others split core reads from feature writes. Preserve snake_case API model fields wherever Go expects them.

## Navigation

`/login` is public. `/app` uses `authGuard`, redirects to dashboard, and loads users, owners, tenants, properties, contracts, payments, and inflation. Each feature declares its child operation routes. Payments open from a contract with `rentalContractId` and `period` query parameters. Contract overview fetches contract, property, and schedule. There are no PDF/chat screens.

## Change a feature

1. Read routes, container component, form/field enum, and feature model.
2. Follow UI service → API service → `ApiUrlService` → Go endpoint.
3. Preserve validators, readonly behavior, saving state, and subscription cleanup (`takeUntilDestroyed` or destroy$).
4. Preserve errors, loading, translations, and navigation after success.
5. Verify create/edit/view/list and direct loading of a deep URL.

From `rent-manager-frontend`: `npm ci`, `npm start` (port 4200), `npm run build`, `npm run format:check`. Inspect `angular.json` before `npm test` or `npm run lint`: a package script alone does not guarantee a configured target. The container builds with Node 22 and serves output through Nginx, with an index.html fallback for SPA routes.
