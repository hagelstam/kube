FROM golang:1.23-alpine

WORKDIR /usr/src/app/

COPY . .

RUN go build -o bin .

ENV PORT=8080

EXPOSE 8080

RUN mkdir -p cache

CMD ["/usr/src/app/bin"]
