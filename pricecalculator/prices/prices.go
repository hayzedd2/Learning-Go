package prices

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type PriceTaxProcess struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *PriceTaxProcess) LoadData() {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	handleFileErr(err, file)
	covertedPrice := make([]float64, len(lines))
	for i, line := range lines {
		priceFloat, err := strconv.ParseFloat(line, 64)
		handleFileErr(err, file)
		covertedPrice[i] = priceFloat
	}
	job.InputPrices = covertedPrice
}

func (job *PriceTaxProcess) Process() {
	job.LoadData()
	result := make(map[string]float64)
	for _, price := range job.InputPrices {
		taxIncludedPrices := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%0.2f", price)] = math.Round(taxIncludedPrices)
	}
	fmt.Println(result)
}

func NewPriceTaxProcess(taxRate float64) *PriceTaxProcess {
	return &PriceTaxProcess{
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}

func handleFileErr(err error, file *os.File) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		file.Close()
		return
	}
}
