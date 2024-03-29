package main

import (
	"aoc-2021/solutions"
	"aoc-2021/util"
	"fmt"
)

func day1() {
	numbers, err := util.ReadInts("inputs/day1")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Day 1 Part 1: %d\n", solutions.CountNumberOfIncreases(numbers))
	fmt.Printf("Day 1 Part 2: %d\n", solutions.CountNumberOfSlidingWindowIncreases(numbers, 3))
}

func day2() {
	lines, err := util.ReadStrings("inputs/day2")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Day 2 Part 1: %d\n", solutions.Day2Part1(lines))
	fmt.Printf("Day 2 Part 2: %d\n", solutions.Day2Part2(lines))
}

func day3() {
	input, err := util.ReadIntSlices("inputs/day3")
	if err != nil {
		panic(err)
	}
	ans, err := solutions.Day3Part1(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Day 3 Part 1: %d\n", ans)

	ans2, err := solutions.Day3Part2(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Day 3 Part 2: %d\n", ans2)
}

func day4() {
	input, err := util.ReadStrings("inputs/day4")
	if err != nil {
		panic(err)
	}

	ans, err := solutions.Day4Part1(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Day 4 Part 1: %d\n", ans)

	ans2, err := solutions.Day4Part2(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Day 4 Part 2: %d\n", ans2)
}

func day5() {
	input, err := util.ReadStrings("inputs/day5")
	if err != nil {
		panic(err)
	}
	ans := solutions.Day5(input, 1)
	fmt.Printf("Day 5 Part 1: %d\n", ans)
	ans2 := solutions.Day5(input, 2)
	fmt.Printf("Day 5 Part 2: %d\n", ans2)
}

func day6() {
	input, err := util.ReadStrings("inputs/day6")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Day 6 Part 1: %d\n", solutions.Day6(input[0], 80))
	fmt.Printf("Day 6 Part 2: %d\n", solutions.Day6(input[0], 256))
}

func day7() {
	input, err := util.ReadStrings("inputs/day7")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Day 7 Part 1: %d\n", solutions.Day7(input[0], 1))
	fmt.Printf("Day 7 Part 2: %d\n", solutions.Day7(input[0], 2))
}

func day8() {
	input, err := util.ReadStrings("inputs/day8")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Day 8 Part 1: %d\n", solutions.Day8Part1(input))
	fmt.Printf("Day 8 Part 2: %d\n", solutions.Day8Part2(input))
}

func day9() {
	input, err := util.ReadIntSlices("inputs/day9")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Day 9 Part 1: %d\n", solutions.Day9Part1(input))
	fmt.Printf("Day 9 Part 2: %d\n", solutions.Day9Part2(input))
}

func day10() {
	input, err := util.ReadStrings("inputs/day10")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Day 10 Part 1: %d\n", solutions.Day10Part1(input))
	fmt.Printf("Day 10 Part 2: %d\n", solutions.Day10Part2(input))
}

func day11() {
	input, err := util.ReadIntSlices("inputs/day11")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Day 11 Part 1: %d\n", solutions.Day11Part1(input, 100))

	input2, err := util.ReadIntSlices("inputs/day11")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Day 11 Part 2: %d\n", solutions.Day11Part2(input2))
}

func day12() {
	input, err := util.ReadStrings("inputs/day12")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Day 12 Part 1: %d\n", solutions.Day12Part1(input))
	fmt.Printf("Day 12 Part 1: %d\n", solutions.Day12Part2(input))
}

func day13() {
	input, err := util.ReadStrings("inputs/day13")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Day 13 Part 1: %d\n", solutions.Day13Part1(input))
	fmt.Println("Day 13 Part 2:")
	solutions.Day13Part2(input)
}

func day14() {
	input, err := util.ReadStrings("inputs/day14")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Day 14 Part 1: %d\n", solutions.Day14(input, 10))
	fmt.Printf("Day 14 Part 1: %d\n", solutions.Day14(input, 40))
}

func day15() {
	input, err := util.ReadIntSlices("inputs/day15")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Day 15 Part 1: %d\n", solutions.Day15Part1(input))
	fmt.Printf("Day 15 Part 2: %d\n", solutions.Day15Part2(input, 5))
}

func main() {
	day15()
}
