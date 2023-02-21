package main

import "fmt"

func nothing() {
	fmt.Println("nothing")
}

func greet(name string) {
	fmt.Println(`Hello, ${name}`)
}

func square(n int) int {
	return n * n
}

var nothing2 = func() {
	fmt.Println("nothing2")
}

var add = func(x, y, z int) int {
	return x + y + z
}

func main() {
	nothing()
	greet("Jaber")
	fmt.Println(square(4))
	nothing2()
	fmt.Println(add(1, 2, 3))
}
