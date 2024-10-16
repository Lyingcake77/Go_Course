package fileManagement

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func SaveFloatToFile(fileName string, newFloat float64) {
	var balanceString string = fmt.Sprintf("%v", newFloat)
	os.WriteFile(fileName, []byte(balanceString), 0644)
}

func GetFloatFromFile(fileName string) (SavedBalance float64, Error error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 0, errors.New("failed to read file")
	}
	SavedBalance, err = strconv.ParseFloat(string(data), 64)
	if err != nil {
		return 0, errors.New("failed to parse number to float")
	}
	return SavedBalance, nil
}
