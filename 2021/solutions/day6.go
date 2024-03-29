package solutions

import "aoc-2021/util"

func Day6(in string, numDays int) int {
	timers := util.NumStringToIntSlice(in)
	return simulate(timers, numDays)
}

func simulate(timers []int, numDays int) int {
	timerCounts := newTimerCounts()
	for _, t := range timers {
		timerCounts[t] = timerCounts[t] + 1
	}
	day := 0
	for day < numDays {
		newCounts := newTimerCounts()
		for timerValue := 0; timerValue < 9; timerValue++ {
			newCounts[timerValue] = timerCounts[(timerValue+1)%9]
		}
		newCounts[6] += timerCounts[0]
		timerCounts = newCounts
		day++
	}
	var count int
	for _, c := range timerCounts {
		count += c
	}
	return count
}

func newTimerCounts() map[int]int {
	timerCounts := map[int]int{}
	for timerValue := 0; timerValue < 9; timerValue++ {
		timerCounts[timerValue] = 0
	}
	return timerCounts
}
