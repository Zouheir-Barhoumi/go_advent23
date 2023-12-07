package part1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	HighHand = iota + 1
	OnePair
	TwoPair
	ThreeOf
	FullHouse
	FourOf
	FiveOf
)

var ranks = []string{
	"HighHand",
	"OnePair",
	"TwoPair",
	"ThreeOf",
	"FullHouse",
	"FourOf",
	"FiveOf",
}

type Hand struct {
	hand string
	bid  int
}

var hands = make([]Hand, 0, 5)

var highHand = make([]Hand, 0)
var onePair = make([]Hand, 0)
var twoPair = make([]Hand, 0)
var threeOf = make([]Hand, 0)
var fullHouse = make([]Hand, 0)
var fourOf = make([]Hand, 0)
var fiveOf = make([]Hand, 0)

var allHandsSorted = make([]Hand, 0)

func mapIntoSortedDict(unord map[string]int) []KeyValue {
	var keyValuePairs = []KeyValue{}
	for k, v := range unord {
		keyValuePairs = append(keyValuePairs, KeyValue{k, v})
	}

	sort.Slice(keyValuePairs, func(m, j int) bool {
		return keyValuePairs[m].Value > keyValuePairs[j].Value
	})
	return keyValuePairs
}

func getCardStren(card string) int {
	switch card {
	case "2":
		return 1
	case "3":
		return 2
	case "4":
		return 3
	case "5":
		return 4
	case "6":
		return 5
	case "7":
		return 6
	case "8":
		return 7
	case "T":
		return 8
	case "J":
		return 9
	case "Q":
		return 10
	case "K":
		return 11
	case "A":
		return 12
	default:
		return 0
	}
}

func getHandType(cards []string) (int, string) {
	var card_counter = map[string]int{}
	strength := 0
	for _, card := range cards {
		strength = getCardStren(card)
		if card_counter[card] > 0 {
			card_counter[card]++
		} else {
			card_counter[card] = 1
		}
		fmt.Printf("Cards: %s, Strength: %d, Repetition: %d\n", card, strength, card_counter[card])
	}

	// []KeyValue
	descHandDict := mapIntoSortedDict(card_counter)

	handType := 0
	handName := ""
	// foundFive := false
	// foundFour := false
	// foundThree := false

	for i, card := range descHandDict {
		switch card.Value {
		case 5:

			handType = FiveOf
			handName = ranks[handType-1]
			fmt.Println("------------------")
			fmt.Printf("Card: %d=%s , Hand Type: %d-%s\n", card.Value, card.Key, handType, handName)
			fmt.Println("------------------")
			return handType, handName
		case 4:

			handType = FourOf
			handName = ranks[handType-1]
			fmt.Println("------------------")
			fmt.Printf("Card: %d=%s , Hand Type: %d-%s\n", card.Value, card.Key, handType, handName)
			fmt.Println("------------------")
			return handType, handName
		case 3:
			if descHandDict[i+1].Value == 2 {
				handType = FullHouse
				handName = ranks[handType-1]
				fmt.Println("------------------")
				fmt.Printf("Card: %d=%s , Hand Type: %d-%s\n", card.Value, card.Key, handType, handName)
				fmt.Println("------------------")
				return handType, handName
			} else {
				handType = ThreeOf
				handName = ranks[handType-1]
				fmt.Println("------------------")
				fmt.Printf("Card: %d=%s , Hand Type: %d-%s\n", card.Value, card.Key, handType, handName)
				fmt.Println("------------------")
				return handType, handName
			}
		case 2:
			if descHandDict[i+1].Value == 2 {
				handType = TwoPair
				handName = ranks[handType-1]
				fmt.Println("------------------")
				fmt.Printf("Card: %d=%s , Hand Type: %d-%s\n", card.Value, card.Key, handType, handName)
				fmt.Println("------------------")
				return handType, handName
			} else if descHandDict[i+1].Key == descHandDict[i].Key {
				handType = OnePair
				handName = ranks[handType-1]
				fmt.Println("------------------")
				fmt.Printf("Card: %d=%s , Hand Type: %d-%s\n", card.Value, card.Key, handType, handName)
				fmt.Println("------------------")
				return handType, handName
			} else {
				handType = HighHand
				handName = ranks[handType-1]
				fmt.Println("------------------")
				fmt.Printf("Card: %d=%s , Hand Type: %d-%s\n", card.Value, card.Key, handType, handName)
				fmt.Println("------------------")
				return handType, handName
			}
		case 1:
			if descHandDict[i+1].Key == descHandDict[i].Key {
				handType = OnePair
				handName = ranks[handType-1]
				fmt.Println("------------------")
				fmt.Printf("Card: %d=%s , Hand Type: %d-%s\n", card.Value, card.Key, handType, handName)
				fmt.Println("------------------")
				return handType, handName
			} else {
				handType = HighHand
				handName = ranks[handType-1]
				fmt.Println("------------------")
				fmt.Printf("Card: %d=%s , Hand Type: %d-%s\n", card.Value, card.Key, handType, handName)
				fmt.Println("------------------")
				return handType, handName
			}
		default:
			fmt.Printf("Type Not FOUND!")

		}
	}
	return handType, handName
}

type KeyValue struct {
	Key   string
	Value int
}

func storeByType(h Hand) {
	// split into cards
	cards := strings.Split(h.hand, "")
	hT, hN := getHandType(cards)
	fmt.Println("getHandType(cards)", hT, hN)
	switch hN {
	case "HighHand":
		highHand = append(highHand, h)
	case "OnePair":
		onePair = append(onePair, h)
	case "TwoPair":
		twoPair = append(twoPair, h)
	case "ThreeOf":
		threeOf = append(threeOf, h)
	case "FullHouse":
		fullHouse = append(fullHouse, h)
	case "FourOf":
		fourOf = append(fourOf, h)
	case "FiveOf":
		fiveOf = append(fiveOf, h)
	}
}

func PartOne() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			// Populate hands array
			stringSlice := strings.Fields(line)
			hand := stringSlice[0]
			bid, _ := strconv.Atoi(stringSlice[1])
			hands = append(hands, Hand{hand, bid})
		}
	}

	for _, h := range hands {
		fmt.Println(h)
		storeByType(h)
	}

	type Card struct {
		Card     string
		Strength int
	}

	// KeyValue
	for _, v := range highHand {
		cards := strings.Split(v.hand, "")

		fmt.Println(cards)
		// stren := getCardStren(v.hand)

	}

}
