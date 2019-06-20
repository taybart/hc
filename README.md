Useage

```Dockerfile
FROM golang:alpine as builder

# Update ssl and git
RUN apk add --update --no-cache ca-certificates git

RUN mkdir /build
# ADD . /build/
WORKDIR /build
COPY go.mod .
COPY go.sum .

# Download deps
RUN go mod download

COPY . .

# Static build, strip DWARF table and debug symbols
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o main .
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o hc github.com/taybart/hc

# Scratch image
FROM scratch

# Copy ssl certificates for https calls
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy static exe
COPY --from=builder /build/hc /app/healthcheck
COPY --from=builder /build/main /app/
WORKDIR /app

HEALTHCHECK --interval=1s --timeout=1s --start-period=2s --retries=3 CMD [ "./healthcheck" ]
ENTRYPOINT ["./main"]
