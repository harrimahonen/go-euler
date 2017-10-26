package main

import (
	"fmt"
	"math"
)

func main(){
desiredNum := 1000
check := true
var sum float64
	var a float64
	var b float64
	for a = 1; check; a++ {
	for b = a+1; b < 400; b++ {
		var c float64 = a*a+b*b
		var cSq float64 = math.Sqrt(c)
		if isPyth(a,b,c) {
			sum = a + b + cSq
			if int(sum) > desiredNum {
				break
			}
			if sum == float64(desiredNum) {
				fmt.Println("Done", a, b, cSq, sum)
				fmt.Println("Product abc: ", int(a*b*cSq))
				check = false
				break
			}
		}
	}

}
}

func isPyth(a float64, b float64, c float64) bool{
	if a*a+b*b == c{
		return true
	} else {
		return false
	}
}