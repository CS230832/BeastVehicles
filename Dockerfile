FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o bin/app.exe cmd/main.go

EXPOSE 8080

CMD ["./bin/app.exe"]