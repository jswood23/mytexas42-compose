# MyTexas42 Deployment

This repository contains the deployment files for the MyTexas42 project.

## Manual Deployment

To manually deploy staging, run the following command:
```bash
docker compose build backend-staging
docker compose up --no-deps -d backend-staging
```

To manually deploy production, run the following command:
```bash
docker compose build backend-production
docker compose up --no-deps -d backend-production
```