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
