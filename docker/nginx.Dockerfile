FROM nginx:1.24-alpine

WORKDIR /usr/src/app

COPY ./docker/nginx.conf /etc/nginx/nginx.conf
COPY ./docker/resolve-localdomain.sh /docker-entrypoint.d/30-resolve-localdomain.sh

RUN chmod +x /docker-entrypoint.d/30-resolve-localdomain.sh
