package part1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var times = make([]int, 0)
var distances = make([]int, 0)

type Race struct {
	time           int
	recordDistance int
}

func PartOne() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// First line is times and Second is Distance
	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()

		// Remove the 'Time:' and 'Distance' parts
		semi := strings.Index(line, ":")
		nOsemi := line[semi+1:]

		// Trim initial and trailing whitespace
		noInitSpace := strings.TrimLeft(nOsemi, " ")
		line = strings.TrimRight(noInitSpace, " ")

		// Get the times
		if lineNumber == 0 {
			stringSlice := strings.Fields(line)
			for _, field := range stringSlice {
				trimmedField := strings.TrimSpace(field)
				newInt, err := strconv.Atoi(trimmedField)
				if err != nil {
					fmt.Println("Failed to convert string to int:")
					log.Fatal(err)
				}
				times = append(times, newInt)

			}
		}
		// Get the distances
		if lineNumber == 1 {
			stringSlice := strings.Fields(line)
			for _, field := range stringSlice {
				trimmedField := strings.TrimSpace(field)
				newInt, err := strconv.Atoi(trimmedField)
				if err != nil {
					fmt.Println("Failed to convert string to int:")
					log.Fatal(err)
				}
				distances = append(distances, newInt)

			}
		}

		fmt.Println(line)
		lineNumber++
	}

	fmt.Printf("\n---------------------------\n")
	fmt.Printf("\n-----------Times-----------\n")
	fmt.Printf("\n---------------------------\n")

	for _, v := range times {
		fmt.Println(v)
	}

	fmt.Printf("\n---------------------------\n")
	fmt.Printf("\n---------Distances---------\n")
	fmt.Printf("\n---------------------------\n")

	for _, v := range distances {
		fmt.Println(v)
	}

	races := make([]Race, 0)
	// Populate Races with times and distances
	// race 0 takes times[0] as race.time and distances[0] as race.recordDistance
	for i := 0; i < len(times); i++ {
		races = append(races, Race{times[i], distances[i]})
	}

	fmt.Printf("\n---------------------------\n")
	fmt.Printf("\n-----------Races-----------\n")
	fmt.Printf("\n---------------------------\n")

	for i, v := range races {
		fmt.Printf("Race %d)\t time: %d \t recordDistance: %d\n", i, v.time, v.recordDistance)
	}

	possibleWays := make([]int, 0)
	// Get the possible ways to win.
	for _, race := range races {
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

	}

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
