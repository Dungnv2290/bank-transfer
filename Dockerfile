FROM golang:1.19 AS base
WORKDIR /app
COPY . .

FROM base AS debugger
WORKDIR /app
COPY . .
RUN go get github.com/go-delve/delve/cmd/dlv
EXPOSE 3001 40000
ENTRYPOINT ["dlv", "debug", "--listen=:40000", "--headless", "--accept-multiclient", "--continue", "--api-version=2"]

FROM base AS development
WORKDIR /app
COPY . .
RUN go get github.com/pilu/fresh
RUN go install -v github.com/pilu/fresh@latest
EXPOSE 3001
ENTRYPOINT ["fresh"]

FROM base AS builder
WORKDIR /app
COPY . .
RUN git config --global --add safe.directory $(pwd)
RUN go build -buildvcs=false -v -a --installsuffix cgo --ldflags="-s" -o main

FROM alpine:latest AS production
RUN apk --no-cache add ca-certificates
COPY --from=builder /app .
EXPOSE 3001
ENTRYPOINT ["./main"]