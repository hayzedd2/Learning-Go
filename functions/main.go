package main

import (
	"fmt"
)

func main() {
	fact := factorial(6)
	numbers := []int{1, 3, 4, 5, 5}
	sum := sumUp(numbers...)
	fmt.Println(fact)
	fmt.Println(sum)
}

func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * (number - 1)
	// result := 1
	// for i := 1; i <= number; i++ {
	// 	result *= i
	// }
	// return result
}

func sumUp(numbers ...int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}
