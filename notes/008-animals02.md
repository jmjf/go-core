# Animal service (v02)

## Problem

The model doesn't do much for us other than define a type and some variables we don't use. What should our model do?

## Get and add animals

We've defined `animals []*Animal` (was `animal` in 006, but that was a typo, so fixing). We want `animals` to hold a set of `Animal` instances but, because the name is lower case, it isn't exported (private to the package). We need a way to return it.

Add two functions to `models/animal.go`:

* `GetAnimals() []*Animal`
* `AddAnimal(animal Animal) (Animal, error)`

**COMMIT: FEAT: add methods to get all animals and add an animal to the list**

## Add a controller (adding methods to an object)

Create a `controllers` directory and `animal.go` inside it.

Add an empty `struct` for `animalController`. We'll add data to it later.

Declare a `ServeHTTP` function bound to `animalController`. It accepts an `http.ResponseWriter` and `*http.Request`. For now, we're stubbing it.

`func (animalCtrl animalController) ServeHTTP (response http.ResponseWriter, request *http.Request)`

## Soapbox: Naming matters

Many folks in Go-land like short (1-character) variable names. For example, `w http.ResponseWriter`, `r *http.Request`. Why have I used names `animalCtrl`, `response`, and `request`?

The challenge is, if you have two different writers (response and file, two files, etc.), what do you call them? Or two readers? Or request and a reader? In one function `r` is a request, in another it's a reader, in a third we have `r1`, `r2` or `httpRequest`, `fileReader`, etc.

I have a ton of respect for Thompson and Pike, but if this extreme focus on brevity comes from them, I respectfully say they got it wrong.

> Code is read much more often that it is written, so plan accordingly.
> [Raymond Chen](https://devblogs.microsoft.com/oldnewthing/20070406-00/?p=27343)

> ...the ratio of time spent reading versus writing is well over 10 to 1. We are constantly reading old code as part of the effort to write new code.
> [Robert C. Martin](https://www.goodreads.com/quotes/835238-indeed-the-ratio-of-time-spent-reading-versus-writing-is)

Add to that list Guido von Rossum of Python fame and a host of others. Code is write once, read many, change occasionally -- and changing requires reading it to understand it.

Clear names reduce cognitive load when reading code. The more mental energy and short term memory space we spend tracking amd translating the meaning of short variable names, especially when those names change meanings over a series of functions, the less mental energy we have for understanding the code and the more likely we are to make mistakes.

We can abbreviate when it makes sense, especially if we have conventions within a code base. For example, `animalCtrl` instead of `animalController` is a reasonable variable name. So is `Id` instead of `Identifier` in most contexts. Depending on the domain in question, other abbreviations probably make sense. Documenting them so people know what they are is wise.

Beware of abbreviation terms that mean different things in different domains. For example, in the financial industry, "repo" means "repossession" in a loans context, but "repurchase agreement" in a financial markets/trading context. When those two contexts cross or come together in a system, data store or report, shorthand can be confusing or misleading to the tune of tens of millions of dollars.

**COMMIT: FEAT: add the ServeHTTP method for the animal controller**

## Add data to the controller struct

The controller will handle resource requests for the whole collection of animals or a specific set of animals that meets query criteria. We'll use a regular expression to do that.

Add a regular expression member to the controller `struct` to hold the regular expression for the id URL path.

We need a constructor function for the controller (`newAnimalController()`). Lower case first letter means it isn't exported from the `controllers` package. The constructor returns `&animalController{<controller definition>}` and initializes the regular expression. The controller is created within the function, which creates a closure. The data isn't lost, it's promoted to higher scope, but every call creates a new one, so they don't overlap.

[Golang Regexp syntax](https://github.com/google/re2/wiki/Syntax)

**COMMIT: FEAT: add a constructor for the controller**

## Interfaces

We can look at [Golang http.Handler type docs](https://pkg.go.dev/net/http#Handler) to see an example of how golang uses interfaces. Notably, the interface spec has a method, `ServeHTTP(ResponseWriter, *Request)` -- which is the same signature we used for the controller's `ServeHTTP()`. So, our controller is a `Handler` and can receive and respond to a request. (Yes, that was intentional.)

We need to create the controller (call the constructor) and register it as a handler. So, we'll add `registerControllers.go` to the `controllers` package and have it do that. (Based on not needing outside the packages, I renamed `NewAnimalController()` to `newAnimalController()`.)

We register a controller with `http.Handle()` ([docs](https://pkg.go.dev/net/http#Handle)), which takes a URL path (`pattern`) and the handler. `http.Handle()` registers with the `DefaultServeMux` (HTTP request multiplexer), which is like a request router in Node and some other languages HTTP server packages. In golang, it's of type `ServeMux` ([docs](https://pkg.go.dev/net/http#ServeMux)).

According to the docs, `ServeMux`, given `"/images/"` as a pattern will match both `/images/` and `/images` (will redirect the latter to the former) unless separate paths are registered. But another reference I'm using says to register both paths. For now, I've commented out the `/`-less path.

**COMMIT: FEAT: add registerControllers()**

## Start the server

Now it's time to pull it all together and start the server. Our controller returns a fixed string, but we can ensure everything is wired up right and go from there.

In `main.go`, replace the code (`fmt.Println()` a fixed object). We need to register controllers from the controllers package and call `http.ListenAndServe()` ([docs](https://pkg.go.dev/net/http#ListenAndServe)). We can pass the port as the address to use localhost (`":9200"` in this case) and use the `DefaultServeMux` by passing `nil` as the handler.

Now we can `go run main.go` and browse to `localhost:9200/animals`. We'll get the "Called the animal controller." message back. Note that the URL changes to `localhost:9200/animals/` with a trailing `/`. And `localhost:9200/animals/987` also works, so we don't need to register both routes (the docs were right).

All the wiring is done. We can work on making the animal controller return animals next.

**COMMIT: FEAT: add code to start the server**
