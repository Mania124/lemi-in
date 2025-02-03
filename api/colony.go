package api

import (
	"log"
	"strconv"
	"strings"
)

type Room struct {
	start       bool
	end         bool
	number      int
	connections []int
	visited     bool
}

type Colony struct {
	NumberOfAnts int
	Rooms        []*Room
}

func NewRoom() *Room {
	return &Room{
		start:       false,
		end:         false,
		number:      0,
		connections: nil,
		visited:     false,
	}
}

func NewColony() *Colony {
	return &Colony{
		NumberOfAnts: 0,
		Rooms:        nil,
	}
}

func ColonY(content string) (*Colony, int, int) {
	// var count int = 0
	colony := NewColony()
	var room *Room
	var start int
	var end int
	var err error
	contentSlice := strings.Split(strings.TrimSpace(content), "\n")
	for i, ch := range contentSlice {
		if strings.HasPrefix(ch, "#") && !strings.HasPrefix(ch, "##") {
			continue
		}
		if i == 0 {

			colony.NumberOfAnts, err = strconv.Atoi(strings.TrimSpace(ch))
			if err != nil {
				log.Fatal(err)
			}
		} else {
			if ch == "##start" {
				room = NewRoom()
				room.start = true
				temp := strings.Fields(strings.TrimSpace(contentSlice[i+1]))
				room.number, err = strconv.Atoi(temp[0])
				start = room.number
				if err != nil {
					log.Fatal(err)
				}
				colony.Rooms = append(colony.Rooms, room)
				// continue
				contentSlice = append(contentSlice[:i], contentSlice[i+1:]...)
				// fmt.Println(room)
			} else if ch == "##end" {
				room = NewRoom()
				room.end = true
				temp := strings.Fields(strings.TrimSpace(contentSlice[i+1]))
				room.number, err = strconv.Atoi(temp[0])
				end = room.number
				if err != nil {
					log.Fatal(err)
				}
				colony.Rooms = append(colony.Rooms, room)
				// fmt.Println(room)
				contentSlice = contentSlice[i+2:]
				// fmt.Println()
				// fmt.Println(strings.Join(contentSlice, "\n"))
				// fmt.Println(contentSlice)
				break

			} else {
				room = NewRoom()
				temp := strings.Fields(strings.TrimSpace(contentSlice[i]))
				room.number, err = strconv.Atoi(temp[0])
				if err != nil {
					log.Fatal(err)
				}
				colony.Rooms = append(colony.Rooms, room)
				// fmt.Println("this works")
				// fmt.Println(room)
			}
		}

	}
	for _, word := range contentSlice {
		temp := Split(strings.TrimSpace(word))
		// fmt.Println(len(temp), "size")
		// break
		if len(temp) < 2 {
			continue
		}
		for _, w := range colony.Rooms {
			rmn, _ := strconv.Atoi(temp[0])
			// fmt.Println(rmn)
			conn, _ := strconv.Atoi(temp[1])
			if w.number == rmn {
				w.connections = append(w.connections, conn)
			}
		}
	}
	return colony, start, end
}

func BuildGraph(colony *Colony) map[int][]int {
	graph := make(map[int][]int)

	for _, room := range colony.Rooms {
		graph[room.number] = room.connections
	}

	return graph
}
