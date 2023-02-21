package main

import "fmt"

func main() {
	var x = 12
	var y int = x + 8
	fmt.Println("y = ", y)

	var z = 12.4
	fmt.Println("z + x =", (z + float64(x)))

	var name = "Ali"
	fmt.Println("Hello,", name)

	var firstName = "Ali"
	var lastName = "Al-Homsi"
	var fullName = fmt.Sprintf("%s %s", firstName, lastName)
	fmt.Println("Hello", fullName)

	canIDrive := false
	fmt.Println("Am I old enough to drive?", canIDrive)
}
