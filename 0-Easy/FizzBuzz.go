package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
 CodeEval's FizzBuzz easy challenge solution

 Details: https://www.codeeval.com/open_challenges/1/

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
		fizzBuzz(scanner.Text())
	}
}

func fizzBuzz(text string) {
	elements := strings.Split(text, " ")
	//
	firstDivider, _ := strconv.Atoi(elements[0])
	secondDivider, _ := strconv.Atoi(elements[1])
	maxCount, _ := strconv.Atoi(elements[2])
	//
	for i := 1; i <= maxCount; i++ {
		if i%firstDivider == 0 {
			fmt.Print("F")
		}
		if i%secondDivider == 0 {
			fmt.Print("B")
		}
		if i%firstDivider != 0 && i%secondDivider != 0 {
			fmt.Print(i)
		}
		fmt.Print(" ")
	}
	fmt.Println()
}
