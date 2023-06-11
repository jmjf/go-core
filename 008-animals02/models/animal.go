package models

type Animal struct {
	Id         int
	Family     string
	Genus      string
	Species    string
	CommonName string
}

var (
	animals []*Animal
	nextId  = 1
)

// GetAnimals returns a slice of pointers to Animal instances
func GetAnimals() []*Animal {
	return animals
}

// AddAnimal adds an animal to the list of animals
// Returns either the added animal or an error
func AddAnimal(animal Animal) (Animal, error) {
	// assign an id to the animal (nextId)
	animal.Id = nextId
	// increment nextId so it's ready for the next animal
	nextId++
	// add the animal to animals
	animals = append(animals, &animal)
	// we don't have any error handling yet, so just return as if okay
	return animal, nil
}
