package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterControllers() {
	animalCtrl := newAnimalController()

	http.Handle("/animals", *animalCtrl)
	http.Handle("/animals/", *animalCtrl)
}

func sendResponseJSON(data interface{}, writer io.Writer) {
	encoder := json.NewEncoder(writer)
	encoder.Encode(data)
}
