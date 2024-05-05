package main

import "fmt"

func main() {
	fmt.Println("Hello Array")
	// First way to declare an array
	// It will initialize each value as zero-value of the type
	var arr [5]int
	print_array_int(arr[:])

	arr_second := [5]int{10, 20, 30, 40, 50}
	print_separater()
	print_array_int(arr_second[:])

	arr_third := [5]int{1: 20, 2: 40}
	print_separater()
	print_array_int(arr_third[:])

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
	print_separater()

	arr_string := [5]string{"Red", "Green", "Blue", "Yellow", "Orange"}
	print_array_string(arr_string)

	// array_string_pointer and copy between two arrays
	arr_string_pointer := [3]*string{new(string), new(string), new(string)}
	*arr_string_pointer[0] = "Blue"
	*arr_string_pointer[1] = "Red"
	*arr_string_pointer[2] = "Yellow"
	//TODO: Review the format to print different data type
	for index, str := range arr_string_pointer {
		fmt.Printf("The address %d of string in array is %o, and the value is %s\n", index+1, str, *str)
		fmt.Printf("The address %d of string in array is 0x%x, and the value is %s\n", index+1, str, *str)
	}
	print_separater()

	var arr_string_pointer_copy = arr_string_pointer
	*arr_string_pointer_copy[1] = "Eddie"
	for index, str := range arr_string_pointer_copy {
		fmt.Printf("The address %d of string in array is %o, and the value is %s\n", index+1, str, *str)
		fmt.Printf("The address %d of string in array is 0x%x, and the value is %s\n", index+1, str, *str)
	}
	print_separater()
	for index, str := range arr_string_pointer {
		fmt.Printf("The address %d of string in array is %o, and the value is %s\n", index+1, str, *str)
		fmt.Printf("The address %d of string in array is 0x%x, and the value is %s\n", index+1, str, *str)
	}

	// copy of two arrays with primitive data type
	var arr_string_copy = arr_string
	arr_string_copy[1] = "Eddie"
	print_separater()
	print_array_string(arr_string)
	print_separater()
	print_array_string(arr_string_copy)

}

// TODO: what is the difference between []int and []string
// Why it allow to pass [5]int as argument to a function with []int as parameter
// However it doesn't work for [5]string
func print_array_int(arr []int) {
	for index, num := range arr {
		fmt.Printf("The number in position %d is %d\n", index+1, num)
	}
}

func print_array_string(arr [5]string) {
	for index, str := range arr {
		fmt.Printf("The string in position %d is %s\n", index+1, str)
	}
}

func print_separater() {
	fmt.Println("======================================")
}
