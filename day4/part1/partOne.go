package part1

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func getNumberParts(part string, scnLineId int, numbers *[][]int) {
	parseDigit := ""
	for i, c := range part {
		if unicode.IsDigit(c) {
			parseDigit += string(c)
		}
		if !unicode.IsDigit(c) || i == len(part)-1 {
			// the number ends if we face a non-digit if a digit already stored OR at end of string
			// store and reset
			fullNumber, err := strconv.Atoi(parseDigit)
			if err == nil && parseDigit != "" {
				// Expand wining numbers to accomodate for new lines
				for scnLineId >= len(*numbers) {
					*numbers = append(*numbers, []int{})
				}
				(*numbers)[scnLineId] = append((*numbers)[scnLineId], fullNumber)
			}
			// reset the number
			parseDigit = ""
		}
	}
}

type WNumbers struct {
	number int
	line   int
}

func Part1() {
	// Split into lines
	// Store all in an slice of type string without the part before :
	// For each line []
	// Multi-deimensional slice of winning numbers [LINE NUBMER][First String]
	// Multi-deimensional slice of owned numbers [LINE NUBMER][Second String]
	// Compare the two with nest for loops i, j
	// Winning cards are the matches
	// If the match is 1: points = 1 for this line
	// If the match is > 1: points = n * 2 for this line
	// Store points for each line in a map[LINE NUMBER]POINTS
	// Loop through the map and sum them all up

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var winningPart [][]int
	dealtPart := make([][]int, 0)
	foundWinning := make([]WNumbers, 0)

	scanner := bufio.NewScanner(file)

	scnLineId := -1
	for scanner.Scan() {
		scnLineId++
		line := scanner.Text()

		semiId := strings.Index(line, ":") + 1

		filteredLine := line[semiId:]

		trimmedLine := strings.Trim(filteredLine, " ")

		fmt.Println(trimmedLine)

		pipeId := strings.Index(trimmedLine, "|")
		firstPart := trimmedLine[:pipeId]
		secondPart := trimmedLine[pipeId+2:] // +2 for the space and the |

		fmt.Println(firstPart)
		fmt.Println(secondPart)

		// winning numbers
		getNumberParts(firstPart, scnLineId, &winningPart)
		// numbers we have
		getNumberParts(secondPart, scnLineId, &dealtPart)

	}
	fmt.Println(winningPart)
	fmt.Println(dealtPart)
	fmt.Println("-------------------------")
	fmt.Println("-------------------------")
	fmt.Println("--------comparison-------")
	fmt.Println("-------------------------")

	for i, wL := range winningPart {
		fmt.Printf("%d - Winning Line %d\n", i, wL)
		for j, wN := range wL {
			fmt.Printf("%d \t %d\n", j, wN)
			for _, dN := range dealtPart[i] {
				if wN == dN {
					fmt.Printf("Found: %d ", dN)
					foundWinning = append(foundWinning, WNumbers{dN, i})
				}
			}
			fmt.Printf("\n")
		}
	}

	fmt.Println("ALL FOUND WINNING")
	for _, fW := range foundWinning {
		fmt.Println(fW)
	}

	lineSum := 0
	totalSum := 0
	// number of lines scanned is scnLineId
	for i := 0; i <= scnLineId; i++ {
		fmt.Printf("\n")
		fmt.Printf("%d - ", i)
		for _, fW := range foundWinning {
			if fW.line == i {
				fmt.Printf("%d ", fW)
				lineSum += 1
				fmt.Println("Line sum:", lineSum)
			}
		}
		if lineSum == 1 {
			totalSum += 1
		}
		if lineSum > 1 {
			power := lineSum - 1
			totalSum += int(math.Pow(2, float64(power)))
		}

		lineSum = 0
		fmt.Println("---")
	}
	fmt.Println("TOTAL SUM:", totalSum)
}
