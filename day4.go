package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./day4Data/day4.csv")

	if err != nil {
		fmt.Println("ERROR", err)
		return
	}

	defer file.Close()

	reader := csv.NewReader(file)

	matches, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Err", err)
		return
	}

	//matches := strings.Split(input, "\n")

	var matrix [][]string

	for _, match := range matches {
		var row []string
		for _, character := range match[0] {
			row = append(row, string(character))
		}
		matrix = append(matrix, row)
	}

	var sum = 0

	for rowIndex, row := range matrix {
		for columnIndex, character := range row {
			//part one
			//if character == "X" {
			//	var xCoords = []int{rowIndex, columnIndex}
			//	sum += checkForXmas(matrix, xCoords)
			//}

			//part two

			if character == "A" {
				var aCoords = []int{rowIndex, columnIndex}
				sum += checkForMas(matrix, aCoords)
			}
		}
	}

	fmt.Println(sum)

}

func checkForXmas(matrix [][]string, xCoords []int) int {
	var directions = [][][]int{
		{{0, 1}, {0, 2}, {0, 3}},       //move right
		{{0, -1}, {0, -2}, {0, -3}},    //move left
		{{1, 0}, {2, 0}, {3, 0}},       //move down,
		{{-1, 0}, {-2, 0}, {-3, 0}},    //move up,
		{{1, -1}, {2, -2}, {3, -3}},    //move up right,
		{{1, 1}, {2, 2}, {3, 3}},       //move down right,
		{{-1, -1}, {-2, -2}, {-3, -3}}, //move up left,
		{{-1, 1}, {-2, 2}, {-3, 3}},    //move down left,
	}

	var sum = 0

	for _, moveSet := range directions {
		var iLetterMCoords = xCoords[0] + moveSet[0][0]
		var jLetterMCoords = xCoords[1] + moveSet[0][1]

		var mCoords = []int{iLetterMCoords, jLetterMCoords}
		var _, errM = checkCoordinateBoundaries(matrix, xCoords[0], mCoords)

		if errM != nil {
			continue
		}

		if matrix[iLetterMCoords][jLetterMCoords] != "M" {
			continue
		}

		var iLetterACoords = xCoords[0] + moveSet[1][0]
		var jLetterACoords = xCoords[1] + moveSet[1][1]

		var aCoords = []int{iLetterACoords, jLetterACoords}
		var _, errA = checkCoordinateBoundaries(matrix, xCoords[0], aCoords)

		if errA != nil {
			continue
		}

		if matrix[iLetterACoords][jLetterACoords] != "A" {
			continue
		}

		var iLetterSCoords = xCoords[0] + moveSet[2][0]
		var jLetterSCoords = xCoords[1] + moveSet[2][1]

		var sCoords = []int{iLetterSCoords, jLetterSCoords}
		var _, errS = checkCoordinateBoundaries(matrix, xCoords[0], sCoords)

		if errS != nil {
			continue
		}

		if matrix[iLetterSCoords][jLetterSCoords] != "S" {
			continue
		}

		sum++
	}

	return sum
}

func checkForMas(matrix [][]string, aCoords []int) int {
	var moveSet = [][]int{
		{-1, 1},  //move up right,
		{1, -1},  //move down left,
		{1, 1},   //move down right,
		{-1, -1}, //move up left,
	}

	var sum = 0

	var iFirstLetterCoords = aCoords[0] + moveSet[0][0]
	var jFirstLetterCoords = aCoords[1] + moveSet[0][1]

	var firstLetterCoords = []int{iFirstLetterCoords, jFirstLetterCoords}
	var _, errFirst = checkCoordinateBoundaries(matrix, aCoords[0], firstLetterCoords)

	fmt.Println(firstLetterCoords, aCoords)

	if errFirst != nil {
		return 0
	}

	var iSecondLetterCoords = aCoords[0] + moveSet[1][0]
	var jSecondLetterCoords = aCoords[1] + moveSet[1][1]

	var secondLetterCoords = []int{iSecondLetterCoords, jSecondLetterCoords}
	var _, errSecond = checkCoordinateBoundaries(matrix, aCoords[0], secondLetterCoords)

	fmt.Println(secondLetterCoords)

	if errSecond != nil {
		return 0
	}

	var iThirdLetterCoords = aCoords[0] + moveSet[2][0]
	var jThirdLetterCoords = aCoords[1] + moveSet[2][1]

	var thirdLetterCoords = []int{iThirdLetterCoords, jThirdLetterCoords}
	var _, errThird = checkCoordinateBoundaries(matrix, aCoords[0], thirdLetterCoords)

	fmt.Println(thirdLetterCoords)

	if errThird != nil {
		return 0
	}

	var iFourthLetterCoords = aCoords[0] + moveSet[3][0]
	var jFourthLetterCoords = aCoords[1] + moveSet[3][1]

	var fourthLetterCoords = []int{iFourthLetterCoords, jFourthLetterCoords}
	var _, errFourth = checkCoordinateBoundaries(matrix, aCoords[0], fourthLetterCoords)
	fmt.Println(fourthLetterCoords)

	if errFourth != nil {
		return 0
	}

	var firstWord = fmt.Sprintf("%sA%s", matrix[iFirstLetterCoords][jFirstLetterCoords], matrix[iSecondLetterCoords][jSecondLetterCoords])
	var secondWord = fmt.Sprintf("%sA%s", matrix[iThirdLetterCoords][jThirdLetterCoords], matrix[iFourthLetterCoords][jFourthLetterCoords])

	fmt.Println(firstWord, secondWord)
	if (firstWord == "MAS" || reverseString(firstWord) == "MAS") && (secondWord == "MAS" || reverseString(secondWord) == "MAS") {
		fmt.Println("ITWASAHIT")
		sum++
	}

	return sum
}

func checkCoordinateBoundaries(matrix [][]string, currentRow int, coords []int) (bool, error) {
	if coords[0] < 0 || coords[0] >= len(matrix) {
		return false, fmt.Errorf("Out of bounds i")
	}

	if coords[1] < 0 || coords[1] >= len(matrix[currentRow]) {
		return false, fmt.Errorf("Out of bounds j")
	}

	return true, nil
}

func reverseString(input string) string {
	// Convert the string to a slice of runes to handle multi-byte characters
	runes := []rune(input)

	// Reverse the slice of runes
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Convert the slice of runes back to a string
	return string(runes)
}
