package main

import (
	"fmt"
	"os"

	"github.com/goinaction/code/chapter3/words"
)

func init() {
	fmt.Println("Hello from init function")
}

func main() {
	fmt.Println("Hello World")
	filename := os.Args[0]
	contents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	text := string(contents)
	count := words.CountWords(text)
	fmt.Printf("There're %d words in your file\n", count)
}
