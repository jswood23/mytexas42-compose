docker stop container-pg
pause
docker system prune -a
pause
docker volume rm mytexas42-compose_postgres-data
pause
docker-compose up -d
pause
