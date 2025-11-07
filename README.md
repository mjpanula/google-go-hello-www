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

# Using Docker Hub

Here’s a step-by-step guide to **log in to Docker Hub and push an image from Windows** (works for both PowerShell and Command Prompt):

---

### 🧩 1. Log in to Docker Hub

Open **PowerShell** or **Command Prompt** and run:

```bash
docker login
```

Then enter your **Docker Hub username** and **password** (or a **personal access token** if you use 2FA).

Alternatively, you can log in non-interactively:

```bash
docker login -u your-username -p your-password
```

> ⚠️ Avoid storing passwords in plain text. It’s safer to use interactive login or a Docker credential helper.

---

### 🏗️ 2. Tag your local image

Docker Hub images follow the format:
`<username>/<repository>:<tag>`

Example:

```bash
docker tag myapp:latest your-username/myapp:latest
```

You can check your local images with:

```bash
docker images
```

---

### 🚀 3. Push the image to Docker Hub

Once tagged, push the image:

```bash
docker push your-username/myapp:latest
```

Docker will upload each layer of your image to Docker Hub.

---

### 🔍 4. Verify upload

You can visit your repository on [https://hub.docker.com/repositories](https://hub.docker.com/repositories) and confirm that the image and tag appear.

---

### 🧰 Example end-to-end workflow

```bash
# Build your image
docker build -t myapp .

# Tag it for Docker Hub
docker tag myapp:latest matti123/myapp:latest

# Log in
docker login

# Push it
docker push matti123/myapp:latest
```

---

### 💡 Optional: Use a Personal Access Token

If you have 2-factor authentication enabled, create a **Personal Access Token (PAT)** in Docker Hub and use it instead of your password when logging in.

---