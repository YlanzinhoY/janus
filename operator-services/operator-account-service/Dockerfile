FROM golang:1.25.5-alpine3.22  AS builder

LABEL authors="lanzinho"


WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -ldflags='-s -w' -o  app cmd/main.go

EXPOSE 3000

FROM alpine:3.23.3 AS final

WORKDIR /app

COPY  --from=builder /app/app .

EXPOSE 3000

ENTRYPOINT ["./app"]