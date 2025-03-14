package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logs information about each incoming request
func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		log.Printf("[REQUEST] %s %s - Started", r.Method, r.URL.Path)

		next(w, r)

		duration := time.Since(startTime)
		log.Printf("[REQUEST] %s %s - Completed in %v", r.Method, r.URL.Path, duration)
	}
}

func main() {
	// Configure logging
	log.SetFlags(log.Ldate | log.Ltime)
	log.Println("Starting server...")

	http.HandleFunc("/", LoggingMiddleware(func(w http.ResponseWriter, r *http.Request) {
		html := `<!DOCTYPE html>
<html>
<head>
    <title>Hello World</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            margin: 0;
            background-color: #f5f5f5;
        }
        .container {
            text-align: center;
            padding: 2rem;
            background-color: white;
            border-radius: 8px;
            box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
        }
        h1 {
            color: #333;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>Hello, Pacvac!</h1>
        <p>This is a simple Go web server.</p>
    </div>
</body>
</html>`

		fmt.Fprintf(w, html)
	}))

	// Health check endpoint
	http.HandleFunc("/up", LoggingMiddleware(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
	}))

	log.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}
