FROM golang:1.24.1-alpine

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN go build -o main main.go

EXPOSE 8080

CMD ["./main"]