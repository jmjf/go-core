package controllers

import "net/http"

func RegisterControllers() {
	animalCtrl := newAnimalController()

	// http.Handle("/animals", *animalCtrl)
	http.Handle("/animals/", *animalCtrl)
}
