# Build a simple program to work with a log file

## Read a specific file and report errors

* Create a working directory (`logfilter01`)
* Put the sample log in the directory (`testapp.log`)

* Create a `go.mod` file with the module name and go version
  * Defines the workspace/module

* Create the file for the program (`main.go`, but "main" could be anything)
* Comments in the file explain what's happening

* In terminal, `go run .` in the `logfilter01` directory and see four lines (no response from identity provider)
  * `go run` will compile and run the code
  * `go run .` will look for a `main` package
  * `go run main.go` and get the same result

## Add filename and log level as command line parameters

* `flag.String()` defines string parameters to get from the command line
* `flag.Parse()` parses the command line and sets the defined parameters (either passed value or default)
* Define two parameters `path` and `level`
* `flag.String()` returns a string pointer, so we need to use values as `*path` and `*level` with functions that want the string

* `go run . -help` will show parameters that can be passed and help text
* If parameters are prefixed with name (`-level INFO`) they're ignored
* `go run .` will use defaults
* `go run . -level INFO` will return INFO log lines
