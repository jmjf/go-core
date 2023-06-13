package main

// VS Code/gopls complains about modules not in workspace, but I'm ignoring it

import (
	"net/http"

	"animals03/controllers"
	"animals03/models"
	// VS Code (gopls) may complain about the lines above, but it works
)

func main() {
	models.InitalizeAnimals()
	controllers.RegisterControllers()
	http.ListenAndServe(":9200", nil)
}
