## Methods
 - Go does not have **~~classes~~**. However, you can define methods on types.
 - A method is a function with a special receiver argument.
 -The receiver appears in its own argument list between the func keyword and the method name.
 ```go
 func receiver_args method_name(args) return_type
```
Example:
```go
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```
- (v Vertex) - receiver "v" has a type of "Vertex"
- Abs method name
- float64 return type

- **Remember:** a method is just a function with a receiver argument.
- You can declare a method on non-struct types, too.
- You can declare methods with pointer receivers.

### Choosing a value or pointer receiver
- There are two reasons to use a pointer receiver.
- The first is so that the method can modify the value that its receiver points to.
- The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct,.
- **Note:** In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both.
