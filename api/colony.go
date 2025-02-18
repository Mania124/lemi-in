package api

import (
	"log"
	"strconv"
	"strings"
)

type Room struct {
	start       bool
	end         bool
	number      string
	connections []string
}

type Colony struct {
	NumberOfAnts int
	Rooms        []*Room
}

// Creates a new room
func NewRoom(number string, isStart, isEnd bool) *Room {
	return &Room{
		start:       isStart,
		end:         isEnd,
		number:      number,
		connections: []string{},
	}
}

// Creates a new colony
func NewColony() *Colony {
	return &Colony{
		NumberOfAnts: 0,
		Rooms:        []*Room{},
	}
}

// Parses file content and builds the colony, returning start and end room numbers
func ColonY(content string) (*Colony, string, string) {
	colony := NewColony()
	var start, end string
	var err error

	contentSlice := strings.Split(strings.TrimSpace(content), "\n")

	for i := 0; i < len(contentSlice); i++ {
		line := contentSlice[i]
		if strings.HasPrefix(line, "#") && !strings.HasPrefix(line, "##") {
			continue
		}

		if i == 0 {
			colony.NumberOfAnts, err = strconv.Atoi(strings.TrimSpace(line))
			if err != nil {
				log.Fatal("Error parsing number of ants:", err)
			}
		} else if (line == "##start" || line == "##end") && i != len(contentSlice)-1 {
			i++
			temp := strings.Fields(strings.TrimSpace(contentSlice[i]))
			roomNumber := temp[0]

			isStart := line == "##start"
			isEnd := line == "##end"
			if isStart {
				start = roomNumber
			} else {
				end = roomNumber
			}
			colony.Rooms = append(colony.Rooms, NewRoom(roomNumber, isStart, isEnd))

		} else if strings.Contains(line, "-") {
			continue // Connection lines are handled later
		} else {
			temp := strings.Fields(strings.TrimSpace(line))
			roomNumber := temp[0]
			colony.Rooms = append(colony.Rooms, NewRoom(roomNumber, false, false))
		}
	}

	// Parse connections
	for _, line := range contentSlice {
		temp := Split(strings.TrimSpace(line))
		if len(temp) < 2 {
			continue
		}

		roomA := temp[0]
		roomB := temp[1]

		for _, room := range colony.Rooms {
			if room.number == roomA {
				room.connections = append(room.connections, roomB)
			}
			if room.number == roomB {
				room.connections = append(room.connections, roomA)
			}
		}
	}

	return colony, start, end
}

// Converts colony into an adjacency list
func BuildGraph(colony *Colony) map[string][]string {
	graph := make(map[string][]string)
	for _, room := range colony.Rooms {
		graph[room.number] = room.connections
	}
	return graph
}
