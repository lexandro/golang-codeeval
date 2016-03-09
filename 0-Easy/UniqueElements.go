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
 CodeEval's Unieque Elements easy challenge solution

 Details: https://www.codeeval.com/open_challenges/29

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
		uniqueElements(scanner.Text())
	}
}

func uniqueElements(text string) {
	elements := strings.Split(text, ",")
	var prev string
	var buffer bytes.Buffer
	for _, ele := range elements {
		if prev != ele {
			prev = ele
			if len(buffer.Bytes()) > 0 {
				buffer.WriteString(",")
			}
			buffer.WriteString(prev)
		}

	}
	fmt.Println(buffer.String())
}
