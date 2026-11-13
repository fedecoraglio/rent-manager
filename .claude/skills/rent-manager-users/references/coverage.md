# Source and operation coverage

Reviewed: 2026-11-13. Resolve paths from the repository root. Use this map to open the implementation; excerpts are navigation indexes, not replacements for complete source bodies. Line numbers refer to this review. Do not interpret internal operations as public endpoints.

## Contents

- File inventory
- Operations, data contracts, and rule locations

## File inventory

- `rent-manager-backend/internal/adapter/http/role_response.go`
- `rent-manager-backend/internal/adapter/http/role.go`
- `rent-manager-backend/internal/adapter/http/user_request.go`
- `rent-manager-backend/internal/adapter/http/user_response.go`
- `rent-manager-backend/internal/adapter/http/user.go`
- `rent-manager-backend/internal/adapter/store/mysql/repository/role.go`
- `rent-manager-backend/internal/adapter/store/mysql/repository/user.go`
- `rent-manager-backend/internal/core/domain/role.go`
- `rent-manager-backend/internal/core/domain/user.go`
- `rent-manager-backend/internal/core/policy/user/user_create_policy.go`
- `rent-manager-backend/internal/core/policy/user/user_update_policy.go`
- `rent-manager-backend/internal/core/port/role.go`
- `rent-manager-backend/internal/core/port/user.go`
- `rent-manager-backend/internal/core/usecase/role/list_roles_use_case.go`
- `rent-manager-backend/internal/core/usecase/user/create_user_use_case.go`
- `rent-manager-backend/internal/core/usecase/user/delete_user_use_case.go`
- `rent-manager-backend/internal/core/usecase/user/get_user_by_email_use_case.go`
- `rent-manager-backend/internal/core/usecase/user/get_user_by_id_use_case.go`
- `rent-manager-backend/internal/core/usecase/user/list_user_use_case.go`
- `rent-manager-backend/internal/core/usecase/user/login_use_case.go`
- `rent-manager-backend/internal/core/usecase/user/search_user_use_case.go`
- `rent-manager-backend/internal/core/usecase/user/update_user_use_case.go`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/role/role-api.service.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/role/role.model.ts`
- `rent-manager-frontend/projects/rent-manager-frontend/src/app/core/role/role.service.ts`
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

## Operations, data contracts, and rule locations

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

### rent-manager-backend/internal/core/usecase/role/list_roles_use_case.go

```text
9: type ListRolesUseCase struct {
13: func NewListRolesUseCase(roleRepository port.RoleRepository) *ListRolesUseCase {
17: func (lr *ListRolesUseCase) ListRoles(ctx context.Context, page uint64, limit uint64) ([]domain.Role, error) {
19: if err != nil {
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
