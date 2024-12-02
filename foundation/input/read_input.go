package input

import (
	"log"
	"os"
	"strconv"
	"strings"
)

// filename includes path
func ReadIntInput(filename string) ([]int, []int, error) {
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
