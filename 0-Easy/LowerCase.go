package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
 CodeEval's Bit Positions easy challenge solution

 Details: https://www.codeeval.com/open_challenges/19/

 Author: Robert Stern (lexandro2000@gmail.com)
 Github: https://github.com/lexandro/golang-codeeval
*/
func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lowerCase(scanner.Text())
	}
}

func lowerCase(text string) {
	fmt.Println(strings.ToLower(text))
}
