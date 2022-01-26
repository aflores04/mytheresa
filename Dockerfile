FROM golang:1.17-alpine

WORKDIR /app

COPY . .

RUN go mod download
RUN go build .

EXPOSE 8080

CMD ["/app/mytheresa"]