package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

/*
 CodeEval's Sum Of Integers From File easy challenge solution

 Details: https://www.codeeval.com/open_challenges/24

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
	sum := 0
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		sum += number
	}
	fmt.Println(sum)
}
