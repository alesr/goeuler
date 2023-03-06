package exercises

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// Problem 8

// Largest product in a series
// The four adjacent digits in the 1000-digit number that have the greatest product are 9 × 9 × 8 × 9 = 5832.

// Find the thirteen adjacent digits in the 1000-digit number that have the greatest product. What is the value of this product?
func largestProductInSeries(inputPath string, adjacents int) int {
	f, err := os.Open(inputPath)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer f.Close()

	var data []int

	fileScanner := bufio.NewScanner(f)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		for i := 0; i < len(line); i++ {
			intValue, err := strconv.Atoi(string(line[i]))
			if err != nil {
				log.Fatalf("failed converting byte to int: %s", err)
			}
			data = append(data, intValue)
		}
	}

	var maxProduct int

	for i := 0; i < len(data)-adjacents; i++ {
		var product int = 1

		for j := 0; j < adjacents; j++ {
			product *= data[i+j]

			if product == 0 {
				break
			}

			if product > maxProduct {
				maxProduct = product
			}
		}
	}
	return maxProduct
}
