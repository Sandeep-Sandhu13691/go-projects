package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", helloHandler)

	// Create a ServeMux
	mux := http.NewServeMux()

	// Associate the ServeMux with your routes
	mux.Handle("/", http.DefaultServeMux)

	// Create an HTTP server with the ServeMux
	server := &http.Server{
		Addr:    ":8080", // You can choose a different port if you prefer
		Handler: mux,
	}

	// Start the server
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "<html><head><style>body { text-align: center; padding: 150px; font-size: 20px; }</style></head><body>Hello, Go!</body></html>")
}
