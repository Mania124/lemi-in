package api

import "fmt"

func ValidateStartEndRooms(start, end string) error {
	switch {
	case start == "" && end == "":
		return fmt.Errorf("ERROR: invalid data format, no start and end room found")
	case start == "":
		return fmt.Errorf("ERROR: invalid data format, no start room found")
	case end == "":
		return fmt.Errorf("ERROR: invalid data format, no end room found")
	default:
		return nil
	}
}

// validateRooms checks for invalid coordinates, missing rooms, and duplicate coordinates
func ValidateRooms(colony *Colony) error {
	if len(colony.Rooms) == 0 {
		return fmt.Errorf("ERROR: no rooms defined")
	}

	coordsMap := make(map[string]string) // Stores coordinates as "x,y" -> room name

	for _, room := range colony.Rooms {
		if room.Number == "" || room.X < 0 || room.Y < 0 {
			return fmt.Errorf("ERROR: invalid room definition for '%s'", room.Number)
		}

		// Check for duplicate coordinates
		coordKey := fmt.Sprintf("%d,%d", room.X, room.Y)
		if existingRoom, exists := coordsMap[coordKey]; exists {
			return fmt.Errorf("ERROR: duplicate coordinates found for rooms '%s' and '%s'", existingRoom, room.Number)
		}
		coordsMap[coordKey] = room.Number
	}

	return nil
}

// validateRoomLinks ensures all rooms have at least one link
func ValidateRoomLinks(graph map[string][]string, colony *Colony) error {
	for _, room := range colony.Rooms {
		if _, exists := graph[room.Number]; !exists || len(graph[room.Number]) == 0 {
			return fmt.Errorf("ERROR: room '%s' has no links", room.Number)
		}
	}
	return nil
}
