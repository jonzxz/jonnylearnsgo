# Control Statements, Declaration, Types

## Control Structures
- Whereever braces are allowed, bracres are required.

### if-else
 ```golang
  if a == b {
    // do this
  } else {
    // do this
  }
  ```

- Conditionals can start with a short declaration
```golang
// run only if a is true
if a := 2*4; a <= 8 {
  return a
}
```

### For loops
-  There is only `for`, no `do` or `while`
#### Explicit `for`
```golang
for i := 0; i < 10; i++ {
  // do something from 1 to 10
}
```
#### Array/slice iteration
  - `for` retrieving index
```golang
for i := range myArr {
  // do myArr[i]
}

```
- `for` retrieving index, value
  - In this case, `v` gets  **copied**, 
```golang
for i, v := range myArr {
  // index, value
  // v gets copied into someTempVar
  someTempVar := v
}
```

#### Map iteration
- Maps in Golang have **no order** which is based on a hashtable 
- If we want to get something in order, eg. alphabetical, extract the keys, sort them then get the values based on the key.
```golang
for k := range myMap {
  // do myMap[k]
}
```

```golang
for k, v := range myMap {
  // do print k, v
}
```

#### Infinite loop
```golang
for {
  // do something
  // for will continue indefinitely until a break statement is found
  a := 0
  v := 10
  a++
  if a > v {
    break
  }
}
```

#### Iteration retrieving value without index
-  `_` is an untyped, reusable "variable" placeholder
- This is generally used when we want the value but not the index
```golang
for _, v := range myArray {
  // do something wih v
}
```

### Switch case
- Shortcut for replacing `if-then`
- Cases **break automatically**, so there's no fallthrough by default, unless explicitly defined
```golang
switch a := f.Get(); a {
  case 0, 1, 2:
    // return if a = 1/2/3
  case 3, 5, 6, 7, 8:
    // NOOP, this does not do anything
  default:
    // return default 
}
```

```golang
a := f.Get()

switch {
  case a <= 2:
    // do something
  case a <= 8:
    // do something
  default: 
    // do something
}
```

## Packages
- Every standalone program has a `main` package, nothing is "global"
- Everything is either "package" scope, or "function" scope
```golang
package main

import "fmt"
....
```

### Package-level declarations
- Everything can be declared at package scope
- **Short declarations do not work in package-level**, i.e. a := 2 does not work in package scope
```golang
package myPackage

const defaultID = "123"
var key string

type secretStruct struct {

}

func do() error {

}
```

### Package control visibility
- Every name that is **capitalized is exported**
- Same concept as "public", "private". Capitalized variables/functions are all "public" / exported
```golang
package secrets

type internal struct {

}

// this is exported
func GetAll() {

}
```

### Imports 
- All source files must import everything it uses 
- Unused imports will result in compilation error
- **Files of the same package lives in the same directory**
- Cyclic dependencies are not allowed
  - Move common dependencies to a third package, or eliminate them

### Initialization
- Items within a package gets initialized before `main`
- There's a private function called `init()` which gets called when your package gets started up
  - Think of it like a "package constructor"

# Declarations

## Variable Declarations
- Valid variable declarations
- If we use `var x`, a `type` or value **must be given**
```golang
var a int // 0 by default 
var b uint = 1
var c = 1 // implicit
var d = 1.0 // implicit

var (
  x, y  int
  z float64
  s string
)

var a // DOES NOT WORK, missing type OR value
```

### Short Declarations
- Cannot be used outside of a function (i.e. package)
- **Must** be used in-place of `var` in a control statement like  `if`
- Must declare at least **one new variable**
- It will not reuse an existing declaration from an outer-scope
```golang
err := doSomething()
err := doSomethingElse() // ERROR, err is already declared
x, err := getSomeValue() // this is ok, err is reassigned and not redeclared
```

### Shadowing Short Declarations
- To emphasize on the point above on not reusing existing declaration from an outer-scope
- The below code will result in a compilation error because the first `err` is unused. The `err` from `doSomething()` is different from the `err` in `doSomethingElse()` because the latter `err` is only local to the scope of that block (if) 
```golang
func main() {
  // this err here is NEVER used
  n, err := doSomething()
  // this err is not the same as the one above
  if _, err := doSomethingElse(); err != nil {

  }
}
```

# Named Typing
- Go uses named typing for non-function user-declared types
  - eg. custom types
```golang
// creates a type called "x" that has a int base
type x int

var a x // creates variable a of type x
b: = 12

a = b // type mismatch, because b is int and a is x

a = 12 // this is ok because type x is base int

// this is ok because we are typecasting b to x type and assigning to a
b: = 12
a = x(b)

```