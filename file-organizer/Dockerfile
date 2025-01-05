FROM golang:1.23.4-alpine

WORKDIR /app

# Install Air for hot-reload
RUN go install github.com/air-verse/air@latest

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Expose port
EXPOSE 3000

WORKDIR /app/src/cmd/api

# Command to run Air
CMD ["air"]