# Functions

## Declaring functions

Basic syntax is:

* `func` keyword
* function name
* `()` (optionally containing parameters)
* optional return type (required if returning anything)
* `{}` (wrapping the function body)

```golang
func startServer() {
   fmt.Println("Starting server...")
   // TODO: start the web server
   fmt.Printf("Server started on %v:%v", serverUrl, serverPort)
   // will not be happy because we aren't passing serverUrl and serverPort
}
```

## Parameters

Technical details:

* Parameters are in the parentheses in the function declaration
* Arguments are in the parentheses in the function call

Let's pass the port to the function.

```golang
// fully spelled out syntax
func startServer(serverPort int, maxRetries int) {
   fmt.Println("Starting server...", serverPort, maxRetries)
   // TODO: start the web server
   fmt.Printf("Server started on port %v", serverPort)
}

// shortened syntax -- don't duplicate consecutive types
func startServer(serverPort, maxRetries int) {
   fmt.Println("Starting server...", serverPort, maxRetries)
   // TODO: start the web server
   fmt.Printf("Server started on port %v", serverPort)
}

// call: startServer(9200, 5)
```

## Return values

Let's have `startServer()` tell us if it works.

* Add the return type(s) to the function declaration
* Return the value

```golang
func startServer(serverPort, maxRetries int) bool {
   fmt.Println("Starting server...", serverPort, maxRetries)
   // TODO: start the web server
   fmt.Printf("Server started on port %v", serverPort)

   return true
}

// call ignoring return: startServer(9200, 5)

// call catching value: isStarted := startServer(9200, 5)
```

But, if something fails, how did it fail? (So we can log it.) So, let's return an `error`.

```golang
func startServer(serverPort, maxRetries int) error {
   fmt.Println("Starting server...", serverPort, maxRetries)
   // TODO: start the web server
   fmt.Printf("Server started on port %v", serverPort)

   return nil // all is well
}

func main() {
   err := startServer(9200, 5)
   if (err != nil) {
      // handle error
   }
}
```

So, how would we return the error?

```golang
import (
   "fmt"
   "error" 
)

func startServer(serverPort, maxRetries int) error {
   fmt.Println("Starting server...", serverPort, maxRetries)
   // TODO: start the web server
   fmt.Printf("Server started on port %v", serverPort)

   if iWantAnError {
      return errors.New("Something went wrong")
   }

   return nil // all is well
}

func main() {
   err := startServer(9200, 5)
   if (err != nil) {
      // handle error
   }
}
```

Go favors string errors, which is a logging problem if we want structured logs. We can do custom errors by declaring a struct and adding an `Error()` method. To get to the members, we need some extra work, it seems.

```golang
import (
   "errors"
   "fmt"
   "net/http"
)

// The custom error.
// Imagine we have other, useful information here that we want to log or use for recovery
type RequestError struct {
   StatusCode int

   Err error
}

// required Error() function
func (r *RequestError) Error() string {
   return r.Err.Error()
}

// We can write functions to identify specific statuses or status families.
// This approach can be useful for deciding how to handle an error.
func (r *RequestError) Temporary() bool {
   return r.StatusCode == http.StatusServiceUnavailable // 503
}

func doRequest() error {
   return &RequestError{
      StatusCode: 503,
      Err:        errors.New("unavailable"),
   }
}

func main() {
   err := doRequest()
   fmt.Println(err)  // prints "unavailable" (error text only)
   // fmt.Println(err.StatusCode) fails
   // we could write Error to return a string with all the fields, but
   // that doesn't play nicely with structured logging systems

   // To get what's in the error we need to extract it like this.
   // or we could do this inside the err != nil block and skip ok
   re, ok := err.(*RequestError)
   if ok {
      // now we can use error members from re
      fmt.Println("re", re, re.StatusCode)   // unavailable 503
      // or call methods
      fmt.Println(re.Temporary())
   }
}
```

Functions can return more than one value.

```golang
import (
   "fmt"
   "error" 
)

func startServer(serverPort, maxRetries int) (int, error) {
   fmt.Println("Starting server...", serverPort, maxRetries)
   // TODO: start the web server
   fmt.Printf("Server started on port %v", serverPort)

   if iWantAnError {
      return nil, errors.New("Something went wrong")
   }

   return serverPort, nil // all is well
}

func main() {
   port, err := startServer(9200, 5)
   if (err != nil) {
      // handle error
   }
   fmt.Println(port, err)
}
```

If we don't care about one of the return values, we can replace it with _.

```golang
func main() {
   _, err := startServer(9200, 5)
   if (err != nil) {
      // handle error
   }
}
```
