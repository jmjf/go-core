# Animal service (v03)

## Add features to the Animal model

If the provided animal has a non-0 id, return an error (and an empty `Animal` structure).

Add functions for:

* `GetAnimalById()` -- returns an animal that has the id. If the id isn't found, return an error.
* `DeleteAnimalById()` -- finds an animal by id and removes it from the list, returns an error if it isn't found.
* `UpdateAnimal()` -- accepts an animal, finds the matching animal by id and updates it, or an error if it isn't found.

In the future, it may be useful to write a function like this.

```golang
// findAnimalById finds an Animal in the list by id and returns it's index.
// If the animal isn't found, it returns -1.
func findAnimalById(id int) int {
 for i, animal := range animals {
  if animal.Id == id {
   return i
  }
 }
 return -1
}
```

Then the functions above can be like:

```golang
   if i := findAnimalById, i >= 0 {
      // action to do when found
      // return
   }
   return // error return
```

For example, `FindAnimalById()` becomes

```golang
   i := findAnimalById(id)
   if i >= 0 {
      return *animals[i], nil
   }
   return Animals{}, fmt.Errorf("animal with id '%v' not found", id)
```

**COMMIT: FEAT: add methods to model for basic CRUD operations**

##
