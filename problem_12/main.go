package main

import (
	"fmt"
	"math"
)

func main(){
	result := 1
	triangle := 0
	x := 1
	for result < 500 {
		result = getDivisors(triangle)
		fmt.Println("trying ", result, triangle)

		triangle = triangle + x
		x++
	}
	fmt.Println("done", result, triangle)
}

func getDivisors(num int) int{
	divs := 0
	sqrt := int(math.Sqrt(float64(num)))

	for x := 1; x <= sqrt; x++ {
		if num % x == 0 {
			divs = divs + 2
		}
	}
	if sqrt * sqrt == num {
		divs--
	}
	return divs
}
func calculateTriangleNum(x int) int {
	result := 0
	counter := 0
	for i := 1; counter < x; i++ {
		result = result+i


		counter++
	}
	return result
}
