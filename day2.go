package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./day2Data/day2.csv")

	if err != nil {
		fmt.Println("ERROR", err)
		return
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Err gete", err)
		return
	}

	var matrix [][]int

	for _, record := range records {
		var row []int

		for _, number := range record {
			num, errParse1 := strconv.Atoi(number)

			if errParse1 != nil {
				fmt.Println("Err", errParse1)
				return
			}
			row = append(row, num)
		}
		matrix = append(matrix, row)
	}

	var unSafeCounter int = 0
	var safeCounter int = 0

	for _, row := range matrix {
		var decreasing bool = row[0]-row[1] > 0

		var index, err = checkRow(row, decreasing)

		if err == nil {
			safeCounter++
			continue
		}
		fmt.Println("INDEX", index)

		var newRow = make([]int, len(row))
		copy(newRow, row)
		newRow = remove(newRow, index)
		fmt.Println("ROW BEFORE REMOVAL", row)

		decreasing = newRow[0]-newRow[1] > 0

		var _, err2 = checkRow(newRow, decreasing)

		fmt.Println("SECOND ROW AFTER REMOVAL: ", newRow)

		if err2 == nil {
			safeCounter++
			continue
		}

		fmt.Println("ROW BEFORE REMOVAL", row)
		var newRow2 = make([]int, len(row))
		copy(newRow2, row)
		newRow2 = remove(newRow2, index+1)
		fmt.Println("THIRD ATTEMPT ROW AFTER REMOVAL", newRow2)

		decreasing = newRow2[0]-newRow2[1] > 0

		var _, err3 = checkRow(newRow2, decreasing)

		if err3 == nil {
			safeCounter++
		}
	}

	fmt.Println(len(matrix), safeCounter, unSafeCounter)
}

func checkRow(row []int, decreasing bool) (int, error) {
	var counter int = 0
	for i := 0; i < len(row)-1; i++ {
		var distance int

		distance = row[i] - row[i+1]
		if distance < 0 {
			if math.Abs(float64(distance)) > 3 || decreasing {
				counter++
				fmt.Println("FAILED INCREASING", row)
				return i, fmt.Errorf("Error increasing")
			}
		} else if distance > 0 {
			if math.Abs(float64(distance)) > 3 || !decreasing {
				counter++
				fmt.Println("FAILED DECREASING", row)
				return i, fmt.Errorf("Error decresing")
			}
		} else {
			fmt.Println("FAILED EVEN", row)
			counter++
			return i, fmt.Errorf("Error even")
		}

	}

	return -1, nil
}

func remove(slice []int, index int) []int {
	// Check if the index is within bounds
	if index < 0 || index >= len(slice) {
		fmt.Println("Index out of range")
		return slice
	}
	// Remove the element at the specified index
	return append(slice[:index], slice[index+1:]...)
}
