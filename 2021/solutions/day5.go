package solutions

import (
	"strconv"
	"strings"
)

func Day5(in []string, part int) int {
	lines := parseDay5Input(in)
	field := make(map[coordinate]int)
	for _, line := range lines {
		line.populatePoints(part == 1)
		for _, p := range line.points {
			if _, ok := field[*p]; ok {
				field[*p]++
			} else {
				field[*p] = 1
			}
		}
	}
	var count int
	for _, v := range field {
		if v > 1 {
			count++
		}
	}
	return count
}

type line struct {
	start, end *coordinate
	points     []*coordinate
}

func (l *line) populatePoints(excludeDiagonals bool) {
	// north/south
	if l.start.i == l.end.i {
		if l.start.j < l.end.j {
			// north
			for j := l.start.j; j < l.end.j+1; j++ {
				l.points = append(l.points, &coordinate{i: l.start.i, j: j})
			}
		}
		if l.start.j > l.end.j {
			//south
			for j := l.start.j; j > l.end.j-1; j-- {
				l.points = append(l.points, &coordinate{i: l.start.i, j: j})
			}
		}
	}

	// east/west
	if l.start.j == l.end.j {
		if l.start.i < l.end.i {
			// east
			for i := l.start.i; i < l.end.i+1; i++ {
				l.points = append(l.points, &coordinate{i: i, j: l.start.j})
			}
		}
		if l.start.i > l.end.i {
			// west
			for i := l.start.i; i > l.end.i-1; i-- {
				l.points = append(l.points, &coordinate{i: i, j: l.start.j})
			}
		}
	}

	if excludeDiagonals {
		return
	}

	// north-east
	if l.start.i < l.end.i && l.start.j < l.end.j {
		for i, j := l.start.i, l.start.j; i < l.end.i+1 && j < l.end.j+1; i, j = i+1, j+1 {
			l.points = append(l.points, &coordinate{i: i, j: j})
		}
	}

	// south-east
	if l.start.i < l.end.i && l.start.j > l.end.j {
		for i, j := l.start.i, l.start.j; i < l.end.i+1 && j > l.end.j-1; i, j = i+1, j-1 {
			l.points = append(l.points, &coordinate{i: i, j: j})
		}
	}

	// south-west
	if l.start.i > l.end.i && l.start.j > l.end.j {
		for i, j := l.start.i, l.start.j; i > l.end.i-1 && j > l.end.j-1; i, j = i-1, j-1 {
			l.points = append(l.points, &coordinate{i: i, j: j})
		}
	}

	// north-west
	if l.start.i > l.end.i && l.start.j < l.end.j {
		for i, j := l.start.i, l.start.j; i > l.end.i-1 && j < l.end.j+1; i, j = i-1, j+1 {
			l.points = append(l.points, &coordinate{i: i, j: j})
		}
	}
}

func parseDay5Input(in []string) []*line {
	var lines []*line
	for _, inputLine := range in {
		p1, p2 := pointsFromInputLine(inputLine)
		line := &line{start: p1, end: p2}
		lines = append(lines, line)
	}
	return lines
}

func pointsFromInputLine(line string) (*coordinate, *coordinate) {
	strFields := strings.Fields(line)
	p1Raw, p2Raw := strings.Split(strFields[0], ","), strings.Split(strFields[2], ",")
	p1x, _ := strconv.Atoi(p1Raw[0])
	p1y, _ := strconv.Atoi(p1Raw[1])
	p2x, _ := strconv.Atoi(p2Raw[0])
	p2y, _ := strconv.Atoi(p2Raw[1])
	return &coordinate{i: p1x, j: p1y}, &coordinate{i: p2x, j: p2y}
}
