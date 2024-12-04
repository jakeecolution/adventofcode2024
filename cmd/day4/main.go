package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	problem        = flag.Bool("problem", true, "true for problem 1 and `-program=false` for problem 2")
	inputFN        = flag.String("inputfn", "test.input", "input filename for the problem to process")
	ErrOutOfBounds = errors.New("array out of bounds")
)

const (
	X = rune('X')
	M = rune('M')
	A = rune('A')
	S = rune('S')
)

func main() {
	flag.Parse()
	// log.Printf("Problem=%v; inputFN='%s'", *problem, *inputFN)
	log.Printf("'M'=%v, 'S'=%v", M, S)
	fbytes, err := os.ReadFile(*inputFN)
	if err != nil {
		log.Fatalf("Error with opening inputFN=%s - %v", *inputFN, err)
	}
	fstring := string(fbytes)
	lines := strings.Split(fstring, "\n")
	linesb := make([][]rune, len(lines))
	for idx, line := range lines {
		linesb[idx] = []rune(line)
	}
	// log.Println(X, M, A, S)
	total := 0
	for i := 0; i < len(linesb); i++ {
		for j := 0; j < len(linesb[i]); j++ {
			if *problem {
				if linesb[i][j] == X {
					log.Printf("Found an 'X' at linesb[%d][%d]", j, i)
					total += findXMAS(linesb, j, i)
				}
			} else {
				if linesb[i][j] == A {
					total += findBigXMAS(linesb, j, i)
				}
			}
		}
	}
	log.Printf("Total XMAS=%d", total)

}

func findBigXMAS(linearr [][]rune, x, y int) int {
	if y == 0 || y == len(linearr)-1 {
		log.Printf("out of bounds [%d, %d]", y, x)
		return 0
	}
	if x == 0 || x == len(linearr[y])-1 {
		log.Printf("out of bounds [%d, %d]", y, x)
		return 0
	}
	edges := make([][]rune, 2)
	edges[0] = make([]rune, 2)
	edges[1] = make([]rune, 2)
	edges[0][0] = linearr[y-1][x-1]
	edges[0][1] = linearr[y-1][x+1]
	edges[1][0] = linearr[y+1][x-1]
	edges[1][1] = linearr[y+1][x+1]
	for _, edge := range edges {
		if !isEdge(edge[0]) || !isEdge(edge[1]) {
			log.Printf("Items are not an edge %v, %v", edge[0], edge[1])
			return 0
		}
	}
	sCount, mCount := 0, 0
	for _, edgeRow := range edges {
		for _, edge := range edgeRow {
			if edge == S {
				sCount += 1
			} else {
				mCount += 1
			}
		}
	}
	if sCount != mCount {
		return 0
	}
	leftTop := edges[0][0]
	if leftTop != edges[0][1] && leftTop != edges[1][0] {
		log.Printf("Exiting due to no leftTop match: %v", edges)
		return 0
	}
	rightBottom := edges[1][1]
	if rightBottom == edges[0][1] || rightBottom == edges[1][0] {
		log.Printf("Found an X-MAS at [%d, %d]\n%s\n-----------------------", y, x, printBigXmas(linearr, y, x))
		return 1
	}
	log.Printf("Exiting due to no rightBottom match: %v", edges)
	return 0
}

func findXMAS(linearr [][]rune, x, y int) int {
	// find M; determine direction
	directions := findDirection(linearr, x, y)
	log.Printf("Direction=%v", directions)
	if len(directions) == 0 {
		return 0
	}
	total := 0
	for _, direction := range directions {
		curx, cury := x, y
		str := fmt.Sprintf("XMAS Found: X=[%d, %d], M=[%d, %d], ", cury, curx, cury+direction[0], curx+direction[1])
		cury += direction[0] * 2
		curx += direction[1] * 2
		if isLetter(A, linearr, curx, cury) {
			str += fmt.Sprintf("A=[%d, %d], ", cury, curx)
			cury += direction[0]
			curx += direction[1]
			if isLetter(S, linearr, curx, cury) {
				log.Printf("%sS=[%d, %d]", str, cury, curx)
				total += 1
			}
		}
	}
	// look along direction for rest of letters
	return total
}

func findDirection(linearr [][]rune, x, y int) [][]int {
	direction := make([][]int, 0)
	var dir bool = false
	dir = isLetter(M, linearr, x-1, y-1)
	if dir {
		direction = append(direction, []int{-1, -1})
	}
	dir = isLetter(M, linearr, x, y-1)
	if dir {
		direction = append(direction, []int{-1, 0})
	}
	dir = isLetter(M, linearr, x+1, y-1)
	if dir {
		direction = append(direction, []int{-1, 1})
	}
	dir = isLetter(M, linearr, x+1, y)
	if dir {
		direction = append(direction, []int{0, 1})
	}
	dir = isLetter(M, linearr, x-1, y)
	if dir {
		direction = append(direction, []int{0, -1})
	}
	dir = isLetter(M, linearr, x+1, y+1)
	if dir {
		direction = append(direction, []int{1, 1})
	}
	dir = isLetter(M, linearr, x, y+1)
	if dir {
		direction = append(direction, []int{1, 0})
	}
	dir = isLetter(M, linearr, x-1, y+1)
	if dir {
		direction = append(direction, []int{1, -1})
	}
	return direction
}

func isEdge(letter rune) bool {
	if letter == M || letter == S {
		return true
	}
	return false
}

func printBigXmas(linearr [][]rune, y, x int) string {
	output := ""
	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			output += string(linearr[i][j])
		}
		output += "\n"
	}
	return output
}

func isLetter(letter rune, linearr [][]rune, x, y int) bool {
	if y < 0 || y >= len(linearr) {
		return false
	} else if x < 0 || x >= len(linearr[y]) {
		return false
	}
	if letter == linearr[y][x] {
		return true
	}
	return false
}
