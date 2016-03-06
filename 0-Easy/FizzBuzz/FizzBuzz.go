package main

import "fmt"
import "log"
import "strings"
import "strconv"
import "bufio"
import "os"

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
		if i % firstDivider == 0 {
			fmt.Print("F")
		}
		if i % secondDivider == 0 {
			fmt.Print("B")
		}
		if i % firstDivider != 0 && i % secondDivider != 0 {
			fmt.Print(i)
		}
		fmt.Print(" ")
	}
	fmt.Println()
}