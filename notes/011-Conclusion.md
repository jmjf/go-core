## Conclusion

We've covered the basic constructs of programming in golang including

* Data types
  * Values
  * Pointers
  * Constants
  * `iota`
* Collective types
  * Arrays and slices
  * Maps
  * `struct`s
* Functions
  * How to declare them
  * How to pass parameters
  * Returning values and errors
* Looping and branching
  * Using `for` to write a `while` loop
  * Using `for` to write a traditional `for` loop
  * Using `for` to loop forever (infinite loop)
  * Using `for ... range` to loop over a collection (array, slice, map)
  * Using `panic()` to throw an error (will be fatal if not recovered)
  * `if ... else if ... else`
  * `switch ... case ... default`

And we've written a web service that manages animals. It isn't production grade, but it exercises all the language features we used and introduces some common library packages like `fmt`, `net/http`, and `encoding/json` and touches on a few others like `strconv` and `regexp`.

I'm wrapping up this learning repo for now. I'll start a learning project next and get into testing, connecting to a database, maybe concurrency, etc.

**COMMIT: DOCS: that's a wrap**
