package main

import (
	"aoc-2021/solutions"
	"aoc-2021/util"
	"fmt"
)

func main() {
	day6()
}

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
