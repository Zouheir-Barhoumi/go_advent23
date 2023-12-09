package part1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/* UTIL */
func pln(arg ...interface{}) {
	fmt.Println(arg...)
}

func ptf(formatStr string, arg ...interface{}) {
	fmt.Fprintf(os.Stdout, formatStr, arg...)
}

/* END OF UTIL */

const (
	R = 1
	L = 0
)

var (
	instructions []rune
)

type Node struct {
	Name  string
	Left  *Node
	Right *Node
}

func PartOne() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	emptyLineFound := false
	nodeTemp := make([]map[string][]string, 0, 6)

	for scanner.Scan() {
		line := scanner.Text()
		pln(line)

		if line == "" {
			emptyLineFound = true
		}
		if !emptyLineFound {

			instructions = append(instructions, []rune(line)...)
		} else if line != "" {
			pln("======================")

			trimmedLine := strings.Trim(line, " ")
			lineFields := strings.Split(trimmedLine, " = ")
			nodeName := strings.Trim(lineFields[0], " ")
			nextNodes := strings.Split(strings.Trim(strings.Trim(strings.Trim(lineFields[1], "("), " "), ")"), ", ")

			pln("Fields:")

			for i, field := range lineFields {
				ptf("Field: %s, index: %d\n", field, i)
			}

			ptf("First: %s, Length: %d\n", nodeName, len(nodeName))
			pln("Second")
			for i, f := range nextNodes {
				ptf("Sfield: %s , Index: %d length: %d\n", f, i, len(f))
			}
			pln("======================")
			pln("Creating NODE %s", nodeName)
			nodeTemp = append(nodeTemp, map[string][]string{nodeName: nextNodes})

			pln("Node TEMP : ", nodeTemp)
		}
	}
	pln(string(instructions))

}
