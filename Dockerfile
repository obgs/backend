FROM golang:1.19-bullseye AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY ./ ./

RUN mkdir /build \
  && CGO_ENABLED=0 go build -ldflags="-s -w" -o /build ./...

FROM debian:bullseye

COPY --from=builder /build/backend /usr/local/bin/

RUN ["backend", "-h"]

CMD ["backend"]
