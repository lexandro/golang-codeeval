package main

import "fmt"
import "strings"
import "strconv"
import "math"

func main() {
	for i := 1000; i > 0; i-- {
		if isPrime(i) && isPalindrome(i) {
			fmt.Println(i)
			break
		}
	}
}

func isPrime(number int) bool {
	for i := 2; float64(i) < math.Sqrt(float64(number)); i++ {
		if number % i == 0 {
			return false
		}
	}
	return true
}

func isPalindrome(number int) bool {
	n1 := strconv.Itoa(number)
	n2 := reverse(n1)
	return strings.Compare(n1, n2) == 0
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r) - 1; i < len(r) / 2; i, j = i + 1, j - 1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}