package solutions

import (
	"fmt"
	"strconv"
	"strings"

	color "github.com/TwiN/go-color"
)

type bingoCard struct {
	numbers map[int]*bingoNumber
	grid    [][]int
	size    int
}

type bingoNumber struct {
	value    int
	called   bool
	position *coordinate
}

type coordinate struct {
	i, j int
}

func Day4Part1(in []string) (int, error) {
	calledNumbers, bingoCards, err := parseInput(in)
	if err != nil {
		return 0, err
	}

	for i, number := range calledNumbers {
		for _, card := range bingoCards {
			position := card.callNumber(number)
			if i > card.size-2 && position != nil && card.positionCausedBingo(position) {
				return number * card.sumOfUncalledNumbers(), nil
			}
		}
	}
	return 0, nil
}

func Day4Part2(in []string) (int, error) {
	calledNumbers, bingoCards, err := parseInput(in)
	if err != nil {
		return 0, err
	}

	for _, number := range calledNumbers {
		var i int
		for i < len(bingoCards) {
			currentCard := bingoCards[i]
			position := currentCard.callNumber(number)
			if position != nil && currentCard.positionCausedBingo(position) {
				bingoCards = append(bingoCards[:i], bingoCards[i+1:]...)
				if len(bingoCards) == 0 {
					return number * currentCard.sumOfUncalledNumbers(), nil
				}
				i = 0
			} else {
				i++
			}
		}
	}
	return 0, nil
}

func newBingoCard(grid [][]int) *bingoCard {
	bc := &bingoCard{grid: grid, size: len(grid), numbers: make(map[int]*bingoNumber)}
	for i, row := range grid {
		for j, number := range row {
			bc.numbers[number] = &bingoNumber{
				value:  number,
				called: false,
				position: &coordinate{
					i: i,
					j: j,
				},
			}
		}
	}
	return bc
}

func (c *bingoCard) callNumber(n int) *coordinate {
	if _, ok := c.numbers[n]; ok {
		c.numbers[n].called = true
		return c.numbers[n].position
	}
	return nil
}

func (c *bingoCard) bingoNumbersInRow(n int) []*bingoNumber {
	var result []*bingoNumber
	for _, num := range c.grid[n] {
		result = append(result, c.numbers[num])
	}
	return result
}

func (c *bingoCard) bingoNumbersInColumn(n int) []*bingoNumber {
	var result []*bingoNumber
	for i := range c.grid {
		result = append(result, c.numbers[c.grid[i][n]])
	}
	return result
}

func (c *bingoCard) sumOfUncalledNumbers() int {
	var sum int
	for val, num := range c.numbers {
		if !num.called {
			sum += val
		}
	}
	return sum
}

func (c *bingoCard) positionCausedBingo(p *coordinate) bool {
	return allCalled(c.bingoNumbersInRow(p.i)) ||
		allCalled(c.bingoNumbersInColumn(p.j))
}

func (c *bingoCard) view() {
	for _, row := range c.grid {
		for _, number := range row {
			if c.numbers[number].called {
				fmt.Print(color.Ize(color.Green, strconv.Itoa(number)))
				fmt.Print(" ")
			} else {
				fmt.Print(number, " ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func allCalled(bingoNumbers []*bingoNumber) bool {
	for _, n := range bingoNumbers {
		if !n.called {
			return false
		}
	}
	return true
}

func parseInput(in []string) ([]int, []*bingoCard, error) {
	calledNumbersRaw := strings.Split(in[0], ",")
	var calledNumbers []int
	for _, rawNum := range calledNumbersRaw {
		num, err := strconv.Atoi(rawNum)
		if err != nil {
			return nil, nil, err
		}
		calledNumbers = append(calledNumbers, num)
	}

	var bingoCards []*bingoCard
	var currentCardInput [][]int
	for _, line := range in[2:] {
		if line == "" {
			bingoCards = append(bingoCards, newBingoCard(currentCardInput))
			currentCardInput = [][]int{}
			continue
		} else {
			rawNums := strings.Fields(line)
			var nums []int
			for _, rawNum := range rawNums {
				num, err := strconv.Atoi(rawNum)
				if err != nil {
					return nil, nil, err
				}
				nums = append(nums, num)
			}
			currentCardInput = append(currentCardInput, nums)
		}
	}
	bingoCards = append(bingoCards, newBingoCard(currentCardInput))
	return calledNumbers, bingoCards, nil
}
