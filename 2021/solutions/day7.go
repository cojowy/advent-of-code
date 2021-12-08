package solutions

import (
	"aoc-2021/util"
	"math"
	"sort"
)

func Day7(in string, part int) int {
	positions := util.NumStringToIntSlice(in)
	sort.Ints(positions)
	minFuelConsumption := calculateFuel(positions, positions[0], part)
	for i := positions[1]; i < positions[len(positions)-1]; i++ {
		fuelConsumption := calculateFuel(positions, i, part)
		if fuelConsumption < minFuelConsumption {
			minFuelConsumption = fuelConsumption
		}
	}
	return minFuelConsumption
}

func calculateFuel(positions []int, candidate int, part int) int {
	var sum int
	for _, p := range positions {
		if part == 1 {
			sum += int(math.Abs(float64(p - candidate)))
		}
		if part == 2 {
			sum += cumulativeFuel(int(math.Abs(float64(p - candidate))))
		}
	}
	return sum
}

func cumulativeFuel(distance int) int {
	var sum int
	for i := 1; i < distance+1; i++ {
		sum += i
	}
	return sum
}
