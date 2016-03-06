package main

import (
	"fmt"
	"math"
)

/*
 CodeEval's Sum of Primes easy challenge solution

 Details: https://www.codeeval.com/open_challenges/4/

 Author: Robert Stern (lexandro2000@gmail.com)
 Github: https://github.com/lexandro/golang-codeeval
*/
func main() {
	sum := 0
	primeCounter := 0
	for i := 2; primeCounter < 1000; i++ {
		if isPrimeNumber(i) {
			sum += i
			primeCounter++
		}
	}
	fmt.Println(sum)
}

func isPrimeNumber(number int) bool {
	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 && number != i {
			return false
		}
	}
	return true
}
