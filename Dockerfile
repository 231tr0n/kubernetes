FROM alpine:latest

ENV PORT=":8080"

RUN apk add gcompat

WORKDIR /root/trial

COPY trial .

CMD ["./trial"]
