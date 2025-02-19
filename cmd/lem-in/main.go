package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"lem-in/api"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: go run main.go <filename>")
	}
	fcontent, err := api.ReadFile(os.Args[1])
	if err != nil || len(strings.TrimSpace(fcontent)) == 0 {
		fmt.Println("ERROR: invalid data format")
		return
	}

	// Updated to handle the 4 return values from ColonY
	colony, start, end, err := api.ColonY(fcontent)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	if colony.NumberOfAnts <= 0 {
		fmt.Println("ERROR: invalid number of Ants")
		return
	}
	if err := api.ValidateStartEndRooms(start, end); err != nil {
		fmt.Println(err)
		return
	}
	// 4. Validate room definitions
	if err := api.ValidateRooms(colony); err != nil {
		fmt.Println(err)
		return
	}
	graph := api.BuildGraph(colony)
	if err := api.ValidateRoomLinks(graph, colony); err != nil {
		fmt.Println(err)
		return
	}

	paths := api.FindAllPaths(graph, start, end)
	if len(paths) == 0 {
		fmt.Println("ERROR: invalid data format")
		return
	}
	fmt.Println(fcontent)
	moves := api.DistributeAnts(paths, colony.NumberOfAnts)
	tog := api.AssignGroups(moves)
	api.Move(tog, moves)
}
