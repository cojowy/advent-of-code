package solutions

import (
	"strconv"
	"strings"
)

const (
	forward = "forward"
	up      = "up"
	down    = "down"
)

type position struct {
	depth      int
	horizontal int
	aim        int
}

func Day2Part1(lines []string) int {
	p := position{}
	p.move(lines)
	return p.depth * p.horizontal
}

func Day2Part2(lines []string) int {
	p := position{}
	p.moveWithAim(lines)
	return p.depth * p.horizontal
}

// day 2 part 1
func (p *position) move(lines []string) {
	for _, line := range lines {
		instruction := strings.Split(line, " ")
		command, numberRaw := instruction[0], instruction[1]
		number, _ := strconv.Atoi(numberRaw)
		switch command {
		case forward:
			p.horizontal += number
		case up:
			p.depth -= number
		case down:
			p.depth += number
		}
	}
}

// day 2 part 2
func (p *position) moveWithAim(lines []string) {
	for _, line := range lines {
		instruction := strings.Split(line, " ")
		command, numberRaw := instruction[0], instruction[1]
		number, _ := strconv.Atoi(numberRaw)
		switch command {
		case forward:
			p.horizontal += number
			p.depth += p.aim * number
		case up:
			p.aim -= number
		case down:
			p.aim += number
		}
	}
}
