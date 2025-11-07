# Developing in Docker

This guide explains how to develop your Go application using Docker with your local code directory mounted into the container.

## Quick Start: VS Code Attach to Container (Recommended)

This is the simplest workflow for developing with Docker and VS Code.

### Step 1: Start Development Container in PowerShell

```powershell
docker run -d --name go-dev -v ${PWD}:/app -w /app -p 8080:8080 golang:1.21-alpine sleep infinity
```

**What this does:**
- `-d` - Run container in detached (background) mode
- `--name go-dev` - Names the container "go-dev" for easy reference
- `-v ${PWD}:/app` - Mounts your current code directory to `/app` in the container
- `-w /app` - Sets working directory to `/app`
- `-p 8080:8080` - Maps port 8080 from container to host
- `sleep infinity` - Keeps the container running

### Step 2: Attach VS Code to the Container

1. Install the **Dev Containers** extension in VS Code (if not already installed)
2. Press `F1` or `Ctrl+Shift+P` to open the command palette
3. Type and select: **Dev Containers: Attach to Running Container...**
4. Select the `go-dev` container from the list
5. A new VS Code window will open, connected to the container

### Step 3: Open Your Code in the Container

In the new VS Code window:
1. Click **File → Open Folder**
2. Navigate to `/app`
3. Click **OK**

You're now editing code inside the container! All changes sync to your host machine.

### Step 4: Run Your Application

Open a terminal in VS Code (`` Ctrl+` ``) and run:

```sh
go run main.go
```

Or build and run:

```sh
go build -o main .
./main
```

Visit `http://localhost:8080` in your browser to see your app.

### Stop and Clean Up

When you're done developing:

```powershell
docker stop go-dev
docker rm go-dev
```

Or in one command:

```powershell
docker rm -f go-dev
```

---

## Alternative Quick Start Options

### Option 1: Interactive Development Container

Start a container with the Go compiler and mount your code directory:

```powershell
docker run -it --rm `
  -v ${PWD}:/app `
  -w /app `
  -p 8080:8080 `
  golang:1.21-alpine `
  /bin/sh
```

**Explanation:**
- `-it` - Interactive terminal
- `--rm` - Automatically remove container when it exits
- `-v ${PWD}:/app` - Mount current directory to `/app` in container
- `-w /app` - Set working directory to `/app`
- `-p 8080:8080` - Map port 8080 from container to host
- `golang:1.21-alpine` - Use Go 1.21 Alpine image
- `/bin/sh` - Start a shell

### Option 2: Run and Auto-Rebuild

For automatic rebuilding and running when you save changes, you can use:

```powershell
docker run -it --rm `
  -v ${PWD}:/app `
  -w /app `
  -p 8080:8080 `
  golang:1.21-alpine `
  sh -c "go run *.go"
```

### Option 3: Using Docker Compose (Recommended)

Create a `docker-compose.yml` file:

```yaml
version: '3.8'

services:
  dev:
    image: golang:1.21-alpine
    working_dir: /app
    volumes:
      - .:/app
    ports:
      - "8080:8080"
    command: sh -c "go run *.go"
    stdin_open: true
    tty: true
```

Then run:

```powershell
docker-compose up
```

## Development Workflow

### 1. Start the Development Container

```powershell
docker run -it --rm -v ${PWD}:/app -w /app -p 8080:8080 golang:1.21-alpine /bin/sh
```

### 2. Inside the Container

Once inside the container, you can:

**Install dependencies:**
```sh
go mod download
```

**Run the application:**
```sh
go run main.go
```

**Build the application:**
```sh
go build -o main .
```

**Run tests:**
```sh
go test ./...
```

**Install additional tools (like air for hot reload):**
```sh
go install github.com/cosmtrek/air@latest
air
```

### 3. Edit Code on Host

Edit your code files on your Windows host using VS Code or any editor. Changes are immediately reflected in the container because the directory is mounted.

### 4. Test in Browser

Open your browser and navigate to:
```
http://localhost:8080
```

## Hot Reload with Air

For automatic reloading when you save changes:

### 1. Create `.air.toml` configuration:

```toml
root = "."
testdata_dir = "testdata"
tmp_dir = "tmp"

[build]
  args_bin = []
  bin = "./tmp/main"
  cmd = "go build -o ./tmp/main ."
  delay = 1000
  exclude_dir = ["assets", "tmp", "vendor", "testdata"]
  exclude_file = []
  exclude_regex = ["_test.go"]
  exclude_unchanged = false
  follow_symlink = false
  full_bin = ""
  include_dir = []
  include_ext = ["go", "tpl", "tmpl", "html"]
  include_file = []
  kill_delay = "0s"
  log = "build-errors.log"
  poll = false
  poll_interval = 0
  rerun = false
  rerun_delay = 500
  send_interrupt = false
  stop_on_error = false

[color]
  app = ""
  build = "yellow"
  main = "magenta"
  runner = "green"
  watcher = "cyan"

[log]
  main_only = false
  time = false

[misc]
  clean_on_exit = false

[screen]
  clear_on_rebuild = false
  keep_scroll = true
```

### 2. Run with Air:

```powershell
docker run -it --rm `
  -v ${PWD}:/app `
  -w /app `
  -p 8080:8080 `
  golang:1.21-alpine `
  sh -c "go install github.com/cosmtrek/air@latest && air"
```

## Troubleshooting

### Port Already in Use

If port 8080 is already in use, change the port mapping:
```powershell
-p 9090:8080
```

### Permission Issues

On Linux/Mac, you might need to run with user mapping:
```bash
-u $(id -u):$(id -g)
```

### Changes Not Reflecting

Ensure the volume mount is correct and the path uses forward slashes on Windows when needed.

## Best Practices

1. **Use .dockerignore** - Exclude unnecessary files from being watched
2. **Keep containers ephemeral** - Use `--rm` flag to clean up
3. **Use named volumes for Go cache** - Speed up builds:
   ```powershell
   -v go-mod-cache:/go/pkg/mod
   ```
4. **Combine with production Dockerfile** - Keep this separate from your production build

## Production Build

When ready to build for production, use the multi-stage Dockerfile:

```powershell
docker build -t google-go-hello-www .
docker run -p 8080:8080 google-go-hello-www
```

This creates an optimized production image without development tools.
