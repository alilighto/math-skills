package utils

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func ReadFile(InputFile string) []float64 {

	FloatConverted := []float64{}

	File, err := os.Open(InputFile)
	if err != nil {
		fmt.Println("Error: Incorrect file name !")
		os.Exit(1)
	}
	defer File.Close()

	scanner := bufio.NewScanner(File)
	if !scanner.Scan() {
		fmt.Println("Error: Input file is empty!")
		os.Exit(1)
	}
	for scanner.Scan() {
		Line := scanner.Text()

		if Line == "" {
			continue
		}

		floatNumber, err := strconv.ParseFloat(Line, 64)
		if err != nil {
			fmt.Println("Error: converting to float64!", err)
			os.Exit(1)
		}

		if floatNumber >= math.MaxInt64 {
			fmt.Println("Error: Float value is out of range!")
			os.Exit(1)
		}

		if floatNumber < 0 {
			fmt.Println("Error: Float value should not be a negative value!")
			os.Exit(1)
		}
		FloatConverted = append(FloatConverted, floatNumber)
	}
	if len(FloatConverted) == 1 {
		fmt.Println("Error: Float value should not be a 1\nyou cant do the operations on one number !")
		os.Exit(1)
	}
	return float
}
