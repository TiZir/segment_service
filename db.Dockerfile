FROM mysql:8.0

COPY ./init/*.sql /docker-entrypoint-initdb.d/

RUN chmod 777 /docker-entrypoint-initdb.d/app.sql
