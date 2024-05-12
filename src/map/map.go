package main

import "fmt"

type MapType interface {
	int | string
}

func main() {
	fmt.Printf("Hello MotherFucker!\n")

	//* How to declare a map in golang
	//? use make function
	//? use iteral declaration
	dict := make(map[string]int) // -> nil map
	dict_color := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}

	value, exist := dict["Red"]
	if exist {
		fmt.Printf("dict has value %d for key %s\n", value, "Red")
	} else {
		fmt.Printf("dict has not value for key %s\n", "Red")
	}

	if value != 0 { // if there is no value, then the value returned will be zero-value of the type
		fmt.Printf("dict has value %d for key %s\n", value, "Red")
	} else {
		fmt.Printf("dict has no value for key %s, therefore, zero-value is assigned\n", "Red")
	}

	value_color, exist_color := dict_color["Red"]
	if exist_color {
		fmt.Printf("dict color has value %s for key %s\n", value_color, "Red")
	}

	print_separater()

	dict_more_color := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}
	print_map(dict_color)
	print_map(dict_more_color)
	// and every time, it won't be in the same order

	// how to delete element in a map
}

// function using generic in golang
func print_map[K comparable, V MapType](m map[K]V) {
	for key, value := range m {
		fmt.Print("Key: ")
		fmt.Print(key)
		fmt.Print(" -> ")
		fmt.Print("Value: ")
		fmt.Print(value)
		fmt.Printf("\n")
	}
}

func print_separater() {
	fmt.Printf("------------------------------------------------------------------------\n")
}
