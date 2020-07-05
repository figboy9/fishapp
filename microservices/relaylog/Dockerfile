FROM golang:1.13-alpine AS builder
WORKDIR /src

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o main .

FROM alpine AS prod
WORKDIR /app

COPY --from=builder /src/main .
COPY conf/conf.yml /app/conf/conf.yml

CMD ["./main"]