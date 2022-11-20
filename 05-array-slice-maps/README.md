# Array, Slice, Maps
## Composite types in Golang
### Array
  - fixed length
  - arrays are typed by size, **fixed at compiled time**
  - Some ways to declare arrays
  ```golang
    var a [4]int  // zero-init array a of fixed size 4
    var b = [4]int {0, 0, 0, 0}
    b :=  4[int] {0,0,0,0}
    var c [...]{0, 0 ,0} // the ... is interpreted by the number of elements in the {} block 
  ```
  - if arrays are assigned to another array variable, the values are **copied**
  ```golang
  var a = [3]int {0, 1, 2}
  var b [3]int
  b = a // Values / bytes of a are copied into b, not referenced
  ```
  - but you **cannot** reassign arrays of different sizes
  ```golang
  var c = [3]int {0, 1, 2}
  var d [4]int
  d = c // THIS DOES NOT WORK due to size mismatch
  ```
### Slice
  - aka arraylist, dynamic sizing
  - Slices have descriptors like strings and it will point to another memory
  - Slices will always have an array behind it
  - Slices have a descriptors containing the `length` and `capacity`
  - Some ways to declare slices
  ```golang
    var a []int // zero init, no storage (0 elements)
    var b = []int{1, 2}
    c := []int{1,2}
    // this makes a slice of int of size 5, zero initialised
    d := make([]int, 5)
  ```  
  - Appending to slices
    - One point to note here is that if the initial slice is only allocated eg. enough memory for 4 elements, and a 5th element is appended, the whole chunk of memory containing the 4 elements are copied to a new place and with the 5th element added, and its assigned back to the original slice by modifying the slice descriptor (similar to the string example)
    ```golang
      a := []int
      a = append(a, 1)
    ```
  - Slice assignments
    - In the below example, the descriptor of a is going to be modified to point to the descriptor of b, so it's a reference and not a copy
    - But if you mutate b via an append, a only holds the memory of the "old" b so the changes are not applied. Likewise if it's not used it'll be GC-ed
    - In the below assignment, all copied elements of b into a will result in a true when compared. 
      - eg. b = a, b[0] == a[0]
    ```golang
    a := []int
    b := []int {1, 2, 3, 4}

    a = b
    ```
  - Subslices
    - Slices work the same way as in Python, eg a[2:4], etc etc
    - Slices are a[`[inclusive]`:`[exclusive]`]
    - From 0th to n-1th is a[:n]
    - From nth to end is a[n:]

### Arrays vs. Slices
- Most Go APIs take slices as inputs and not arrays

| Slice | Array |
| ----- | ----- |
| Variable Length | Fixed length at **compile time** |
| Passed by reference | Passed by value (copied) |
| Not comparable | comparable (==) |
| Cannot be used as map key |  Can be used as map key |
| Has copy & append helpers | N.A. |
| Useful as function params | Useful as pseudo constants (you can't actually const an array) |

### Maps
- dictionaries / hashtable equivalent
- indexed by key and returns a value
- we can read from a nil map, but inserting/write will `panic`
  - when we read a key where it doesn't exist in a map, we get a default value of "zero" of the relevant value type
- Maps are **passed by reference**, there's no copying
  - which also means the same behavior as slices, if you modify a map through a function it modifies the original map
  - likewise if you have a map `a` and `b = a`, modifying `b` will change `a`
    ```golang
	    a := map[string]int{
		    "keyOne": 1,
	    }
	    b := make(map[string]int)

	    b = a
	    b["keyOne"]++
	    fmt.Printf("a: %v", a) // a's keyOne now becomes 2 because of modification in b
    ``` 
  - so what happens if we want to "copy" a map?
    - we need to iterate the original map and put each value accordingly
- We can update maps
- Types used for key **must** have `==` and `!=` comparators defined (aka not slice, map or funcs)
  - The only way to compare map is whether a map `== nil`
- Some ways to declare maps
  ```golang
  // this declaration can be read, but cannot be written because there's no map created
  
  var m map[string]int // nil, no memory allocated?
  
  // when we use "make", it actually creates the map in the background so there's a place to store the values. which means you can insert in a map from "make" but not the var above 
  m := make(map[string]int) // non-nil, but empty
  m := map[string]int {
    "Key1": 1,
    "Key2": 2,
  }
  ```
- Retrieving from maps
  ```golang
  m := map[string]int {
    "Key1": 1,
    "Key2": 2,
  }
  a := m["Key3"] // returns 0 because nothing is found
  m["Key3"] = 3 // PANIC because nil map?
  ```
### Special two-result lookup in maps
- There's an idea of looking up a map and returning 2 values, where
  - the first is the value of key, and second is a `boolean` if value is found
- Example
  ```golang
    // "ok" is actually a go convention to check for truthy
    a := make(map[string]int)
    val, ok := a["abc"] // this returns 0, false, because "abc" key is not in the map
    a["abc"]++
    valNow, ok := a["abc"] // this returns 1, true 
  
  // declares two variables w, ok
  // by indexing a with key "the"
  // so this loop only goes in if ok is true, otherwise the default value is returned and this skips the conditional
  if w, ok := a["abc"]; ok {}
  ```

## Common Built-in functions for each type
| func | structure type | remarks |
| ---- | -------------- | ------- |
| `len(s)` | string | length of string |
| `len(a)` | array | length of array |
| `cap(a)` | array | capacity of array (which is constant) |
| `make(T, x)` | slice | create slice of type T with length and capacity x, eg. make([]int, 4) |
| `make(T, x, y)` |  slice | create slice of type T with length x and capacity y. eg make([]int, 4, 4) |
| `copy` | slice | copy(c, d) | copies d into c |
| `c = append(c, d)` | slice | appends d after c and reassign it to c |
| `len(s)` | slice | length of slice |
| `cap(s)` | slice | capacity of slice (how much **more space** there is in the slice before the slice gets reallocated inanothe memory) |
| `make(T)` | map | create map of type T eg. make(map[string]int) |
| `make(T, x)` | map | create map of type T with space hint for x elements. eg make(map[string]int, 4) |
| `delete(m, k)` | map | delete key `k` from map `m` if found, else no change |
| `len(m)` | map | length of map `m` |

### Making nil useful
- Nil is a type of zero, indicating absence of something (like `null`)
- Many built-ins are safe like `len`, `cap`, `range`
```golang
s := make([]int)
m := make(map[string]int)

lengthS := len(s) // returns 0 since slice empty
lengthM := len(m)  // returns 0 since map is empty

i, ok := m["abc"] // returns 0, false because of missing key

// this does not iterate because s is empty
for _, v := range s {
  ...  
}
```