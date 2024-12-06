package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jakeecolution/adventofcode2024/foundation"
)

var (
	problem = flag.Bool("problem", true, "problem is true for the 1st part and false for the 2nd part")
	inputFN = flag.String("inputfn", "test.input", "input file name for giving the problem")
	stdout  = flag.Bool("stdout", true, "Should I print to the terminal or to a file (-stdout=false for file.)")
)

const (
	Up          = '^'
	Down        = 'v'
	Right       = '>'
	Left        = '<'
	Obstacle    = '#'
	Empty       = '.'
	Beenhere    = 'X'
	NewObstacle = 'O'
	VertMove    = '|'
	HoriMove    = '-'
	BothMove    = '+'
)

func main() {
	flag.Parse()
	if !*stdout {
		outfile, err := os.Create("day6.log")
		if err != nil {
			log.Fatalf("Cannot open 'day6.log': %v", err)
		}
		defer outfile.Close()
		log.SetOutput(outfile)
	}
	fbytes, err := os.ReadFile(*inputFN)
	if err != nil {
		log.Fatalf("Error opening file '%s': %v", *inputFN, err)
	}
	mapStr := strings.Split(string(fbytes), "\n")
	mapArea := MapArea{
		Area:      make([][]rune, len(mapStr)),
		GuardPos:  make([]int, 2),
		Direction: Up,
		WalkCount: 0,
		LoopCount: 0,
	}
	for i, item := range mapStr {
		mapArea.Area[i] = []rune(item)
		check, pos, value := foundation.ContainsItemInList(mapArea.Area[i], []rune{Up, Down, Left, Right})
		if !check {
			continue
		}
		mapArea.Direction = value
		mapArea.GuardPos[0] = i
		mapArea.GuardPos[1] = pos
		log.Printf("Found guard at [%d, %d]=%v", mapArea.GuardPos[0], mapArea.GuardPos[1], mapArea.Direction)
	}
	// Walk through logic
	if *problem {
		for {
			cury, curx := mapArea.GuardPos[0], mapArea.GuardPos[1]
			endWalk := false
			switch mapArea.Direction {
			case Up:
				if cury-1 < 0 {
					mapArea.Area[cury][curx] = Beenhere
					endWalk = true
					break
				} else if mapArea.Area[cury-1][curx] == Obstacle {
					mapArea.Direction = Right
					continue
				}
				mapArea.GuardPos[0] = cury - 1
				mapArea.Area[cury][curx] = Beenhere
				mapArea.Area[cury-1][curx] = Up
			case Down:
				if cury+1 >= len(mapArea.Area) {
					mapArea.Area[cury][curx] = Beenhere
					endWalk = true
					break
				} else if mapArea.Area[cury+1][curx] == Obstacle {
					mapArea.Direction = Left
					continue
				}
				mapArea.GuardPos[0] = cury + 1
				mapArea.Area[cury][curx] = Beenhere
				mapArea.Area[cury+1][curx] = Down
			case Right:
				if curx+1 >= len(mapArea.Area[cury]) {
					mapArea.Area[cury][curx] = Beenhere
					endWalk = true
					break
				} else if mapArea.Area[cury][curx+1] == Obstacle {
					mapArea.Direction = Down
					continue
				}
				mapArea.GuardPos[1] = curx + 1
				mapArea.Area[cury][curx] = Beenhere
				mapArea.Area[cury][curx+1] = Right
			case Left:
				if curx-1 < 0 {
					mapArea.Area[cury][curx] = Beenhere
					endWalk = true
					break
				} else if mapArea.Area[cury][curx-1] == Obstacle {
					mapArea.Direction = Up
					continue
				}
				mapArea.GuardPos[1] = curx - 1
				mapArea.Area[cury][curx] = Beenhere
				mapArea.Area[cury][curx-1] = Left
			}
			if endWalk {
				break
			}
		}

		mapArea.CountPlacesGuardBeen()
		log.Println(mapArea.PrintArea())
		log.Printf("Places the guard has been=%d", mapArea.WalkCount)

	} else {
		startingPos := []int{mapArea.GuardPos[0], mapArea.GuardPos[1]}
		countBox := 0
		topBlock, rightBlock, downBlock, leftBlock := []int{mapArea.GuardPos[0], mapArea.GuardPos[1]}, []int{mapArea.GuardPos[0], mapArea.GuardPos[1]}, []int{mapArea.GuardPos[0], mapArea.GuardPos[1]}, []int{mapArea.GuardPos[0], mapArea.GuardPos[1]}
		for {
			cury, curx := mapArea.GuardPos[0], mapArea.GuardPos[1]
			endWalk := false
			switch mapArea.Direction {
			case Up:
				if cury-1 < 0 {
					mapArea.Area[cury][curx] = VertMove
					endWalk = true
					break
				} else if mapArea.Area[cury-1][curx] == Obstacle {
					mapArea.Direction = Right
					topBlock[0] = cury
					topBlock[1] = curx
					mapArea.Area[cury][curx] = VertMove
					countBox += 1
					continue
				}
				if countBox == 3 {
					if cury == rightBlock[0] && curx == leftBlock[1] {
						tmp := mapArea.Area[cury-1][curx]
						mapArea.Area[cury-1][curx] = NewObstacle
						mapArea.Area[startingPos[0]][startingPos[1]] = Up
						log.Println(mapArea.PrintArea())
						mapArea.LoopCount += 1
						mapArea.Area[cury-1][curx] = tmp
						countBox = 0
					}
				}
				if mapArea.Area[cury][curx] == Empty {
					mapArea.Area[cury][curx] = VertMove
				} else if mapArea.Area[cury][curx] == HoriMove {
					mapArea.Area[cury][curx] = BothMove
				}
				mapArea.GuardPos[0] = cury - 1
			case Down:
				if cury+1 >= len(mapArea.Area) {
					mapArea.Area[cury][curx] = VertMove
					endWalk = true
					break
				} else if mapArea.Area[cury+1][curx] == Obstacle {
					mapArea.Direction = Left
					downBlock[0] = cury
					downBlock[1] = curx
					mapArea.Area[cury][curx] = VertMove
					countBox += 1
					continue
				}

				if countBox == 3 {
					if cury == leftBlock[0] && curx == rightBlock[1] {
						tmp := mapArea.Area[cury+1][curx]
						mapArea.Area[startingPos[0]][startingPos[1]] = Up
						mapArea.Area[cury+1][curx] = NewObstacle
						log.Println(mapArea.PrintArea())
						mapArea.LoopCount += 1
						mapArea.Area[cury+1][curx] = tmp
						countBox = 0
					}
				}
				if mapArea.Area[cury][curx] == Empty {
					mapArea.Area[cury][curx] = VertMove
				} else if mapArea.Area[cury][curx] == HoriMove {
					mapArea.Area[cury][curx] = BothMove
				}
				mapArea.GuardPos[0] = cury + 1
			case Right:
				if curx+1 >= len(mapArea.Area[cury]) {
					mapArea.Area[cury][curx] = HoriMove
					endWalk = true
					break
				} else if mapArea.Area[cury][curx+1] == Obstacle {
					mapArea.Direction = Down
					rightBlock[0] = cury
					rightBlock[1] = curx
					countBox += 1
					mapArea.Area[cury][curx] = HoriMove
					continue
				}
				if countBox >= 3 {
					if cury == topBlock[0] && curx == downBlock[1] {
						tmp := mapArea.Area[cury][curx+1]
						mapArea.Area[startingPos[0]][startingPos[1]] = Up
						mapArea.Area[cury][curx+1] = NewObstacle
						log.Println(mapArea.PrintArea())
						mapArea.LoopCount += 1
						mapArea.Area[cury][curx+1] = tmp
						countBox = 0
					}
				}
				if mapArea.Area[cury][curx] == Empty {
					mapArea.Area[cury][curx] = HoriMove
				} else if mapArea.Area[cury][curx] == VertMove {
					mapArea.Area[cury][curx] = BothMove
				}
				mapArea.GuardPos[1] = curx + 1
			case Left:
				if curx-1 < 0 {
					mapArea.Area[cury][curx] = HoriMove
					endWalk = true
					break
				} else if mapArea.Area[cury][curx-1] == Obstacle {
					mapArea.Direction = Up
					leftBlock[0] = cury
					leftBlock[1] = curx
					countBox += 1
					mapArea.Area[cury][curx] = HoriMove
					continue
				}
				if countBox == 3 {
					if cury == downBlock[0] && curx == topBlock[1] {
						tmp := mapArea.Area[cury][curx-1]
						mapArea.Area[startingPos[0]][startingPos[1]] = Up
						mapArea.Area[cury][curx-1] = NewObstacle
						log.Println(mapArea.PrintArea())
						mapArea.LoopCount += 1
						mapArea.Area[cury][curx-1] = tmp
						countBox = 0
					}
				}
				if mapArea.Area[cury][curx] == Empty {
					mapArea.Area[cury][curx] = HoriMove
				} else if mapArea.Area[cury][curx] == VertMove {
					mapArea.Area[cury][curx] = BothMove
				}
				mapArea.GuardPos[1] = curx - 1
			}
			if endWalk {
				break
			}
		}
		log.Printf("New obstacles options=%d", mapArea.LoopCount)
	}
}

type MapArea struct {
	Area           [][]rune
	GuardPos       []int
	Direction      rune
	WalkCount      int
	LoopCount      int
	AddedObstacles [][]int
}

func (ma *MapArea) CountPlacesGuardBeen() {
	count := 0
	for _, line := range ma.Area {
		count += foundation.Count(line, Beenhere)
	}
	ma.WalkCount = count
}

func (ma *MapArea) PrintArea() string {
	out := "Final Map area\n------------------------------\n"
	for i, line := range ma.Area {
		out += fmt.Sprintf("Line #%3d %v\n", i, string(line))
	}
	return out
}
