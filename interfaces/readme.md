## Interfaces
- An interface is two things
  1. It is a set of methods
  2. It is also a type
-  Core concept in Go’s type system:  Talking about "what kind of data our types can hold".
- But interface talking about "what actions our types can execute".

### The interfac a set of methods
```go
type Animal interface {
    Speak() string
}
```
- **pretty simple:** we define an Animal as being any type that has a method named Speak
- The Speak method takes no arguments and returns a string.
- Any type that defines this method is said to satisfy the Animal interface.  There is no ~~implements~~ keyword in Go; whether or not a type satisfies an interface is determined automatically.
Let’s create a couple of types that satisfy this interface:

```go
type Myint int
type MyFloat float64

type Dog struct {
}

type Cat struct {
}

type Llama struct {
}


type JavaProgrammer struct {
}

// implement methods
func (d Dog) Speak() string {
    return "Woof!"
}

func (c Cat) Speak() string {
    return "Meow!"
}

func (l Llama) Speak() string {
    return "?????"
}

func (j JavaProgrammer) Speak() string {
    return "Design patterns!"
}

func (mi Myint) Speak() string {
   return "Number Number!"
}
```

### The interface{} type
- The interface{} type, the empty interface, is the source of much confusion. The interface{} type is the interface that has no methods. Since there is no implements keyword, all types implement at least zero methods, and satisfying an interface is done automatically, all types satisfy the empty interface.
-  That means that if you write a function that takes an interface{} value as a parameter, you can supply that function with any value. So, this function:

```go
func DoSomething(v interface{}) {
   // ...
}
```
It will accept any parameter whatsoever.

- **can I convert a []T to an []interface{}**
```go
package main

import (
    "fmt"
)

func PrintAll(vals []interface{}) {
    for _, val := range vals {
        fmt.Println(val)
    }
}

func main() {
    names := []string{"stanley", "david", "oscar"}
    PrintAll(names)
}
```
Run it here:[http://play.golang.org/p/4DuBoi2hJU](http://play.golang.org/p/4DuBoi2hJU)

By running this, you can see that we encounter the following error: cannot use names (type []string) as type []interface {} in function argument. If we want to actually make that work, we would have to convert the []string to an []interface{}:

```go
package main

import (
    "fmt"
)

func PrintAll(vals []interface{}) {
    for _, val := range vals {
        fmt.Println(val)
    }
}

func main() {
    names := []string{"stanley", "david", "oscar"}
    vals := make([]interface{}, len(names))
    for i, v := range names {
        vals[i] = v
    }
    PrintAll(vals)
}
```
Run it here:[http://play.golang.org/p/Dhg1YS6BJS](http://play.golang.org/p/Dhg1YS6BJ)


-
```shell
cmd > go build main.go
cmd > main
```