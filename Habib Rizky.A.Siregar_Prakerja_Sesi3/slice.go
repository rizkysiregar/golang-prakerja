package main

import "fmt"

func main() {
	// Creating a slice using a slice literal
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Original slice:", numbers)

	// Accessing elements of a slice
	fmt.Println("First element:", numbers[0])
	fmt.Println("Third element:", numbers[2])

	// Modifying an element of a slice
	numbers[1] = 10
	fmt.Println("Modified slice:", numbers)

	// Slicing a slice to create a new slice
	sliced := numbers[1:4]
	fmt.Println("Sliced slice:", sliced)

	// Modifying the sliced slice modifies the original slice
	sliced[1] = 20
	fmt.Println("Modified sliced slice:", sliced)
	fmt.Println("Original slice after modification:", numbers)

	// Appending elements to a slice
	numbers = append(numbers, 6, 7, 8)
	fmt.Println("Slice after appending elements:", numbers)

	// Copying a slice
	copySlice := make([]int, len(numbers))
	copy(copySlice, numbers)
	fmt.Println("Copied slice:", copySlice)
}
