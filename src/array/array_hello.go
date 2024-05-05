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

	// get access of pointer value of array
	arr_pointer := [5]*int{0: new(int), 1: new(int)}
	*arr_pointer[0] = 10
	*arr_pointer[1] = 20
	for index, num := range arr_pointer {
		// if variable is declared by new, there will be no zero-value initialized
		// instead, it will be initialized as nil
		// and if it's nil, then it can't use *
		fmt.Println(index, num)
	}
}

func print_array(arr []int) {
	for index, num := range arr {
		fmt.Printf("The number in position %d is %d\n", index+1, num)
	}
}

func print_separater() {
	fmt.Println("======================================")
}
