package main

import "fmt"

func main() {
	var toThisNumber int = 100
	var sumOfSq int
	var SqOfSum int

	for i := 1; i <= toThisNumber; i++ {
		sumOfSq += i * i
		fmt.Println("Current sumOfSq: ", sumOfSq)

		SqOfSum += i
	}
	fmt.Println("sum of squares: ", sumOfSq)
	fmt.Println(SqOfSum, "square is ", SqOfSum*SqOfSum)
	fmt.Println("Answer is ", SqOfSum*SqOfSum-sumOfSq)
}
