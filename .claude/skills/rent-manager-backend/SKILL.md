---
name: rent-manager-backend
description: Understand and modify the Rent Manager Go backend architecture, hexagonal layers, Gin API, policies, calculations, MySQL repositories, and migrations.
---

# Go backend

Read `rent-manager-backend/cmd/api/main.go` for dependency assembly and `internal/adapter/http/router.go` for the route contract. Read [coverage](references/coverage.md) for the technical inventory.

## Structure and dependencies

| Directory | Role |
| --- | --- |
| cmd/api | Composition root: configuration, DB, migrations, security, calculators, policies, use cases, handlers, router |
| internal/core/domain | Entities, result structures, business error codes and errors |
| internal/core/port | Persistence, security, and calculation interfaces |
| internal/core/policy | Normalization, constraints, and queries before writes |
| internal/core/usecase | Feature orchestration |
| internal/core/service | IPC, adjustment, interest, and total calculations |
| internal/adapter/http | Gin, binding, DTOs, conversions, middleware, responses |
| internal/adapter/security | Bcrypt and JWT |
| internal/adapter/store/mysql | database/sql connection, Squirrel, repositories, embedded migrations |
| internal/adapter/config | Environment variable loading |

Use Go 1.25.4 as declared in `go.mod`. Write flow: JSON → DTO/binding → entity → use case → policy → port/SQL repository → response. Reads do not require a write policy. Keep business logic outside handlers and SQL outside use cases.

## Startup and persistence

`main` loads configuration, pings MySQL, and runs `Migrate()` before serving. Errors abort startup. SQL is embedded with `//go:embed migrations/*.sql`, so migration changes require rebuilding the binary. `migrate.Up()` accepts `ErrNoChange`. Add migrations for schema evolution; do not rewrite an applied migration.

## Shared HTTP behavior

The API uses `/v1`. Most groups require JWT; login and inflation do not use the middleware. Do not assume role authorization or owner isolation: inspect actual query filters and middleware.

Normal responses use `{data: ...}` and HTTP 200 through `handleSuccess`. Errors usually use `{code,message}` through `statusFromErrorCode`; some handlers return `{error: ...}`. Do not assume all errors map to 400/404/409: the default is 500. Common pagination: `page >= 1`, `1 <= limit <= 100`; search has its own DTO.

## Develop and verify

From `rent-manager-backend`: `go run ./cmd/api`, `go build ./cmd/api`, `go test ./...`. Prepare MySQL and variables required by `config.go`; do not copy secrets into documentation. Validate real policy and persistence cases when rules, dates, money, or uniqueness change. A build alone does not validate business behavior.
