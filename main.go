package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("./day1Data/day1.csv")

	if err != nil {
		fmt.Println("ERROR", err)
		return
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Err", err)
		return
	}

	var numbers1 []float64
	var numbers2 []float64

	for _, record := range records {
		number1, errParse1 := strconv.ParseFloat(record[0], 64)

		if errParse1 != nil {
			return
		}

		numbers1 = append(numbers1, number1)

		number2, errParse2 := strconv.ParseFloat(record[1], 64)

		if errParse2 != nil {
			return
		}

		numbers2 = append(numbers2, number2)

	}

	sort.Float64s(numbers1)
	sort.Float64s(numbers2)

	var sum float64 = 0

	//First star
	//for index := range numbers1 {
	//	sum += math.Abs(numbers1[index] - numbers2[index])
	//}

	//Second star
	for _, number1 := range numbers1 {
		var occcurrences float64 = 0

		for _, number2 := range numbers2 {
			if number1 == number2 {
				occcurrences++
			}
		}

		sum += number1 * occcurrences
	}

	fmt.Printf("Hello world: %f\n", sum)
}
