package main

import (
	"fmt"
	"strconv"
)

func main() {
	// start at 999 * 999, then decrease latter number until 900 then decrease former to until 900, check if this results in a palindrome number
	var index []int
	for i := 999; 100 < i; i-- {
		for o := 999; 100 < o; o-- {
			var resultInt int = i * o
			var result string = strconv.Itoa(i * o)
			if len(result) == 6 {
				if result[0] == result[5] && result[1] == result[4] && result[2] == result[3] {
					index = append(index, resultInt)
				}
			}
			if len(result) == 5 {
				break
			}
		}
	}

	// GET MAX Int value in []
	var currMax int
	for _, e := range index {
		if e >= currMax {
			currMax = e
		}
	}
	fmt.Println(currMax)

}
