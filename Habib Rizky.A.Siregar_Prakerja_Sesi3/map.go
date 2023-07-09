package main

import "fmt"

func main() {
	// Creating a map to store employee information
	employee := make(map[string]string)

	// Adding key-value pairs to the map
	employee["name"] = "John Doe"
	employee["position"] = "Software Engineer"
	employee["department"] = "Engineering"

	// Accessing values from the map
	fmt.Println("Name:", employee["name"])
	fmt.Println("Position:", employee["position"])
	fmt.Println("Department:", employee["department"])

	// Updating a value in the map
	employee["position"] = "Senior Software Engineer"
	fmt.Println("Updated Position:", employee["position"])

	// Deleting a key-value pair from the map
	delete(employee, "department")
	fmt.Println("Map after deletion:", employee)

	// Checking if a key exists in the map
	_, exists := employee["name"]
	fmt.Println("Does 'name' exist in the map?", exists)

	// Iterating over key-value pairs in the map
	for key, value := range employee {
		fmt.Printf("%s: %s\n", key, value)
	}
}
