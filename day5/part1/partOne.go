package part1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func sliceShow[T any](slice []T, name string) {
	if len(slice) > 0 {
		fmt.Printf(name + "s" + ":\n")
		for _, el := range slice {
			fmt.Println(el)
		}
	}
}

func mapShow[K comparable, V any](m map[K]V, name string) {
	if len(m) > 0 {
		fmt.Println(name + ":")
		for k, v := range m {
			fmt.Printf("%s - %v: %v\n", name, k, v)
		}
	}
}

func parseNumber(s string) (int, error) {
	return strconv.Atoi(s)
}

func createMap(input []string) map[int]int {
	resultMap := make(map[int]int)

	for _, line := range input {
		fields := strings.Fields(line)
		if len(fields) != 3 {
			fmt.Println("Invalid input line", line)
			continue
		}

		destStart, _ := parseNumber(fields[0])
		sourStart, _ := parseNumber(fields[1])
		rangeLen, _ := parseNumber(fields[2])

		for i := 0; i < rangeLen; i++ {
			resultMap[sourStart] = destStart
			sourStart++
			destStart++
		}
	}
	fmt.Println("Map created!")
	return resultMap
}

func mapSeedToLocation(tempSeed int, sToSoil, sToFert, sToWat, sToLight, sToTemp, sToHum, sToLoc map[int]int) (int, int) {
	// chennel to receive location results
	locationChannel := make(chan int)

	// goroutine to calc the location
	go func() {

		tempSoil, found := sToSoil[tempSeed]
		if !found {
			// fmt.Println("Seed mapping not found for seed number:", tempSeed)
			tempSoil = tempSeed
		} else {
			fmt.Printf("Seed mapping found %d for seed number %d\n", tempSoil, tempSeed)
		}

		tempFert, found := sToFert[tempSoil]
		if !found {
			// fmt.Println("Fertilizer mapping not found for soil number:", tempSoil)
			tempFert = tempSoil
		}

		tempWater, found := sToWat[tempFert]
		if !found {
			// fmt.Println("Water mapping not found for fertilizer number:", tempFert)
			tempWater = tempFert
		}

		tempLight, found := sToLight[tempWater]
		if !found {
			// fmt.Println("Light mapping not found for water number:", tempWater)
			tempLight = tempWater
		}

		tempTemp, found := sToTemp[tempLight]
		if !found {
			// fmt.Println("Temperature mapping not found for light number:", tempLight)
			tempTemp = tempLight
		}

		tempHumid, found := sToHum[tempTemp]
		if !found {
			// fmt.Println("Humidity mapping not found for temperature number:", tempTemp)
			tempHumid = tempTemp
		}

		tempLocation, found := sToLoc[tempHumid]
		if !found {
			// fmt.Println("Location mapping not found for humidity number:", tempHumid)
			tempLocation = tempHumid
		}
		fmt.Printf("seed: %d -> soil: %d -> Fert: %d -> Water: %d -> Light: %d -> Temp: %d -> Hmid: %d -> Location: %d\n", tempSeed, tempSoil, tempFert, tempWater, tempLight, tempTemp, tempHumid, tempLocation)
		locationChannel <- tempLocation
	}()

	// Wait for location result
	tempLocation := <-locationChannel

	fmt.Println("Got Location!")
	return tempSeed, tempLocation
}

