FROM golang:1.23-alpine AS api-builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .


RUN go build -o main .

FROM alpine:latest

COPY --from=api-builder /app/main ./
COPY --from=api-builder /app/.env ./


EXPOSE $API_PORT

CMD ["/main"]