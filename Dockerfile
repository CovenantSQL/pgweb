FROM alpine:3.6
LABEL maintainer="Dan Sosedoff <qi.xiao@covenantsql.com>"
ENV PGWEB_VERSION 0.11.2

RUN \
  apk update && \
  apk add --no-cache ca-certificates openssl wget && \
  update-ca-certificates && \
  rm -rf /var/cache/apk/* && \
  cd /tmp && \
  wget -q https://github.com/CovenantSQL/pgweb/releases/download/v$PGWEB_VERSION/pgweb_linux_amd64.zip && \
  unzip pgweb_linux_amd64.zip -d /usr/bin && \
  mv /usr/bin/pgweb_linux_amd64 /usr/bin/pgweb && \
  rm -f pgweb_linux_amd64.zip

EXPOSE 8081
CMD ["/usr/bin/pgweb", "--bind=0.0.0.0", "--listen=8081"]
