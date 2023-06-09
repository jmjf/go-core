package main

// [net/http](https://pkg.go.dev/net/http) is the golang http package
// it provides both http call methods and http server methods
import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	// HandleFunc registers a handler function for a route
	// Handlers accept a ResponseWriter and a *Request
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		// Query parameters are in the request
		// Request includes a lot of useful information including Method, URL, protocol information, headers, etc.
		// The URL is an object that includes query parameters (RawQuery)
		// Query() returns RawQuery as a map[string][string]
		queryParams := request.URL.Query()
		// let's show the URI and query parameters so we know what we're getting
		log.Println("INFO Received", request.RequestURI, queryParams)

		name := queryParams.Get("name")
		// alternate styles
		// name := request.URL.Query().Get("name")
		// name := queryParams["name"]
		// name := request.URL.Query()["name"][0] // could have more than one with this style

		// Write writes data to the HTTP connection.
		// It accepts a slice of bytes, so we need to cast here.
		response.Write([]byte("<h1>Hello, " + name + ".</h1>"))
	})

	// Let's do a JSON result
	http.HandleFunc("/json", func(response http.ResponseWriter, request *http.Request) {
		queryParams := request.URL.Query()

		log.Println("INFO Received", request.RequestURI, queryParams)

		name := queryParams.Get("name")

		// To return JSON, we need to build something we can convert to a JSON string
		jsonMap := map[string]string{
			"message": "Hello",
			"name":    name,
		}
		// And we need a JSON encoder
		// The JSON encoder accepts a writer. Anything we encode to it is written to the writer.
		jsonEncoder := json.NewEncoder(response)
		// Now we can encode the JSON map as a string
		jsonEncoder.Encode(jsonMap)
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
