package main

import "fmt"

func main() {
	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("all array's elements", arr)

	fmt.Println("element by element")
	// using counter loop
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// using for range
	for index, value := range arr {
		fmt.Println(index, value)
	}

	arr[0] = 99
	fmt.Println(arr)

}
