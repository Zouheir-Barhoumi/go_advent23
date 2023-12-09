package part1

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	strengths = map[rune]int{
		'A': 13,
		'K': 12,
		'Q': 11,
		'J': 10,
		'T': 9,
		'9': 8,
		'8': 7,
		'7': 6,
		'6': 5,
		'5': 4,
		'4': 3,
		'3': 2,
		'2': 1,
	}

	strengthList = []rune{
		'A',
		'K',
		'Q',
		'J',
		'T',
		'9',
		'8',
		'7',
		'6',
		'5',
		'4',
		'3',
		'2',
	}

	ranks = []string{
		"fiveofakind",
		"fourofakind",
		"fullhouse",
		"threeofakind",
		"twopair",
		"onepair",
		"highcard",
	}
)

type Card struct {
	Value       string
	SortedValue string
	Bid         int
	rank        int
	strength    int
}

func sortHand(v string) string {
	ra := []rune(v)
	sort.Slice(ra, func(i, j int) bool {
		return strengths[ra[i]] > strengths[ra[j]]
	})

	return string(ra)
}

func calcHand(s string) int {
	v := []rune(s)
	// five of a kind
	if v[0] == v[1] && v[1] == v[2] && v[2] == v[3] && v[3] == v[4] {
		return 7
	}

	// four of a kind
	if v[0] == v[1] && v[1] == v[2] && v[2] == v[3] {
		return 6
	}

	if v[1] == v[2] && v[2] == v[3] && v[3] == v[4] {
		return 6
	}

	// full house
	if v[0] == v[1] && v[1] == v[2] && v[3] == v[4] {
		return 5
	}

	if v[0] == v[1] && v[2] == v[3] && v[3] == v[4] {
		return 5
	}

	// three of a kind
	if v[0] == v[1] && v[1] == v[2] {
		return 4
	}

	if v[1] == v[2] && v[2] == v[3] {
		return 4
	}

	if v[2] == v[3] && v[3] == v[4] {
		return 4
	}

	// two pair
	if v[0] == v[1] && v[2] == v[3] {
		return 3
	}

	if v[0] == v[1] && v[3] == v[4] {
		return 3
	}

	if v[1] == v[2] && v[3] == v[4] {
		return 3
	}

	// one pair
	if v[0] == v[1] || v[1] == v[2] || v[2] == v[3] || v[3] == v[4] {
		return 1
	}

	return 0
}

func calcStrength(s string) int {
	v := []rune(s)
	str := float64(0)
	for i, r := range v {
		str += float64(strengths[r]) * math.Pow(100, float64(len(v)-i))
	}

	return int(str)
}

func NewCard(value string, bid int) Card {
	c := Card{
		Value:       value,
		SortedValue: sortHand(value),
		Bid:         bid,
		rank:        0,
		strength:    0,
	}

	c.rank = calcHand(c.SortedValue)
	c.strength = calcStrength(c.Value)

	return c
}

func PartOne() {
	dat, _ := os.ReadFile("input.txt")

	rows := strings.Split(string(dat), "\n")
	cards := []Card{}
	for _, row := range rows {
		if len(row) == 0 {
			continue
		}

		parts := strings.Split(row, " ")
		value := parts[0]
		bid, _ := strconv.Atoi(parts[1])
		cards = append(cards, NewCard(value, bid))
	}

	sort.Slice(cards, func(i, j int) bool {
		return cards[i].rank < cards[j].rank || (cards[i].rank == cards[j].rank && cards[i].strength < cards[j].strength)
	})

	total := 0
	for i, c := range cards {
		fmt.Printf("%v %v\n", c.Value, c.SortedValue)
		total += c.Bid * (i + 1)
	}

	fmt.Printf("total: %d\n", total)
}
