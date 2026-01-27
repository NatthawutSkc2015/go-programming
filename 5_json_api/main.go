package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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
	// Register handler
	http.HandleFunc("/hello", helloHandler)

	// Start server
	fmt.Println("=== JSON API Server ===")
	fmt.Println("Server starting on http://localhost:8080")
	fmt.Println("Endpoint: POST /hello")
	fmt.Println("\nExample usage with curl:")
	fmt.Println(`  curl -X POST http://localhost:8080/hello \`)
	fmt.Println(`    -H "Content-Type: application/json" \`)
	fmt.Println(`    -d '{"name": "Somchai"}'`)
	fmt.Println("\nPress Ctrl+C to stop the server")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
