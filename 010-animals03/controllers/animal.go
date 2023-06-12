package controllers

import (
	"net/http"
	"regexp"
)

type animalController struct {
	animalIdRegexp *regexp.Regexp
}

func newAnimalController() *animalController {
	return &animalController{
		animalIdRegexp: regexp.MustCompile(`^/animals/(\d+)/?`),
	}
	// ^ 				-> at the beginning of the line
	// /animals/ 	-> literal
	// (\d+) 		-> capturing group for one or more digits
	// /? 			-> 0 or 1 /
	// So, it looks for /animals/ followed by digits followed by an optional /
}

func (animalCtrl animalController) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Called the animal controller."))
}
