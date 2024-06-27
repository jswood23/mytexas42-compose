# specify image name
FROM postgres AS initdb
# get os args
ARG POSTGRESPW
# set environment variables
ENV POSTGRES_USER=jswood-admin
ENV POSTGRES_PASSWORD=$POSTGRESPW
ENV POSTGRES_DB=mytexas42
# set db schema
RUN rm -rf /docker-entrypoint-initdb.d/*
COPY mytexas42-schema.sql /docker-entrypoint-initdb.d/