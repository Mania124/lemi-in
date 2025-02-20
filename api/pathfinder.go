package api

import (
	"container/list"
	"fmt"
	"sort"
)

// var count = make(map[string]bool)

// moved = make(map[int]int)

// BFS finds the shortest path while avoiding blocked edges
func BFS(graph map[string][]string, start, end string, blockedEdges map[string]map[string]bool) []string {
	queue := list.New()
	queue.PushBack([]string{start})

	for queue.Len() > 0 {
		currentPath := queue.Remove(queue.Front()).([]string)
		lastNode := currentPath[len(currentPath)-1]

		if lastNode == end {
			return currentPath // Found a valid path
		}

		// Explore neighbors, skipping blocked edges and cycles
		for _, neighbor := range graph[lastNode] {
			if !blockedEdges[lastNode][neighbor] && !contains(currentPath, neighbor) {
				newPath := append([]string{}, currentPath...)
				newPath = append(newPath, neighbor)
				queue.PushBack(newPath)
			}
		}
	}
	return nil // No valid path found
}

// Finds all non-overlapping paths from start to end
func FindAllPaths(graph map[string][]string, start, end string) [][]string {
	var paths [][]string
	blockedEdges := make(map[string]map[string]bool)

	// Initialize blockedEdges for all nodes
	for node := range graph {
		blockedEdges[node] = make(map[string]bool)
	}

	for {
		path := BFS(graph, start, end, blockedEdges)
		if path == nil {
			break // No more paths available
		}
		paths = append(paths, path)

		// Block edges instead of entire nodes to allow more alternative paths
		for j := 0; j < len(path)-1; j++ {
			blockedEdges[path[j]][path[j+1]] = true
		}
	}

	return paths
}

// Distributes ants fairly, prioritizing shorter paths when possible
func DistributeAnts(paths [][]string, numAnts int) map[int][]string {
	antAssignments := make(map[int][]string) // Tracks assigned paths per ant
	antsPerPath := make([]int, len(paths))   // Tracks number of ants per path
	pathLengths := make([]int, len(paths))   // Store each path’s length

	// Get the length of each path
	for i, path := range paths {
		pathLengths[i] = len(path) - 1 // Exclude starting node "0"
	}

	// Assign ants dynamically, always preferring the least burdened shortest path
	for antID := 1; antID <= numAnts; antID++ {
		minIndex := 0
		minLoad := antsPerPath[0] + pathLengths[0] // Consider path length + current load

		for i := 1; i < len(paths); i++ {
			currentLoad := antsPerPath[i] + pathLengths[i]
			if currentLoad < minLoad {
				minIndex = i
				minLoad = currentLoad
			}
		}

		// Assign ant to the selected shortest available path
		antAssignments[antID] = paths[minIndex][1:]
		antsPerPath[minIndex]++ // Increase load for this path
	}

	return antAssignments
}

// group ants with similar paths together
func AssignGroups(assignment map[int][]string) [][]int {
	var tog [][]int
	var gr []int
	found := make(map[int]bool)
	for _, v := range assignment {
		for key := range assignment {
			if compareSlices(assignment[key], v) && !found[key] {
				gr = append(gr, key)
				found[key] = true
			}
		}
		if len(gr) != 0 {
			sort.Ints(gr)
			tog = append(tog, gr)
			gr = []int{}
		}
	}

	sort.Slice(tog, func(i, j int) bool {
		return len(tog[i]) > len(tog[j]) || tog[i][0] < tog[j][0]
	})
	return tog
}

// Move prints elements cumulatively row by row.
func Move(input [][]int, myMap map[int][]string) {
	// Track how many elements we've used from each key in myMap
	moved := make(map[int]int)

	// Find max depth (longest list in input)
	maxLen := 0
	for _, v := range input {
		if len(v) > maxLen {
			maxLen = len(v)
		}
	}

	// Print row by row, accumulating elements
	for count := 0; ; count++ {
		line := ""   // Collect output for this row
		done := true // Track if all values are exhausted

		for _, v := range input {
			for j := 0; j <= count && j < len(v); j++ {
				id := v[j] // Current list ID

				// Ensure we do not exceed the mapped values
				if moved[id] < len(myMap[id]) {
					if line != "" {
						line += " " // Add space before next value
					}

					// Append value to line
					line += fmt.Sprintf("L%d-%s", id, myMap[id][moved[id]])

					// Move to next element in `myMap`
					moved[id]++
					done = false // Not done yet, there is more data
				}
			}
		}

		if done {
			break // Stop if all elements are exhausted
		}

		fmt.Println(line) // Print the accumulated row
	}
}

// CompareSlices checks if two slices have the same elements and are of the same length.
func compareSlices[T comparable](slice1, slice2 []T) bool {
	// Check if the lengths are the same
	if len(slice1) != len(slice2) {
		return false
	}

	// Compare each element
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	// If all elements are the same, return true
	return true
}

// Helper function to check if a value exists in a slice
func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
