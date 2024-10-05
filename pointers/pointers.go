package main

import "fmt"

func main() {
	age := 32
	adultYear := getAdultYears(&age)
	fmt.Println(adultYear)
}

func getAdultYears (age *int) int{
	return *age - 18

}