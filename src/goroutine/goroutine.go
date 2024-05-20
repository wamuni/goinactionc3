package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Printf("Goroutine, bro!\n")
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Printf("Start Goroutine...\n")

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	fmt.Printf("Waiting to Finish\n")
	wg.Wait()
	fmt.Printf("\nTerminating Program...")
}
