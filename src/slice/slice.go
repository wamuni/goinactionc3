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

	var slice_nil []string           // this will declare a null slice
	slice_empty := make([]string, 0) // this will declare an empty slice
	slice_empty_iteral := []string{} // this will declare an empty slice
	//TODO: what will be the difference between a nil slice and empty slice
	//! nil slice is used for return value of handling errors
	//! empty slice is used for indicating 0 value returned
	print_slice_len_cap(slice_nil)
	print_slice_len_cap(slice_empty)
	print_slice_len_cap(slice_empty_iteral)

	print_separater()
	slice_to_slice := []int{10, 20, 30, 40, 50} // cap: 5, len: 5
	print_slice_len_cap_int(slice_to_slice)
	new_slice := slice_to_slice[1:3]
	print_slice_len_cap_int(new_slice) // cap: 4, len: 2
	//! the new slice cap will be the old cap - start index
	// which means new slice won't have access to the elements before start index
	//TODO: Does that mean the new slice are pointing at the same array?
	slice_to_slice[2] = 100
	fmt.Printf("New change in slice: %d\n", slice_to_slice[2])
	fmt.Printf("Element in new slice: %d\n", new_slice[1])
	//? The new slice are pointing as the same array address with the old slice

	print_separater()
	// append algorithm check
	// append_algorithm_check()

}

func append_algorithm_check() {
	//* Through the observation
	//? 1. it will start to double
	//? 2. then 1.75 around 500
	//? 3. around 1.4 around 1000
	// TODO: Check the append algorithm
	slice := []int{0, 1, 2, 3, 4, 5}
	for i := 0; i < 2000; i += 1 {
		slice = append(slice, i)
		print_slice_len_cap_int(slice)
	}
}

func print_slice_len_cap_int(slice []int) {
	fmt.Printf("The int slice has length of %d and cap of %d\n", len(slice), cap(slice))
}

func print_slice_len_cap(slice []string) {
	fmt.Printf("The slice has length of %d and cap of %d\n", len(slice), cap(slice))
}

func print_separater() {
	fmt.Println("=================================================")
}
