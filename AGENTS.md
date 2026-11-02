# AGENTS.md — StudySmart Project Rules for AI Agents

## Project Overview
This is a full-stack web application built with **Vue 3** (frontend), **Go/Gin** (backend), and **PostgreSQL 16** (database), orchestrated via **Docker Compose**. The existing app and Vite package live at the repository root (`src/`, `package.json`); `frontend/` holds the frontend Dockerfile and Nginx configuration.

---

## Architecture

### Frontend (Vue 3 + Vite + Tailwind CSS)
- **Framework**: Vue 3 with Composition API
- **Build Tool**: Vite
- **Styling**: Tailwind CSS
- **State Management**: Pinia
- **Routing**: Vue Router
- **API Client**: Axios

### Backend (Go 1.22+ / Gin)
- **Framework**: Gin HTTP framework
- **Architecture**: Clean Architecture (3-tier: Handler → Service → Repository)
- **ORM**: GORM with PostgreSQL driver
- **Logging**: Structured logging via `log/slog`
- **Config**: Environment variable-based configuration

### Database
- **Engine**: PostgreSQL 16
- **Migrations**: Managed via GORM AutoMigrate (dev) / manual SQL (production)

---

## Coding Standards

### Vue / Frontend Rules
1. Always use `<script setup>` syntax (Composition API) for Vue components. Do not use Options API.
2. Use `defineProps` and `defineEmits` for component contracts.
3. Keep components small and single-responsibility.
4. Place reusable logic in composables (`src/composables/`).
5. Use Pinia stores for shared state — avoid prop drilling beyond 2 levels.
6. Use `axios` via the centralized API service (`src/services/api.js`) for all HTTP calls.
7. All routes are defined in `src/router/index.js`. Use lazy-loaded components for views.
8. Tailwind CSS utility classes are preferred over custom CSS. Use `@apply` sparingly.

### Go / Backend Rules
1. Follow idiomatic Go: short variable names in small scopes, descriptive names in larger scopes.
2. **Always** handle errors explicitly — never ignore returned errors.
3. Use the 3-tier Clean Architecture pattern:
   - `handler/` — HTTP request parsing, response writing, input validation.
   - `service/` — Business logic, orchestration between repositories.
   - `repository/` — Data access, database queries, model definitions.
4. Use `internal/` package to prevent external imports.
5. Configuration is loaded from environment variables via `internal/config/`.
6. All API responses use a consistent JSON envelope:
   ```json
   { "success": true, "data": {}, "message": "ok" }
   ```
7. Use `log/slog` for structured logging. No `fmt.Println` in production code.
8. Run `go fmt`, `go vet`, and `golint` before committing.

### Docker & Environment Rules
1. All services are defined in `docker-compose.yml` at the project root.
2. Environment variables are the single source of truth for configuration.
3. Backend Dockerfile uses multi-stage builds (builder → runner).
4. Frontend Dockerfile supports both dev (hot reload) and production (nginx) modes.
5. Never hardcode connection strings, ports, or secrets in source code.
6. PostgreSQL data is persisted via a named Docker volume.

---

The active Vue app lives in the repository root (`src/`, `package.json`, and `vite.config.js`). `frontend/` holds the Dockerfile and Nginx config. The following tree is a general architecture reference, not a list of current files.

## Directory Structure Reference
```
study-smart/
├── AGENTS.md
├── docker-compose.yml
├── backend/
│   ├── Dockerfile
│   ├── main.go
│   ├── go.mod
│   └── internal/
│       ├── config/config.go
│       ├── handler/health.go
│       ├── service/health.go
│       └── repository/db.go
└── frontend/
    ├── Dockerfile
    ├── vite.config.js
    ├── tailwind.config.js
    ├── postcss.config.js
    ├── package.json
    ├── index.html
    └── src/
        ├── main.js
        ├── App.vue
        ├── style.css
        ├── router/index.js
        ├── stores/health.js
        ├── services/api.js
        └── views/
            ├── HomeView.vue
            └── HealthView.vue
```

---

## API Endpoints

| Method | Path              | Description          |
|--------|-------------------|----------------------|
| GET    | /api/v1/health    | Health check status  |

---

## Running the Project
```bash
docker compose up --build
```
- Frontend: http://localhost:5173
- Backend API: http://localhost:8080
- PostgreSQL: localhost:5433
