# Looping and branching

## Loops

All loops in golang are `for` loops. We can accomplish any type of loop with `for`.

## Condition-based (while loop)

```golang
   var x int
   for x < 10 {
      // loops while x >= 10
      // while x < 10 {}
      println(x)  // simpler println method that's good enough
      x++  // increment x -- if we don't change x, we have an infinite loop
      
      // we can stop a loop early with break, for example, when x is 7
      if x == 7 {
         break
      }

      // we can stop the current iteration of the loop with continue, for example, when x is 5
      if x == 5 {
         continue
      }

      println("end of iteration") // will not print for x == 5
   }
```

## Post clauses (traditional for loop)

```golang
   var x int
   // adding ; x++ increments x after each loop iteration ends
   for ; x < 10; x++ {
      println(x)

      // other stuff

      println("end of iteration")
   }
```

Or, we can merge the declaration.

```golang
   // adding x := 0; declares and initializes x; we could initialize to any value
   // adding ; x++ increments x after each loop iteration ends; the condition can be more complex if needed
   for x := 0; x < 10; x++ {
      println(x)

      // other stuff

      println("end of iteration")
   }
   // because x is defined in the scope of the for block, x is not declared here
   // println(x) will fail outside the loop
   // if you need x to exist outside the loop, declare it outside the loop (first example in this section)
```

## Infinite loops (forever loops)

```golang
   var x int
   // remove the conditions, but keep the semicolons
   for ; ; {
      // reverse the condition from < 10 to >= 10 and break on the reversed condition
      // the condition could be more complex
      if x >= 10 {
         break
      }
      println(x)
      // increment the loop variable
      // the "increment" could be more complex and may not be a true increment
      // may call some function and loop for specific results, etc.
      x++
   }
```

But, we really don't need the semicolons.

```golang
   var x int
   // remove the semicolons
   for {
      if x >= 10 {
         break
      }
      println(x)
      x++
   }
```

## Looping over collections

```golang
   canids := []string{"wolf", "fox", "coyote", "dog"}

   // we can loop over the members of canids with a counter
   for i := 0 ; i < len(slice); i++ {
      println(slice[i])
   }

   // but there's a simpler way
   // we can loop over the slice itself; range returns two values, the key/index and the value
   for i, value := range slice {
      println(i, value)
   }

   // the same thing works for maps
   statuses := map[string]int{"OK": 0, "ERROR": 1, "FATAL": 2}
   for key, value := range status {
      println(key, value)
   }

   // if you only need keys, syntax below ignores value
   for key := range status {
      println(key)
   }

   // if you only need values, use _
   for _, value := range status {
      println(value)
   }
```

## Panic

Panic is like a thrown exception.

```golang
   println("Starting")

   // something unrecoverable happens -- will print the message and some information about what failed where
   // not quite a stack trace, but conceptually similar
   panic("Something unrecoverable happened")
   // applications can recover from panics (like catching a thrown error and handling it), but outside immediate scope

   println("Started")
```

## If

Tests a logical condition and chooses behavior based on truth or falsehood of that condition.

```golang
   id1, id2 := 1, 2

   if id1 == id2 {
      println("equal")
   }
   
   if id1 != id2 {
      println("not equal")
   }
   // we can also use other and more complex conditions

   // we can also use else
   if id1 == id2 {
      println("equal")
   } else {
      println("not equal")
   }

   // and we can use (and chain) else if
      if id1 == id2 {
      println("equal")
   } else if id1 < id2 {
      println("id1 less than id2")
   } else { // else is optional
      println("id1 greater than id2")
   }
```

## Switch

Chooses from a set of options, similar to `if ... else if`, but not exactly. This section discusses the basics of `switch`. It has some [other syntaxes](https://gobyexample.com/switch) that give it other uses.

```golang
   // assume status is set by some process and could be any of several values
   status := 1

   switch status {
      // switch chooses the first case that matches and only that case
      // no break keywords needed, which means no fallthrough (good because fallthrough is confusing)
      case 0:
         println("ok")
      case 1:
         println("warning")
      case 2:
         println("error")
      default:
         // default executes if none of the cases match
         println("unknown status")
   }
```

**COMMIT: DOCS: add notes on looping and branching**
