package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	problem = flag.Bool("problem", true, "problem is true for the 1st part and false for the 2nd part")
	inputFN = flag.String("inputfn", "test.input", "input file name for giving the problem")
	stdout  = flag.Bool("stdout", true, "Should I print to the terminal or to a file (-stdout=false for file.)")
)

func main() {
	flag.Parse()

	if !*stdout {
		outfile, err := os.Create("day7.log")
		if err != nil {
			log.Fatalf("Cannot open 'day7.log': %v", err)
		}
		defer outfile.Close()
		log.SetOutput(outfile)
	}

	fbytes, err := os.ReadFile(*inputFN)
	if err != nil {
		log.Fatalf("Error opening file '%s': %v", *inputFN, err)
	}
	operations := strings.Split(string(fbytes), "\n")
	total := 0
	for _, operartion := range operations {
		if calculation := CalculateOp(operartion, !*problem); calculation > 0 {
			total += calculation
		}
	}
	log.Printf("Total of passed calculations is '%d'", total)
}

func CalculateOp(line string, isPart2 bool) int {
	operation := strings.Split(line, ":")
	equals, err := strconv.Atoi(operation[0])
	if err != nil {
		log.Fatalf("1. Error converting string to integer: %v", err)
	}

	rightSide := strings.Fields(operation[1])
	rightInts := make([]int, len(rightSide))
	for idx, item := range rightSide {
		nItem, err := strconv.Atoi(item)
		if err != nil {
			log.Fatalf("2. Error converting string to integer: %v", err)
		}
		rightInts[idx] = nItem
	}
	passThru := calcOp(equals, 0, rightInts, isPart2)
	if passThru {
		log.Printf("%d= %v passed calculation test", equals, rightInts)
		return equals
	}
	log.Printf("%d= %v failed calculation test", equals, rightInts)
	return 0
}

func calc(a, b int, operation byte) int {
	calculation := 0
	var err error
	switch operation {
	case '+':
		calculation = a + b
		// log.Printf("%d = %d + %d", calculation, a, b)
	case '*':
		calculation = a * b
		// log.Printf("%d = %d * %d", calculation, a, b)
	case '|':
		calculation, err = strconv.Atoi(fmt.Sprintf("%d%d", a, b))
		if err != nil {
			log.Fatalf("Error converting string to integer: %v", err)
		}
	}
	// log.Printf("calculation=%d, a=%d, b=%d", calculation, a, b)
	return calculation
}

func calcOp(expectedSum, sum int, input []int, isPart2 bool) bool {
	if len(input) == 0 {
		// log.Printf("%d == %d", sum, expectedSum)
		return sum == expectedSum
	}

	if sum > expectedSum {
		return false
	}

	if calcOp(expectedSum, calc(sum, input[0], '+'), input[1:], isPart2) {
		return true
	}

	if isPart2 && calcOp(expectedSum, calc(sum, input[0], '|'), input[1:], isPart2) {
		return true
	}

	return calcOp(expectedSum, calc(sum, input[0], '*'), input[1:], isPart2)
}
