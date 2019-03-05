## Switch
- A switch statement is a shorter way to write a sequence of if - else statements. It runs the first case whose value is equal to the condition expression.
- Go automatically break statement end of the case.
- Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.
- Switch cases evaluate cases from top to bottom, stopping when a case succeeds.
ex.

```go
switch i {
case 0:
case f():
}
```
**does not call f if i==0.)**

```shell
cmd > go build main.go
cmd > main
```