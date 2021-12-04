package solutions

import (
	"aoc-2021/util"
	"strconv"
)

func Day3Part1(input [][]int) (int, error) {
	g, e := computeGammaAndEpsilonBits(input)
	gDecimal, err := strconv.ParseInt(g, 2, 64)
	if err != nil {
		return 0, err
	}
	eDecimal, err := strconv.ParseInt(e, 2, 64)
	if err != nil {
		return 0, err
	}
	return int(gDecimal) * int(eDecimal), nil
}

func Day3Part2(input [][]int) (int, error) {
	ogr := util.IntSliceToNumString(oxygenGeneratorRatingValue(input, 0))
	csr := util.IntSliceToNumString(co2ScrubberRatingValue(input, 0))

	ogrDecimal, err := strconv.ParseInt(ogr, 2, 64)
	if err != nil {
		return 0, err
	}
	csrDecimal, err := strconv.ParseInt(csr, 2, 64)
	if err != nil {
		return 0, err
	}

	return int(ogrDecimal) * int(csrDecimal), nil
}

func computeGammaAndEpsilonBits(input [][]int) (string, string) {
	cols := util.Transpose(input)
	gammaBits, epsilonBits := "", ""
	for _, c := range cols {
		bitCounts := util.IntCounts(c)
		gammaBits += strconv.Itoa(mostFrequentBit(bitCounts))
		epsilonBits += strconv.Itoa(leastFrequentBit(bitCounts))
	}
	return gammaBits, epsilonBits
}

func oxygenGeneratorRatingValue(input [][]int, bitIndex int) []int {
	var selected [][]int
	cols := util.Transpose(input)
	colCount := util.IntCounts(cols[bitIndex])
	mf := mostFrequentBit(colCount)
	for _, line := range input {
		if line[bitIndex] == mf {
			selected = append(selected, line)
		}
	}
	if len(selected) == 1 {
		return selected[0]
	}
	return oxygenGeneratorRatingValue(selected, bitIndex+1)
}

func co2ScrubberRatingValue(input [][]int, bitIndex int) []int {
	var selected [][]int
	cols := util.Transpose(input)
	colCount := util.IntCounts(cols[bitIndex])
	lf := leastFrequentBit(colCount)
	for _, line := range input {
		if line[bitIndex] == lf {
			selected = append(selected, line)
		}
	}
	if len(selected) == 1 {
		return selected[0]
	}
	return co2ScrubberRatingValue(selected, bitIndex+1)
}

func mostFrequentBit(in map[int]int) int {
	if in[0] > in[1] {
		return 0
	}
	return 1
}

func leastFrequentBit(in map[int]int) int {
	if in[1] < in[0] {
		return 1
	}
	return 0
}
