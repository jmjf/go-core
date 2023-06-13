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

		switch request.Method {
		case http.MethodGet:
			animalCtrl.getAll(response, request)
		case http.MethodPost:
			animalCtrl.post(response, request)
		default:
			response.WriteHeader(http.StatusNotImplemented)
		}

	} else {
		// we have an id
		reMatches := animalCtrl.animalIdRegexp.FindStringSubmatch(request.URL.Path)
		if len(reMatches) == 0 {
			response.WriteHeader(http.StatusBadRequest) // malformed request because required id is missing
		}

		id, err := strconv.Atoi(reMatches[1])
		if err != nil {
			response.WriteHeader(http.StatusBadRequest) // malformed request because id is not an int (invalid)
		}

		switch request.Method {
		case http.MethodGet:
			animalCtrl.get(id, response)
		case http.MethodPut:
			animalCtrl.put(id, response, request)
		case http.MethodDelete:
			animalCtrl.delete(id, response)
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
// If the id isn't found, it returns an internal server error (HTTP status 500).
func (animalCtrl animalController) get(id int, response http.ResponseWriter) {
	animal, err := models.GetAnimalById(id)
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		return
	}
	sendResponseJSON(animal, response)
}

// animalController.post() adds an animal using the data in the request body and returns the animal
// If the request body can't be parsed, it returns a bad request error (HTTP status 400).
// If the add fails, it returns an internal server error.
func (animalCtrl animalController) post(response http.ResponseWriter, request *http.Request) {
	animal, err := animalCtrl.parseRequest(request)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest) // cannot parse body, so it's malformed, missing or invalid
		return
	}

	animal, err = models.AddAnimal(animal)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError) // add failed
		return
	}

	sendResponseJSON(animal, response)
}

// animalController.put() updates an animal using the data in the request body and returns the animal.
// If the request body can't be parsed, it returns a bad request error (HTTP status 400).
// If the request body id doesn't match the id in the URL, it returns a bad request error.
// If the update fails, it returns an internal server error.
func (animalCtrl animalController) put(id int, response http.ResponseWriter, request *http.Request) {
	animal, err := animalCtrl.parseRequest(request)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest) // cannot parse body, so it's malformed, missing or invalid
		return
	}

	if animal.Id != id {
		response.WriteHeader(http.StatusBadRequest) // id on URL doesn't match id on body, so it's malformed
		return
	}

	_, err = models.UpdateAnimal(animal)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError) // update failed
		return
	}

	sendResponseJSON(animal, response)
}

// animalController.delete() adds an animal using the data in the request body.
// if the
func (animalCtrl animalController) delete(id int, response http.ResponseWriter) {
	err := models.DeleteAnimalById(id)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError) // delete failed
		return
	}

	response.WriteHeader(http.StatusNoContent) // 204 -> not supplying any information
}
