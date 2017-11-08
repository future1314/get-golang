package main

import "fmt"

type Vertex2 struct {
	Lat, Long float64
}

var m map[string]Vertex2

func main() {
	m = make(map[string]Vertex2)
	m["Bell Lab"] = Vertex2{
		40.68433, -74.39967,
	}
	fmt.Println(m)

	m2 := map[string]Vertex2{
		"Bell Labs": Vertex2{
			40.68433, -74.39967,
		},
		"Google": Vertex2{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m2)

	m3 := map[string]Vertex2{
		"Bell Labs": {
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}
	delete(m3, "Google")
	v1, ok1 := m3["Bell Labs"]
	v2, ok2 := m3["Google"]
	fmt.Println(m3)
	fmt.Println(v1, ok1)
	fmt.Println(v2, ok2)
}
