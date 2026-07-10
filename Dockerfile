# syntax=docker/dockerfile:1

# ---- build stage ----
FROM golang:1.24-alpine AS build
WORKDIR /src

# Download modules first so this layer caches unless go.mod/go.sum change.
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# The MySQL driver is pure Go, so CGO can stay off -> a fully static binary
# that runs on the distroless "static" base with no libc.
RUN CGO_ENABLED=0 GOOS=linux go build \
      -trimpath \
      -ldflags="-s -w" \
      -o /out/exporter ./cmd/exporter

# ---- runtime stage ----
FROM gcr.io/distroless/static-debian12:nonroot
LABEL org.opencontainers.image.source="https://github.com/Kapel/prom-azerothcore-exporter" \
      org.opencontainers.image.description="Prometheus exporter for AzerothCore (WoW WotLK) private servers" \
      org.opencontainers.image.licenses="MIT"

COPY --from=build /out/exporter /exporter

# Metrics endpoint (override with PORT).
EXPOSE 7000
USER nonroot:nonroot
ENTRYPOINT ["/exporter"]
