package part2

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

type NumberDetail struct {
	number int
	line   int
}
type Card struct {
	numbers []NumberDetail
	line    int
	// copyCards *[]Card
}

type OgCard struct {
	Card
	cardsMatched []int
}

func detailNumbers(numbers []int, line int) []NumberDetail {
	tempSlice := []NumberDetail{}
	// add details to each number
	if len(numbers) != 0 {
		for _, n := range numbers {
			tempSlice = append(tempSlice, NumberDetail{n, line})
		}
	}
	return tempSlice
}

func Part2() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var winningPart [][]int
	dealtPart := make([][]int, 0)
	foundWinning := make([]NumberDetail, 0)
	// copyCards := make([]Card, 0)
	cards := make([]Card, 0)
	ogCards := make([]OgCard, 0)

	// scratchCards := make([]SCard, 0)

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
					foundWinning = append(foundWinning, NumberDetail{dN, i})
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

	var lineSums []int

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

		lineSums = append(lineSums, lineSum)
		lineSum = 0
		fmt.Println("---")
	}
	fmt.Println("TOTAL SUM:", totalSum)

	// tempSlice := make([]int, 0)
	// create a card out of each line
	for i, line := range dealtPart {
		cards = append(cards, Card{detailNumbers(line, i), i})
	}
	// create a card out of each line

	for j, s := range lineSums {
		fmt.Printf("LINE SUM %d: %d\n", j, s)
		nOfMatchesAtI := s
		fmt.Printf("Number of Matches at %d is %d, \n", j, nOfMatchesAtI)
		if nOfMatchesAtI > 0 {
			for cardId, card := range cards {
				if cardId == j {
					fmt.Println("We have a match: ", cardId, card)
					fmt.Printf("YOU WIN CARDS:\n")
					for x := cardId + 1; x < cardId+1+nOfMatchesAtI; x++ {
						fmt.Printf("%d\t", x)
						// tempSlice = append(tempSlice, x)
					}
					fmt.Println("")
					// ogCards = append(ogCards, OgCard{card, tempSlice})
				}
			}
			continue
		}

		// for i, card := range cards {
		// 	ogCards = append(ogCards, OgCard{card, *new([]int)})
		// 	fmt.Printf("Appended OgCard: %v\n", ogCards[i])
		// }

	}

	totalScratchCards := make(map[int]int, scnLineId+1)
	fmt.Printf("\n\n\n\n\n")
	fmt.Printf("\n\n OUR MAP IS READY\n\n\n")

	for i := 1; i <= scnLineId+1; i++ {
		totalScratchCards[i] = 0
		fmt.Println(totalScratchCards[i])
	}

	fmt.Printf("\n\n\n\n\n")
	fmt.Printf("\n|=============================\n")
	fmt.Printf("\n|===========LineSums==========\n")
	fmt.Printf("\n|=============================\n")
	allMatchedCards := make([][]int, 0)
	if len(lineSums) > 0 {
		for i := 0; i < len(lineSums); i++ {
			nOfMatchesAtI := lineSums[i]
			fmt.Printf("Number of Matches at %d is %d, \n", i, nOfMatchesAtI)
			if nOfMatchesAtI > 0 {
				for cardId, card := range cards {
					if cardId == i {
						fmt.Println("We have a match: ", cardId, card)
						fmt.Printf("YOU WIN CARDS:\n")
						for x := cardId + 1; x < cardId+1+nOfMatchesAtI; x++ {
							fmt.Printf("%d\t", x)
							// totalScratchCards += x
							if value, ok := totalScratchCards[x]; ok {
								totalScratchCards[x] = value + 1
							}
							for len(lineSums) > len(allMatchedCards) {
								allMatchedCards = append(allMatchedCards, []int{})
							}
							allMatchedCards[i] = append(allMatchedCards[i], x)
						}
						fmt.Println("")
					}
				}

				continue
			}
		}
	}

	fmt.Printf("\n************\nCopy CARDS ARRAY\n************\n")

	for _, v := range cards {
		fmt.Println(v)
	}
	fmt.Printf("\n************\nOG CARDS ARRAY\n************\n")

	for _, v := range ogCards {
		fmt.Println(v)
	}
	fmt.Printf("\n************\nALL MATCHED CARDS\n************\n")

	for _, v := range allMatchedCards {
		fmt.Println(v)
	}

	var cout int

	// Double each sub-slice based on the corresponding number from the map
	for i, subSlice := range allMatchedCards {
		if count, ok := totalScratchCards[i+1]; ok && count > 0 {
			// Double the sub-slice count count
			for j := 1; j < count; j++ {
				allMatchedCards[i] = append(allMatchedCards[i], subSlice...)
			}
		}
	}
	println(cout)

	println("Updated 2D Slice:")
	for _, subSlice := range allMatchedCards {
		fmt.Println(subSlice)
	}

	for _, v := range allMatchedCards {
		for _, q := range v {
			q += 1
			cout++
			if val, ok := totalScratchCards[q]; ok {
				fmt.Printf("VAL %d, q %d \n", val, q)
			}

		}
	}

	println(cout)

	fmt.Println("NUMBER OF LINES: ", scnLineId+1)
	accum := 0
	for _, v := range totalScratchCards {
		accum += v
	}
	fmt.Println("Total ScratchCards: ", totalScratchCards)
	fmt.Println("Total ScratchCards: ", accum)
}
