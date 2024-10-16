package prices

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

type PriceTaxProcess struct {
	TaxRate           float64            `json:"tax_rate"`
	InputPrices       []float64          `json:"input_prices"`
	TaxIncludedPrices map[string]float64 `json:"tax_included_prices"`
}

func (job *PriceTaxProcess) LoadData() error {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
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
	return nil
}

func writeJson(path string, data interface{}) error {
	file, err := os.Create(path)
	if err != nil {
		return errors.New("failed to create file")
	}
	defer file.Close()
	time.Sleep(3 * time.Second)
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return errors.New("failed to convert data to jSON")
	}
	return nil
}
func (job *PriceTaxProcess) Process(chanStatus chan bool, errStatus chan error) {
	err := job.LoadData()
	if err != nil {
		errStatus <- err
		return
	}
	result := make(map[string]float64)
	for _, price := range job.InputPrices {
		taxIncludedPrices := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%0.2f", price)] = math.Round(taxIncludedPrices)
	}
	job.TaxIncludedPrices = result
	writeJson(fmt.Sprintf("result_%.0f.json", job.TaxRate*100), job)
	chanStatus <- true
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
