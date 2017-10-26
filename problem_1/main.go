package main

import "fmt"

func main() {
	var sum int
	for i := 0; i < 1000; i++ {
		if i%15 == 0 {
			sum += i
		} else if i%5 == 0 {
			sum = sum + i
		} else if i%3 == 0 {
			sum = sum + i
		}

	}
	fmt.Print(sum)
}
