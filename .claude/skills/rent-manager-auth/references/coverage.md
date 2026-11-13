# Source and operation coverage

Reviewed: 2026-11-13. Resolve paths from the repository root. Use this map to open the implementation; excerpts are navigation indexes, not replacements for complete source bodies. Line numbers refer to this review. Do not interpret internal operations as public endpoints.

## Contents

- File inventory
- Operations, data contracts, and rule locations

## File inventory

- `rent-manager-backend/internal/adapter/http/auth.go`
- `rent-manager-backend/internal/adapter/http/middleware_auth.go`
- `rent-manager-backend/internal/adapter/security/bcrypt_password_hasher.go`
- `rent-manager-backend/internal/adapter/security/jwt_token_provider.go`
- `rent-manager-backend/internal/core/domain/login.go`
- `rent-manager-backend/internal/core/domain/token.go`
- `rent-manager-backend/internal/core/port/token.go`
- `rent-manager-backend/internal/core/usecase/user/login_use_case.go`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/auth/auth-api.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/auth/auth.guard.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/auth/auth.interceptor.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/auth/auth.models.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/auth/auth.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/auth/token-storage.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/login/login.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/login/login.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/login/login.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/feature/login/login.routes.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/layout/auth/auth-layout.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/layout/auth/auth-layout.component.scss`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/layout/auth/auth-layout.component.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/layout/main/main-layout.component.html`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/layout/main/main-layout.component.ts`

## Operations, data contracts, and rule locations

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

### rent-manager-backend/internal/adapter/http/middleware_auth.go

```text
19: func AuthMiddleware(tokenProvider port.TokenProvider) gin.HandlerFunc {
22: if header == "" {
28: if !strings.HasPrefix(header, bearerPrefix) {
35: if tokenValue == "" {
42: if err != nil {
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
