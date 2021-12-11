package solutions

import (
	"sort"
)

var (
	openCloseMapping = map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}
	illegalCharacterPoints = map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	autocompletePoints = map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
)

func Day10Part1(in []string) int {
	var score int
	for _, line := range in {
		illegalChar := findIllegalCharacter(line)
		if illegalChar != -1 {
			score += illegalCharacterPoints[illegalChar]
		}
	}
	return score
}

func Day10Part2(in []string) int {
	var scores []int
	for _, line := range in {
		illegalChar := findIllegalCharacter(line)
		if illegalChar == -1 {
			scores = append(scores, autocompleteScore(line))
		}
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func findIllegalCharacter(in string) rune {
	stack := []rune{}
	for _, char := range in {
		if _, ok := openCloseMapping[char]; ok {
			stack = append(stack, char)
		} else {
			if openCloseMapping[stack[len(stack)-1]] != char {
				return char
			}
			stack = stack[:len(stack)-1]
		}
	}
	return -1
}

func autocompleteScore(in string) int {
	stack := []rune{}
	for _, char := range in {
		if _, ok := openCloseMapping[char]; ok {
			stack = append(stack, char)
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	var score int
	for len(stack) > 0 {
		score = score*5 + autocompletePoints[openCloseMapping[stack[len(stack)-1]]]
		stack = stack[:len(stack)-1]
	}
	return score
}
