package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()
	var prime float64 = 600851475143
	var primesqrt int = int(math.Sqrt(prime))
	for i := 2; i < primesqrt; i++ {
		if int(prime)%i == 0 {
			fmt.Print(i)
			for o := 2; o < int(float64(i)); o++ {
				if i%o == 0 {
					fmt.Print(" not prime")
					break
				}

			}
			fmt.Print("\n")
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("%s", elapsed)
}
