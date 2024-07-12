#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
    CREATE DATABASE mytexas42;
    GRANT ALL PRIVILEGES ON DATABASE mytexas42 TO jswood-admin;
EOSQL
