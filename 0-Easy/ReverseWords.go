package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
 CodeEval's Reverse Words easy challenge solution

 Details: https://www.codeeval.com/open_challenges/8/

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
		reverseWords(scanner.Text())
	}
}

func reverseWords(text string) {
	var buffer bytes.Buffer
	elements := strings.Split(text, " ")

	for i := len(elements) - 1; i >= 0; i-- {
		buffer.WriteString(elements[i])
		buffer.WriteString(" ")
	}
	fmt.Println(buffer.String())
}
