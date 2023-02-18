FROM alpine:3.14
COPY bin/webhook-sample /usr/bin
ENTRYPOINT ["webhook-sample"]
