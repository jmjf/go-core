package main

// This package is named main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// get parameters from the command line
	// flag.String takes the name, default value and help text for the parameter
	path := flag.String("path", "testapp.log", "Path to log file to read")
	level := flag.String("level", "ERROR", "Log level to show. [DEBUG | INFO | WARNING | ERROR | CRITICAL]")

	// parse the flags and get the values (or default)
	flag.Parse()

	// os.Open() will return either f (open file handle) or err (an error)
	// path is a string pointer, so need to use *path
	f, err := os.Open(*path)

	// if we have an error (it isn't nil/null), log it and fail
	if err != nil {
		log.Fatal(err)
	}

	// If we're here, we have a file handle
	// ensure we close the file when the function ends, but note it now so we don't forget
	defer f.Close()
	// create a buffered i/o reader for the file so we can get data from it
	r := bufio.NewReader(f)
	// loop over the reader forever (or until we bail)
	for {
		// read a string from the reader until we hit \n (newline)
		s, err := r.ReadString('\n')
		// if we get an error (including reading past end of file), break out of the loop
		if err != nil {
			break
		}
		// if the line is an error (contains 'ERROR'), print it
		// level is a string pointer
		if strings.Contains(s, *level) {
			fmt.Println(s)
		}
	}
}
