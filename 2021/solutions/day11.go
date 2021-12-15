package solutions

import (
	"fmt"
	"strconv"

	"github.com/TwiN/go-color"
)

func Day11Part1(in [][]int, steps int) int {
	grid := &grid{
		numbers:    in,
		numRows:    len(in),
		numColumns: len(in[0]),
	}
	var flashes int
	for i := 0; i < steps; i++ {
		grid.step()
		toFlash := grid.findPositionsToFlash()
		for len(toFlash) > 0 {
			flashes += len(toFlash)
			grid.setPositionsToZero(toFlash)
			grid.increaseNeighbourValues(toFlash)
			toFlash = grid.findPositionsToFlash()
		}
		if grid.countZeroes() == 100 {
			fmt.Println(i)
		}
	}
	return flashes
}

func Day11Part2(in [][]int) int {
	grid := &grid{
		numbers:    in,
		numRows:    len(in),
		numColumns: len(in[0]),
	}

	var steps int
	for {
		grid.step()
		toFlash := grid.findPositionsToFlash()
		for len(toFlash) > 0 {
			grid.setPositionsToZero(toFlash)
			grid.increaseNeighbourValues(toFlash)
			toFlash = grid.findPositionsToFlash()
		}
		if grid.countZeroes() == grid.numColumns*grid.numRows {
			return steps + 1
		}
		steps++
	}
}

func (g *grid) step() [][]int {
	var toFlash [][]int
	for i := 0; i < g.numRows; i++ {
		for j := 0; j < g.numColumns; j++ {
			g.numbers[i][j]++
			if g.numbers[i][j] > 9 {
				toFlash = append(toFlash, []int{i, j})
			}
		}
	}
	return toFlash
}

func (g *grid) getValidNeighboursWithDiagonals(i, j int) [][]int {
	result := [][]int{}
	for _, neighbour := range [][]int{
		{i, j - 1},
		{i, j + 1},
		{i - 1, j},
		{i + 1, j},
		{i + 1, j + 1},
		{i + 1, j - 1},
		{i - 1, j + 1},
		{i - 1, j - 1},
	} {
		if g.validGridPosition(neighbour) {
			result = append(result, neighbour)
		}
	}
	return result
}

func (g *grid) findPositionsToFlash() [][]int {
	var toFlash [][]int
	for i := 0; i < g.numRows; i++ {
		for j := 0; j < g.numColumns; j++ {
			if g.numbers[i][j] > 9 {
				toFlash = append(toFlash, []int{i, j})
			}
		}
	}
	return toFlash
}

func (g *grid) setPositionsToZero(positions [][]int) {
	for _, position := range positions {
		g.numbers[position[0]][position[1]] = 0
	}
}

func (g *grid) increaseNeighbourValues(positions [][]int) {
	for _, position := range positions {
		for _, neighbour := range g.getValidNeighboursWithDiagonals(position[0], position[1]) {
			if g.numbers[neighbour[0]][neighbour[1]] > 0 {
				g.numbers[neighbour[0]][neighbour[1]]++
			}
		}
	}
}

func (g *grid) viewFlashes() {
	for i := 0; i < g.numRows; i++ {
		for j := 0; j < g.numColumns; j++ {
			if g.numbers[i][j] == 0 {
				fmt.Print(color.Ize(color.Green, strconv.Itoa(g.numbers[i][j])), " ")
			} else {
				fmt.Print(strconv.Itoa(g.numbers[i][j]), " ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g *grid) countZeroes() int {
	var count int
	for i := 0; i < g.numRows; i++ {
		for j := 0; j < g.numColumns; j++ {
			if g.numbers[i][j] == 0 {
				count++
			}
		}
	}
	return count
}
