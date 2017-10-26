package main

import "fmt"

func main() {
	var input int = 20
	//var result int
	var keepCounting bool = true
	var keepScore int
	var intro string = `2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
	What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?`
	fmt.Println(intro)

	for i := 20; keepCounting; i++ {
		for counter := input; 1 <= counter; counter-- { // while counter is bigger or exactly 1
			if i%counter != 0 { // i / counter, not evenly divisible
				keepScore = 0
				break
			}
			keepScore++
			if keepScore == 20 {
				fmt.Println("DONE", i)
				keepCounting = 0 != 0
			}
		}

	}
}
