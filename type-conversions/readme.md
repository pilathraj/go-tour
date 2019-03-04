# Type conversions
The expression T(v) converts the value v to the type T.

Some numeric conversions:

```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```

Or, put more simply:
```go
i := 42
f := float64(i)
u := uint(f)
```

```shell
cmd > go build main.go
cmd > main
cmd > 3 4 5
```