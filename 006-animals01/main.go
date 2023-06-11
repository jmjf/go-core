package main

// VS Code/gopls complains about modules not in workspace, but I'm ignoring it

import "fmt"

type animal struct {
	Id         int
	Family     string
	Genus      string
	Species    string
	CommonName string
}

func main() {
	animals := []animal{
		{
			Id:         1,
			Family:     "Canidae",
			Genus:      "Canis",
			Species:    "lupus",
			CommonName: "Gray wolf",
		},
		{
			Id:         2,
			Family:     "Canidae",
			Genus:      "Vulpes",
			Species:    "vulpes",
			CommonName: "Red fox",
		},
		{
			Id:         3,
			Family:     "Canidae",
			Genus:      "Canis",
			Species:    "familiaris",
			CommonName: "Dog",
		},
		{
			Id:         4,
			Family:     "Canidae",
			Genus:      "Canis",
			Species:    "latrans",
			CommonName: "Coyote",
		},
	}

	for _, a := range animals {
		fmt.Println(a)
	}
}
