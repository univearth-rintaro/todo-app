# Development stage
FROM golang:1.22-bullseye as dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go install github.com/cosmtrek/air@latest

EXPOSE 5050

CMD ["air"]
