package main

import (
	"fmt"
	"example.com/investmentCalculator/basics/filehooks"
	"github.com/Pallinder/go-randomdata"

	// "math"
)

// const inflationRate = 5.5
const accountBalanceFile = "balance.txt"
func main() {
	var accountBalance, err = filehooks.ReadFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Printf("Error: %v \n", err)
		fmt.Println("--------")
		// panic(err)
	}
	// var investmentAmount, expectedReturnRate, years float64

	// outputText("Investment Amount:")
	// fmt.Scan(&investmentAmount)

	// outputText("Expected Return Rate: ")
	// fmt.Scan(&expectedReturnRate)

	// outputText("Years: ")
	// fmt.Scan(&years)
	// futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	// fmt.Printf("Your futureValue is %.2f and your future real value is %.2f", futureValue, futureRealValue)
	fmt.Println("Welcome to Z bank!")
	fmt.Println("Reach us @", randomdata.Email())
	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1- to check balance")
		fmt.Println("2- to deposit")
		fmt.Println("3- to withdraw")
		fmt.Println("4- to exit ")
		fmt.Print("Your choice:")
		var userInput int
		fmt.Scan(&userInput)
		switch userInput {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Deposit amount: ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0!")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Balance updated! New account balance is ", accountBalance)
			filehooks.SaveFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			var withdrawalAmount float64
			fmt.Print("withdrawal amount: ")
			fmt.Scan(&withdrawalAmount)
			if withdrawalAmount < 0 {
				fmt.Println("Invalid amount. Must be greater than 0!")
				continue
			}
			if withdrawalAmount > accountBalance {
				fmt.Println("Insufficient funds!")
				continue
			}
			accountBalance -= withdrawalAmount
			fmt.Println("Balance updated! New account balance is", accountBalance)
			filehooks.SaveFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Exiting the program.")
			fmt.Println("Thanks for banking with us")
			return
		}

	}

}


// func outputText(text string) {
// 	fmt.Print(text)
// }

// func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (calculatedFutureValue float64, calculatedFutureRealValue float64) {
// 	calculatedFutureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
// 	calculatedFutureRealValue = calculatedFutureValue / math.Pow(1+inflationRate, years)
// 	return calculatedFutureValue, calculatedFutureRealValue

// }
