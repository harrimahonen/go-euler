package main

import "fmt"

func main(){
	result := 0
	x := 2
	triangle := 1
	for result < 500 {
		result = getDivisors(triangle)
		fmt.Println("trying ", result, triangle)

		triangle = triangle + x
		x++
	}
	fmt.Println("done", result, triangle)
}

func getDivisors(num int) int{
	divs := 1
	if num == 1 {return 1}

	for x := 1; x <= num/2+1; x++ {
		if num % x == 0 {
			divs++
		}
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
