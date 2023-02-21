FROM golang:latest
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -a /app/cmd/Everlinq/
ENTRYPOINT exec go run cmd/Everlinq/main.go