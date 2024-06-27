# MyTexas42 Deployment

This repository contains the deployment files for the MyTexas42 project.

## Automated Deployment

The web service in this repository manages CI/CD for the project.

To deploy to staging, use this endpoint:
```
GET /deploy/staging
```

To deploy to production, use this endpoint:
```
GET /deploy/production
```

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
