package main

// VS Code/gopls complains about modules not in workspace, but I'm ignoring it

import (
	"net/http"

	// the module is animals02, must be prefixed to avoid issues
	"animals02/controllers"
	// VS Code (gopls) may complain about the line above, but it works
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":9200", nil)

	// animals := []models.Animal{
	// 	{
	// 		Id:         1,
	// 		Family:     "Canidae",
	// 		Genus:      "Canis",
	// 		Species:    "lupus",
	// 		CommonName: "Gray wolf",
	// 	},
	// 	{
	// 		Id:         2,
	// 		Family:     "Canidae",
	// 		Genus:      "Vulpes",
	// 		Species:    "vulpes",
	// 		CommonName: "Red fox",
	// 	},
	// 	{
	// 		Id:         3,
	// 		Family:     "Canidae",
	// 		Genus:      "Canis",
	// 		Species:    "familiaris",
	// 		CommonName: "Dog",
	// 	},
	// 	{
	// 		Id:         4,
	// 		Family:     "Canidae",
	// 		Genus:      "Canis",
	// 		Species:    "latrans",
	// 		CommonName: "Coyote",
	// 	},
	// }

	// for _, a := range animals {
	// 	fmt.Println(a)
	// }
}
