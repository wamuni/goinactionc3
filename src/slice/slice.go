package main

import "fmt"

func main() {
	slice := make([]int, 3, 5)
	fmt.Printf("Declare slice with make function, giving len: %d and cap: %d\n", len(slice), cap(slice))

	//the other way to declare slice, len and cap will be the same
	slice_string := []string{"Blue", "Red", "Yellow", "Orange"}
	print_slice_len_cap(slice_string)

	slice_string_index := []string{99: ""} // cap and len are both 100
	print_slice_len_cap(slice_string_index)

	//what is the difference between slice and array
	// from declaration perspective:
	arr := [3]int{0, 1, 2}
	fmt.Printf("The length of arr: %d\n", len(arr))
	arr[2] = 3
	slice_arr := []int{1, 2, 3}
	fmt.Printf("The length of slice: %d\n", len(slice_arr))

	var slice_nil []string
	print_slice_len_cap(slice_nil)
}

func print_slice_len_cap(slice []string) {
	fmt.Printf("The slice has length of %d and cap of %d\n", len(slice), cap(slice))
}
