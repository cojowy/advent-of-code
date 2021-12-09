package solutions

import (
	"sort"
	"strconv"
	"strings"
)

func Day8Part1(in []string) int {
	var count int
	for _, line := range in {
		for _, val := range strings.Fields(strings.Split(line, "|")[1]) {
			if len(val) == 2 || len(val) == 4 || len(val) == 3 || len(val) == 7 {
				count++
			}
		}
	}
	return count
}

func Day8Part2(in []string) int {
	var result int
	for _, line := range in {
		patternMapping := map[int]string{}
		patterns := uniquePatterns(strings.Fields(strings.Join(strings.Split(line, "|"), " ")))
		populateKnownLengthPatterns(patterns, patternMapping)
		populateLengthFivePatterns(patternsWithLength(patterns, 5), patternMapping)
		populateLengthSixPatterns(patternsWithLength(patterns, 6), patternMapping)

		reverseMap := map[string]int{}
		for k, v := range patternMapping {
			reverseMap[v] = k
		}
		resultPatterns := strings.Fields(strings.Split(line, "|")[1])
		lineResultString := ""
		for _, pattern := range resultPatterns {
			lineResultString += strconv.Itoa(reverseMap[sortPattern(pattern)])
		}
		lineResult, _ := strconv.Atoi(lineResultString)
		result += lineResult
	}
	return result
}

func populateKnownLengthPatterns(patterns map[string]string, mapping map[int]string) {
	for pattern := range patterns {
		switch len(pattern) {
		case 2:
			mapping[1] = pattern
		case 4:
			mapping[4] = pattern
		case 3:
			mapping[7] = pattern
		case 7:
			mapping[8] = pattern
		}
		if len(mapping) == 4 {
			return
		}
	}
}

func populateLengthFivePatterns(patterns []string, mapping map[int]string) {
	var lettersInOnePattern []rune
	for _, r := range mapping[1] {
		lettersInOnePattern = append(lettersInOnePattern, r)
	}
	var i int
	for i < len(patterns) {
		// 3 contains both letters from 1
		if _, ok := mapping[3]; !ok && nLettersInPattern(lettersInOnePattern, patterns[i], 2) {
			mapping[3] = patterns[i]
			patterns = append(patterns[:i], patterns[i+1:]...)
		}
		i++
	}

	var lettersInFourPattern []rune
	for _, r := range mapping[4] {
		lettersInFourPattern = append(lettersInFourPattern, r)
	}
	i = 0
	for i < len(patterns) {
		// 2 contains 2 of 4's letters
		if nLettersInPattern(lettersInFourPattern, patterns[i], 2) {
			mapping[2] = patterns[i]
			patterns = append(patterns[:i], patterns[i+1:]...)
		}
		i++
	}
	// 2 is only 5 letter pattern left
	mapping[5] = patterns[0]
}

func populateLengthSixPatterns(patterns []string, mapping map[int]string) {
	var lettersInOnePattern []rune
	for _, r := range mapping[1] {
		lettersInOnePattern = append(lettersInOnePattern, r)
	}
	var i int
	for i < len(patterns) {
		// 6 contains only 1 letter from 1
		if nLettersInPattern(lettersInOnePattern, patterns[i], 1) {
			mapping[6] = patterns[i]
			patterns = append(patterns[:i], patterns[i+1:]...)
		}
		i++
	}

	var lettersInFivePattern []rune
	for _, r := range mapping[5] {
		lettersInFivePattern = append(lettersInFivePattern, r)
	}
	i = 0
	for i < len(patterns) {
		// 9 contains 5 letters from 5
		if nLettersInPattern(lettersInFivePattern, patterns[i], 5) {
			mapping[9] = patterns[i]
			patterns = append(patterns[:i], patterns[i+1:]...)
		}
		i++
	}

	// 0 is the only 6-letter pattern left
	mapping[0] = patterns[0]
}

func uniquePatterns(patterns []string) map[string]string {
	sortedPatterns := map[string]string{}
	for _, pattern := range patterns {
		sortedPattern := sortPattern(pattern)
		if _, ok := sortedPatterns[sortedPattern]; !ok {
			sortedPatterns[sortedPattern] = ""
		}
	}
	return sortedPatterns
}

func sortPattern(pattern string) string {
	intLetterValues := []int{}
	for _, letter := range pattern {
		intLetterValues = append(intLetterValues, int(letter))
	}
	sort.Ints(intLetterValues)
	sortedRunes := []rune{}
	for _, intValue := range intLetterValues {
		sortedRunes = append(sortedRunes, rune(intValue))
	}
	return string(sortedRunes)
}

func patternsWithLength(patterns map[string]string, n int) []string {
	var result []string
	for pattern := range patterns {
		if len(pattern) == n {
			result = append(result, pattern)
		}
	}
	return result
}

func nLettersInPattern(letters []rune, pattern string, n int) bool {
	var matches int
	letterMap := map[rune]string{}
	for _, l := range pattern {
		letterMap[l] = ""
	}
	for _, letter := range letters {
		if _, ok := letterMap[letter]; ok {
			matches++
		}
	}
	return matches == n
}
