# mytexas42-compose

Docker Compose orchestration + nginx + Let's Encrypt + a tiny Go deploy server for the self-hosted MyTexas42 stack. Brings up the backend (`../my-texas-42-backend`), the frontend (`../my-texas-42-frontend`), and PostgreSQL — one of each per environment (staging + production), behind a single nginx reverse proxy.

## Services (`compose.yaml`)

| Service | Image / Build | Port (host) | Notes |
|---|---|---|---|
| `db-staging` | built from `postgres-build/` | `${POSTGRES_EXTERNAL_PORT}` | volume `postgres-staging-data` |
| `db-production` | built from `postgres-build/` | not exposed | volume `postgres-production-data` |
| `backend-staging` | `${BACKEND_STAGING_REPO}` (`../my-texas-42-backend`) | `${BACKEND_STAGING_PORT}` | `ENVIRONMENT=staging` |
| `backend-production` | same source | `${BACKEND_PRODUCTION_PORT}` | `ENVIRONMENT=production` |
| `frontend-staging` | `${FRONTEND_STAGING_REPO}` (`../my-texas-42-frontend`) | `${FRONTEND_STAGING_PORT}` | `REACT_APP_ENVIRONMENT=staging` |
| `frontend-production` | same source | `${FRONTEND_PRODUCTION_PORT}` | `REACT_APP_ENVIRONMENT=production` |
| `nginx` | `./nginx` | 80, 443 | reverse proxy, SSL termination, SNI routing |
| `certbot` | `certbot/certbot` | — | webroot ACME for `mytexas42.com` certs |

Both backends share the same source tree but build into different containers via `ENVIRONMENT`.

## nginx routing (`nginx/nginx.conf`)

```
mytexas42.com                → 301 → www
www.mytexas42.com / app.…    → frontend-production:3000
staging-app.mytexas42.com    → frontend-staging:3000
api.mytexas42.com            → backend-production:8080
staging-api.mytexas42.com    → backend-staging:8080
```

Port 80 redirects to 443 except for `/.well-known/acme-challenge/*` (certbot webroot).

## Database

`postgres-build/Dockerfile` extends the official `postgres:latest` and runs `init-db.sh` to create the `mytexas42` DB and grant perms. The schema in `postgres-build/mytexas42-schema.sql` is **not** the source of truth — the backend creates tables itself from `sql_scripts/table-schema.go`. Treat the SQL file as historical; if you change the schema, change the Go file.

## Deploy server (`main.go`)

A bare `net/http` server (port from `system.GetPort()`) that exposes:

```
GET /health                 → 200 OK
GET /deploy/staging         → docker compose up -d --build backend-staging frontend-staging
GET /deploy/production      → same for production
GET /deploy/all
GET /stop/all
```

Authentication is via the `ADMIN_PASSWORD` env var (compared as a query / header — see the handler). It's intentionally minimal — don't expose this server to the open internet without an additional layer in front.

## `.env`

The compose file pulls everything from `.env`. Required keys (see the file for current values):

```
POSTGRES_USER, POSTGRES_PASSWORD, DB_NAME
POSTGRES_PRODUCTION_HOST_NAME, POSTGRES_STAGING_HOST_NAME
POSTGRES_PRODUCTION_PORT, POSTGRES_EXTERNAL_PORT
BACKEND_PRODUCTION_PORT, BACKEND_STAGING_PORT
FRONTEND_PRODUCTION_PORT, FRONTEND_STAGING_PORT
BACKEND_STAGING_REPO, BACKEND_PRODUCTION_REPO       # paths to ../my-texas-42-backend
FRONTEND_STAGING_REPO, FRONTEND_PRODUCTION_REPO     # paths to ../my-texas-42-frontend
STAGING_USER_POOL_NAME, STAGING_USER_POOL_APP_KEY
PRODUCTION_USER_POOL_NAME, PRODUCTION_USER_POOL_APP_KEY
REACT_APP_*_API_PATH, REACT_APP_*_WEBSOCKET_API_PATH
GIN_MODE=release
ADMIN_PASSWORD
```

`.env` is committed and contains real credentials — **rotate the secrets and `.gitignore` it before this repo goes public**.

## Run

```bash
# Bring up everything
docker compose up -d

# Just one environment
docker compose up -d db-staging backend-staging frontend-staging

# Rebuild after backend/frontend code changes
docker compose up -d --build backend-staging frontend-staging

# Logs
docker compose logs -f backend-production
```

Initial cert issuance: run `certbot` once with the right hostnames before flipping nginx to HTTPS-only. The `certbot` service is designed to renew, not to issue from scratch.

## Conventions / gotchas

- The two database services share the same `POSTGRES_USER` / `POSTGRES_PASSWORD`; if you need to isolate, fork the env-file pattern.
- `db-staging` exposes its port externally for ad-hoc `psql` debugging; `db-production` does not.
- nginx mounts `/etc/nginx` **read-only** from `./nginx`, so config changes require a `docker compose restart nginx`.
- The frontend builds bake `REACT_APP_ENVIRONMENT` and the API/WS URLs into the bundle at build time — changing those env values requires a rebuild, not a restart.
- The backend Dockerfile builds to `scratch`, so the binary cannot shell out and there is no `/etc/passwd`. Don't add code that calls `exec.Command` from there.

## Related repos

- `../my-texas-42-backend` — Go server source
- `../my-texas-42-frontend` — React client source
- `../my-texas-42-react-app` — legacy AWS Lambda implementation (reference only)
