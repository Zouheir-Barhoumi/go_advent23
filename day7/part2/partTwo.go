package part2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var time = 0
var distance = 0

type Race struct {
	time           int
	recordDistance int
}

func PartTwo() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// First line is time and Second is Distance
	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()

		// Remove the 'Time:' and 'Distance' parts
		semi := strings.Index(line, ":")
		nOsemi := line[semi+1:]

		// Trim initial and trailing whitespace
		noInitSpace := strings.TrimLeft(nOsemi, " ")
		line = strings.TrimRight(noInitSpace, " ")

		// Get the time
		if lineNumber == 0 {
			stringSlice := strings.Fields(line)
			newString := ""
			for _, field := range stringSlice {
				trimmedField := strings.TrimSpace(field)
				newString += trimmedField
			}
			newInt, err := strconv.Atoi(newString)
			if err != nil {
				fmt.Println("Failed to convert string to int:")
				log.Fatal(err)
			}
			time = newInt
		}
		// Get the distance
		if lineNumber == 1 {
			stringSlice := strings.Fields(line)
			newString := ""
			for _, field := range stringSlice {
				trimmedField := strings.TrimSpace(field)
				newString += trimmedField
			}
			newInt, err := strconv.Atoi(newString)
			if err != nil {
				fmt.Println("Failed to convert string to int:")
				log.Fatal(err)
			}
			distance = newInt
		}

		fmt.Println(line)
		lineNumber++
	}

	fmt.Printf("\n---------------------------\n")
	fmt.Printf("\n-----------Time------------\n")
	fmt.Printf("\n---------------------------\n")

	fmt.Println(time)

	fmt.Printf("\n---------------------------\n")
	fmt.Printf("\n---------Distance----------\n")
	fmt.Printf("\n---------------------------\n")

	fmt.Println(distance)

	race := Race{time, distance}

	fmt.Printf("\n---------------------------\n")
	fmt.Printf("\n-----------Race------------\n")
	fmt.Printf("\n---------------------------\n")

	fmt.Printf("Race %d)\t time: %d \t recordDistance: %d\n", 1, time, distance)

	possibleWays := make([]int, 0)
	// Get the possible ways to win.
	btnHold := 0          // duration to hold the button to charge the toy boat (takes away from race time)
	distanceTraveled := 0 // distance the boat travels after charging
	possibleWay := 0      // number of possible ways to win

	// First race: Number of possible ways to win
	for j := 0; j <= race.time; j++ {
		btnHold = j
		timeRemaining := race.time - btnHold // the remaining time after holding the button for btnHold duration
		distanceTraveled = btnHold * timeRemaining

		if distanceTraveled > race.recordDistance {
			possibleWay++
		}
	}
	possibleWays = append(possibleWays, possibleWay)

	fmt.Printf("\n---------------------------\n")
	fmt.Printf("\n-----------Ways------------\n")
	fmt.Printf("\n---------------------------\n")

	for i, v := range possibleWays {
		fmt.Printf("Race %d)\t possibleWays: %d\n", i, v)
	}

	marginOfError := 1
	for _, v := range possibleWays {
		fmt.Printf("VALUE: %d\n", v)

		if v != 0 {
			marginOfError *= v
		}
	}

	fmt.Println("RESULT:", marginOfError)
}
