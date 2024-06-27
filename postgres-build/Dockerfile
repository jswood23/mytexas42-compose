# specify image name
FROM postgres AS initdb
# set db schema
RUN rm -rf /docker-entrypoint-initdb.d/*
COPY mytexas42-schema.sql /docker-entrypoint-initdb.d/