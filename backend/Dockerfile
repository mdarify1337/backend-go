FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Explicit build
RUN go build -o /app/main ./main.go

EXPOSE 3001

CMD ["/app/main"]
