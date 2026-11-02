# StudySmart

StudySmart uses Vue 3 and Vite for the frontend, Go/Gin for the API, and PostgreSQL 16 for persistence.

## Run locally

```sh
docker compose up --build
```

Open http://localhost:5173. The API is available at http://localhost:8080, and PostgreSQL is exposed on port 5433. The health endpoint is `GET /api/v1/health`; the frontend health page is `/health`.

Copy `.env.example` to `.env` to customize local settings. `.env` is ignored by Git. Use a unique password outside local development.

The main frontend app lives in `src/`; `frontend/` contains its development/production Dockerfile and Nginx configuration. Backend handlers, services, repositories, and configuration live under `backend/internal/`.
