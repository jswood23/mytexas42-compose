docker stop container-pg
pause
docker system prune -a
pause
docker volume rm mytexas42-compose_postgres-data
pause
set /p POSTGRESPW=Set PostgreSQL password:
docker-compose up -d -e POSTGRESPW=%POSTGRESPW%
pause
