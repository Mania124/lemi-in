package api

import (
	"fmt"
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
	nummberOfAnts int
	rooms         []*Room
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
		nummberOfAnts: 0,
		rooms:         nil,
	}
}

func ColonY(content string) *Colony {
	// var count int = 0
	colony := NewColony()
	var room *Room
	var err error
	contentSlice := strings.Split(strings.TrimSpace(content), "\n")
	for i, ch := range contentSlice {
		if ch[0] == '#' && ch[1] != '#' {
			continue
		}
		if i == 0 {

			colony.nummberOfAnts, err = strconv.Atoi(strings.TrimSpace(ch))
			if err != nil {
				log.Fatal(err)
			}
		} else {
			if ch == "##start" {
				room = NewRoom()
				room.start = true
				temp := strings.Fields(strings.TrimSpace(contentSlice[i+1]))
				room.number, err = strconv.Atoi(temp[0])
				if err != nil {
					log.Fatal(err)
				}
				colony.rooms = append(colony.rooms, room)
				contentSlice = append(contentSlice[:i], contentSlice[i+2:]...)
				fmt.Println(room)
			} else if ch == "##end" {
				room = NewRoom()
				room.end = true
				temp := strings.Fields(strings.TrimSpace(contentSlice[i+1]))
				room.number, err = strconv.Atoi(temp[0])
				if err != nil {
					log.Fatal(err)
				}
				colony.rooms = append(colony.rooms, room)
				fmt.Println(room)
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
				colony.rooms = append(colony.rooms, room)
				fmt.Println(room)
			}
		}

	}
	for _, word := range contentSlice {
		temp := Split(strings.TrimSpace(word))
		// fmt.Println(len(temp), "size")
		// break
		for _, w := range colony.rooms {
			rmn, _ := strconv.Atoi(temp[0])
			// fmt.Println(rmn)
			conn, _ := strconv.Atoi(temp[1])
			if w.number == rmn {
				w.connections = append(w.connections, conn)
			}
		}
	}
	return colony
}
