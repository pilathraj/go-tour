## Function as value (anonymous function)
 - A function in go can also be a value. This means you can assign a function to a variable.
### Not Immediately invoked function


```go
package main

import "fmt"

//declare global
var add = func(a int, b int) int {
	return a + b
}

func main() {
    //declare in main
	var sub = func(a int, b int) int {
		return a - b
	}
	fmt.Println("5 + 3 =", add(5, 3))
	fmt.Println("5 - 3 =", sub(5, 3))
}


```

###  Immediately invoked function

```go
package main

import "fmt"

var add = func(a int, b int) int {
	return a + b
}(5, 3)

func main() {
	var sub = func(a int, b int) int {
		return a - b
	}(5, 3)
	fmt.Println("5 + 3 =", add)
	fmt.Println("5 - 3 =", sub)
}

```


