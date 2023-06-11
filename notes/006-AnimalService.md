# Animal service

Let's build a simple service that returns animals.

For version 1, we're just going to define the type and show we can use it.

## Setup

Set up the directory and `go.mod` file.

```bash
mkdir 006-animals01
cd 006-animals01
go mod init .
touch main.go
```

Define a struct for the animal object

```golang
   type animal struct {
      Id int
      Family string
      Genus string
      Species string
      CommonName string
   }
```

Put animals in a slice

```golang
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
```

**COMMIT: FEAT: define animals**

## Move struct definition to models

Create a `models` directory and package to hold the struct.

Move `animal` to `models/animal.go`, rename to `Animal` so it will export, and import it in `main.go`. Also define `animals []*Animal` and `nextId = 1`.

`animal.go`

```golang
package models

type Animal struct {
   Id         int
   Family     string
   Genus      string
   Species    string
   CommonName string
}

var (
   animal []*Animal
   nextId = 1
)
```

In `main.go`, comment out the original code, import `animals01/models`, and define and print an animal.

```golang
package main

// VS Code/gopls complains about modules not in workspace, but I'm ignoring it

import (
   "fmt"

   "animals01/models" // the module is animals01, must be prefixed to avoid issues
   // VS Code (gopls) may complain about the line above, but it works
)

func main() {
   animal := models.Animal{
      Id:         2,
      Family:     "Canidae",
      Genus:      "Vulpes",
      Species:    "vulpes",
      CommonName: "Red fox",
   }
   fmt.Println(animal)

   // commented code omitted
}
```

**COMMIT: REFACTOR: move models into a separate package so they're more usable**
