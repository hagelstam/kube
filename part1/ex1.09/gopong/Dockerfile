FROM golang:1.23-alpine

WORKDIR /app

COPY . .

RUN go build -o bin .

ENV PORT=8081

EXPOSE 8081

CMD ["/app/bin"]
