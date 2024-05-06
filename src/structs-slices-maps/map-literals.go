package main

import "fmt"

/**
映射的字面量和结构体类似，只不过必须有键名。
*/

type VertexMap struct {
	Lat, Long float64
}

var maps = map[string]VertexMap{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(maps)
}
