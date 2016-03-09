package main

import (
	"fmt"
)

/*
 CodeEval's Odd Numbers easy challenge solution

 Details: https://www.codeeval.com/open_challenges/25/

 Author: Robert Stern (lexandro2000@gmail.com)
 Github: https://github.com/lexandro/golang-codeeval
*/
func main() {
	for i := 1; i < 100; i += 2 {
		fmt.Println(i)
	}
}
