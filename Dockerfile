FROM golang:1.21.0

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o ./docker-gs-ping

CMD ["./docker-gs-ping"]
