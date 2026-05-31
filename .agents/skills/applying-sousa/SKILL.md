---
name: applying-sousa
description: Apply SOUSA architecture discipline (Onion/Clean layering with one-way inward dependencies) when working in a repo that declares it follows SOUSA. Use when placing new code, reviewing changes, refactoring, or wiring adapters in projects like Diacous, EC, GemGem, Hokey, or any repo containing a SOUSA.md/AGENTS.md that references SOUSA.
---

# Applying SOUSA

SOUSA is a strict **placement and import discipline** inspired by Onion / Clean Architecture. In a SOUSA repo, **dependencies flow inward only**:

```diagram
╭────────╮     ╭─────────────╮     ╭──────────────────────────╮     ╭─────────╮
│ Domain │◀────│ Application │◀────│ Infrastructure / Delivery│◀────│ Runtime │
╰────────╯     ╰─────────────╯     ╰──────────────────────────╯     ╰─────────╯
   (pure)        (use cases,         (DB, HTTP, CLI, JWT,            (wiring,
                  ports/             filesystem — concrete            composition
                  interfaces)        adapters)                        root)
```

Outer layers may import inner layers. Inner layers must **never** import outer layers. Infrastructure and Delivery are peer outer layers — neither imports the other; create a port instead.

For a coding agent, SOUSA means three things:

1. Put behavior in the **correct layer**.
2. Keep dependencies pointing **inward only**.
3. Do **not** bypass boundaries for convenience.

## When this skill applies

Use this skill whenever any of these are true:

- The repo contains a `SOUSA.md`, or an `AGENTS.md` / `README.md` that references SOUSA.
- The repo layout matches the structure below (`internal/domain`, `internal/app`, `internal/infra/*`, `internal/delivery/*`, `cmd/*` runtime).
- The user mentions one of the SOUSA projects by name: **Diacous**, **EC** (Epimethean Challenge), **GemGem** (OTTOMAP), **Hokey**.

Before doing anything else in a SOUSA repo, **read its project-specific SOUSA document if one exists** (e.g. `docs/SOUSA.md`, `SOUSA.md`, or the relevant section of `AGENTS.md`). Per-project rules override anything here. See [reference/per-project-glossary.md](reference/per-project-glossary.md) for the known project name/path mappings.

## The five layers

Layer names vary by project (`Domain` vs `Core`, `Application` vs `Usecase`, etc.). The roles are the same.

### 1. Domain (a.k.a. Core)

Pure business entities, value objects, invariants, deterministic transformations, sentinel errors.

- **No** I/O, no framework types, no SQL, no JWT, no HTTP, no filesystem, no `time.Now`, no randomness.
- Easiest code in the repo to unit test.
- Typical locations: `internal/domain/`, `internal/cerr/`, `core/...`, `hexes/`, etc.

### 2. Application (a.k.a. Use Cases / Usecase)

System behaviors and orchestration: commands, queries, authorization decisions, transaction intent. Owns the **ports** (interfaces) it needs from outer layers.

- May import: Domain, sentinel errors, stdlib.
- Must **not** import: concrete infrastructure or delivery packages, Echo, SQLite driver packages, JWT library internals, CLI frameworks.
- Typical location: `internal/app/`.

### 3. Infrastructure (a.k.a. Store)

Concrete adapters that **implement** application ports: SQLite repositories, JWT manager, filesystem stores, magic-link delivery, bcrypt, etc.

- May import: Application (for the port interfaces it implements), Domain, driver packages.
- Must not contain HTTP semantics, route logic, or business authorization policy.
- Typical locations: `internal/infra/sqlite/`, `internal/infra/filestore/`, `internal/infra/auth/`, `internal/jwtmgr/`.

### 4. Delivery (a.k.a. Interface)

Translates external input/output into application calls. HTTP handlers, CLI commands, file-format adapters. Handlers stay thin: **parse → call use case → serialize**.

- May import: Application, narrow support packages needed for delivery concerns.
- Must **not**: execute raw SQL, open DB connections directly, import infrastructure packages directly, implement business rules.
- `delivery/http` and `delivery/cli` are peers — neither is privileged.
- Typical locations: `internal/delivery/http/`, `internal/delivery/cli/`.

