Example dockerfile

```Dockerfile
FROM golang:alpine as builder

# Update ssl and git
RUN apk add --update --no-cache ca-certificates git
RUN addgroup -S app && adduser -S -G app app

RUN mkdir /build
WORKDIR /build

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o hc github.com/taybart/hc

# Download deps
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# Static build, strip DWARF table and debug symbols
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o main .

# Scratch image
FROM scratch
COPY --from=builder /etc/passwd /etc/passwd
USER app

# Copy ssl certificates for https calls
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/


# Copy static exe
COPY --from=builder /build/hc /app/hc
COPY --from=builder /build/main /app/
WORKDIR /app

HEALTHCHECK --interval=1s --timeout=1s --start-period=2s --retries=3 CMD [ "./hc" ]
ENTRYPOINT ["./main"]
```
