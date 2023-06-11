package controllers

import "net/http"

type animalController struct{}

func (animalCtrl animalController) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Called the animal controller."))
}
