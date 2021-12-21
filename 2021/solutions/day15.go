package solutions

import (
	"aoc-2021/util"
)

type graph struct {
	nodes map[string]*node
	edges map[string][]*edge
}

type node struct {
	key string
}

type edge struct {
	start  *node
	end    *node
	weight int
}

func Day15Part1(in [][]int) int {
	grid := &grid{
		numbers:    in,
		numRows:    len(in),
		numColumns: len(in[0]),
	}

	graph := &graph{nodes: make(map[string]*node), edges: make(map[string][]*edge)}
	// populate nodes
	for i := 0; i < grid.numRows; i++ {
		for j := 0; j < grid.numColumns; j++ {
			nodeKey := util.IntSliceToNumString([]int{i, j}, ",")
			graph.nodes[nodeKey] = &node{key: nodeKey}
		}
	}

	// populate edges
	for i := 0; i < grid.numRows; i++ {
		for j := 0; j < grid.numColumns; j++ {
			nodeKey := util.IntSliceToNumString([]int{i, j}, ",")
			neighbours := grid.getValidNeighbours(i, j)
			for _, n := range neighbours {
				neighbourKey := util.IntSliceToNumString(n, ",")
				graph.edges[nodeKey] = append(
					graph.edges[nodeKey],
					&edge{
						start:  graph.nodes[nodeKey],
						end:    graph.nodes[neighbourKey],
						weight: grid.numbers[n[0]][n[1]],
					},
				)
			}
		}
	}

	return graph.djikstra(
		util.IntSliceToNumString([]int{0, 0}, ","),
		util.IntSliceToNumString([]int{grid.numRows - 1, grid.numColumns - 1}, ","),
	)
}

func Day15Part2(in [][]int, steps int) int {
	incrementTiles := [][][]int{in}
	for s := 0; s < steps-1; s++ {
		new := incrementAll(incrementTiles[s])
		incrementTiles = append(incrementTiles, new)
	}

	var slices [][][]int
	for s := 0; s < steps; s++ {
		row := incrementTiles[s]
		newTile := incrementAll(row)
		for i := s + 1; i < s+steps; i++ {
			row = append(row, newTile...)
			newTile = incrementAll(newTile)
		}
		slices = append(slices, row)
	}
	return Day15Part1(mergeSlices(slices))
}

func (g *graph) djikstra(source, end string) int {
	unvisitedNodeKeys := make(map[string]string)
	nodeDistances := make(map[string]int)

	for nodeKey := range g.nodes {
		unvisitedNodeKeys[nodeKey] = ""
	}
	nodeDistances[source] = 0
	for len(unvisitedNodeKeys) > 0 {
		node := smallestDistanceUnvistedNode(unvisitedNodeKeys, nodeDistances)
		delete(unvisitedNodeKeys, node)
		for _, edge := range g.edges[node] {
			dist := nodeDistances[node] + edge.weight
			if _, ok := nodeDistances[edge.end.key]; !ok || dist < nodeDistances[edge.end.key] {
				nodeDistances[edge.end.key] = dist
			}
		}
	}
	return nodeDistances[end]
}

func smallestDistanceUnvistedNode(unvisited map[string]string, distances map[string]int) string {
	if len(distances) == 1 {
		for k := range distances {
			return k
		}
	}
	var minDistKey string
	minDist := 999999999999999
	if len(distances) < len(unvisited) {
		for k, v := range distances {
			if _, ok := unvisited[k]; ok && v < minDist {
				minDist = v
				minDistKey = k
			}
		}
	} else {
		for k := range unvisited {
			if val, ok := distances[k]; ok && val < minDist {
				minDist = val
				minDistKey = k
			}
		}
	}
	return minDistKey
}

func incrementAll(in [][]int) [][]int {
	var result [][]int
	for i := 0; i < len(in); i++ {
		result = append(result, []int{})
		for j := 0; j < len(in[i]); j++ {
			result[i] = append(result[i], (in[i][j])%9+1)
		}
	}
	return result
}

func mergeSlices(slices [][][]int) [][]int {
	var result [][]int
	for i := range slices[0] {
		var row []int
		for _, slice := range slices {
			row = append(row, slice[i]...)
		}
		result = append(result, row)
	}
	return result
}
