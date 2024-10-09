package main

import "fmt"

type marketList struct {
	title string
	id    int
	price float64
}

func main() {
	hobbies := [3]string{"games", "Coding", "gisting"}
	restHobbies := hobbies[1:3]
	initialHobbies := hobbies[:2]
	initialHobbies = initialHobbies[1:3]

	fmt.Println(hobbies)
	fmt.Println(hobbies[0])
	fmt.Println(restHobbies)
	fmt.Println(initialHobbies)

	courseGoals := []string{"Become a pro", "Get jobs"}
	courseGoals[1] = "Get foreign jobs"
	courseGoals = append(courseGoals, "Proceed on django")
	fmt.Println(courseGoals)

	products := []marketList{
		{
			"Fan",
			1,
			10.88,
		}, {
			"Car",
			2,
			20.99,
		},
	}
	newProduct := marketList{
		"Bike",
		3,
		40.99,
	}
	prices := []float64{1.99, 2.99}
	discountPrices := []float64{10.99, 8.99}
	prices = append(prices, discountPrices...)
	products = append(products, newProduct)
	fmt.Println(products)
	fmt.Println(prices)
}
