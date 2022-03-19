FROM golang:1.17.8-alpine3.14

WORKDIR /app

COPY . .

RUN go build -o bareksa_test app.go

EXPOSE 4456

CMD ["/app/bareksa_test"]