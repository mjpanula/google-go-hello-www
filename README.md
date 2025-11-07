# Hello World Go Web Server

A simple HTTP web server written in Go that responds with "Hello, World!"

## Running Locally

```bash
go run main.go
```

Visit http://localhost:8080 in your browser.

## Running with Docker

### Build the Docker image:
```bash
docker build -t hello-world-go .
```

### Run the container:
```bash
docker run -p 8080:8080 hello-world-go
```

Visit http://localhost:8080 in your browser.

### Stop the container:
```bash
docker ps
docker stop <container_id>
```

## Features

- Simple HTTP server on port 8080
- Returns "Hello, World! 🌍" for all requests
- Logs each incoming request
- Multi-stage Docker build for optimized image size
