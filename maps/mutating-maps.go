package main

import "fmt"

func mutatingMaps() {
	m := make(map[string]int)
	m["Answer"] = 42                        //set value
	fmt.Println("The value: ", m["Answer"]) // get value

	m["Answer"] = 48                       //reset value
	fmt.Println("The value:", m["Answer"]) //get value

	delete(m, "Answer")                    // delete value from map
	fmt.Println("The value:", m["Answer"]) // 0

	v, ok := m["Answer"] // check key exists
	fmt.Println("The value:", v, "Present?", ok)

}
