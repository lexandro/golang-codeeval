package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*
 CodeEval's  XXXXX XXXXXX easy challenge solution

 Details: https://www.codeeval.com/open_challenges/XXXXXXXXXXXXXX

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
		problemSolver(scanner.Text())
	}
}

func problemSolver(text string) {
	fmt.Println(text)
}
