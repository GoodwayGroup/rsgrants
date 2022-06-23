FROM alpine:3.16.0

COPY rsgrants /usr/local/bin/rsgrants
RUN chmod +x /usr/local/bin/rsgrants

RUN mkdir /workdir
WORKDIR /workdir

ENTRYPOINT [ "/usr/local/bin/rsgrants" ]