package main

import (
	"fmt"
	"strconv"
)

type Student struct {
	Name  string
	Phone string
	Gpa   float64
}

func (s *Student) ConvertPhoneNumberToInteger() int64 {
	numberWithoutPlus := s.Phone[1:]

	intNumber, err := strconv.ParseInt(numberWithoutPlus, 10, 64)
	if err != nil {
		panic(err)
	}

	return intNumber
}

func main() {
	ali := Student{Name: "Ali Al-Homsi", Phone: "+962791234567", Gpa: 90.2}
	fmt.Println("ali's data", ali)
}
