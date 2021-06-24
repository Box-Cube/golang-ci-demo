FROM alpine:3.14.0

EXPOSE 80
COPY golang-demo /usr/local/bin/

ENTRYPOINT [ "golang-demo" ]