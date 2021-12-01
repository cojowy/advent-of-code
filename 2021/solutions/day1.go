package solutions

import "aoc-2021/util"

// day 1 part 1
func CountNumberOfIncreases(in []int) int {
	depth := in[0]
	count := 0
	for _, num := range in[1:] {
		if num > depth {
			count++
		}
		depth = num
	}
	return count
}

// day 1 part 2
func CountNumberOfSlidingWindowIncreases(in []int, window int) int {
	depthSum := util.Sum(in[0:window])
	count := 0
	for i := 1; i < len(in)-window+1; i++ {
		currWindowDepth := util.Sum(in[i : i+window])
		if currWindowDepth > depthSum {
			count++
		}
		depthSum = currWindowDepth
	}
	return count
}
