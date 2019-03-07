// Package main struct
package main

import "fmt"

type Student struct {
	id   int
	name string
	age  int
	dob  string
}

var (
	s1 = Student{1, "ragu", 28, "01-01-2000"}  // has type Student
	s2 = Student{id: 10, name: "ravi"}         // age:0 and dob:"" are implicit
	s3 = Student{}                             //id:0, name:"",age:0, dob:""
	p1 = &Student{1, "main", 28, "11-12-2000"} // has type *Student
)

func main() {
	s := Student{1, "pilathraj", 28, "01-01-2000"}
	fmt.Println(s)                   // {1 pilathraj 28 01-01-2000}
	fmt.Printf("Name: %s\n", s.name) // pilathraj
	s.age = 20
	fmt.Printf("Age is: %d\n", s.age) // 20

	p := &s
	// *p.dob = "12-12-2000"  error -bcz- language permits us instead to write just p.dob
	p.dob = "12-12-2000"
	fmt.Printf("Date of birth is: %s\n", s.dob) // retrive form struct 12-12-2000
	s.age = 30
	fmt.Printf("Date of birth is: %d\n", p.age) // read age through pointer 30

	fmt.Println(s1, s2, s3, p1) // {1 ragu 28 01-01-2000} {10 ravi 0 } {0  0 } &{1 main 28 11-12-2000}

}
