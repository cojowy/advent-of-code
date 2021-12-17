package solutions

import (
	"strings"
)

func Day14(in []string, steps int) int {
	template, rules := parseDay14Input(in)
	pairCounts, charCounts := make(map[string]int), make(map[string]int)

	chars := strings.Split(template, "")
	for i := 0; i < len(chars)-1; i++ {
		pair := chars[i] + chars[i+1]
		incrementCount(pairCounts, pair, 1)
		incrementCount(charCounts, chars[i], 1)
	}
	incrementCount(charCounts, chars[len(chars)-1], 1)

	for i := 0; i < steps; i++ {
		toMerge := []map[string]int{}
		for k, v := range pairCounts {
			chars := strings.Split(k, "")
			insertionChar := rules[chars[0]+chars[1]]
			toMerge = append(
				toMerge,
				map[string]int{
					chars[0] + insertionChar: v,
					insertionChar + chars[1]: v,
				},
			)
			incrementCount(charCounts, insertionChar, v)
		}
		pairCounts = merge(toMerge)
	}

	var maxCount, minCount int
	for _, v := range charCounts {
		if v > maxCount {
			maxCount = v
		}
		if minCount == 0 || v < minCount {
			minCount = v
		}
	}
	return maxCount - minCount
}

func incrementCount(counts map[string]int, key string, increment int) {
	if _, ok := counts[key]; !ok {
		counts[key] = increment
	} else {
		counts[key] += increment
	}
}

func merge(maps []map[string]int) map[string]int {
	result := make(map[string]int)
	for _, m := range maps {
		for k, v := range m {
			if _, ok := result[k]; !ok {
				result[k] = v
			} else {
				result[k] += v
			}
		}
	}
	return result
}

func parseDay14Input(in []string) (string, map[string]string) {
	template := in[0]
	pairInsertionRules := make(map[string]string)
	for _, line := range in[2:] {
		fields := strings.Fields(line)
		pairInsertionRules[fields[0]] = fields[2]
	}
	return template, pairInsertionRules
}

func naiveComputeNewTemplate(template string, rules map[string]string) string {
	var newTemplate string
	chars := strings.Split(template, "")
	for i := 0; i < len(chars)-1; i++ {
		insertionChar := rules[chars[i]+chars[i+1]]
		newTemplate += chars[i] + insertionChar
	}
	newTemplate += chars[len(chars)-1]
	return newTemplate
}
