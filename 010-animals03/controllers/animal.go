package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"animals03/models"
)

type animalController struct {
	animalIdRegexp *regexp.Regexp
}

// newAnimalController() initializes the animal controller.
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

// animalController.parseRequest() decodes JSON from the request body and puts it into an Animal.
// If the JSON can't be converted to an Animal, it returns the error from the decoder.
func (animalCtrl animalController) parseRequest(request *http.Request) (models.Animal, error) {
	// create a JSON decoder and an animal
	decoder := json.NewDecoder(request.Body)
	var animal models.Animal

	// decode the request body into animal -- handle error
	err := decoder.Decode(&animal)
	if err != nil {
		return models.Animal{}, err
	}

	// parsed animal okay, so return it
	return animal, nil
}

// animalController.ServeHTTP() handles the route for animals and dispatches requests to the correct action handlers.
func (animalCtrl animalController) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// does the request URL have an id?
	if request.URL.Path == "/animals" || request.URL.Path == "/animals/" {
		// request does not have an id -> getAll()
		switch request.Method {
		case http.MethodGet:
			animalCtrl.getAll(response, request)
		default:
			response.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		// we have an id

		// get the id using the regexp -- findStringSubmatch() returns a slice of matches
		reMatches := animalCtrl.animalIdRegexp.FindStringSubmatch(request.URL.Path)
		if len(reMatches) == 0 {
			// no matches -- bad request because the required id value is missing or invalid (not a number)
			response.WriteHeader(http.StatusBadRequest)
		}
		// convert the string to an integer
		id, err := strconv.Atoi(reMatches[1])
		if err != nil {
			response.WriteHeader(http.StatusBadRequest)
		}
		// should be get()
		switch request.Method {
		case http.MethodGet:
			animalCtrl.get(id, response)
		default:
			response.WriteHeader(http.StatusNotImplemented)
		}

	}
}

// animalController.getAll() gets a list of all animals and sends it to the caller as JSON.
func (animalCtrl animalController) getAll(response http.ResponseWriter, request *http.Request) {
	sendResponseJSON(models.GetAnimals(), response)
}

// animalController.get() gets an animal by id (from the request URL) and sends it to the caller as JSON.
// if the id isn't found, it returns an internal server error (HTTP status 500)
func (animalCtrl animalController) get(id int, response http.ResponseWriter) {
	// get the animal by id
	animal, err := models.GetAnimalById(id)
	if err != nil {
		// respond with an internal server error
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	sendResponseJSON(animal, response)
}
