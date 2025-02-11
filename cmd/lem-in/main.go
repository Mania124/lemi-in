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
	fmt.Println(fcontent)
	colony, start, end := api.ColonY(fcontent)
	// fmt.Println("Number of Ants:", colony.NumberOfAnts)

	graph := api.BuildGraph(colony)
	paths := api.FindAllPaths(graph, start, end)

	// fmt.Println("Paths Found:")
	// for _, path := range paths {
	// 	fmt.Println(path)
	// }

	moves := api.DistributeAnts(paths, colony.NumberOfAnts)
	api.MoveAnts(moves)
}
