package solutions

import (
	"aoc-2021/util"
	"sort"
)

type grid struct {
	numRows    int
	numColumns int
	numbers    [][]int
	visited    map[string]string
}

func Day9Part1(numbers [][]int) int {
	grid := &grid{
		numbers:    numbers,
		numRows:    len(numbers),
		numColumns: len(numbers[0]),
	}
	lowPoints := grid.findLowPoints()
	var sumRiskFactors int
	for _, p := range lowPoints {
		sumRiskFactors += 1 + grid.numbers[p[0]][p[1]]
	}
	return sumRiskFactors
}

func Day9Part2(numbers [][]int) int {
	grid := &grid{
		numbers:    numbers,
		numRows:    len(numbers),
		numColumns: len(numbers[0]),
		visited:    make(map[string]string),
	}
	lowPoints := grid.findLowPoints()
	basinSizes := []int{}
	for _, p := range lowPoints {
		basinCoords := grid.findBasinCoordinates(p)
		basinCoords = removeDuplicates(basinCoords)
		basinSizes = append(basinSizes, len(basinCoords))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	return basinSizes[0] * basinSizes[1] * basinSizes[2]
}

func (g *grid) findBasinCoordinates(centre []int) [][]int {
	g.visited[util.IntSliceToNumString(centre, ",")] = ""

	neighbours := g.getValidNeighbours(centre[0], centre[1])
	neighbours = g.removeAlreadyVisited(neighbours)
	neighbours = g.removeNines(neighbours)

	points := [][]int{centre}
	for _, p := range neighbours {
		points = append(points, g.findBasinCoordinates(p)...)
	}
	return points
}

func (g *grid) findLowPoints() [][]int {
	lowPoints := [][]int{}
	for i := 0; i < g.numRows; i++ {
		for j := 0; j < g.numColumns; j++ {
			if g.positionLessThanAllOthers([]int{i, j}, g.getValidNeighbours(i, j)) {
				lowPoints = append(lowPoints, []int{i, j})
			}
		}
	}
	return lowPoints
}

func (g *grid) getValidNeighbours(i, j int) [][]int {
	result := [][]int{}
	for _, neighbour := range [][]int{
		{i, j - 1},
		{i, j + 1},
		{i - 1, j},
		{i + 1, j},
	} {
		if g.validGridPosition(neighbour) {
			result = append(result, neighbour)
		}
	}
	return result
}

func (g *grid) validGridPosition(p []int) bool {
	return p[0] > -1 && p[0] < g.numRows && p[1] > -1 && p[1] < g.numColumns
}

func (g *grid) positionLessThanAllOthers(p []int, others [][]int) bool {
	for _, other := range others {
		if !(g.numbers[p[0]][p[1]] < g.numbers[other[0]][other[1]]) {
			return false
		}
	}
	return true
}

func (g *grid) removeAlreadyVisited(points [][]int) [][]int {
	newPoints := [][]int{}
	for _, p := range points {
		key := util.IntSliceToNumString(p, ",")
		if _, ok := g.visited[key]; !ok {
			newPoints = append(newPoints, p)
		}
	}
	return newPoints
}

func (g *grid) removeNines(points [][]int) [][]int {
	newPoints := [][]int{}
	for _, p := range points {
		if g.numbers[p[0]][p[1]] != 9 {
			newPoints = append(newPoints, p)
		}
	}
	return newPoints
}

func removeDuplicates(input [][]int) [][]int {
	result := [][]int{}
	uniques := map[string]string{}
	for _, in := range input {
		if _, ok := uniques[util.IntSliceToNumString(in, ",")]; !ok {
			uniques[util.IntSliceToNumString(in, ",")] = ""
			result = append(result, in)
		}
	}
	return result
}
