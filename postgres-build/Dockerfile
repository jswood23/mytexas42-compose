# specify image name
FROM postgres AS initdb

RUN rm -rf /docker-entrypoint-initdb.d/*

COPY init-db.sh /docker-entrypoint-initdb.d/
RUN chmod +x /docker-entrypoint-initdb.d/init-db.sh
