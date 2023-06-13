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

**COMMIT: REFACTOR: use findAnimalById() to simplify CRUD functions**

## Support different CRUD actions from the controller -- GET methods

Currently, the controller returns a fixed response. Now we need to build it out so it can accept different methods and perform different actions. We'll write functions to handle each of the actions and have the controller choose the correct function based on the HTTP method and information in the request.

Actions to support and the functions we'll write:

* GET all -> `getAll()`
* GET by id -> `get()`
* POST (add new) -> `post()`
* PUT (update) -> `put()`
* DELETE by id -> `delete()`

We'll also need a function to JSON-ify the response data (`encodeAsJSON()`).

The controller will pass the `response` (`http.ResponseWriter`) to the action handler. The action handler will be responsible for sending the response. The controller can also pass the `request` (`*http.Request`) or specific data from it to the action handler.

The `getAll()` function is simple enough. It sends all the animals in the response. No real magic here, it just needs to call the model's `GetAnimals()`.

The `get()` function calls `GetAnimalById()` and needs to handle a possible error.

**NOTE:** I renamed `RegisterControllers.go` to `controllers.go` because it's the base of the controllers module and holds common methods -- notably, `sendResponseJSON()`.

In the animal controller's `ServeHTTP()` method (in `controllers/animal.go`), we need to decide which function we should call. We do that by asking if the request URL has an id.

```golang
   if request.URL.Path == "/animals" || request.URL.Path == "/animals/"
   // we have no id
```

If the URL has no id, we use a `switch` on the HTTP method to decide what to do. If it's a `GET`, call `getAll()`. For any other method, we don't know what it is, so return an error. (We'll add `post()` to this `switch` later.)

If the URL has an id, use the regular expression to get the id.

```golang
   reMatches := animalCtrl.animalIdRegexp.FindStringSubmatch(request.URL.Path)
```

`FindStringSubmatch()` returns a slice of strings that match the regular expression. The way the regexp is written, we can get only one match and it will be the second element of the slice. Convert it to an `int` with `strconv.Atoi()` and, if that doesn't return an error (too big numeric string), call `get()` with it.

If no matches, return an error because the client has added something that doesn't fit the id pattern after the base URL. That's a bad request.

Or if the `Atoi()` returns an error, return a bad request too (invalid id).

I could rush ahead and build the rest of the methods, but let's see these work first.

I added an `InitializeAnimals()` function to the model that loads four values into `animals` and sets `nextId` to 5. I call that in `main.go`.

Run the program and test browsing to `/animals`. It returns the array. `/animals/1` returns id 1. `/animals/3` returns id 3.

Testing with `curl` shows an issue. `/animals` returns "Moved Permanently" with a redirect. So, let's add `/animals` to the routes in `controllers.go` and try again. And now `curl http://localhost:9200/animals` returns the array. So, that's why we need both paths, to avoid the redirect. The browser follows it automatically, but `curl` doesn't.

The advanced controller is working. Next I'll add the other action handlers and include them in the `switch`es.

**COMMIT: FEAT: add action handlers and wire controller to respond to GET requests with and without an id**

## Support POST, PUT, DELETE methods

The `post()` function needs to convert the request body from JSON to an `Animal` (`parseRequest()`). If that succeeds, it calls `AddAnimal()` to add the animal. In the controller, we call `post()` from the "no id" branch of the `if`.

The `put()` function also calls `parseRequest()`. `UpdateAnimal()` takes the animal, not the id, but `put()` should confirm the ids match.

The `delete()` function uses the id only, so doesn't call `parseRequest()`. On success, we'll return a 204 (No Content) because we aren't describing status but succeeded.

**COMMIT: FEAT: add actions for POST, PUT and DELETE**
