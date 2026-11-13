---
name: rent-manager-auth
description: Work on login, credential validation, JWT, session persistence, protected routes, and logout in Rent Manager.
---

# Authentication and session

Read [sources and operations](references/coverage.md).

## Business behavior

Users log in with email/password to access administration. Public signup, password recovery, and refresh tokens are not implemented. Logout removes the local token; it does not revoke a JWT already issued by the server.

## Technical flow

1. LoginComponent sends credentials through AuthService → AuthApiService to `POST /v1/auth/login`.
2. LoginUseCase retrieves the user and compares the password through PasswordHasher; JwtTokenProvider issues a token with the 24-hour lifetime configured in main.
3. The UI extracts `data.token` and stores it in localStorage under `rent_manager_access_token`.
4. authGuard allows `/app` when a token exists. It only checks presence, not expiration or signature.
5. authInterceptor adds Bearer to requests when a token exists; its current implementation does not restrict the destination domain.
6. Go AuthMiddleware validates the token and places authenticated data in context. Check keys and claims in code before consuming them.
7. AuthService.logout clears storage. Follow MainLayoutComponent for logout navigation.

The error interceptor displays a message and rethrows; it does not implement token renewal or automatic logout on 401. Role preloading after login is commented out in AuthService. Do not assume RBAC merely because roles exist in JWT-related discussions or catalogs.

## Verify changes

Check valid/invalid credentials, missing or malformed Bearer, expired JWT, browser reload, and logout. Separate UI protection from actual server authorization. Inspect the router: inflation and login are public; AI is a separate service without this middleware.
