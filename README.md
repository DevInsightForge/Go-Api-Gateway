# API Gateway

This is a simple API gateway implemented in Go, which works as a reverse proxy to route requests based on URL prefixes.

## Project Structure

```plaintext
project-root/
├── cmd/
│   └── api-gateway/
│       └── main.go         # Application entry point
├── internal/
│   └── proxy/
│       ├── proxy.go        # Proxy logic
│       └── config.go       # Configuration logic
├── configs/                # Configuration files for different environments
│   ├── config.yaml         # Common config for testing
├── go.mod                  # Go module file
├── go.sum                  # Go module dependencies file
└── README.md               # Project README
