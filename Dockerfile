FROM golang:1.15-alpine AS builder

WORKDIR /go/src/tinyurl

# Leverage go modules cache to speed up builds
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ./
RUN go build ./cmd/tinyurl-service
# ---
FROM alpine

COPY --from=builder /go/src/tinyurl/tinyurl-service .
EXPOSE 8099

# ENTRYPOINT ["./recipe-service"]

# CMD ["web"]
ENTRYPOINT ["./tinyurl-service"]
