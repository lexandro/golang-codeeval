package main

import (
	"math"
	"fmt"
)

func main() {
	sum := 0;
	primeCounter := 0;
	for i := 2; primeCounter < 1000; i++ {
		if isPrime(i) {
			sum += i;
			primeCounter++;
		}
	}
	fmt.Println(sum)
}

func isPrime(number int) bool {
	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number % i == 0 && number != i {
			return false
		}
	}
	return true
}