### 5. Runtime (a.k.a. Composition Root)

Process startup, dependency construction, server/CLI lifecycle, config and flag parsing.

- May import **all** layers for wiring only.
- Must not host business rules.
- No inner layer imports Runtime.
- Typical locations: `cmd/<binary>/`, `internal/runtime/`, `internal/server/`.

## Frontend is also delivery

If the project has a frontend (React/Next/Vite/Tailwind), it is a delivery layer around the backend — **not** a second source of truth.

| Concern                                | Goes in           |
| -------------------------------------- | ----------------- |
| API client, auth/session, shared types | `src/lib/`        |
| Reusable presentational components     | `src/components/` |
| Route/page composition                 | `src/pages/` or `src/app/` |

Rules:

- Centralize API and auth in `src/lib/`. Do not scatter `fetch` and token handling across pages.
- Components stay presentational. No hidden business rules or API orchestration in leaf components.
- The frontend may reflect permissions in the UI, but the **backend is the enforcement boundary**. Hidden buttons or route guards are not authorization.
- Do not duplicate domain rules in TypeScript unless needed for UX, and keep duplication minimal.

## Hard rules (canonical)

1. Domain must not import Application, Infrastructure, Delivery, or Runtime — or any framework / driver package.
2. Application must not import Infrastructure or Delivery packages, framework packages (Echo, Cobra), or driver packages (SQLite, JWT libs).
3. Infrastructure implements ports declared by Application; Application never depends on Infrastructure concrete types.
4. Delivery layers (`http`, `cli`, etc.) are peers and stay thin.
5. Delivery must not execute raw SQL, open DB connections, or implement business rules.
6. Runtime owns wiring; no inner layer imports Runtime.
7. The backend is authoritative for business rules. Frontend code is not.
8. Sentinel errors live in their own small package (commonly `internal/cerr/`) and express business meaning, not transport formatting.

## Coding-agent workflow

When you are about to add or change code in a SOUSA repo:

1. **Read the project's SOUSA doc first** if one exists, plus any nearby `AGENTS.md`.
2. **Identify the layer that owns the behavior.** Ask: is this a rule about *what is true* (Domain), *what the system does* (Application), *how it persists or talks to the outside* (Infrastructure / Delivery), or *how it starts up* (Runtime)?
3. **Place new code in the innermost layer that can own it cleanly.** Prefer adding or refining a small interface in Application over importing a concrete adapter.
4. **Check imports**. Open the file you are editing and verify the import block does not point outward.
5. **Keep handlers thin.** In Delivery: parse → call use case → serialize. No SQL, no orchestration.
6. **Keep adapters dumb.** In Infrastructure: row mapping, queries, driver setup — no business policy.
7. **Add or update tests in the layer where the logic lives.** Domain tests need no DB; Application tests use fakes or a narrow harness; Infrastructure tests use a real driver against a temp DB.
8. **Greenfield vs incremental.** If the project is greenfield (e.g. Diacous), comply fully from day one. If it is incremental (e.g. Hokey, GemGem, EC), apply SOUSA fully to new files, and for edited legacy files move only the touched behavior toward the correct layer — do not undertake unrelated rewrites in the same change.
9. **Do not use existing violations as permission to add new ones.** If you must add an exception, call it out explicitly with what / why / planned cleanup.

## Pre-merge review checklist

### Backend

- [ ] Business logic lives in Application or Domain, not in handlers or DB adapters.
- [ ] DB driver access is isolated to Infrastructure.
- [ ] JWT / auth concrete details are kept out of Application behind a port (e.g. `app.JWTIssuer`).
- [ ] All imports flow inward only.
- [ ] Handlers follow parse → call use case → serialize.
- [ ] No new direct repo / adapter imports were added in Delivery.
- [ ] Tests cover the changed Domain / Application behavior.

### Frontend (if applicable)

- [ ] API and auth concerns are centralized in `src/lib/`.
- [ ] Components stay mostly presentational.
- [ ] No business rules were hidden in route components or UI widgets.
- [ ] The frontend is not the only enforcement point for permissions or validation.
- [ ] Shared types/helpers are extracted, not copied across pages.

