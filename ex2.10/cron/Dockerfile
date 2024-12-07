FROM alpine:latest

RUN apk add --no-cache curl jq bash

WORKDIR /app

COPY ./job.sh /app/job.sh

RUN chmod +x /app/job.sh

ENTRYPOINT [ "/app/job.sh" ]
