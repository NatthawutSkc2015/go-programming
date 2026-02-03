package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// Request structure for incoming JSON
type HelloRequest struct {
	Name string `json:"name"`
}

// Response structure for outgoing JSON
type HelloResponse struct {
	Message string `json:"message"`
}

// ErrorResponse structure for error messages
type ErrorResponse struct {
	Error string `json:"error"`
}

// loggerMiddleware logs request information and processing time
func loggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Record start time
		startTime := time.Now()

		// Print request start time, method, and URL path
		fmt.Printf("[%s] %s %s - Started\n",
			startTime.Format("2006-01-02 15:04:05"),
			r.Method,
			r.URL.Path)

		// Call the next handler
		next(w, r)

		// Calculate duration
		duration := time.Since(startTime)

		// Print processing duration
		fmt.Printf("[%s] %s %s - Completed in %v\n",
			time.Now().Format("2006-01-02 15:04:05"),
			r.Method,
			r.URL.Path,
			duration)
	}
}

// helloHandler handles the /hello endpoint
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Check if method is POST
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error: "Method not allowed. Please use POST",
		})
		return
	}

	// Read request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error: "Failed to read request body",
		})
		return
	}
	defer r.Body.Close()

	// Parse JSON
	var req HelloRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error: "Invalid JSON format",
		})
		return
	}

	// Validate that name is not empty
	if req.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error: "Name field is required",
		})
		return
	}

	// Create response
	response := HelloResponse{
		Message: fmt.Sprintf("Hello %s", req.Name),
	}

	// Send success response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Register handler with middleware
	http.HandleFunc("/hello", loggerMiddleware(helloHandler))

	// Start server
	fmt.Println("=== JSON API Server ===")
	fmt.Println("Server starting on http://localhost:8080")
	fmt.Println("Endpoint: POST /hello")
	fmt.Println("\nExample usage with curl:")
	fmt.Println(`  curl -X POST http://localhost:8080/hello \`)
	fmt.Println(`    -H "Content-Type: application/json" \`)
	fmt.Println(`    -d '{"name": "Somchai"}'`)
	fmt.Println("\nPress Ctrl+C to stop the server")
	fmt.Println("\n--- Server Logs ---")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
