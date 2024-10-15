package fn

import (
	"fmt"
)

type operationFn func(int) int

func fn() {
	num := []int{1, 6, 9, 16}
	// num2 := []int{2, 6, 9, 16}
	// surName := []string{"Azeez", "Lawal", "Shittu"}
	// firstName := []string{"Alhameen", "Taiwo", "Aisha"}
	// fullName := concantenateName(&surName, &firstName)
	// initials := getInitials(&surName, &firstName)
	double := createOperation(2)
	triple := createOperation(3)
	doubledNumbers := numberOperations(&num, double)
	tripledNumbers := numberOperations(&num, triple)
	// operationType1 := getOperationFunction(&num)
	// operationType2 := getOperationFunction(&num2)
	// operationResult1 := numberOperations(&num, operationType1)
	// operationResult2 := numberOperations(&num, operationType2)
	fmt.Println(doubledNumbers)
	fmt.Println(tripledNumbers)
	// fmt.Println(fullName)
	// fmt.Println(initials[0])

}

func numberOperations(nums *[]int, operation operationFn) []int {
	dNumbers := []int{}
	for _, v := range *nums {
		dNumbers = append(dNumbers, operation(v))
	}
	return dNumbers
}

// func getOperationFunction(num *[]int) operationFn {
// 	if (*num)[0] == 1 {
// 		return double
// 	}
// 	return triple

// }
// func concantenateName(surName, firstName *[]string) []string {
// 	fullName := []string{}
// 	minLength := len(*surName)
// 	if len(*firstName) < minLength {
// 		minLength = len(*firstName)
// 	}
// 	for i := 0; i < minLength; i++ {
// 		fn := (*surName)[i] + " " + (*firstName)[i]
// 		fullName = append(fullName, fn)
// 	}
// 	return fullName
// }

// func getInitials(surName, firstName *[]string) []string {
// 	initials := []string{}
// 	minLength := len(*surName)
// 	if len(*firstName) < minLength {
// 		minLength = len(*firstName)
// 	}
// 	for i := 0; i < minLength; i++ {
// 		in := string((*surName)[i][0]) + "." + string((*firstName)[i][0])
// 		initials = append(initials, in)
// 	}
// 	return initials
// }
// func double(num int) int {
// 	return num * 2

// }

// func triple(num int) int {
// 	return num * 3

// }

func createOperation(factor int) operationFn  {
	return func(number int) int {
		return number * factor
	}
}
