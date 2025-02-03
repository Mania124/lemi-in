package api

import (
	"container/list"
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
	for {
		path := BFS(graph, start, end)
		if path == nil {
			break
		}
		paths = append(paths, path)
		for _, node := range path {
			if node != start && node != end {
				graph[node] = []int{}
			}
		}
	}
	return paths
}

func DistributeAnts(paths [][]int, numAnts int) map[int][]int {
	antAssignments := make(map[int][]int)
	pathUsage := make([]int, len(paths))

	for antID := 1; antID <= numAnts; antID++ {
		for i, path := range paths {
			if i+1 <len(paths) && len(path)-1+pathUsage[i] < len(paths[i+1])-1+pathUsage[i+1] {
				antAssignments[antID] = path
				pathUsage[i]++
				break
			}
		}
	}
	return antAssignments
}
