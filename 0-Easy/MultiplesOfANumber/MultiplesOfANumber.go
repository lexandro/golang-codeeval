package main

import "log"
import "strings"
import "bufio"
import "os"
import (
	"strconv"
	"fmt"
)

/*
 CodeEval's Multiples of a Number easy challenge solution

 Details: https://www.codeeval.com/open_challenges/18/

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
		solveChallenge(scanner.Text())
	}
}

func solveChallenge(text string) {
	elements := strings.Split(text, ",")
	x, _ := strconv.Atoi(elements[0])
	n, _ := strconv.Atoi(elements[1])
	m := n
	for m < x {
		m += n
	}
	fmt.Println(m)
}