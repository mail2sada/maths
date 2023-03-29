package airthmatic

import "fmt"

var (
	gloablToPkg int
	Exported    int = 10000
)

func init() {
	fmt.Println("Init from airthmatic1..")
}

func init() {
	fmt.Println("Init from airthmatic2..")
}

func Add(element ...int) int {
	gloablToPkg = 0
	sum := 0

	for _, val := range element {
		sum += val
	}
	return sum
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) int {
	return a / b
}
