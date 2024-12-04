FROM golang:1.23-alpine

WORKDIR /usr/src/app/

COPY . .

RUN go build -o bin .

ENV PORT=3000

EXPOSE 3000

RUN mkdir -p cache

CMD ["/usr/src/app/bin"]
