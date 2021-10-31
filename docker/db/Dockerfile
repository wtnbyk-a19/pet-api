FROM mysql:8.0

RUN chown -R mysql /var/lib/mysql && \
    chgrp -R mysql /var/lib/mysql

COPY ./etc/my.cnf /etc/mysql/conf.d/my.cnf

COPY ./docker/db/init/* /docker-entrypoint-initdb.d/