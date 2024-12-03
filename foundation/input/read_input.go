package input

import (
	"log"
	"os"
	"strconv"
	"strings"
)

// filename includes path
func ReadIntColumns(filename string) ([]int, []int, error) {
	var err error = nil
	filecontents, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("Error reading file '%s': %v", filename, err)
		return []int{}, []int{}, err
	}
	contents := string(filecontents)
	sliceLen := strings.Count(contents, "\n") + 1
	input1 := make([]int, sliceLen)
	input2 := make([]int, sliceLen)
	for i, item := range strings.Split(contents, "\n") {
		items := strings.Split(item, "   ")
		tmp, err := strconv.Atoi(items[0])
		if err != nil {
			log.Printf("Error parsing int at line %d", i+1)
			return []int{}, []int{}, err
		}
		input1[i] = tmp
		tmp2, err := strconv.Atoi(items[1])
		if err != nil {
			log.Printf("Error parsing int at line %d", i+1)
			return []int{}, []int{}, err
		}
		input2[i] = tmp2
	}
	return input1, input2, err
}

func ReadIntMatrix(filename string) ([][]int, error) {
	var err error = nil
	var errOutput [][]int = make([][]int, 0)
	fcontents, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("Error reading output from filename '%s' - Error = '%v", filename, err)
		return errOutput, err
	}
	fstring := string(fcontents)
	lineCount := strings.Count(fstring, "\n") + 1
	input := make([][]int, lineCount)
	for idx, item := range strings.Split(fstring, "\n") {
		items := strings.Split(item, " ")
		line := make([]int, len(items))
		for i, itm := range items {
			t, err := strconv.Atoi(itm)
			if err != nil {
				log.Printf("Failed converting a string to an integer - %v", err)
				return errOutput, err
			}
			line[i] = t
		}
		input[idx] = line
	}
	return input, err
}
