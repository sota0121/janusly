# compose config service.build.context should be '.' that means the root of the project

# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.20 as builder
WORKDIR /app

# Copy go mod and sum files
# Download dependencies and cache them
COPY server/go.mod server/go.sum ./
RUN go mod download

# Copy the rest of the files
COPY server/ ./

# Build the application
RUN make linux-build

# Host stage
FROM alpine:latest
COPY --from=builder /app/bin/ /app/bin/

CMD ["/app/bin/server"]
