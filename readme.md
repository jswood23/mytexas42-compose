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

## Helpful Server Commands

To start the CD server in a new process, run the following command:
```bash
sudo ./start.sh
```

To start the CD server in the current terminal, run the following command:
```bash
sudo go run . [SSH passphrase]
```

If you get a `command not found` error after that, run this command first:
```bash
chmod +x start.sh
```

