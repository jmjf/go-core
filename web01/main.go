package main

// [net/http](https://pkg.go.dev/net/http) is the golang http package
// it provides both http call methods and http server methods
import (
	"log"
	"net/http"
)

func main() {
	// HandleFunc registers a handler function for a route
	// Handlers accept a ResponseWriter and a *Request
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Write writes data to the HTTP connection.
		// It accepts a slice of bytes, so we need to cast here.
		w.Write([]byte("Hello, world."))
	})

	// ListenAndServe starts listening for HTTP requests.
	// It accepts an address (just a port here, so will be on localhost) and a handler.
	// The handler is where actually the router (in golang terms mux) where the handlers are registered.
	// We're using the default handler so can pass nil
	err := http.ListenAndServe(":9070", nil)
	// as usual, we want to check for errors
	if err != nil {
		log.Fatal(err)
	}
}
