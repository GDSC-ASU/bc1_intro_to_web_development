package main

import "fmt"

func main() {
	var shallIPass = true
	if shallIPass {
		fmt.Println("You shall pass")
	} else {
		fmt.Println("You shall not pass")
	}

	x, y, z := 12, 10, 13

	if x > y && x > z {
		fmt.Println("x is the greatest")
	} else if y > x && y > z {
		fmt.Println("y is the greatest")
	} else if z > x && z > y {
		fmt.Println("z is the greatest")
	}

	var mark = "B"
	switch mark {
	case "A":
		fmt.Println("big nerd")
	case "B":
		fmt.Println("nerd")
	case "C":
		fmt.Println("ain't much")
	default:
		fmt.Println("I don't know what is that!")
	}

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	var count = 10
	for count > 0 {
		fmt.Println(count)
		count--
	}

}
