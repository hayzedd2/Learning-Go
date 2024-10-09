package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google": "google.com",
		"Amazon": "amazon.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Google"])

	websites["Hayzedd"] = "alhameen.com"
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)

	userNames := make([]string, 2, 5)
	userNames = append(userNames, "Hayzedd")
	userNames = append(userNames, "Tek")
	fmt.Println(userNames)

	courseRatings := make(map[string]float64, 3)

	courseRatings["go"] = 4.7
	fmt.Println(courseRatings)

	for i, v := range userNames {
		fmt.Printf("Index is %v and value is %v \n", i, v)
	}
	for k, v := range courseRatings {
		fmt.Printf("Key is %v and value is %v \n", k, v)
	}
}
