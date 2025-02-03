package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"lem-in/api"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	// Readfile
	fcontent, err := api.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(os.Stdout, fcontent)

	// extract number of ants and room-numbers with connections
	colony, start, end := api.ColonY(fcontent)
	// fmt.Println(start, end)
	// fmt.Printf("number of ant:%d\n", colony.NumberOfAnts)
	graph := api.BuildGraph(colony)
	// fmt.Println(graph)
	// find paths then filter most efficient paths
	paths := api.FindAllPaths(graph, start, end)
	// fmt.Println(paths)
	// move the ants and display result
	fmt.Println(api.DistributeAnts(paths, colony.NumberOfAnts))
}
