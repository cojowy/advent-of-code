package main

import (
	"aoc-2021/solutions"
	"aoc-2021/util"
	"fmt"
)

func main() {
	day1()
}

func day1() {
	numbers, err := util.ReadInts("inputs/day1")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Day 1 Part 1: %d\n", solutions.CountNumberOfIncreases(numbers))
	fmt.Printf("Day 1 Part 2: %d\n", solutions.CountNumberOfSlidingWindowIncreases(numbers, 3))
}
