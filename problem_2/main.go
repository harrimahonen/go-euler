package main

import "fmt"

func main() {
	var prev int = 2
	var i int = 1
	var counter int
	for i <= 4000000 {
		fmt.Println(i, prev)
		if i%2 == 0 {
			counter += i
		}
		if prev%2 == 0 {
			counter += prev
		}
		i = prev + i    // i == 3 && prev == 2
		prev = prev + i // 2+3 = prev == 5
	}
	fmt.Println(counter)
}
