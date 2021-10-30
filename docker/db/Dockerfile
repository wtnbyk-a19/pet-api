FROM mysql:8.0

RUN chown -R mysql /var/lib/mysql && \
    chgrp -R mysql /var/lib/mysql

COPY ./docker/db/init/* /docker-entrypoint-initdb.d/