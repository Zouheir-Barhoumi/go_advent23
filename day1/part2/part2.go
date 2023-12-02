package part2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var newString string

var spelledNumbers = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func PartTwo() {
	list := make([]string, 0, 1000)

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scn := bufio.NewScanner(file)

	for scn.Scan() {
		if err := scn.Err(); err != nil {
			log.Fatal(err)
		}
		list = append(list, scn.Text())
	}
	total := 0

	for k, line := range list {
		fmt.Printf("key: %v\t value: %v\n", k, line)
		fmt.Printf("\t|\n")

		firstDigit := 0
		lastDigit := 0
		for _, r := range line {
			fmt.Printf("\t\t %c\n", r)
			if unicode.IsDigit(r) {
				newString = ""
				digit, _ := strconv.Atoi(string(r))
				fmt.Printf("\t\tFound digit: %d\n", digit)
				lastDigit = digit
				if firstDigit == 0 {
					firstDigit = digit
				}
			} else {
				newString += string(r)

				if index := stringExists(newString); index != -1 {
					lastDigit = index + 1
					fmt.Printf("\t\tFound spelled digit: %d whose index is %d\n", lastDigit, index)
					if firstDigit == 0 {
						firstDigit = index + 1
					}
					newString = ""
				}
			}
			fmt.Println("\t\tThis line's total:", firstDigit, lastDigit)
		}
		total += (firstDigit*10 + lastDigit)

		fmt.Printf("\tTotal: %d\n", total)
	}
}

func stringExists(target string) int {
	for i, v := range spelledNumbers {
		// checks for partial matches while incrementally called on target
		if strings.Contains(target, v) {
			return i
		}
	}
	return -1
}
