package main

import (
	"errors"
	"flag"
	"log"

	"github.com/jakeecolution/adventofcode2024/foundation/input"
)

/*
 *
The unusual data (your puzzle input) consists of many reports, one report per line. Each report is a list of numbers called levels that are separated by spaces. For example:

7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
This example data contains six reports each containing five levels.

The engineers are trying to figure out which reports are safe. The Red-Nosed reactor safety systems can only tolerate levels that are either gradually increasing or gradually decreasing. So, a report only counts as safe if both of the following are true:

The levels are either all increasing or all decreasing.
Any two adjacent levels differ by at least one and at most three.
In the example above, the reports can be found safe or unsafe by checking those rules:

7 6 4 2 1: Safe because the levels are all decreasing by 1 or 2.
1 2 7 8 9: Unsafe because 2 7 is an increase of 5.
9 7 6 2 1: Unsafe because 6 2 is a decrease of 4.
1 3 2 4 5: Unsafe because 1 3 is increasing but 3 2 is decreasing.
8 6 4 4 1: Unsafe because 4 4 is neither an increase or a decrease.
1 3 6 7 9: Safe because the levels are all increasing by 1, 2, or 3.
So, in this example, 2 reports are safe.

Analyze the unusual data from the engineers. How many reports are safe?
*/

var (
	// Example for problem `go run main.go -problem=false`
	problem                   *bool   = flag.Bool("problem", true, "True for problem 1 and False for problem 2")
	inputFN                   *string = flag.String("inputfn", "test.input", "Input filename for the problems")
	ErrReactorUnsafeDecrease          = errors.New("reactor unsafe because report decreases too much")
	ErrReactorUnsafeIncrease          = errors.New("reactor unsafe because report increases too much")
	ErrReactorUnsafeStaysSame         = errors.New("reactor unsafe because reports stays the same from one report to another")
	ErrReactorUnsafeSwitches          = errors.New("reactor unsafe because reports switch direction from inc to dec or vice versa")
)

func main() {
	flag.Parse()
	log.Println(*inputFN)
	problemIn, err := input.ReadIntMatrix(*inputFN)
	if err != nil {
		log.Fatalf("%v", err)
	}
	if *problem {
		safeReports := 0
		for idx, line := range problemIn {
			_, err := ReactorReportSafety(line)
			if err != nil {
				log.Printf("Line[%d]= '%v' and Error = '%v'", idx, line, err)
				continue
			}
			log.Printf("Line[%d] = '%v' is Safe", idx, line)
			safeReports += 1
		}
		log.Println("Safe reports =", safeReports)
	} else {
		/*
					--- Part Two ---
			The engineers are surprised by the low number of safe reports until they realize they forgot to tell you about the Problem Dampener.

			The Problem Dampener is a reactor-mounted module that lets the reactor safety systems tolerate a single bad level in what would otherwise be a safe report. It's like the bad level never happened!

			Now, the same rules apply as before, except if removing a single level from an unsafe report would make it safe, the report instead counts as safe.

			More of the above example's reports are now safe:

			7 6 4 2 1: Safe without removing any level.
			1 2 7 8 9: Unsafe regardless of which level is removed.
			9 7 6 2 1: Unsafe regardless of which level is removed.
			1 3 2 4 5: Safe by removing the second level, 3.
			8 6 4 4 1: Safe by removing the third level, 4.
			1 3 6 7 9: Safe without removing any level.
			Thanks to the Problem Dampener, 4 reports are actually safe!

			Update your analysis by handling situations where the Problem Dampener can remove a single level from unsafe reports. How many reports are now safe?
		*/
		safeReports := 0
		for idx, line := range problemIn {
			index, err := ReactorReportSafety(line)
			if err != nil {
				initErr := err
				myflag := false
				for i := 0; i < len(line); i++ {
					var newline []int = make([]int, 0)
					for idy, item := range line {
						if idy == i {
							continue
						}
						newline = append(newline, item)
					}
					err = nil
					_, err = ReactorReportSafety(newline)
					if err == nil {
						log.Printf("Line[%d] = '%v' is Safe by removing the item at this level %d\n", idx, newline, index+1)
						safeReports += 1
						myflag = true
						break
					}
				}
				if myflag {
					continue
				}
				log.Printf("Line[%d] = %v and Error = '%v'", idx, line, initErr)
			} else {
				log.Printf("Line[%d] = '%v' is Safe", idx, line)
				safeReports += 1
			}
		}
		log.Println("Safe reports =", safeReports)
	}
}

func ReactorReportSafety(myline []int) (int, error) {
	var direction bool
	for i := 0; i < len(myline)-1; i++ {
		cur, next := myline[i], myline[i+1]
		diff := next - cur
		if i == 0 {
			if diff > 0 {
				direction = true
			} else if diff == 0 {
				return 0, ErrReactorUnsafeStaysSame
			} else {
				direction = false
			}
		}
		if direction {
			if diff <= 3 && diff > 0 {
				continue
			} else if diff > 3 {
				return i, ErrReactorUnsafeIncrease
			} else if diff == 0 {
				return i, ErrReactorUnsafeStaysSame
			} else {
				return i, ErrReactorUnsafeSwitches
			}
		} else {
			if diff >= -3 && diff < 0 {
				continue
			} else if diff < -3 {
				return i, ErrReactorUnsafeDecrease
			} else if diff == 0 {
				return i, ErrReactorUnsafeStaysSame
			} else {
				return i, ErrReactorUnsafeSwitches
			}
		}

	}
	return len(myline) - 1, nil
}
