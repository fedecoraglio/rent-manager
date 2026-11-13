---
name: rent-manager-users
description: Modify user creation, listing, search, retrieval, updates, deletion, passwords, and role lookup flows in Rent Manager.
---

# Users and roles

Read [detailed coverage](references/coverage.md) for DTOs, operations, and validations. UI: `feature/user`; API: `/v1/users`, protected by JWT. Roles: `GET /v1/roles`.

## Business behavior and operations

Support creating administrative users, searching/listing, retrieving by ID, updating, and deleting. The role catalog supplies forms; the router exposes no role CRUD. Do not assume different permissions for administrators and other roles.

## Flow

List/form component → UserService/UserApiService → UserHandler → usecase/user → policy/user → user repository → users table. Creation normalizes name/email, requires name, valid email, and password, and checks for duplicate email. CreateUserUseCase hashes the password with bcrypt; main configures cost 12. Inspect response DTOs to prevent hash exposure.

Updates have their own policy and use case: UpdateUserUseCase calls PathUser without PasswordHasher and does not rehash as creation does. user_update_policy.go contains an anomalous assignment to Email when PasswordHash is empty; document it as an observed limitation, not intended behavior. Inspect the DTO and PathUser columns before offering password changes.

Listing uses pagination; search has its own contract. GetUserByEmailUseCase exists internally, but there is no dedicated email endpoint: follow the router for actual exposed routes. DeleteUser is the only administrative entity DELETE in the current router.

## Verify

Cover duplicate email on create/update (excluding the same ID), invalid input, missing user, creation hashing, responses without passwords, and deletion effects on SQL relationships. Follow actual HTTP error handling rather than assuming every error is cataloged.
