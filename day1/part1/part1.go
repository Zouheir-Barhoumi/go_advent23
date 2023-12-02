package part1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

/** Day1
* Calibration document 'input.txt" contains lines of text
* Each line has a specific calibration value to be recovered
* The value is a two-digit number formed by combining the first and the last digits
that appear in each line
*/

// - read the document
// - store the lines as elements in an array
// - read each line and:
// 	* check if the char is a digit
// 	* if it is a digit store in firstDigit
// 	* if it's not move on to the next
// 	* if firstDigit != 0 then next digit encountered is lastDigit
// 	* if lastDigit != 0 override lastDigit with the new digit
// 	* if no other digit encountered then lastDigit := firtDigit

func PartOne() {
	list := make([]string, 0, 1000)

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scn := bufio.NewScanner(file)

	for scn.Scan() { // By default, scanner advances to next the token by newline
		list = append(list, scn.Text())
	}
	total := 0

	for k, line := range list {
		fmt.Printf("key: %v\t value: %v\n", k, line)
		fmt.Printf("\t|\n")

		firstDigit := 0
		lastDigit := 0
		for i, r := range line {
			fmt.Printf("\t\tIndex: %d Character: %c\n", i, r)
			if unicode.IsDigit(r) {
				digit, _ := strconv.Atoi(string(r))
				fmt.Printf("\t\tFound first digit: %d\n", digit)
				lastDigit = digit
				if firstDigit == 0 {
					firstDigit = digit
				}
			}
		}
		total += (firstDigit*10 + lastDigit)

		fmt.Printf("\tTotal: %d\n", total)
	}

}
