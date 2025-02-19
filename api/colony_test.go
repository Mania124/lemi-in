package api

import (
	"testing"
)

func TestColonY(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		wantAnts int
		wantErr  bool
	}{
		{
			name:     "Valid input",
			content:  "3\n##start\n0 1 2\n##end\n1 5 6\n0-1\n1-2",
			wantAnts: 3,
			wantErr:  false,
		},
		{
			name:     "Invalid number of ants",
			content:  "invalid\n##start\n0 1 2\n##end\n1 5 6\n0-1\n1-2",
			wantAnts: 0,
			wantErr:  true,
		},
		{
			name:     "No start or end room",
			content:  "3\n0 1 2\n1 5 6\n0-1\n1-2",
			wantAnts: 3,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			colony, _, _, err := ColonY(tt.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("ColonY() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && colony.NumberOfAnts != tt.wantAnts {
				t.Errorf("ColonY() got = %v, want %v", colony.NumberOfAnts, tt.wantAnts)
			}
		})
	}
}

func TestBuildGraph(t *testing.T) {
	colony := &Colony{
		Rooms: []*Room{
			{Number: "0", Connections: []string{"1"}},
			{Number: "1", Connections: []string{"0", "2"}},
			{Number: "2", Connections: []string{"1"}},
		},
	}

	graph := BuildGraph(colony)
	expected := map[string][]string{
		"0": {"1"},
		"1": {"0", "2"},
		"2": {"1"},
	}

	if !compareGraphs(graph, expected) {
		t.Errorf("BuildGraph() got = %v, want %v", graph, expected)
	}
}

// Helper function to compare two graphs
func compareGraphs(a, b map[string][]string) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if !compareSlices(v, b[k]) {
			return false
		}
	}
	return true
}
