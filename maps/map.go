//Package main - map example
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex
var m1 = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}
var m2 = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	m = make(map[string]Vertex)
	m["Pilath Raj"] = Vertex{40.23232, -74.343434}
	fmt.Println(m["Pilath Raj"]) // {40.23232 -74.343434}
	fmt.Println(m)               // map[Pilath Raj:{40.23232 -74.343434}]
	fmt.Println(m1)              //map[Google:{37.42202 -122.08408} Bell Labs:{40.68433 -74.39967}]
	fmt.Println(m2)              //map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]
	mutatingMaps()
	wordCountinfo()
}
