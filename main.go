package main

import (
	"fmt"
	am "maths/airthmatic"
	_ "maths/airthmatic/advanced"
	"maths/geometry"
)

func init() {
	fmt.Println("Init in main1")
}
func init() {
	fmt.Println("Init in main2")
}

func main() {
	fmt.Println("Demo: Packages...")
	a := am.Add(10, 20, 30, 40, 50)
	fmt.Println("Result is ", a)
	fmt.Println(am.Subtract(10, 20))
	fmt.Println(am.Multiply(10, 20))
	fmt.Println(am.Exported)
	//fmt.Println(advanced.Squre(100))

	fmt.Println(geometry.AreaRectangle(10, 20))
	fmt.Println(geometry.AreaSqure(50))
	fmt.Println(geometry.AreaTriangle(50, 10))
}
