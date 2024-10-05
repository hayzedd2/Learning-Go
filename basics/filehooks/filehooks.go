package filehooks

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	// "math"
)


func SaveFloatToFile(float float64, fileName string) {
	floatText := fmt.Sprint(float)
	os.WriteFile(fileName, []byte(floatText), 0644)
}

func ReadFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil{
		return 1000, errors.New("failed to read file")
	}
	floatText := string(data)
	float, err := strconv.ParseFloat(floatText, 64)
	if err != nil{
		return 1000, errors.New("failed to parse float")
	}
	return float, nil
}
