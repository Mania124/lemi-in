package main

import (
	"fmt"
	"log"
	"os"

	"lem-in/api"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: go run main.go <filename>")
	}
	fcontent, err := api.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	colony, start, end := api.ColonY(fcontent)
	// fmt.Println("Number of Ants:", colony.NumberOfAnts)

	graph := api.BuildGraph(colony)
	paths := api.FindAllPaths(graph, start, end)
	if len(paths) == 0 || colony.NumberOfAnts == 0 {
		fmt.Println("ERROR: invalid data format")
		return
	}
	fmt.Println(fcontent)
	fmt.Println("Paths Found:")
	for _, path := range paths {
		fmt.Println(path)
	}

	moves := api.DistributeAnts(paths, colony.NumberOfAnts)
	fmt.Println(moves)
	api.Ass(moves)
	// api.MoveAnts(moves)
}
