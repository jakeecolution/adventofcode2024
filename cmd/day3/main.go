package main

import (
	"flag"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	problem    = flag.Bool("problem", true, "true for problem 1 and false for problem 2")
	inputFN    = flag.String("inputfn", "test.input", "input filename for ingestion and running against the problem")
	prob1Regex = regexp.MustCompile(`(?s)mul\((\s*\d{1,3},\s*\d{1,3})\)`)
	prob1Test  = regexp.MustCompile(`(?s)mul\((\d{1,3},\d{1,3})\)`)
	prob2Regex = regexp.MustCompile(`(?s)(don't\(\))|(do\(\))|mul\((\s*\d{1,3},\s*\d{1,3})\)`)
)

const (
	DoesCount   = "do()"
	DoesntCount = "don't()"
)

func main() {
	flag.Parse()
	log.Printf("Problem=%v; inputFN=%s", *problem, *inputFN)
	if *problem {
		fstring := readInput(*inputFN)
		total := 0
		for i, match := range prob1Test.FindAllStringSubmatch(fstring, -1) {
			tmp := Multiply(match[1])
			if tmp == -1 {
				log.Fatal("Error in Multiply function.")
			}
			total += tmp
			log.Printf("Match: %v = %d; found at index=%d", match[1], tmp, i)

		}
		log.Printf("Total sum is %d", total)
	} else {
		fstring := readInput(*inputFN)
		total := 0
		do := true
		for i, match := range prob2Regex.FindAllStringSubmatch(fstring, -1) {
			check := match[0]
			bracket := match[len(match)-1]
			// log.Printf("match[%d]=%v; capture=%s; check=%s", i, match, bracket, check)
			if check == DoesCount {
				log.Println("do() found so counting Multiplys")
				do = true
			} else if check == DoesntCount {
				log.Println("don't() found so not counting Multiplys")
				do = false
			} else {
				if do {
					tmp := Multiply(bracket)
					if tmp == -1 {
						log.Fatal("Error in Multiply function.")
					}
					total += tmp
					log.Printf("Match: %v = %d; found at index=%d", bracket, tmp, i)
				} else {
					log.Printf("Match: %v does not count because of the last `don't()` statement...", bracket)
				}
			}
		}
		log.Printf("Total sum is %d", total)
	}
}

func readInput(fn string) string {
	fbytes, err := os.ReadFile(fn)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	return string(fbytes)
}

func Multiply(in string) int {
	numbers := strings.Split(in, ",")
	i1, err := strconv.Atoi(numbers[0])
	if err != nil {
		log.Printf("Error converting string '%s' to integer: %v", numbers[0], err)
		return -1
	}
	i2, err := strconv.Atoi(numbers[1])
	if err != nil {
		log.Printf("Error converting string '%s' to integer: %v", numbers[1], err)
		return -1
	}
	return i1 * i2
}
