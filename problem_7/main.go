package main

import (
	"fmt"
	"time"
)

func main() {
	var start = time.Now()
	var keepScore int = 0
	for i := 2; keepScore < 1000000; i++ {
		if checkPrime(i) {
			keepScore++
		}
	}
	fmt.Print(time.Since(start))
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
