package part1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

type NumerInfo struct {
	number int
	prevId int
	nextId int
	lineN  int
}

type SymbolInfo struct {
	symbol string
	index  int
	lineN  int
}

func scanData(file *os.File) *bufio.Scanner {
	return bufio.NewScanner(file)
}

func checkDigit(c rune) bool {
	return unicode.IsDigit(rune(c))
}

func existInSlice(slice []NumerInfo, n NumerInfo) int {
	for i, v := range slice {
		if n == v {
			return i
		}
	}
	return -1
}

func Part1() {
	numbersFound := make([]NumerInfo, 0)
	symbolsFound := make([]SymbolInfo, 0)
	file, err := os.OpenFile("input.txt", os.O_RDONLY, 0555)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scn := scanData(file)

	lineId := 0

	for scn.Scan() {
		// fmt.Println(len(scn.Bytes())) // 140
		lineId++

		line := scn.Text()
		fmt.Println(line)
		fmt.Println(lineId)

		nString := ""
		nId := 0
		for i, c := range line {
			// Check for digits
			if checkDigit(c) {
				nString += string(c)
				nId = i
			}
			if !checkDigit(c) || i == len(line)-1 {
				if nString != "" {
					// number ended: store number and reset
					nInt, _ := strconv.Atoi(nString)
					numbersFound = append(numbersFound, NumerInfo{nInt, nId - len(nString), i, lineId})
					// reset
					nString = ""
					nId = 0
				}
				// Check for symbols
				if string(c) != "." && string(c) != "0" {
					// it's a symbol
					// sId := i + 1

					symbolsFound = append(symbolsFound, SymbolInfo{string(c), i, lineId})
				}
			}
		}
	}

	sameLineMatches := make([]int, 0)
	nextLineMatches := make([]int, 0)
	previousLineMatches := make([]int, 0)
	shortList := make([]NumerInfo, 0)
	sum := 0
	// found := false
	for _, number := range numbersFound {
		fmt.Printf("Number found: %d \t Index: %d-%d \t Line: %d\n", number.number, number.prevId, number.nextId, number.lineN)
		// same line
		for _, symbol := range symbolsFound {
			if number.lineN == symbol.lineN {
				// fmt.Printf("Symbol found: %s \t Line: %d\n", symbol.symbol, symbol.lineN)
				fmt.Printf("Symbol %s index: %d\t\n", symbol.symbol, symbol.index)
				if symbol.index >= number.prevId-1 && symbol.index <= number.nextId {
					fmt.Printf("%s Matches %d\n", symbol.symbol, number.number)
					sameLineMatches = append(sameLineMatches, number.number)
					if existInSlice(shortList, number) == -1 {
						sum += number.number
						shortList = append(shortList, number)
					}
					// found = true
				}
			}
		}
		// if found {
		// 	continue
		// }
		// next line
		for _, symbol := range symbolsFound {
			if number.lineN+1 == symbol.lineN {
				// fmt.Printf("Symbol found: %s \t Line: %d\n", symbol.symbol, symbol.lineN)
				fmt.Printf("Symbol %s index: %d\t\n", symbol.symbol, symbol.index)
				if symbol.index >= number.prevId-1 && symbol.index <= number.nextId+1 {
					fmt.Printf("%s Matches %d\n", symbol.symbol, number.number)
					nextLineMatches = append(nextLineMatches, number.number)
					if existInSlice(shortList, number) == -1 {
						sum += number.number
						shortList = append(shortList, number)
					}
					// found = true
				}
			}
		}
		// if found {
		// 	continue
		// }
		// previous line
		for _, symbol := range symbolsFound {
			if number.lineN-1 == symbol.lineN {
				// fmt.Printf("Symbol found: %s \t Line: %d\n", symbol.symbol, symbol.lineN)
				fmt.Printf("Symbol %s index: %d\t\n", symbol.symbol, symbol.index)
				if symbol.index >= number.prevId && symbol.index-1 <= number.nextId+1 {
					fmt.Printf("%s Matches %d\n", symbol.symbol, number.number)
					previousLineMatches = append(previousLineMatches, number.number)
					if existInSlice(shortList, number) == -1 {
						sum += number.number
						shortList = append(shortList, number)
					}
					// found = true
				}
			}
		}
	}
	fmt.Printf("Same Line Matches: \n")
	for _, v := range sameLineMatches {
		fmt.Printf("%d\n", v)
	}
	fmt.Printf("Next Line Matches: \n")
	for _, v := range nextLineMatches {
		fmt.Printf("%d\n", v)
	}
	fmt.Printf("Previous Line Matches: \n")
	for _, v := range previousLineMatches {
		fmt.Printf("%d\n", v)
	}
	fmt.Printf("ShortList: \n")
	for _, v := range shortList {
		fmt.Printf("%d\n", v.number)
	}
	fmt.Println("TOTAL: ", sum)
	// prevSymbols := make([]SymbolInfo, 0)
	// nextSymbols := make([]SymbolInfo, 0)
	// sum := 0

	// shortList := make([]int, 0)

	// for nId, n := range numbersFound {
	// 	fmt.Println("Iteration n:", nId)
	// 	fmt.Println("-----------------------------------------------")
	// 	for _, s := range symbolsFound {
	// 		// store symbols of last line as prevSymbols
	// 		if s.lineN == nId {
	// 			prevSymbols = append(prevSymbols, s)
	// 		}
	// 		// check for adjacent symbols on the same line
	// 		if s.lineN == n.lineN {
	// 			// before or after n
	// 			if s.index == n.prevId || s.index == n.nextId {
	// 				fmt.Printf("FOUND PART: %d\n", n.number)
	// 				if existInSlice(shortList, n.number) == -1 {
	// 					sum += n.number
	// 					shortList = append(shortList, n.number)
	// 				}
	// 			}
	// 		}
	// 	}
	// 	for _, s := range symbolsFound {
	// 		// store symbols of last line as nextSymbols
	// 		if s.lineN == nId+2 {
	// 			fmt.Printf("%d YES", nId+1)
	// 			nextSymbols = append(nextSymbols, s)
	// 		}

	// 	}

	// 	fmt.Printf("\nSymbols from last line %v\n", prevSymbols)
	// 	fmt.Printf("\nSymbols from next line %v\n", nextSymbols)

	// 	// check of adjacent symbols on the previous line
	// 	if n.lineN > 1 {
	// 		for _, s := range prevSymbols {
	// 			if s.index >= n.prevId && s.index <= n.nextId {
	// 				fmt.Printf("PREV PART FOUND: %d\n", n.number)
	// 				if existInSlice(shortList, n.number) == -1 {
	// 					sum += n.number
	// 					shortList = append(shortList, n.number)
	// 				}
	// 				break
	// 			}
	// 		}
	// 		fmt.Printf("\nNumber %d - %d\n", n.prevId, n.nextId)
	// 	}

	// 	// check of adjacent symbols on the next line
	// 	if nId != len(numbersFound)-1 {
	// 		for _, s := range nextSymbols {
	// 			if s.index >= n.prevId-1 && s.index <= n.nextId {
	// 				fmt.Printf("NEXT PART FOUND: %d\n", n.number)
	// 				if existInSlice(shortList, n.number) == -1 {
	// 					sum += n.number
	// 					shortList = append(shortList, n.number)
	// 				}
	// 				break
	// 			}
	// 		}
	// 		fmt.Printf("\nNumber %d - %d\n", n.prevId, n.nextId)
	// 	}

	// 	// reset prevSymbols on each line
	// 	if n.lineN != nId {
	// 		prevSymbols = prevSymbols[:0]
	// 	}
	// 	// reset nextSymbols on each line
	// 	if n.lineN != nId+1 {
	// 		nextSymbols = nextSymbols[:0]
	// 	}
	// 	fmt.Println("Total sum:", sum)
	// }

	// for _, v := range shortList {
	// 	fmt.Println(v)
	// }
}
