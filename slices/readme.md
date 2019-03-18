## Slices
- An array has a fixed size. A slice, on the other hand, is a **dynamically-sized**, flexible view into the elements of an array. In practice, slices are much more common than arrays.

The type []T is a slice with elements of type T.

A slice is formed by specifying two indices, a low and high bound, separated by a colon:
a[low : high]
This selects a half-open range which includes the first element, but excludes the last one.

The following expression creates a slice which includes elements 1 through 3 of a:
a[1:4]

### Slices are like references to arrays
- A slice does not store any data, it just describes a section of an underlying array.
- Changing the elements of a slice modifies the corresponding elements of its underlying array.
- Other slices that share the same underlying array will see those changes.

### Slice literals
- A slice literal is like an array literal **without the length**.
-This is an array literal:
```go
[3]bool{true, true, false}
```
And this creates the same array as above, then builds a slice that references it:
```go
[]bool{true, true, false}
```

### Slice defaults
- When slicing, you may omit the high or low bounds to use their defaults instead.
- The default is **zero** for the **low bound** and the **length** of the slice for the **high bound**.

### Slice length and capacity
- The length and capacity of a slice s can be obtained using the expressions **len(s)** and **cap(s)**.

### Nil slices
- The zero value of a slice is nil.
- A nil slice has a length and capacity of 0 and has no underlying array.

### Creating a slice with make
- Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.

The make function allocates a zeroed array and returns a slice that refers to that array:

```go
a := make([]int, 5)  // len(a)=5
```
To specify a capacity, pass a third argument to make:

```go
b := make([]int, 0, 5) // len(b)=0, cap(b)=5
b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
```

### Slices of slices
- Slices can contain any type, including other slices.

### Range
- The range form of the for loop iterates over a **slice or map**.
- When ranging over a slice, two values are returned for each iteration. The first is the **index**, and the second is a copy of the element's **value** at that index.
- You can skip the index or value by assigning to _.
- If you only want the index, you can omit the second variable.

```shell
cmd > go build
cmd > slices
```