func PartOne() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	seg_strings := make([]string, 8)

	empty := 0
	for scanner.Scan() {
		line := scanner.Text()

		// count # of empty lines
		if line == "" {
			empty++
			continue
		} else {
			switch empty {
			case 0:
				semi := strings.Index(line, ":")
				seg_strings[0] += strings.TrimLeft(line[semi+1:], " ")
				// fmt.Printf("String: %s\n", seg_strings[0])
			case 1:
				semi := strings.Index(line, ":")
				seg_strings[1] += strings.TrimLeft(line[semi+1:], " ") + ","
				// fmt.Printf("String: %s\n", seg_strings[1])
			case 2:
				semi := strings.Index(line, ":")
				seg_strings[2] += strings.TrimLeft(line[semi+1:], " ") + ","
				// fmt.Printf("String: %s\n", seg_strings[2])
			case 3:
				semi := strings.Index(line, ":")
				seg_strings[3] += strings.TrimLeft(line[semi+1:], " ") + ","
				// fmt.Printf("String: %s\n", seg_strings[3])
			case 4:
				semi := strings.Index(line, ":")
				seg_strings[4] += strings.TrimLeft(line[semi+1:], " ") + ","
				// fmt.Printf("String: %s\n", seg_strings[4])
			case 5:
				semi := strings.Index(line, ":")
				seg_strings[5] += strings.TrimLeft(line[semi+1:], " ") + ","
				// fmt.Printf("String: %s\n", seg_strings[5])
			case 6:
				semi := strings.Index(line, ":")
				seg_strings[6] += strings.TrimLeft(line[semi+1:], " ") + ","
				// fmt.Printf("String: %s\n", seg_strings[6])
			default:
				semi := strings.Index(line, ":")
				seg_strings[7] += strings.TrimLeft(line[semi+1:], " ") + ","
				// fmt.Printf("String: %s\n", seg_strings[6])
			}
		}
		// fmt.Println(line)
	}

	seeds := make([]int, 0)

	seedsString := strings.Fields(seg_strings[0])

	for _, v := range seedsString {
		// fmt.Println("Assiging to index:", k)
		n, err := parseNumber(v)
		if err != nil {
			fmt.Println("Error assiging seedsS to seeds slice")
		} else {
			seeds = append(seeds, n)
		}
	}

	// remove commas from all strings in slice
	for i, seg := range seg_strings {
		if strings.Contains(seg, ",") {
			// fmt.Println("Does contain comma:")
			// fmt.Println(seg)
			seg_strings[i] = strings.TrimLeft(seg, ",")
			seg_strings[i] = strings.TrimRight(seg_strings[i], ",")
			// fmt.Println("Comma removed:")
			// fmt.Println(seg)
		}
	}

	// create string slice
	segmentLength := 3
	sToSoils := make([]string, 0, (len(seg_strings[1])/6)*segmentLength-4)
	sToFerts := make([]string, 0, (len(seg_strings[2])/6)*segmentLength-4)
	sToWats := make([]string, 0, (len(seg_strings[3])/6)*segmentLength-4)
	sToLights := make([]string, 0, (len(seg_strings[4])/6)*segmentLength-4)
	sToTemps := make([]string, 0, (len(seg_strings[5])/6)*segmentLength-4)
	sToHums := make([]string, 0, (len(seg_strings[6])/6)*segmentLength-4)
	sToLocs := make([]string, 0, (len(seg_strings[7])/6)*segmentLength-4)

	// create int slices
	for i, s := range seg_strings {
		switch i {
		case 1:
			sToSoils = append(sToSoils, strings.Split(s, ",")...)
			// fmt.Println("Success!!!", sToSoils)
		case 2:
			sToFerts = append(sToFerts, strings.Split(s, ",")...)
			// fmt.Println("Success!!!", sToFerts)
		case 3:
			sToWats = append(sToWats, strings.Split(s, ",")...)
			// fmt.Println("Success!!!", sToWats)
		case 4:
			sToLights = append(sToLights, strings.Split(s, ",")...)
			// fmt.Println("Success!!!", sToLights)
		case 5:
			sToTemps = append(sToTemps, strings.Split(s, ",")...)
			// fmt.Println("Success!!!", sToTemps)
		case 6:
			sToHums = append(sToHums, strings.Split(s, ",")...)
			// fmt.Println("Success!!!", sToHums)
		case 7:
			sToLocs = append(sToLocs, strings.Split(s, ",")...)
			// fmt.Println("Success!!!", sToLocs)
		}
	}

	// create maps
	sToSoil := createMap(sToSoils)
	sToFert := createMap(sToFerts)
	sToWat := createMap(sToWats)
	sToLight := createMap(sToLights)
	sToTemp := createMap(sToTemps)
	sToHum := createMap(sToHums)
	sToLoc := createMap(sToLocs)

	// display maps
	// mapShow(sToSoil, "SEED TO SOIL MAP")
	// mapShow(sToFert, "SOIL TO FERTILIZER MAP")
	// mapShow(sToWat, "FERTILIZER TO WATER MAP")
	// mapShow(sToLight, "WATER TO LIGHT")
	// mapShow(sToTemp, "LIGHT TO TEMPERATURE")
	// mapShow(sToHum, "TEMPERATURE TO HUMIDITY")
	// mapShow(sToLoc, "HUMIDITY TO LOCATION")

	// tempSeed := 13

	// location
	// seed, location := mapSeedToLocation(tempSeed, sToSoil, sToFert, sToWat, sToLight, sToTemp, sToHum, sToLoc)

	// fmt.Printf("Locatoin for seed number %d is %d\n", seed, location)

	fmt.Printf("\n\n----------------------------------------------\n\n")
	minimum := int(^uint(0) >> 1)

	for i, seed := range seeds {
		fmt.Printf("SEED: %d  -  %d\n", i, seed)
		// GET THE LOWEST LOCATION
		// Goroutine to map the seed and get the location

		_, location := mapSeedToLocation(seed, sToSoil, sToFert, sToWat, sToLight, sToTemp, sToHum, sToLoc)

		if location < minimum {
			minimum = location
		}
	}
	for _, seg := range seg_strings {
		fmt.Println(seg)
	}

	fmt.Println("The lowest location is:", minimum)
}
