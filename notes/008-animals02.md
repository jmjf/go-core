# Animal service (v02)

## Problem

The model doesn't do much for us other than define a type and some variables we don't use. What should our model do?

## Get and add animals

We've defined `animals []*Animal` (was `animal` in 006, but that was a typo, so fixing). We want `animals` to hold a set of `Animal` instances but, because the name is lower case, it isn't exported (private to the package). We need a way to return it.

Add two functions to `models/animal.go`:

* `GetAnimals() []*Animal`
* `AddAnimal(animal Animal) (Animal, error)`

**COMMIT: FEAT: add methods to get all animals and add an animal to the list**
