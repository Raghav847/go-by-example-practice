# ---- build stage ----
FROM golang:1.27-alpine AS builder

# go-sqlite3 is a cgo package, so we need a C compiler + musl headers
RUN apk add --no-cache gcc musl-dev

WORKDIR /src

# cache deps first
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# static-ish binary, sqlite3 file will live at /data in the runtime stage
RUN CGO_ENABLED=1 GOOS=linux go build -o /out/shoppinglist .

# ---- runtime stage ----
FROM alpine:3.20

RUN apk add --no-cache ca-certificates tzdata wget \
 && adduser -D -u 10001 app

WORKDIR /app

COPY --from=builder /out/shoppinglist /app/shoppinglist

# sqlite file is opened as ./database.db -> needs a writable dir
RUN mkdir -p /data && chown app:app /data /app
VOLUME ["/data"]
WORKDIR /data

USER app
EXPOSE 8888

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s \
  CMD wget -qO- http://127.0.0.1:8888/lists || exit 1

CMD ["/app/shoppinglist"]
