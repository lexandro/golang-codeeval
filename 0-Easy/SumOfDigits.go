package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

/*
 CodeEval's Sum of Digits easy challenge solution

 Details: https://www.codeeval.com/open_challenges/21/

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
		sumOfDigits(scanner.Text())
	}
}

func sumOfDigits(text string) {
	sum := 0
	for _, r := range text {
		value, _ := strconv.Atoi(string(r))
		sum += value
	}
	fmt.Println(sum)
}
