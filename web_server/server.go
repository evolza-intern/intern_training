package main

import (
	"fmt"
	"net/http"
)

func main() {

	// Register handler function to serve the index.html file for all requests
	http.HandleFunc("/", handler)

	fmt.Println("Server started at http://localhost:8080")

	// Start HTTP server listening on port 8080, nil uses the default router http.ServeMux
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed:", err)
	}
}

// handler serves the index.html file for all incoming HTTP requests
func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
