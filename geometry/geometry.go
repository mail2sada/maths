package geometry

import "fmt"

func init() {
	fmt.Println("Init from geometry..")
}

func AreaSqure(side int) int {
	return side * side
}

func AreaRectangle(lenght int, breadth int) int {
	return lenght * breadth
}

func AreaTriangle(base int, height int) int {
	return (base * height) / 2
}
