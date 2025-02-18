package api

import (
	"fmt"
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

func NewRoom(number string, isStart, isEnd bool) *Room {
	return &Room{
		start:       isStart,
		end:         isEnd,
		number:      number,
		connections: []string{},
	}
}

func NewColony() *Colony {
	return &Colony{
		NumberOfAnts: 0,
		Rooms:        []*Room{},
	}
}

func ColonY(content string) (*Colony, string, string, error) {
	colony := NewColony()
	var start, end string

	contentSlice := strings.Split(strings.TrimSpace(content), "\n")

	for i := 0; i < len(contentSlice); i++ {
		line := contentSlice[i]
		if strings.HasPrefix(line, "#") && !strings.HasPrefix(line, "##") {
			continue
		}

		if i == 0 {
			numberOfAnts, err := strconv.Atoi(strings.TrimSpace(line))
			if err != nil {
				return nil, "", "", fmt.Errorf("invalid number of ants: %w", err)
			}
			colony.NumberOfAnts = numberOfAnts
		} else if (line == "##start" || line == "##end") && i != len(contentSlice)-1 {
			i++
			temp := strings.Fields(strings.TrimSpace(contentSlice[i]))
			if len(temp) < 1 {
				return nil, "", "", fmt.Errorf("invalid room format after start/end")
			}
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
			continue
		} else {
			temp := strings.Fields(strings.TrimSpace(line))
			if len(temp) < 1 {
				return nil, "", "", fmt.Errorf("invalid room format")
			}
			roomNumber := temp[0]
			colony.Rooms = append(colony.Rooms, NewRoom(roomNumber, false, false))
		}
	}

	for _, line := range contentSlice {
		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			if len(parts) != 2 {
				return nil, "", "", fmt.Errorf("invalid connection format")
			}
			roomA, roomB := parts[0], parts[1]

			for _, room := range colony.Rooms {
				if room.number == roomA {
					room.connections = append(room.connections, roomB)
				}
				if room.number == roomB {
					room.connections = append(room.connections, roomA)
				}
			}
		}
	}

	return colony, start, end, nil
}

func BuildGraph(colony *Colony) map[string][]string {
	graph := make(map[string][]string)
	for _, room := range colony.Rooms {
		graph[room.number] = room.connections
	}
	return graph
}
