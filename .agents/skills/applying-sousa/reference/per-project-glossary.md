# SOUSA Per-Project Glossary

Quick lookup for the four known SOUSA projects. Always confirm against the repo's own SOUSA / AGENTS document; this is a summary, not a substitute.

## Diacous — Empyrean Challenge game engine

- **Status:** greenfield. All code must comply with SOUSA from day one. No legacy exemption.
- **Module path:** `github.com/mdhender/diacous`
- **Backend root:** `apps/backend/` (its own `go.mod`)
- **Frontend:** React + Vite + Tailwind, SPA (no SSR), under `apps/frontend/`
- **SQLite driver:** `zombiezen.com/go/sqlite` + `sqlitex` (pool-based, no `database/sql`)
- **HTTP framework:** Echo v5
- **Layer naming:** Domain / Application / Infrastructure / Delivery / Runtime

| Role            | Path                              |
| --------------- | --------------------------------- |
| Domain          | `internal/domain/`                |
| Sentinel errors | `internal/cerr/`                  |
| Application     | `internal/app/`                   |
| Infra (SQLite)  | `internal/infra/sqlite/`          |
| Infra (JWT)     | `internal/jwtmgr/`                |
| Delivery (HTTP) | `internal/delivery/http/`         |
| Runtime         | `cmd/server/`, `internal/server/` |

## EC — Epimethean Challenge

- **Status:** SOUSA from the start; ships both an API server and a CLI sharing one SQLite DB.
- **Backend root:** `backend/` (its own `go.mod`)
- **Two delivery peers:** `delivery/http` (Echo) and `delivery/cli` are peers; neither is privileged.
- **Frontend:** React + Vite + TailwindCSS under `apps/web/`
- **Docs site:** Hugo + Hextra under `apps/site/`
- **SQLite operational model:** API server is stopped or in maintenance mode while the CLI runs turn processing. CLI holds exclusive write access during turn execution.
- **Layer naming:** domain / cerr / app / infra / delivery / runtime

| Role               | Path                                          |
| ------------------ | --------------------------------------------- |
| Domain             | `backend/internal/domain/`                    |
| Sentinel errors    | `backend/internal/cerr/`                      |
| Application        | `backend/internal/app/`                       |
| Infra              | `backend/internal/infra/{sqlite,filestore,auth}/` |
| Delivery (HTTP)    | `backend/internal/delivery/http/`             |
| Delivery (CLI)     | `backend/internal/delivery/cli/`              |
| Runtime (server)   | `backend/internal/runtime/server/`, `backend/cmd/api/` |
| Runtime (CLI)      | `backend/internal/runtime/cli/`, `backend/cmd/cli/`    |

Hard rules of note:

- The CLI must not bypass `app` for core game behavior.
- The API must not implement game rules independently of the core.
- Shared-DB operational mode is enforced in `runtime`, deploy scripts, and operator docs — not in domain / app / delivery.

## GemGem — OTTOMAP monorepo

- **Status:** incremental adoption. SOUSA applies to all new files and to any file you edit; no big-bang rewrites.
- **Module path:** `github.com/mdhender/gemgem`
- **Backend:** Go API under `apps/api/`
- **Frontend:** Next.js under `apps/web/`
- **SQLite driver:** `modernc.org/sqlite` with standard `database/sql` patterns
- **HTTP framework:** Echo v5
- **CLI / entry framework:** Cobra
- **Layer naming:** Domain / Application / Infrastructure / Delivery / Runtime

| Role            | Path                          |
| --------------- | ----------------------------- |
| Domain          | `internal/domain/`            |
| Sentinel errors | `internal/cerr/`              |
| Application     | `internal/app/`               |
| Infra (SQLite)  | `internal/infra/sqlite/`      |
| Infra (JWT)     | `internal/jwtmgr/`            |
| Delivery (HTTP) | `internal/delivery/http/`     |
| Runtime         | `cmd/api/main.go`, `internal/server/` |

Known transitional violation (do not extend): `internal/delivery/http/routes.go` currently receives `app.UserRepo` directly. New routes must depend on a focused use case / application service, not on raw repos.

## Hokey — Go + SQLite game engine

- **Status:** incremental adoption (alpha). SOUSA applies to all new files and to any file you edit.
- **SQLite driver:** `zombiezen.com/go/sqlite` + `sqlitex` (pool-based, no `database/sql`)
- **No HTTP server** (yet) — primary delivery is CLI / REPL
- **Layer naming uses Hokey's own vocabulary:** Core → Usecase → Store → Interface → Runtime

| Role                | Path                                                |
| ------------------- | --------------------------------------------------- |
| Core (Domain)       | `hexes/`, `core/terrain/`, `cerr/`                  |
| Usecase (App)       | `usecase/mapgen/`, future `usecase/engine/`         |
| Store (Infra)       | `gamedb/` (adapters like `MapStoreAdapter`, `RegionLandStoreAdapter`) |
| Interface (Deliv.)  | `cmd/hokey/main.go` REPL & dispatch, `mapload`, `tiles` rendering |
| Runtime             | `cmd/hokey/main.run(...)` composition               |

Known mixed concerns to be aware of:

- `cmd/hokey/main.go` mixes Interface and Runtime — acceptable for now, do not deepen.
- `tiles` mixes Core types and rendering — rendering is Interface and should eventually move.
- `gamedb.LoadMap` was removed; all callers use `MapStoreAdapter`.

## Quick cross-project name map

| Concept                  | Diacous / GemGem / EC         | Hokey                 |
| ------------------------ | ----------------------------- | --------------------- |
| Pure types & invariants  | Domain                        | Core                  |
| Use cases & ports        | Application (`internal/app`)  | Usecase               |
| DB / external adapters   | Infrastructure (`internal/infra`) | Store (`gamedb`)  |
| Handlers / CLI / format  | Delivery (`internal/delivery`)| Interface             |
| Wiring & lifecycle       | Runtime (`cmd/*`, `internal/server` or `internal/runtime`) | Runtime (`main.run`) |
