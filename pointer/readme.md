## Pointers
- A pointer holds the memory address of a value.
- The type *T is a pointer to a T value. Its zero value is nil.

```go
var p *int
```

The & operator generates a pointer to its operand.

```go
i := 42
p = &i
```
The * operator denotes the pointer's underlying value.
```go
fmt.Println(*p) // read i through the pointer p
*p = 21         // set i through the pointer p
```


```shell
cmd > go build main.go
cmd > main
```