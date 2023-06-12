package models

import (
	"fmt"
)

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

// GetAnimals gets all Animal instances.
func GetAnimals() []*Animal {
	return animals
}

// AddAnimal adds an Animal to the list of animals.
// If the animal has a non-zero Id, it returns an error.
func AddAnimal(animal Animal) (Animal, error) {
	if animal.Id != 0 {
		return Animal{}, fmt.Errorf("new animals must have id = 0; animal has id '%v'", animal.Id)
	}

	animal.Id = nextId
	nextId++
	animals = append(animals, &animal)

	return animal, nil
}

// GetAnimalById finds an Animal in the list by id and returns it.
// If the id isn't found, it returns an error.
func GetAnimalById(id int) (Animal, error) {
	i, err := findAnimalById(id)
	if err != nil {
		return Animal{}, err
	}

	return *animals[i], nil
}

// DeleteAnimalById finds an Animal in the list by id and removes it from the list.
// If the id isn't found, it returns an error.
func DeleteAnimalById(id int) error {
	i, err := findAnimalById(id)
	if err != nil {
		return err
	}
	// build a new slice
	// take all elements in animals up to, but not including i (i matches the id)
	// and add all elements in animals after i (spreading them so append() will be happy)
	animals = append(animals[:i], animals[i+1:]...)
	return nil

}

// UpdateAnimal finds an Animal in the list with the id of the passed Animal and replaces it with the passed Animal.
// If the id is found, it returns the Animal, otherwise it returns an error
func UpdateAnimal(animal Animal) (Animal, error) {
	i, err := findAnimalById(animal.Id)
	if err != nil {
		return Animal{}, nil
	}

	animals[i] = &animal
	return animal, nil

}

// findAnimalById finds an Animal in the list by id and returns it's index.
// If the animal isn't found, it returns -1.
func findAnimalById(id int) (int, error) {
	for i, animal := range animals {
		if animal.Id == id {
			return i, nil
		}
	}
	return -1, fmt.Errorf("animal with id '%v' not found", id)
}
