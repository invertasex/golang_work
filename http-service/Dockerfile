FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o http-service main.go

EXPOSE 8080
CMD ["./http-service"]
