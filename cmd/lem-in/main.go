package main

import (
	"fmt"
	"os"

	"lem-in/pkg/lemin"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR: invalid data format, please provide input file")
		os.Exit(1)
	}

	// Parse input file
	colony, err := lemin.ParseFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Create solver and find solution
	solver := lemin.NewSolver(colony)
	moves, err := solver.Solve()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Print solution
	solver.PrintSolution(moves)
}
