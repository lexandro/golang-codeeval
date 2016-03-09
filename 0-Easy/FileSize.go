package main

import (
	"fmt"
	"log"
	"os"
)

/*
 CodeEval's File Size easy challenge solution

 Details:https://www.codeeval.com/open_challenges/26/

 Author: Robert Stern (lexandro2000@gmail.com)
 Github: https://github.com/lexandro/golang-codeeval
*/
func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fileStat, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileStat.Size())
}
