FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum if they exist
COPY go.* ./

# Initialize module if go.mod doesn't exist
RUN if [ ! -f go.mod ]; then go mod init go-release-tour; fi

# Copy source code
COPY . .

# Build the application
RUN go build -o main .

# Final stage
FROM golang:1.25-alpine

WORKDIR /app

# Install curl for health checks
RUN apk add --no-cache curl

# Copy the binary and static files
COPY --from=builder /app/main .
COPY --from=builder /app/static ./static
COPY --from=builder /app/releases ./releases

EXPOSE 8080

CMD ["./main"]