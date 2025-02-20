package api

import (
	"testing"
)

func TestBFS(t *testing.T) {
	graph := map[string][]string{
		"0": {"1"},
		"1": {"0", "2"},
		"2": {"1", "3"},
		"3": {"2"},
	}
	found:=make(map[string]map[string]bool)
	path := BFS(graph, "0", "3",found)
	expected := []string{"0", "1", "2", "3"}

	if !compareSlices(path, expected) {
		t.Errorf("BFS() got = %v, want %v", path, expected)
	}
}

func TestFindAllPaths(t *testing.T) {
	graph := map[string][]string{
		"0": {"1", "2"},
		"1": {"0", "3"},
		"2": {"0", "3"},
		"3": {"1", "2"},
	}

	paths := FindAllPaths(graph, "0", "3")
	expected := [][]string{
		{"0", "1", "3"},
		{"0", "2", "3"},
	}

	if len(paths) != len(expected) {
		t.Errorf("FindAllPaths() got = %v, want %v", paths, expected)
	}
}

func TestDistributeAnts(t *testing.T) {
	paths := [][]string{
		{"0", "1", "3"},
		{"0", "2", "3"},
	}

	assignments := DistributeAnts(paths, 5)
	if len(assignments) != 5 {
		t.Errorf("DistributeAnts() got = %v, want 5 assignments", assignments)
	}
}

func TestMove(t *testing.T) {
	assignments := map[int][]string{
		1: {"1", "3"},
		2: {"2", "3"},
	}

	groups := [][]int{
		{1, 2},
	}

	// Capture output (optional)
	// You can use a buffer to capture and verify printed output if needed.
	Move(groups, assignments)
}
