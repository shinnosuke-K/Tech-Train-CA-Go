FROM mysql:latest

EXPOSE 3306

COPY ./my.cnf /etc/mysql/conf.d/my.cnf
COPY ./db /docker-entrypoint-initdb.d
