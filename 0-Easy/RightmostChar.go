package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
 CodeEval's Rightmost Char easy challenge solution

 Details: https://www.codeeval.com/open_challenges/31

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
		rightmostChar(scanner.Text())
	}
}

func rightmostChar(text string) {
	elements := strings.Split(text, ",")
	fmt.Println(strings.LastIndex(elements[0], elements[1]))
}
