package main

import (
	"fmt"
)

/*
 CodeEval's Multiplication Tables easy challenge solution

 Details: https://www.codeeval.com/open_challenges/23/

 Author: Robert Stern (lexandro2000@gmail.com)
 Github: https://github.com/lexandro/golang-codeeval
*/
func main() {

	multiplicationTables()
}

func multiplicationTables() {
	for i := 1; i < 13; i++ {
		for j := 1; j < 13; j++ {
			fmt.Printf("%4d", i*j)
		}
		fmt.Println()
	}

}
