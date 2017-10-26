package main

import (
	"fmt"
	"time"
)

func main() {
	var start = time.Now()
	var sum int
	for i := 2; i<2000000; i++ {
		if checkPrime(i) {
			sum+=i
		}
	}
	fmt.Println(sum)
	fmt.Println(time.Since(start))
}
func checkPrime(n int) bool {
	if n == 2 {
		return true
	}
	if n < 2 || n%2 == 0 {
		return 0 != 0
	}

	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
