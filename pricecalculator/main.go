package main

import (
	"example.com/investmentCalculator/pricecalculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		priceJob := prices.NewPriceTaxProcess(taxRate)
		priceJob.Process()
	}

}
