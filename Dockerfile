FROM golang:1.14-buster as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o server

FROM alpine:3

ARG BUILD_VERSION='0.0.0'

RUN apk add --no-cache ca-certificates

ENV PORT 8080
EXPOSE 8080

ENV BUILD_VERSION=${BUILD_VERSION}

COPY --from=builder /app/server /server

CMD ["/server"]
