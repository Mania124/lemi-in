package api

import (
	"container/list"
	"fmt"
)

// Edmonds-Karp BFS for finding shortest paths
func BFS(graph map[int][]int, start, end int) []int {
	visited := make(map[int]bool)
	queue := list.New()
	queue.PushBack([]int{start})

	for queue.Len() > 0 {
		path := queue.Remove(queue.Front()).([]int)
		last := path[len(path)-1]

		if last == end {
			fmt.Println("Found path:", path)
			return path
		}

		if visited[last] {
			continue
		}
		visited[last] = true

		for _, neighbor := range graph[last] {
			if !visited[neighbor] {
				newPath := append([]int{}, path...) // Copy path
				newPath = append(newPath, neighbor)
				queue.PushBack(newPath)
			}
		}
	}
	return nil
}

// Finds all shortest paths from start to end
func FindAllPaths(graph map[int][]int, start, end int) [][]int {
	var paths [][]int
	originalGraph := make(map[int][]int)

	// ✅ Make a deep copy of the original graph
	for key, val := range graph {
		originalGraph[key] = append([]int{}, val...)
	}

	for {
		path := BFS(graph, start, end)
		if path == nil {
			break // No more paths found
		}
		paths = append(paths, path)

		// ✅ Instead of wiping out nodes, only disable the most recently found path
		for i := 1; i < len(path)-1; i++ {
			graph[path[i]] = []int{} // Prevent reuse of this exact path
		}
	}

	// ✅ Restore the original graph after finding paths
	for key := range graph {
		graph[key] = originalGraph[key]
	}

	return paths
}

func DistributeAnts(paths [][]int, numAnts int) map[int][]int {
	antAssignments := make(map[int][]int)
	pathUsage := make([]int, len(paths)) // Tracks how many ants are assigned to each path

	// ✅ Debugging: Print available paths
	fmt.Println("Available Paths:", paths)

	for antID := 1; antID <= numAnts; antID++ {
		assigned := false // ✅ Ensure every ant is assigned a path

		for i := 0; i < len(paths); i++ {
			// ✅ If this is the first ant, assign it to the first path
			if antID == 1 || i == 0 {
				antAssignments[antID] = paths[i]
				pathUsage[i]++
				assigned = true
				break
			}

			// ✅ Ensure `i+1` is within bounds before checking next path
			if i+1 < len(paths) && (len(paths[i])-1+pathUsage[i] < len(paths[i+1])-1+pathUsage[i+1]) {
				antAssignments[antID] = paths[i]
				pathUsage[i]++
				assigned = true
				break
			}
		}

		// ✅ If no path was selected (should not happen), assign the first path
		if !assigned {
			antAssignments[antID] = paths[0]
			pathUsage[0]++
		}
	}

	// ✅ Debugging: Print final assignments
	// fmt.Println("Ant Assignments:", antAssignments)

	return antAssignments
}
