package util

import (
	"bufio"
	"os"
	"strconv"
)

// ReadInts reads newline-separated integers from a file into an []int
func ReadInts(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	return numbers, scanner.Err()
}

// ReadStrings reads lines from a file into a []string
func ReadStrings(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func ReadIntSlices(path string) ([][]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var numbers [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rawNumber := scanner.Text()
		b := make([]int, len(rawNumber))
		for i, v := range rawNumber {
			b[i] = int(v - '0')
		}
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, b)
	}
	return numbers, scanner.Err()
}

func Sum(in []int) int {
	sum := in[0]
	for _, num := range in[1:] {
		sum += num
	}
	return sum
}

// Transpose turns an n*m 2d integer slice into an m*n one
func Transpose(input [][]int) [][]int {
	n, m := len(input[0]), len(input)
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			result[i][j] = input[j][i]
		}
	}
	return result
}

// IntSliceToNumString converts a []int to string
func IntSliceToNumString(in []int) string {
	var result string
	for _, num := range in {
		result += strconv.Itoa(num)
	}
	return result
}

// IntCounts returns {int: frequency count} for []int
func IntCounts(in []int) map[int]int {
	elemCounts := map[int]int{}
	for _, elem := range in {
		if _, ok := elemCounts[elem]; ok {
			elemCounts[elem]++
		} else {
			elemCounts[elem] = 1
		}
	}
	return elemCounts
}
