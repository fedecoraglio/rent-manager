# Source and operation coverage

Reviewed: 2026-11-13. Resolve paths from the repository root. Use this map to open the implementation; excerpts are navigation indexes, not replacements for complete source bodies. Line numbers refer to this review. Do not interpret internal operations as public endpoints.

## Contents

- File inventory
- Operations, data contracts, and rule locations

## File inventory

- `rent-manager-frontend/projects/rent-manager-frontend/public/assets/config.json`
- `rent-manager-frontend/projects/rent-manager-frontend/public/assets/i18n/en.json`
- `rent-manager-frontend/projects/rent-manager-frontend/public/assets/i18n/es.json`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/app.config.ts`
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
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/error/api-error.interceptor.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/error/api-error.model.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/i18n/language.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/loading/loading.interceptor.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/loading/loading.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/notification/notification.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/layout/auth/auth-layout.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/layout/auth/auth-layout.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/layout/auth/auth-layout.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/layout/main/main-layout.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/layout/main/main-layout.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/main.ts`

## Operations, data contracts, and rule locations

### rent-manager-frontend/projects/rent-manager-frontend/src/app/app.config.ts

```text
32: provideAppInitializer(() => {
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

### rent-manager-frontend/projects/rent-manager-frontend/src/app/core/loading/loading.service.ts

```text
9: show(): void {
13: hide(): void {
```

### rent-manager-frontend/projects/rent-manager-frontend/src/app/core/notification/notification.service.ts

```text
18: showError(message: string): void {
25: showSuccess(message: string): void {
32: showInfo(message: string): void {
39: clear(): void {
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
