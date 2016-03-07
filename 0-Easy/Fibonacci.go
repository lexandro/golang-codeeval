package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

/*
 CodeEval's Fibonacci easy challenge solution

 Details: https://www.codeeval.com/open_challenges/22/

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
	n, _ := strconv.Atoi(text)
	f1 := 1;
	if n > 0 {
		f0 := 0;

		for i := 1; i < n; i++ {
			f1 = f0 + f1;
			f0 = f1 - f0;
		}

	} else {
		f1 = 0;
	}
	fmt.Println(f1)
}
