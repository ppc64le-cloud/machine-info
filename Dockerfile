FROM golang:1.20

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download


COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /machine-info

EXPOSE 8080

# Run
CMD ["/machine-info"]