package part1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func checkDigit(d byte) bool {
	// unicode.IsDigit takes rune only
	return unicode.IsDigit(rune(d))
}
func Part1() {
	/**
	**Determine which games would have been possible if the bag had been loaded with only 12 red cubes,
	** 13 green cubes, and 14 blue cubes. What is the sum of the IDs of those games?
	Scan input file line by line
	*/

	colors := make(map[string]int)

	colors["red"] = 0
	colors["green"] = 0
	colors["blue"] = 0

	games := make([][]string, 0, 128)

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scn := bufio.NewScanner(file)

	for scn.Scan() {

		scannedLine := scn.Text()
		titleIndex := strings.Index(scannedLine, ":") + 1
		filtered := scannedLine[titleIndex:]

		trimmed := strings.Trim(filtered, " ")

		tokenized := strings.Split(trimmed, ";")

		games = append(games, tokenized)

		// fmt.Println(len(scn.Bytes()))
	}

	sum := 0
	possible := true
	for gId, game := range games {
		fmt.Println(game)

		for setId := 0; setId < len(game) && possible; setId++ {
			set := game[setId]
			fmt.Println(set)
			for i := 0; i+3 < len(set); i++ {
				if checkDigit(set[i]) {
					if checkDigit(set[i+1]) {
						// We have a two-digit number
						doubleDigit := string(set[i]) + string(set[i+1])
						color := set[i+3]
						fmt.Printf("The Color Initial is %c\n", color)

						// INCREMENT COLORs FREQUENCIES
						switch string(color) {
						case "r":
							redDigit, _ := strconv.Atoi(doubleDigit)
							colors["red"] += redDigit
						case "g":
							greenDigit, _ := strconv.Atoi(doubleDigit)
							colors["green"] += greenDigit
						case "b":
							blueDigit, _ := strconv.Atoi(doubleDigit)
							colors["blue"] += blueDigit
						}
						i++
						continue
					} else {
						// we have a one-digit number
						singleDigit := string(set[i])
						color := set[i+2]
						fmt.Printf("The Color Initial is %c\n", color)

						switch string(color) {
						case "r":
							redDigit, _ := strconv.Atoi(singleDigit)
							colors["red"] += redDigit
						case "g":
							greenDigit, _ := strconv.Atoi(singleDigit)
							colors["green"] += greenDigit
						case "b":
							blueDigit, _ := strconv.Atoi(singleDigit)
							colors["blue"] += blueDigit
						}

						// INCREMENT COLORs FREQUENCIES
					}
				}

			}

			fmt.Println("---------------------------------")

			fmt.Printf("RED: %d \t GREEN: %d \t BLUE: %d \t \n", colors["red"], colors["green"], colors["blue"])
			if colors["red"] > 12 || colors["green"] > 13 || colors["blue"] > 14 {
				// fmt.Printf("\n\n\n\n\nThis line %d is POSSIBLE!\n\n\n\n", gId)
				possible = false
			}
			fmt.Printf("----------End Of Set-------------\n\n")
			// reset colors with each set
			colors["red"] = 0
			colors["green"] = 0
			colors["blue"] = 0
		}
		fmt.Printf("RED: %d \t GREEN: %d \t BLUE: %d \n ", colors["red"], colors["green"], colors["blue"])

		colors["red"] = 0
		colors["green"] = 0
		colors["blue"] = 0
		fmt.Println("============================")
		if possible {
			fmt.Printf("\n\n\nThis line %d is possible\n\n\n", gId+1)
			sum += gId + 1

		} else {
			fmt.Printf("\n\n\nThis line %d is IMPOSSIBLE\n\n\n", gId+1)
		}
		fmt.Printf("\n\n\nTotal Possible: %d\n\n\n", sum)
		possible = true
	}

}
