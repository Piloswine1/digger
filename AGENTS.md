# AGENTS.md

## Architecture

- **Monorepo** with independent Go backend (`cmd/`, `inner/`) and Vue 3 frontend (`front/`).
- Backend is a Gin HTTP server on `:8080`. Entrypoint: `cmd/main.go` → `rest.InitApp()`.
- Frontend Vite dev server proxies `/api` → `http://localhost:8080`. Entrypoint: `front/src/main.ts`.

## Commands

### Backend (Go)

```bash
# Requires mise (see mise.toml). Go tool version managed by mise.
mise run dev            # go run cmd/main.go (watches inner/**/*.go)
go build ./...           # type-check / compile check
go vet ./...             # static analysis
```

### Frontend (pnpm, from `front/`)

```bash
pnpm dev                # Vite dev server (hot reload)
pnpm build              # type-check + build (run-p type-check build-only)
pnpm lint               # oxlint + eslint (run-s lint:*)
pnpm format             # oxfmt src/
pnpm type-check         # vue-tsc --build (standalone type-check)
```

### Running both

Start backend first (`mise run dev`), then frontend (`cd front && pnpm dev`).

## Key quirks

- **Podman, not Docker**: The backend talks to Podman via Docker API. `DOCKER_HOST` env is set in `mise.toml` to `unix:///run/user/1000/podman/podman.sock`. The podman socket must be running (`mise run podman:socket` if needed).
- **Basic auth** on `/api/v1/containers/*` routes. Env vars `USER_LOGIN`/`USER_PASS`; defaults to `diggerlogin`/`diggerpass` if unset.
- `GIN_MODE=debug` drives tint-colored structured logging; unset → JSON logs.
- **Vue 3 beta**: `pnpm-workspace.yaml` overrides all `@vue/*` packages to `beta`.
- **Frontend formatting**: oxfmt with `semi: false`, `singleQuote: true`. No semicolons in TS/Vue files.

## Package layout

| Directory | Purpose |
|-----------|---------|
| `cmd/main.go` | Server entrypoint |
| `inner/config/` | Auth middleware, logger setup |
| `inner/docker/` | Docker (Podman) client helper |
| `inner/rest/` | Routes, handlers, DTOs, models |
| `front/src/app/` | Vue app shell |
| `front/src/features/` | Feature modules |
| `front/src/components/` | Shared components |
| `front/src/lib/` | Shared utilities |

## No tests yet

There are no test files in the repo. If you add tests:
- Go: standard `go test ./...`
- Frontend: no test runner configured yet.
