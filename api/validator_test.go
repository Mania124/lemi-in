package api

import (
	"testing"
)

func TestValidateStartEndRooms(t *testing.T) {
	tests := []struct {
		name      string
		start, end string
		expectErr bool
	}{
		{"Both missing", "", "", true},
		{"Missing start", "", "end", true},
		{"Missing end", "start", "", true},
		{"Valid start and end", "start", "end", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateStartEndRooms(tt.start, tt.end)
			if (err != nil) != tt.expectErr {
				t.Errorf("got error %v, expected error: %v", err, tt.expectErr)
			}
		})
	}
}

func TestValidateRooms(t *testing.T) {
	colony := &Colony{
		Rooms: []*Room{
			{Number: "A", X: 0, Y: 0},
			{Number: "B", X: 1, Y: 1},
		},
	}

	if err := ValidateRooms(colony); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	invalidColony := &Colony{Rooms: []*Room{}}
	if err := ValidateRooms(invalidColony); err == nil {
		t.Errorf("expected error for empty rooms, got nil")
	}
}

func TestValidateRoomLinks(t *testing.T) {
	colony := &Colony{
		Rooms: []*Room{
			{Number: "A", X: 0, Y: 0},
			{Number: "B", X: 1, Y: 1},
		},
	}
	graph := map[string][]string{
		"A": {"B"},
		"B": {"A"},
	}

	if err := ValidateRoomLinks(graph, colony); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	graphNoLinks := map[string][]string{}
	if err := ValidateRoomLinks(graphNoLinks, colony); err == nil {
		t.Errorf("expected error for unlinked rooms, got nil")
	}
}
