package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// This package is named main

func main() {
	// os.Open() will return either f (open file handle) or err (an error)
	f, err := os.Open("testapp.log")

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
		if strings.Contains(s, "ERROR") {
			fmt.Println(s)
		}
	}
}
