FROM golang:1.23-alpine

WORKDIR /app

COPY . .

RUN go build -o bin .

ENV PORT=3001

EXPOSE 3001

CMD ["/app/bin"]
