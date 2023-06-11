# Animal service

Let's build a simple service that returns animals.

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
