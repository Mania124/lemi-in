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
	rooms         []Room
}

func ColonY(content string) interface{} {
	// var count int = 0
	var colony Colony
	var room Room
	var err error
	contentSlice := strings.Split(content, "\n")
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
				room.start = true
				temp := strings.Fields(strings.TrimSpace(contentSlice[i+1]))
				room.number, err = strconv.Atoi(temp[0])
				if err != nil {
					log.Fatal(err)
				}
				colony.rooms = append(colony.rooms, room)
				room.start = false
				contentSlice = append(contentSlice[:i], contentSlice[i+2:]...)
			} else if ch == "##end" {
				room.end = true
				temp := strings.Fields(strings.TrimSpace(contentSlice[i+1]))
				room.number, err = strconv.Atoi(temp[0])
				if err != nil {
					log.Fatal(err)
				}
				colony.rooms = append(colony.rooms, room)
				contentSlice = contentSlice[i+2:]
				fmt.Println()
				fmt.Println(strings.Join(contentSlice, "\n"))
				fmt.Println(contentSlice[0])
				break

			} else {
				temp := strings.Fields(strings.TrimSpace(contentSlice[i]))
				room.number, err = strconv.Atoi(temp[0])
				if err != nil {
					log.Fatal(err)
				}
				colony.rooms = append(colony.rooms, room)
			}
		}

	}
	for _, word := range contentSlice {
		temp := Split(strings.TrimSpace(word))
		fmt.Println(temp)
		// break
		for _, w := range colony.rooms {
			rmn, _ := strconv.Atoi(temp[0])
			fmt.Println(rmn)
			conn, _ := strconv.Atoi(temp[1])
			if w.number == rmn {
				w.connections = append(w.connections, conn)
			}
		}
	}
	return colony
}
