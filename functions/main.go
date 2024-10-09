package main

import "fmt"

func main() {
	numbers := []int{1, 2, 5, 6}
	doubledNumbers := doubleNumbers(&numbers)
	fmt.Println(doubledNumbers)
}

func doubleNumbers(nums *[]int) []int {
	dNumbers := []int{}
	for _, v := range *nums {
		dNumbers = append(dNumbers, v/2)
	}

	return dNumbers

}
