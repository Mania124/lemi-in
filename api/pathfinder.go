package api

import (
	"container/list"
	"fmt"
	"sort"
	"strings"
)

// Finds the shortest path using BFS (no global visited)
func BFS(graph map[int][]int, start, end int) []int {
	queue := list.New()
	queue.PushBack([]int{start})

	for queue.Len() > 0 {
		currentPath := queue.Remove(queue.Front()).([]int)
		lastNode := currentPath[len(currentPath)-1]

		if lastNode == end {
			return currentPath
		}

		// Explore neighbors not already in the current path (prevents cycles)
		for _, neighbor := range graph[lastNode] {
			if !contains(currentPath, neighbor) {
				newPath := append([]int{}, currentPath...)
				newPath = append(newPath, neighbor)
				queue.PushBack(newPath)
			}
		}
	}
	return nil
}

// Finds all non-overlapping paths by blocking nodes in found paths
func FindAllPaths(graph map[int][]int, start, end int) [][]int {
	var paths [][]int
	tempGraph := copyGraph(graph) // Create a modifiable copy of the graph

	for {
		path := BFS(tempGraph, start, end)
		if path == nil {
			break // No more paths
		}
		paths = append(paths, path)

		// Block intermediate nodes in this path (node-disjoint paths)
		for i := 1; i < len(path)-1; i++ {
			nodeToBlock := path[i]
			delete(tempGraph, nodeToBlock) // Remove the node from the graph

			// Remove references to the blocked node in other nodes' connections
			for u := range tempGraph {
				var newNeighbors []int
				for _, v := range tempGraph[u] {
					if v != nodeToBlock {
						newNeighbors = append(newNeighbors, v)
					}
				}
				tempGraph[u] = newNeighbors
			}
		}
	}
	return paths
}

// Helper to deep-copy the graph
func copyGraph(original map[int][]int) map[int][]int {
	copied := make(map[int][]int)
	for u, neighbors := range original {
		copied[u] = append([]int{}, neighbors...)
	}
	return copied
}

// Helper function to check if a value exists in a slice
func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// Distributes ants optimally across paths
func DistributeAnts(paths [][]int, numAnts int) map[int][]int {
	// Sort paths by length (ascending)
	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})

	antAssignments := make(map[int][]int)
	pathLoads := make([]int, len(paths))

	for antID := 1; antID <= numAnts; antID++ {
		bestPath := 0
		minCost := len(paths[0]) + pathLoads[0]

		// Find the path with the lowest cost (length + current load)
		for i := 1; i < len(paths); i++ {
			cost := len(paths[i]) + pathLoads[i]
			if cost < minCost {
				minCost = cost
				bestPath = i
			}
		}

		antAssignments[antID] = paths[bestPath]
		pathLoads[bestPath]++
	}

	return antAssignments
}

// Simulates ant movements step-by-step
func MoveAnts(antAssignments map[int][]int) {
	maxSteps := 0
	antSteps := make(map[int]int) // Tracks each ant's current step

	// Determine max steps needed
	for _, path := range antAssignments {
		if len(path)-1 > maxSteps {
			maxSteps = len(path) - 1
		}
	}

	// Simulate each step
	for step := 0; step < maxSteps; step++ {
		var moves []string
		occupied := make(map[int]bool) // Tracks occupied intermediate rooms for this step

		// First pass: Collect valid moves without conflicts
		tentativeMoves := make(map[int]int) // antID -> nextNode
		for antID, path := range antAssignments {
			if antSteps[antID] < len(path)-1 {
				nextNode := path[antSteps[antID]+1]
				tentativeMoves[antID] = nextNode
			}
		}

		// Second pass: Commit moves if the room is not occupied (except start/end)
		for antID, nextNode := range tentativeMoves {
			path := antAssignments[antID] // Retrieve the path for this ant
			isStartOrEnd := antSteps[antID] == 0 || nextNode == path[len(path)-1]
			if isStartOrEnd || !occupied[nextNode] {
				if !isStartOrEnd {
					occupied[nextNode] = true // Block intermediate room
				}
				antSteps[antID]++
				moves = append(moves, fmt.Sprintf("L%d-%d", antID, nextNode))
			}
		}

		if len(moves) > 0 {
			fmt.Println(strings.Join(moves, " "))
		}
	}
}
