package main

import (
	"flag"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

var (
	problem = flag.Bool("problem", true, "problem is true for the 1st part and false for the 2nd part")
	inputFN = flag.String("inputfn", "test.input", "input file name for giving the problem")
	stdout  = flag.Bool("stdout", true, "Should I print to the terminal or to a file (-stdout=false for file.)")
)

type Update struct {
	Original      []int
	Pass          bool
	PassingUpdate []int
}

func (u Update) MiddlePageNumber() int {
	return u.PassingUpdate[len(u.PassingUpdate)/2]
}

func (u *Update) FixMe(pOrders map[int][]int, errorIndex int) {
	copy(u.PassingUpdate, u.Original)
restartSearch:
	for i := errorIndex; i >= 0; i-- {
		tmp := make([]int, len(u.PassingUpdate[:i]))
		copy(tmp, u.PassingUpdate[:i])
		befores, ok := pOrders[u.PassingUpdate[i]]
		if !ok {
			continue
		}
		for _, before := range befores {
			check, pos := Contains(tmp, before)
			if check {

				t := u.PassingUpdate[pos]
				u.PassingUpdate[pos] = u.PassingUpdate[i]
				u.PassingUpdate[i] = t
				errorIndex = i
				goto restartSearch
			}
		}
	}
	return
}

func Contains(listly []int, comparator int) (bool, int) {
	for i, item := range listly {
		if comparator == item {
			return true, i
		}
	}
	return false, -1
}

func NewUpdate(original []int, passing bool) Update {
	if passing {
		return Update{
			Original:      original,
			Pass:          true,
			PassingUpdate: original,
		}
	}
	return Update{
		Original:      original,
		Pass:          false,
		PassingUpdate: make([]int, len(original)),
	}
}

func main() {
	flag.Parse()
	if !*stdout {
		outfile, err := os.Create("day5.log")
		if err != nil {
			log.Fatalf("Cannot open 'day5.log': %v", err)
		}
		defer outfile.Close()
		log.SetOutput(outfile)
	}
	log.Printf("Problem=%v; inputFN='%s'", *problem, *inputFN)
	fbytes, err := os.ReadFile(*inputFN)
	if err != nil {
		log.Fatalf("Error opening filename '%s': %v", *inputFN, err)
	}
	bOrders := make(map[int][]int)
	fstring := string(fbytes)
	fstrArr := strings.Split(fstring, "\n")
	var pageOrders, printUpdates []string
	for _, item := range fstrArr {
		if strings.Contains(item, "|") {
			pageOrders = append(pageOrders, item)
		} else if strings.Contains(item, ",") {
			printUpdates = append(printUpdates, item)
		}
	}
	pUpdates := make([][]int, len(printUpdates))
	for _, item := range pageOrders {
		subs := strings.Split(item, "|")
		left, err := strconv.Atoi(subs[0])
		if err != nil {
			log.Fatalf("Couldn't convert %s to integer: %v", subs[0], err)
		}
		right, err := strconv.Atoi(subs[1])
		if err != nil {
			log.Fatalf("Couldn't convert %s to integer: %v", subs[1], err)
		}
		if _, ok := bOrders[left]; !ok {
			bOrders[left] = []int{right}
		} else {
			bOrders[left] = append(bOrders[left], right)
		}
	}
	for idy, item := range printUpdates {
		tstr := strings.Split(item, ",")
		intArr := make([]int, len(tstr))
		for idx, num := range tstr {
			iNum, err := strconv.Atoi(num)
			if err != nil {
				log.Fatalf("Error converting '%s' to integer: %v", num, err)
			}
			intArr[idx] = iNum
		}
		// log.Println(intArr)
		pUpdates[idy] = make([]int, len(intArr))
		copy(pUpdates[idy], intArr)
	}
	// log.Println(bOrders, "\n\n-----------------\n", pUpdates, "\n----------------")
	allUpdates := make([]Update, len(pUpdates))
	total := 0
	notCorrectTotal := 0
	for index, update := range pUpdates {
		// log.Println(update)
		for i := len(update) - 1; i >= 0; i-- {
			updateValid := true
			// log.Println("i=", i)
			befores, ok := bOrders[update[i]]
			if !ok {
				continue
			}
			if i == 0 {
				log.Printf("Update passed: '%v' = %d", update, update[len(update)/2])
				allUpdates[index] = NewUpdate(update, true)
				total += allUpdates[index].MiddlePageNumber()
			}
			tmp := make([]int, len(update[:i]))
			copy(tmp, update[:i])
			slices.Sort(tmp)
			for _, before := range befores {
				exists := slices.Contains(tmp, before)
				if exists {
					log.Printf("Failed at pOrder[%d]=%v using %d, tmp=%v", update[i], befores, before, tmp)
					updateValid = false
					allUpdates[index] = NewUpdate(update, false)
					allUpdates[index].FixMe(bOrders, i)
					notCorrectTotal += allUpdates[index].MiddlePageNumber()
					break
				}
			}
			if !updateValid {
				break
			}
		}
	}
	// log.Println(bOrders)
	log.Println("Total =", total)
	log.Println("Not correct total=", notCorrectTotal)
}
