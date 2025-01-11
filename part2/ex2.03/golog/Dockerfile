FROM golang:1.23-alpine

WORKDIR /app

COPY . .

RUN go build -o bin .

ENV PORT=3000

EXPOSE 3000

CMD ["/app/bin"]
