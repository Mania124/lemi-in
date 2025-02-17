package api

import (
	"container/list"
	"fmt"
	"sort"
	"strings"
)

var count = make(map[string]bool)

// moved = make(map[int]int)

// Finds the shortest path using BFS (no global visited)
func BFS(graph map[string][]string, start, end string) []string {
	queue := list.New()
	queue.PushBack([]string{start})

	for queue.Len() > 0 {
		currentPath := queue.Remove(queue.Front()).([]string)
		lastNode := currentPath[len(currentPath)-1]

		if lastNode == end {
			str := strings.Join(currentPath, "")
			if !count[str] {
				count[str] = true
				return currentPath
			}

		}

		// Explore neighbors not already in the current path (prevents cycles)
		for _, neighbor := range graph[lastNode] {
			if !contains(currentPath, neighbor) {
				newPath := append([]string{}, currentPath...)
				newPath = append(newPath, neighbor)
				queue.PushBack(newPath)
			}
		}
	}
	return nil
}

// Finds all non-overlapping paths by blocking nodes in found paths
func FindAllPaths(graph map[string][]string, start, end string) [][]string {
	var paths [][]string
	tempGraph := copyGraph(graph) // Create a modifiable copy of the graph

	// First try to find shortest paths
	for {
		path := BFS(tempGraph, start, end)
		if path == nil {
			break // No more paths
		}
		paths = append(paths, path)

		// Block nodes in this path except start and end
		for i := 1; i < len(path)-1; i++ {

			nodeToBlock := path[i]
			// Remove the node from the graph
			for u := range tempGraph {
				var newNeighbors []string
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

// Distributes ants optimally across paths
func DistributeAnts(paths [][]string, numAnts int) map[int][]string {
	antAssignments := make(map[int][]string)
	antAssignment := make(map[int][]string)

	// Calculate the length of each path
	pathLengths := make([]int, len(paths))
	for i, path := range paths {
		pathLengths[i] = len(path)
	}

	// Calculate the number of ants to assign to each path
	antsPerPath := make([]int, len(paths))
	remainingAnts := numAnts

	// Distribute ants to paths based on their lengths
	for remainingAnts > 0 {
		// Find the path with the minimum (length + ants already assigned)
		minIndex := 0
		minValue := pathLengths[0] + antsPerPath[0]
		for i := 1; i < len(paths); i++ {
			currentValue := pathLengths[i] + antsPerPath[i]
			if currentValue < minValue {
				minIndex = i
				minValue = currentValue
			}
		}

		// Assign one ant to the selected path
		antsPerPath[minIndex]++
		remainingAnts--
	}

	// Assign ants to paths based on the calculated distribution
	antID := 1
	for i, ants := range antsPerPath {
		for j := 0; j < ants; j++ {
			antAssignments[antID] = paths[i]
			antID++
		}
	}
	for k, v := range antAssignments {
		antAssignment[k] = v[1:]
	}

	return antAssignment
}

// group ants with similar paths together
func Ass(assignment map[int][]string) [][]int {
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
		return len(tog[i]) > len(tog[j])
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

// Helper to deep-copy the graph
func copyGraph(original map[string][]string) map[string][]string {
	copied := make(map[string][]string)
	for u, neighbors := range original {
		copied[u] = append([]string{}, neighbors...)
	}
	return copied
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
