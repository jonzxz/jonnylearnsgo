# 03 - Basic Types

## Type systems
### Interpreted systems
- eg. Python
- assigning a = 2, a is actually an object and not a number
- Since python is in C, so it'll "translate" into the value of 2
- Goes through an interpreter before CPU+RAM

### Machine natives (Golang)
- a:= 2
- a  is actually the name / address of a memory location in the machine
- so it's directly into RAM + CPU
- This gives computational advantage (faster)

### Integers
- The basic type is `int`
- But there are also signed, unsigned `uint` ranging 8, 16, 32, 64 bytes

### Floating Point
- There's floating point types in Golang like `float32`, `float64`
  - By default it's `float64`
- Complex (imaginary floating points) like `complex64`, `complex128`
- In Golang, floating point is **not** used for monetary calculation.
  - Use a package like `Go money` instead
  - Floating points are actually for scientifics and not good for money in multiple languages anyway

## Variable declarations
- Declarative
  - `var a int`
    - Declares a variable `a` of `int`, initialized to 0
  - `var a int = 2`
    - Declares a variable `a` of `int` with value 2 
  - Multi declare
    ```golang
    var ( 
        b = 2 
        f = 2.01 
    )
    ```
    - It's implicitly typed based on the values assigned, so `b` is an `int` and `f` is a `float64`
- Inferred
  - `c:= 2`
  - This is called the "short declaration"
  - It's not an **assignment**, it's a short declaration
    - The type is implicitly defined
- Generally both are the same, and  can be used interchangeably
  - The only time we wanna use declarative is when we want to initialize in an outerscope 

## Variable initialization
- All variables are initialized to zero/equivalent if it's not initialized by user
- `int` gets 0, `floats` gets 0.0, `complex` gets 0i
- `bool` gets false
- `string` gets "" aka empty string
- Everything else is nil, which includes
  - pointers, slices, maps, channels, functions, interfaces

## Constants
- Immutable behavior, cannot be changed
- Only numbers (int, float), strings, booleans can be constants
  - Can't have contsant struct, array etc
  - This is because of concurrency in Golang
- Some examples of const declarations
```golang
const (
  a = 1
  b = 2 * 1024
  c = b << 3

  g uint8 = 0x07
  h uint8 = g & 0x03

  s = "a string"
  t = len(s)
  // THIS IS INVALID, since it's a slice
  u =  s[2:]
)
```