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

/*
"Our computers are having issues, so I have no idea if we have any Chief Historians in stock! You're welcome to check the warehouse, though," says the mildly flustered shopkeeper at the North Pole Toboggan Rental Shop. The Historians head out to take a look.

The shopkeeper turns to you. "Any chance you can see why our computers are having issues again?"

The computer appears to be trying to run a program, but its memory (your puzzle input) is corrupted. All of the instructions have been jumbled up!

It seems like the goal of the program is just to multiply some numbers. It does that with instructions like mul(X,Y), where X and Y are each 1-3 digit numbers. For instance, mul(44,46) multiplies 44 by 46 to get a result of 2024. Similarly, mul(123,4) would multiply 123 by 4.

However, because the program's memory has been corrupted, there are also many invalid characters that should be ignored, even if they look like part of a mul instruction. Sequences like mul(4*, mul(6,9!, ?(12,34), or mul ( 2 , 4 ) do nothing.

For example, consider the following section of corrupted memory:

xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))
Only the four highlighted sections are real mul instructions. Adding up the result of each instruction produces 161 (2*4 + 5*5 + 11*8 + 8*5).

Scan the corrupted memory for uncorrupted mul instructions. What do you get if you add up all of the results of the multiplications?
*/

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
