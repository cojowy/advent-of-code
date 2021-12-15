package solutions

import (
	"strings"
)

func Day12Part1(in []string) int {
	graph := populateGraph(in)
	return len(findRoutes(graph, []string{}, "start", "end"))
}

func Day12Part2(in []string) int {
	graph := populateGraph(in)
	return len(findRoutesSingleSmallCaveTwice(graph, []string{}, "start", "end"))
}

func findRoutes(graph map[string][]string, visited []string, start, end string) [][]string {
	var routes [][]string
	visited = append(visited, start)
	eligibleNeighbours := findEligibleNeighbours(graph, visited, start)
	for _, neighbour := range eligibleNeighbours {
		if neighbour == end {
			routes = append(routes, visited)
		} else {
			routes = append(routes, findRoutes(graph, visited, neighbour, end)...)
		}
	}
	return routes
}

func findEligibleNeighbours(graph map[string][]string, visited []string, key string) []string {
	var eligibles []string
	for _, neighbour := range graph[key] {
		if strings.ToUpper(neighbour) == neighbour || !stringSliceContainsString(visited, neighbour) {
			eligibles = append(eligibles, neighbour)
		}
	}
	return eligibles
}

func findRoutesSingleSmallCaveTwice(graph map[string][]string, visited []string, start, end string) [][]string {
	var routes [][]string
	visited = append(visited, start)
	eligibleNeighbours := findEligibleNeighboursSingleSmallCaveTwice(graph, visited, start)
	for _, neighbour := range eligibleNeighbours {
		if neighbour == end {
			routes = append(routes, visited)
		} else {
			routes = append(routes, findRoutesSingleSmallCaveTwice(graph, visited, neighbour, end)...)
		}
	}
	return routes
}

func findEligibleNeighboursSingleSmallCaveTwice(graph map[string][]string, visited []string, key string) []string {
	var eligibles []string
	for _, neighbour := range graph[key] {
		if uppercase(neighbour) || !stringSliceContainsString(visited, neighbour) || canRevisit(visited, neighbour) {
			eligibles = append(eligibles, neighbour)
		}
	}
	return eligibles
}

func canRevisit(visited []string, cave string) bool {
	if cave == "start" || cave == "end" {
		return false
	}
	return !anySmallCaveVisitedTwice(visited)
}

func anySmallCaveVisitedTwice(visited []string) bool {
	counts := make(map[string]int)
	for _, v := range visited {
		if !uppercase(v) {
			if _, ok := counts[v]; !ok {
				counts[v] = 1
			} else {
				return true
			}
		}
	}
	return false
}

func uppercase(s string) bool {
	return strings.ToUpper(s) == s
}

func stringSliceContainsString(slice []string, str string) bool {
	for _, s := range slice {
		if str == s {
			return true
		}
	}
	return false
}

func populateGraph(in []string) map[string][]string {
	graph := make(map[string][]string)
	for _, line := range in {
		fields := strings.Split(line, "-")
		if _, ok := graph[fields[0]]; !ok {
			graph[fields[0]] = []string{fields[1]}
		} else {
			graph[fields[0]] = append(graph[fields[0]], fields[1])
		}
		if _, ok := graph[fields[1]]; !ok {
			graph[fields[1]] = []string{fields[0]}
		} else {
			graph[fields[1]] = append(graph[fields[1]], fields[0])
		}
	}
	return graph
}
