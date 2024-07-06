# go-ptr

![CI](https://github.com/jackcipher/go-ptr/actions/workflows/ci.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/jackcipher/go-ptr)](https://goreportcard.com/report/github.com/jackcipher/go-ptr)

The ptr package provides utility functions for working with pointers in Go. It includes functions to create pointers for various types, check if a pointer is nil, and retrieve values from pointers with default fallbacks.


## Installation

To install the ptr package, use go get:

```sh
go get github.com/jackcipher/go-ptr
```


## Usage

Import the `ptr` package into your Go project:

```go
import (
    "github.com/jackcipher/go-ptr"
    "time"
)
```

### Creating Pointers

The `ptr` package provides functions to create pointers for various types:

```go
v := 42
p := ptr.Int(v)
fmt.Println(*p) // Output: 42

s := "hello"
sp := ptr.String(s)
fmt.Println(*sp) // Output: hello

b := true
bp := ptr.Bool(b)
fmt.Println(*bp) // Output: true

f := 3.14
fp := ptr.Float64(f)
fmt.Println(*fp) // Output: 3.14

t := time.Now()
tp := ptr.Time(t)
fmt.Println(*tp) // Output: current time

slice := []int{1, 2, 3}
slicePtr := ptr.SlicePtr(slice)
fmt.Println(*slicePtr) // Output: [1, 2, 3]

m := map[string]int{"a": 1, "b": 2}
mapPtr := ptr.MapPtr(m)
fmt.Println(*mapPtr) // Output: map[a:1 b:2]
```


### Checking for `nil`

Use `ptr.IsNil` to check if a pointer is `nil`:

```go
var p *int
fmt.Println(ptr.IsNil(p)) // Output: true

v := 42
p = &v
fmt.Println(ptr.IsNil(p)) // Output: false
```


### Retrieving Values with Defaults

Use `ptr.GetValueOrDefault` to retrieve a pointer's value, or return a default value if the pointer is `nil`:

```go
v := 42
p := &v
fmt.Println(ptr.GetValueOrDefault(p, 0)) // Output: 42

var nilP *int
fmt.Println(ptr.GetValueOrDefault(nilP, 0)) // Output: 0
```

Use ptr.GetNonZeroValueOrDefault to retrieve a non-zero value from a pointer, or return a default value if the pointer is nil or points to a zero value:

```go
v := 42
p := &v
fmt.Println(ptr.GetNonZeroValueOrDefault(p, 0)) // Output: 42

zero := 0
zeroP := &zero
fmt.Println(ptr.GetNonZeroValueOrDefault(zeroP, 42)) // Output: 42
```


### Creating Pointers Conditionally

Use `ptr.NewPtrIf` to create a pointer based on a condition:

```go
v := 42
p := ptr.NewPtrIf(true, v)
fmt.Println(*p) // Output: 42

p = ptr.NewPtrIf(false, v)
fmt.Println(p) // Output: <nil>
```


### Comparing Pointers


Use `ptr.Equal` to compare the values pointed to by two pointers:

```go
a, b := 42, 42
fmt.Println(ptr.Equal(&a, &b)) // Output: true

c := 43
fmt.Println(ptr.Equal(&a, &c)) // Output: false
```



### Functional Utilities

Use `ptr.Map` to apply a function to each element of a slice and return a new slice:

```go
input := []int{1, 2, 3}
result := ptr.Map(input, func(v int) int { return v * 2 })
fmt.Println(result) // Output: [2, 4, 6]
```

Use ptr.Filter to filter elements of a slice based on a predicate function:

```go
input := []int{1, 2, 3, 4, 5}
result := ptr.Filter(input, func(v int) bool { return v%2 == 0 })
fmt.Println(result) // Output: [2, 4]
```

## Testing

The `ptr` package includes a comprehensive test suite. To run the tests, use the `go test` command:

```sh
go test github.com/jackcipher/go-ptr
```

## License
This project is licensed under the MIT License. See the LICENSE file for details.

## Contributing
Contributions are welcome! Please submit a pull request or open an issue to discuss any changes.

## Acknowledgements
This project was inspired by the need for utility functions to simplify working with pointers in Go.