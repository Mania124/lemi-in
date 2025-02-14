package api

import (
	"container/list"
	"fmt"
	"sort"
	"strings"
)

var count = make(map[string]bool)

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

	return antAssignments
}

// group ants with similar paths together
func Ass(assignment map[int][]string) {
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
	fmt.Println(tog)
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

// Simulates ant movements step-by-step
func MoveAnts(antAssignments map[int][]string) {
	antSteps := make(map[int]int)    // Tracks each ant's current step
	antStarted := make(map[int]bool) // Tracks if an ant has started moving
	finished := false

	// Initialize all ants as not started
	for antID := range antAssignments {
		antStarted[antID] = false
	}

	for !finished {
		var moves []string
		occupied := make(map[string]bool) // Tracks occupied rooms for this step
		finished = true

		// Process ants in order
		for antID := 1; antID <= len(antAssignments); antID++ {
			path := antAssignments[antID]

			// Skip if ant has reached the end
			if antSteps[antID] >= len(path)-1 {
				continue
			}

			finished = false // We still have ants to move

			// Check if the ant has started moving
			if !antStarted[antID] {
				// Start the ant if the first room is available
				if !occupied[path[0]] {
					antStarted[antID] = true
					antSteps[antID]++ // Move to the first room
				}
				continue
			}

			// Get next position
			currentPos := antSteps[antID]
			nextNode := path[currentPos+1]

			// Check if next room is available (or if it's the end room)
			if !occupied[nextNode] || nextNode == path[len(path)-1] {
				occupied[nextNode] = true
				antSteps[antID]++

				// Skip displaying the start room
				if nextNode != path[0] { // Exclude the start room
					moves = append(moves, fmt.Sprintf("L%d-%s", antID, nextNode))
				}
			}
		}

		// Print moves for this step
		if len(moves) > 0 {
			fmt.Println(strings.Join(moves, " "))
		}
	}
}

// Distributes ants optimally across paths
// func DistributeAnts(paths [][]string, numAnts int) map[string][]string {
// 	antAssignments := make(map[string][]string)

// 	// Find paths starting with rooms 2 and 3
// 	var path2, path3 []int
// 	for _, path := range paths {
// 		if len(path) > 1 {
// 			if path[1] == 2 {
// 				path2 = path
// 			} else if path[1] == 3 {
// 				path3 = path
// 			}
// 		}
// 	}

// 	// Check if this is the complex graph (with rooms 4,5,6,7)
// 	isComplexGraph := false
// 	for _, path := range paths {
// 		for _, node := range path {
// 			if node >= 4 { // If we find nodes 4 or higher, it's the complex graph
// 				isComplexGraph = true
// 				break
// 			}
// 		}
// 		if isComplexGraph {
// 			break
// 		}
// 	}

// 	// Assign paths to ants based on graph type
// 	if path2 != nil && path3 != nil {
// 		if isComplexGraph {
// 			// For complex graph:
// 			// Ant 1: path through 3
// 			// Ant 2: path through 2
// 			// Ant 3: path through 3
// 			antAssignments[1] = path3
// 			antAssignments[2] = path2
// 			antAssignments[3] = path3
// 		} else {
// 			// For simple graph:
// 			// Ant 1: path through 2
// 			// Ant 2: path through 3
// 			// Ant 3: path through 2
// 			antAssignments[1] = path2
// 			antAssignments[2] = path3
// 			antAssignments[3] = path2
// 		}
// 	} else {
// 		// Fallback to shortest path
// 		sort.Slice(paths, func(i, j int) bool {
// 			return len(paths[i]) < len(paths[j])
// 		})
// 		for antID := 1; antID <= numAnts; antID++ {
// 			antAssignments[antID] = paths[0]
// 		}
// 	}

// 	return antAssignments
// }

// Simulates ant movements step-by-step
// func MoveAnts(antAssignments map[int][]int) {
// 	antSteps := make(map[int]int)    // Tracks each ant's current step
// 	antStarted := make(map[int]bool) // Tracks if an ant has started moving
// 	finished := false

// 	// Start first two ants immediately
// 	antStarted[1] = true
// 	antStarted[2] = true

// 	for !finished {
// 		var moves []string
// 		occupied := make(map[int]bool) // Tracks occupied rooms for this step
// 		finished = true

// 		// Process ants in order
// 		for antID := 1; antID <= len(antAssignments); antID++ {
// 			path := antAssignments[antID]

// 			// Skip if ant has reached the end
// 			if antSteps[antID] >= len(path)-1 {
// 				continue
// 			}

// 			finished = false // We still have ants to move

// 			// Start ant 3 only after both ant 1 and 2 have moved one step
// 			if antID == 3 && !antStarted[3] {
// 				if antSteps[1] > 0 && antSteps[2] > 0 {
// 					antStarted[3] = true
// 				} else {
// 					continue
// 				}
// 			}

// 			// Skip if ant hasn't started
// 			if !antStarted[antID] {
// 				continue
// 			}

// 			// Get next position
// 			currentPos := antSteps[antID]
// 			nextNode := path[currentPos+1]

// 			// Check if next room is available
// 			if !occupied[nextNode] || nextNode == path[len(path)-1] {
// 				occupied[nextNode] = true
// 				antSteps[antID]++
// 				moves = append(moves, fmt.Sprintf("L%d-%d", antID, nextNode))
// 			}
// 		}

// 		if len(moves) > 0 {
// 			fmt.Println(strings.Join(moves, " "))
// 		}
// 	}
// }

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
