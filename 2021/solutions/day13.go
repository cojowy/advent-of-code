package solutions

import (
	"aoc-2021/util"
	"fmt"
	"strconv"
	"strings"
)

func Day13Part1(in []string) int {
	var dots [][]int
	dotPositions, foldInputs := parseInstructions(in)
	firstInstruction := foldInputs[0]
	if val, ok := firstInstruction["y"]; ok {
		dots = foldOnY(dotPositions, val)
	}
	if val, ok := firstInstruction["x"]; ok {
		dots = foldOnX(dotPositions, val)
	}
	dots = util.RemoveDuplicates(dots)
	return len(dots)
}

func Day13Part2(in []string) {
	dots, foldInputs := parseInstructions(in)
	for _, instruction := range foldInputs {
		if val, ok := instruction["y"]; ok {
			dots = foldOnY(dots, val)
		}
		if val, ok := instruction["x"]; ok {
			dots = foldOnX(dots, val)
		}
	}
	dots = util.RemoveDuplicates(dots)
	viewDots(dots)
}

func foldOnX(dots [][]int, x int) [][]int {
	for i := range dots {
		if dots[i][0] > x {
			dots[i][0] = x - (dots[i][0] - x)
		}
	}
	return dots
}

func foldOnY(dots [][]int, y int) [][]int {
	for i := range dots {
		if dots[i][1] > y {
			dots[i][1] = y - (dots[i][1] - y)
		}
	}
	return dots
}

func maxX(dots [][]int) int {
	max := dots[0][0]
	for _, dot := range dots[1:] {
		if dot[0] > max {
			max = dot[0]
		}
	}
	return max
}

func maxY(dots [][]int) int {
	max := dots[0][1]
	for _, dot := range dots[1:] {
		if dot[1] > max {
			max = dot[1]
		}
	}
	return max
}

func viewDots(dots [][]int) {
	sizeX, sizeY := maxX(dots), maxY(dots)
	dotMap := make(map[string]string)
	for _, dot := range dots {
		dotMap[util.IntSliceToNumString(dot, ",")] = ""
	}

	for y := 0; y < sizeY+1; y++ {
		for x := 0; x < sizeX+1; x++ {
			if _, ok := dotMap[util.IntSliceToNumString([]int{x, y}, ",")]; ok {
				fmt.Print("# ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func parseInstructions(in []string) ([][]int, []map[string]int) {
	var dotPositions [][]int
	var foldInstructions []map[string]int
	for _, line := range in {
		if strings.Contains(line, "=") {
			fields := strings.Split(strings.Fields(line)[2], "=")
			number, _ := strconv.Atoi(fields[1])
			foldInstructions = append(foldInstructions, map[string]int{fields[0]: number})
		} else {
			if strings.Contains(line, ",") {
				points := strings.Split(line, ",")
				x, _ := strconv.Atoi(points[0])
				y, _ := strconv.Atoi(points[1])
				dotPositions = append(dotPositions, []int{x, y})
			}
		}
	}
	return dotPositions, foldInstructions
}
