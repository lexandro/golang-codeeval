package main

import "fmt"
import "log"
import "strings"
import "bufio"
import "os"
import "bytes"

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

	for i := len(elements)-1; i >= 0; i-- {
		buffer.WriteString(elements[i])
		buffer.WriteString(" ")
	}
	fmt.Println(buffer.String())
}