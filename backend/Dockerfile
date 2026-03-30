FROM golang:1.23-alpine AS builder
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/api ./cmd/api

FROM gcr.io/distroless/static:nonroot
COPY --from=builder /bin/api /api
EXPOSE 8080
ENTRYPOINT ["/api"]
