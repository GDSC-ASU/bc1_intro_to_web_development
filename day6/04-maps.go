package main

import "fmt"

func main() {
	mp := make(map[string]int)
	mp["square"] = 4
	mp["pentagon"] = 5
	mp["triangle"] = 3

	// using println statement
	fmt.Println(mp)

	// using for range
	for key, value := range mp {
		fmt.Println(key, value)
	}

	// delete an element
	delete(mp, "pentagon")
	fmt.Println(mp)

	fmt.Println("length of the map:", len(mp))
}
