package main

import (
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
	// find paths then filter most efficient paths
	// move the ants and display result
}
