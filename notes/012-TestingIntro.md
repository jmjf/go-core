# Testing

It's time to learn how to write tests.

## Testing basics

Golang has a built in `testing` package and `go test` to run tests. Tests accept `*testing.T` and use it for testing specific features.

```golang
// tests accept a *testing.T
// tests do not return a value
func TestExample(t *testing.T) {
   // define the value you expect to receive
   // get the value from the code under test

   // if the values don't match, report an error from the test
   if got != expected {
      t.Errorf("<error message> Expected: %v; Got: %v", expected, got)
   }
}
```

## Packages

Packages available from the standard library include:

* `testing` -- the core testing features [docs](https://pkg.go.dev/testing)
* `testing/quick` -- a few utility functions to help with black box testing; frozen
* `testing/iotest` -- utility readers and writers for IO testing [docs](https://pkg.go.dev/testing/iotest)
* `net/http/httptest` -- utilities for testing HTTP servers including injecting requests and recording responses [docs](https://pkg.go.dev/net/http/httptest)

Other packages from the community that may be of interest:

* Testify -- adds an assertion API, test suites, mocks and other features
* Ginkgo -- adds a BDD-style assertion API (Describe, When, Context, It)
* GoConvey -- browser-based test results interface; API is "different"
* httpexpect -- HTTP testing with a readable API (e.GET("route")...)
* gomock -- mocking framework from Google
* go-sqlmock -- mockable database provider

Related to `pgx`, which I'm using elsewhere, `pgxmock` provides a mock `pgx` for testing.

**COMMIT:** DOCS: outline of native and community testing tools
