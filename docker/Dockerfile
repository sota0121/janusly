FROM golang:1.20

WORKDIR /app

COPY . .

RUN go build -o main ./api/main.go

CMD ["./main"]
