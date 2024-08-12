FROM alpine:latest

ENV PORT=":8080"

WORKDIR /root/trial

COPY trial .

ENTRYPOINT ["trial"]
