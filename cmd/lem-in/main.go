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

	colony, start, end := api.ColonY(fcontent)
	if colony.NumberOfAnts <= 0 {
		fmt.Println("ERROR: invalid number of Ants")
		return
	}
	if start == "" || end == "" {
		if start == "" && len(end) != 0 {
			fmt.Println("ERROR: invalid data format, no start room found")
			return
		}
		if end == "" && len(start) != 0 {
			fmt.Println("ERROR: invalid data format, no end room found")
			return
		}
		fmt.Println("ERROR: invalid data format, no start and end room found")
		return
	}

	// fmt.Println("Number of Ants:", colony.NumberOfAnts)

	graph := api.BuildGraph(colony)
	paths := api.FindAllPaths(graph, start, end)
	if len(paths) == 0 {
		fmt.Println("ERROR: invalid data format")
		return
	}
	fmt.Println(fcontent)
	// fmt.Println("Paths Found:")
	// for _, path := range paths {
	// 	fmt.Println(path)
	// }

	moves := api.DistributeAnts(paths, colony.NumberOfAnts)
	// fmt.Println(moves)
	tog := api.AssignGroups(moves)
	// fmt.Println(tog)
	api.Move(tog, moves)
	// api.MoveAnts(moves)
}
