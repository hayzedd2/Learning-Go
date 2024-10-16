package main

import (
	"fmt"

	"example.com/investmentCalculator/pricecalculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	chanStatuses := make([]chan bool, len(taxRates))
	errStatuses := make([]chan error, len(taxRates))
	for i, taxRate := range taxRates {
		chanStatuses[i] = make(chan bool)
		errStatuses[i] = make(chan error)
		priceJob := prices.NewPriceTaxProcess(taxRate)
		go priceJob.Process(chanStatuses[i], errStatuses[i])
	}
	for i := range taxRates {
		select {
		case err := <-errStatuses[i]:
			if err != nil {
				fmt.Println(err)
			}
		case <-chanStatuses[i]:
			fmt.Println("Suceeded!")
		}

	}

}
