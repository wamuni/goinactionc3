package main

import "fmt"

func main() {
	fmt.Println("Hello Array")
	// First way to declare an array
	// It will initialize each value as zero-value of the type
	var arr [5]int
	print_array(arr[:])

	arr_second := [5]int{10, 20, 30, 40, 50}
	print_separater()
	print_array(arr_second[:])

	arr_third := [5]int{1: 20, 2: 40}
	print_separater()
	print_array(arr_third[:])
}

func print_array(arr []int) {
	for index, num := range arr {
		fmt.Printf("The number in position %d is %d\n", index+1, num)
	}
}

func print_separater() {
	fmt.Println("======================================")
}
