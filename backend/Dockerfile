# ---- Build Stage ----
FROM golang:1.22-alpine AS builder

WORKDIR /build

# Cache dependencies
COPY go.mod ./
RUN go mod download

# Copy source and build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /build/server .

# ---- Run Stage ----
FROM alpine:3.19 AS runner

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app
COPY --from=builder --chown=65532:65532 /build/server .

EXPOSE 8080
USER 65532:65532
ENTRYPOINT ["./server"]
