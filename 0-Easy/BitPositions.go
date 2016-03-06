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
		bitPositions(scanner.Text())
	}
}

func bitPositions(text string) {
	elements := strings.Split(text, ",")
	number, _ := strconv.Atoi(elements[0])
	bit1, _ := strconv.Atoi(elements[1])
	bit2, _ := strconv.Atoi(elements[2])

	fmt.Println(hasBit(number, uint(bit1)) == hasBit(number, uint(bit2)))
}

func hasBit(n int, pos uint) bool {
	val := n & (1 << (pos - 1))
	return (val > 0)
}
