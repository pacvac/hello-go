# Go Hello World Web Server

A simple Go web server that displays a "Hello World" HTML page.

## Running Locally

```bash
go run main.go
```

Then visit http://localhost:8080 in your browser.

## Building and Running with Docker

Build the Docker image:

```bash
docker build -t hello-go .
```

Run the container:

```bash
docker run -p 8080:8080 hello-go
```

Then visit http://localhost:8080 in your browser.
