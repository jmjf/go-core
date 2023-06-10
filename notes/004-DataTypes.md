# Data types

## Note

This document covers basic types and syntax. Use the [Go Playground](https://go.dev/play/) to try out the examples and play with variations to see how the language behaves.

## Declaration

We can declare variables several ways depending on what we need to do.

Separate declaration and initialization

```golang
   var x int
   x = 123
   fmt.Println(x)
```

Initialize on declaration

```golang
   var tau float32 = 6.283185
   fmt.Println(tau)
```

Declare and initialize with inferred type

```golang
   y := 73
   fmt.Println(y)
```

The third option is most common, but the others exist for cases where we need to control the data type or declare and initialize later.

## Value types

Go data types include:

* signed integers (int, int8, int16, int32, int64)
* unsigned integers (uint, uint8, uint16, uint32, uint64, uintptr)
* byte (uint8 alias)
* float (float32, float64)
* complex (complex64, complex128) -- imaginary numbers such as 2 + 4i
* bool -- true/false
* string -- sequence of bytes
* rune -- Unicode characters (int32 alias)

Go also supports arrays, slices (of arrays), pointers, structs, functions (type based on parameters and return values), interfaces (set of functions that may be attached to structs), maps (key/value lookups), and channels (communication between concurrent goroutines).

## Pointer types

Pointers are a reference to the variable.

```golang
   var cityName *string = new(string)  // must create a string to point to
   *cityName = "Tokyo"                 // change the pointer to point to the string (literal)
   fmt.Println(*cityName)              // * -> dereference -> gets the value in the pointer
```

Go does not allow pointer math because it's easy to get it wrong.

```golang
   countryName := "Japan"
   ptr := &countryName              // & -> address of countryName
   fmt.Println(ptr, *ptr)           // ptr -> address; *ptr -> value pointed to
```

Try variations of printing and assignment to see errors so you're familiar with them if you run into them.

* `var cityName *string
* `var cityName *string = "Tokyo"`
* `cityName = "Tokyo"`
* `cityName = &"Tokyo"`
* `ptr := countryName`
* `ptr := *countryName`
* other variations

## Constants

Use `const` to define a constant.

* Go doesn't allow changes to constants.
* Constants must be initialized when declared (because they can't change).
* Values used to initialize a constant must be defined a compile time. Functions are evaluated at run time, so can't be used.

```golang
   const tau = 6.283185
   fmt.Println(tau)
   // tau = 5.1 -> error: cannot assign to tau
   // const tau -> error: missing init expression
```

```golang
   func t() float32 {
      return 6.238185
      // declaring a constant and returning it also fails below
   }

   func main() {
      const tau = t()  // error: t() is not constant
      fmt.Println(tau)
   }
```

```golang
   const v = 17            // no type specified
   fmt.Println(v + 10)     // 27
   // code code code
   fmt.Println(v + 1.5)    // 18.5
   // in the examples above, Go dynamically interprets the type of v because it isn't set
   const r int = 10
   fmt.Println(r + 1.5)    // error: r truncated to int -- because r has a defined type
   fmt.Println(float32(r) + 1.5) // explicit cast works
```

## Constant blocks

```golang
   const tau = 6.238185       // tau is available to the whole package

   func main() {
   fmt.Println(tau)
   }
```

Just as `import` can import many packages using a list in parentheses, `const` can define many constants using a list in parentheses.

```golang
   const (                    // all four constants are available to the whole package
      tau = 6.238185
      yearNumber = 1603
      cityName = "Tokyo"
      aRune = '世'
   )

   func main() {
      // aRune printed as 19990 because it's an int32 under the covers
      fmt.Println(tau, yearNumber, cityName, aRune)
      // aRune formatted as a character, so it prints 世
      fmt.Printf("%v %v %s %c", tau, yearNumber, cityName, aRune)
   }
```

## iota

WARNING: `iota` can be abused (less readable code) or create unexpected problems.

The constant expression `iota` increments by one each time it's used.

```golang
   const (
      c1 = iota         // c1 = 0
      c2 = iota         // c2= 1
      c3 = iota + 6     // c3 = 8 (2 + 6)
      c4 = 2 << iota    // c4 = 16 (2 << 3 -> 0 0010 shift left 3 bits to get 1 0000 )
      c5 = iota << 1    // c5 = 8 (4 << 1 -> 0100 shift left 1 bits to get 1000)
   )

   func main() {
      fmt.Println(c1, c2, c3, c4, c5)
      // fmt.Println(iota) -> error: can't use iota outside a constant declaration
   }
```

Go reuses constant expressions. This option feels unintuitive and seems like it could lead to misunderstanding.

```golang
   const (
      c1 = iota         // c1 = 0
      c2                // c2 = 1 -> reuses iota
      c3                // c3 = 2 -> reuses iota
      c4 = 5 + 5        // c4 = 10
      c5                // c5 = 10 -> reuses 5 + 5
   )

   func main() {
      fmt.Println(c1, c2, c3, c4, c5)
   }
```

The value of `iota` resets for each constant block.

```golang
   const (
      c1 = iota         // c1 = 0
      c2                // c2 = 1
   )

   const (
      c3 = iota         // c3 = 0 -> iota reset in a new constant block
      c4                // c4 = 1
   )

   func main() {
      fmt.Println(c1, c2, c3, c4)
   }
```

Use `iota` carefully.

It's useful for incrementing constant values, but order matters, so be sure to define names consistently. I suspect this would be best done as a package that's used across everything in an application or application suite.

```golang
   const (
      red = iota
      green
      blue
   )

   // gets different values for each constant than

   const (
      blue = iota
      green
      red
   )
```

It's also useful for defining bit masks with the same caveats.

```golang
   const (
      read = 1 << iota  // 1
      write             // 2
      remove            // 4
   )
```

Beware about changes to `iota` based constants that may affect stored data. With great power comes great risk.

For this reason, when shared (exported) or storable constants are set, NEVER change them. (All developers have a development platform. Some are lucky enough to have a separate production platform.)

```golang
   // assume we start with this code
   const (
      red   int = 1
      green int = 2
      blue  int = 3
   )

   // assume we use the values above to write data to a database, 
   // send them in an HTTP data body, publish them on a message bus to other program
   // or otherwise expose them outside the running code
   // (this can include constants in shared modules/packages)

   // now we decide to use iota because it's simpler
   const (
      red   int = iota
      green
      blue
   )

   // now red = 0, green = 1, blue = 2, so any data outside the running code is out of sync
   // if we're using this block of constants in a common reference file, which makes sense from a code management perspective
   // we may pass tests, etc., and not find the issue until we use real saved data
   // we can prevent the issue by setting red = iota + 1, but need to remember to do that
```

For general safety, wisdom argues against `iota` and in favor of explicit constant values for exported constants or constants that may be saved or transmitted outside the running code.

If a set of constants is used inside the running code only, then `iota` can make the code easier to write (though maybe not easier to read).
