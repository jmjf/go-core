# Collections

## Arrays

A collection of similar, same-type members.

Explicit initialization

```golang
   var arr [3]int
   arr[0] = 1              // arrays use 0-based indexing
   arr[1] = 2
   arr[2] = 3
   fmt.Println(arr)
   fmt.Println(arr[1])     // 2 because 0-based
   // arr[3] = 4 -> error (only 3 elements)
   // fmt.Println(arr[3]) -> error for same reason
```

Implicit initialization

```golang
   arr2 := [3]int{1, 2, 3}    // same array as above
   fmt.Println(arr2)          // same output
```

## Slice

A slice is a segment of an array. Arrays are fixed size, but slices can vary in size.

Building a slice from an array

```golang
   arr := [3]int{1, 2, 3}

   slice := arr[:]         // takes a slice that covers the whole array

   fmt.Println(arr, slice) // same values in both

   // slices use the underlying array's storage, don't have their own
   arr[0] = 5
   slice[1] = 6
   fmt.Println(arr, slice) // both are still the same
```

Building a slice from an array literal

```golang
   slice2 := []int{1, 2, 3}
   fmt.Println(slice2)
```

The size of a slice can change

```golang
   slice3 := []int{1, 2, 3}
   fmt.Println(slice3)     // 3 elements -> 1, 2, 3

   slice3 = append(slice3, 5) // append wants a slice, not an array
   fmt.Println(slice3)     // 4 elements -> 1, 2, 3, 5

   // a slice backed by an array can grow, but doesn't change the array
   arr4 := [3]int{1, 2, 3}
   slice4 := arr4[:]
   slice4 = append(slice4, 5)
   fmt.Println(arr4, slice4)  // [1 2 3] [1 2 3 5]
   
   // appending to the slice detaches it from the array
   slice4[0] = 9
   fmt.Println(arr4, slice4)  // [1 2 3] [9 2 3 5]
```

You can slice slices

```golang
   slice6 := []int{1, 2, 3}
   slice7 := slice6[1:]        // slices starting at index 1 to the end -> [2 3]
   slice8 := slice6[:2]        // slices from 0 up to (excluding) 2 -> [1 2]
   slice9 := slice6[1:2]       // slices from 1 up to (excluding) 2 -> [2]
   fmt.Println(slice6, slice7, slice8, slice9)
```

## Maps

Maps are key/value collections. Values are all the same type

```golang
   m1 := map[string]int{}     // keys are strings, values are ints
   m2 := map[string]int{ "k1": 123 }
   fmt.Println(m1, m2, m2["k1"]) // indexing m2 returns the value

   // can add keys to the map and change values of a key
   m1["b"] = 999
   m2["k1"] = 456
   fmt.Println(m1["b"], m2["k1"])   // 999 456

   // can delete keys from a map
   delete(m2, "k1")
   fmt.Println(m2)   // m2 is empty
```

## Structs

A collection of values that may be different types.

Must declare the `struct` structure. It's fixed at compile time.

```golang
   type animal struct {
      Genus string
      Species string
      CommonName string
      AvgWeightGrams int
   }

   var fox animal
   fmt.Println(fox)     // {   0 } because we didn't set any values

   // we can assign values to members of a struct
   fox.Genus = "Vulpes"
   fox.Species = "vulpes"
   fox.CommonName = "Red fox"
   fox.AvgWeightGrams = 7500
   fmt.Println(fox)     // gets more interesting results
   fmt.Println(fox.CommonName)   // Red fox

   // Alternate syntax
   wolf := animal{
      Genus:          "Canis",
      Species:        "lupus",
      CommonName:     "Gray wolf",
      AvgWeightGrams: 45000,        // need a comma here because auto-inserted semicolons
   }
   fmt.Println(wolf)
```

We can declare `struct`s wherever it makes sense -- either within a function or in the package scope.
