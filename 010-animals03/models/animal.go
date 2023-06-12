package models

import (
	"errors"
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
		// return an error
		// the signature requires an Animal too, so return an empty Animal
		return Animal{}, errors.New("new Animals must have Id = 0")
	}

	animal.Id = nextId
	// increment nextId so it's ready for the next animal
	nextId++
	// add the animal to animals
	animals = append(animals, &animal)
	// we don't have any error handling yet, so just return as if okay
	return animal, nil
}

// GetAnimalById finds an Animal in the list by id and returns it.
// If the id isn't found, it returns an error.
func GetAnimalById(id int) (Animal, error) {
	// loop over users and find the one with the id
	for _, animal := range animals {
		if animal.Id == id {
			return *animal, nil
		}
	}
	// not found case
	// fmt.Errorf is like errors.New(fmt.Sprintf())
	return Animal{}, fmt.Errorf("animal with id '%v' not found", id)
}

// DeleteAnimalById finds an Animal in the list by id and removes it from the list.
// If the id isn't found, it returns an error.
func DeleteAnimalById(id int) error {
	for i, animal := range animals {
		if animal.Id == id {
			// this says, build a new slice
			// take all elements in animals up to, but not including i (i matches the id)
			// and add all elements in animals after i (spreading them so append() will be happy)
			animals = append(animals[:i], animals[i+1:]...)
			return nil // we're done
		}
	}
	return fmt.Errorf("animal with id '%v' not found", id)
}

// UpdateAnimal finds an Animal in the list with the id of the passed Animal and replaces it with the passed Animal.
// If the id is found, it returns the Animal, otherwise it returns an error
func UpdateAnimal(animal Animal) (Animal, error) {
	for i, toCheck := range animals {
		if toCheck.Id == animal.Id {
			animals[i] = &animal
			return animal, nil
		}
	}
	return Animal{}, fmt.Errorf("animal with id '%v' not found", animal.Id)
}
