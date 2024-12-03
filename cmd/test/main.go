package main

import (
	"log"
	"strconv"
	"strings"
)

const (
	test = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`
)

func main() {
	lines := strings.Split(test, "\n")
	iLines := make([][]int, len(lines))
	for idx, line := range lines {
		myline := strings.Split(line, " ")
		iMyLine := make([]int, len(myline))
		for ix, stri := range myline {
			temp, err := strconv.Atoi(stri)
			if err != nil {
				panic(err)
			}
			iMyLine[ix] = temp
		}
		iLines[idx] = iMyLine
	}
	for _, iLine := range iLines {
		for i := 0; i < len(iLine); i++ {
			var newLine []int = make([]int, 0)
			for ix, item := range iLine {
				if ix == i {
					continue
				}
				newLine = append(newLine, item)
			}

			log.Println(iLine)
			log.Println(newLine, "\n---------------------")
		}
		return
	}
}
