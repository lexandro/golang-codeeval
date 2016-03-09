package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
 CodeEval's Set Intersection easy challenge solution

 Details: https://www.codeeval.com/open_challenges/30/

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
	elements := strings.Split(text, ";")
	set1 := strings.Split(elements[0], ",")
	set2 := strings.Split(elements[1], ",")
	var buffer bytes.Buffer
	//
	for i1, i2 := 0, 0; i1 < len(set1) || i2 < len(set2); {
		val1, _ := strconv.Atoi(set1[i1])
		val2, _ := strconv.Atoi(set2[i2])
		if val1 == val2 {
			if len(buffer.Bytes()) > 0 {
				buffer.WriteString(",")
			}
			buffer.WriteString(set1[i1])

		}
		if i1 < len(set1)-1 {
			if val1 <= val2 {
				i1++
				continue
			}
		}
		if i2 < len(set2)-1 {
			if val2 <= val1 {
				i2++
				continue
			}
		}
		break

	}
	fmt.Println(buffer.String())
}