## Common scenarios — where does this go?

| Scenario                                                       | Layer          | Notes                                                |
| -------------------------------------------------------------- | -------------- | ---------------------------------------------------- |
| New entity / value object / invariant                          | Domain         | Pure types only.                                     |
| New sentinel error like `ErrGameNotFound`                      | `cerr` (Domain)| Business meaning, not HTTP status.                   |
| New "create user" / "process turn" / "upload document"         | Application    | Define a use-case function and any needed port.     |
| Bcrypt hashing, JWT signing, magic-link storage, file upload   | Infrastructure | Application talks to ports, never to these directly. |
| New REST endpoint                                              | Delivery/http  | Thin handler; call an existing use case.             |
| New CLI subcommand for turn processing                         | Delivery/cli   | Calls the same Application as the HTTP server.       |
| Wiring the new repo into the running server                    | Runtime        | `cmd/<binary>/main.go` or `internal/server/`.        |
| New API call from a React page                                 | `src/lib/`     | Then consume from the page/component.                |
| New shared button or card                                      | `src/components/` | Presentational; no fetch / auth logic inside.     |

## Project-specific details

Per-project SOUSA documents always win when they conflict with this skill. They specify:

- Exact module paths and import rules (e.g. `github.com/mdhender/<project>/internal/...`).
- Which SQLite driver the project uses (this matters — see below).
- Which HTTP framework, CLI framework, and frontend stack.
- Whether the project is greenfield (strict) or incremental.

For a quick lookup of the four known SOUSA projects, see [reference/per-project-glossary.md](reference/per-project-glossary.md).

### SQLite drivers differ — do not mix patterns

Two driver families appear across SOUSA projects. Use the one the project already uses; never mix.

- **`zombiezen.com/go/sqlite` + `sqlitex`** (Diacous, Hokey)
  - Wrap a `*sqlitex.Pool`. Every adapter method does `conn := pool.Take(ctx); defer pool.Put(conn)`.
  - Apply per-connection pragmas via `PrepareConn` in `sqlitex.PoolOptions`, **not** via DSN `_pragma`.
  - Persistent DB: WAL, NORMAL locking, pool size matched to concurrency.
  - In-memory test DB: EXCLUSIVE locking, journal MEMORY, pool size 1.
  - Use `sqlitex.ExecuteTransient` for DDL and PRAGMAs.
  - **Do not** use `database/sql` patterns (`sql.DB`, `sql.Tx`, `rows.Next()`) — they are incompatible.
- **`modernc.org/sqlite` + `database/sql`** (GemGem)
  - Use `*sql.DB`, `*sql.Tx`, `ExecContext`, `QueryContext`, `rows.Next()`, `BeginTx`.
  - Translate `sql.ErrNoRows` and driver-specific errors at the adapter boundary; do not leak them upward.
  - Transactions for use cases flow through an `app.Tx` port, not through handlers.

In both cases: SQL, migrations, PRAGMAs, and row mapping stay in `internal/infra/sqlite/`. Return typed data / typed errors upward — never leak statement or row types.

### Auth conventions

- Bcrypt password hashing → Infrastructure (a SQLite adapter or dedicated auth adapter).
- Magic-link issuance and validation → Application use case. Token storage and email delivery → Infrastructure.
- JWT issuance / validation → Infrastructure adapter (commonly `internal/jwtmgr/`). Application depends only on the port (e.g. `app.JWTIssuer`).
- JWT middleware that parses headers and injects claims → Delivery. Authorization decisions about what a user may do → Application.
- Never store plaintext passwords or signing keys in Domain or Application code.

## If you are unsure

- Choose the layer that minimizes outward dependencies.
- Keep logic pure first, then add adapters.
- Prefer small extraction refactors over broad reorganizations.
- If a handler "needs" a repo directly, stop and ask whether that should be a new application service or use case instead.
- If a frontend page or component is growing API / auth logic, move that logic into `src/lib/`.
- When in real doubt, propose the placement to the user before writing the code.
