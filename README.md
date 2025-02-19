# lem-in

A Go implementation of an ant farm pathfinding algorithm that efficiently moves ants through a network of rooms.

## Description

The program reads a map of an ant farm from a file and finds the optimal paths to move ants from the start room to the end room. It handles various graph configurations and ensures efficient ant movement while avoiding congestion.

## Features

- Reads ant farm configuration from a file
- Finds optimal paths using BFS (Breadth-First Search)
- Handles multiple graph scenarios (simple and complex)
- Efficiently distributes ants across available paths
- Prevents room congestion during ant movement
- Outputs ant movements in a clear, readable format

## Project Structure

```bash

lem-in/
├── api/
│   ├── colony.go       # Colony and room structures
│   ├── pathfinder.go   # Path finding and ant movement logic
│   ├── reader.go       # Input file parsing
│   ├── validator.go    # Validation functions
│   └── customsplit.go  # Custom string splitting functions
├── cmd/
│   └── lem-in/
│       └── main.go     # Main program entry
├── test/
│   ├── api/
│   │   ├── colony_test.go
│   │   ├── pathfinder_test.go
│   │   ├── reader_test.go
│   │   ├── validator_test.go
│   │   └── customsplit_test.go
│   └── test_files/
│       ├── text.txt
│       └── test1.txt
├── LICENSE
└── README.md
```

## Usage

```bash
go run cmd/lem-in/main.go <filename>
```

### Input File Format

The input file should follow this format:

```bash
<number_of_ants>
<room_definitions>
##start
<start_room> <x_coord> <y_coord>
##end
<end_room> <x_coord> <y_coord>
<room_connections>
```

Example:

```bash
3
##start
0 1 2
##end
1 9 2
3 5 4
0-2
0-3
2-1
3-1
2-3
```

## Contributing

Feel free to submit issues and enhancement requests.

## Contributors

- mmoffat
- hshikuku
- oragwelr

## Repository

[lem-in Repository](https://learn.zone01kisumu.ke/git/hshikuku/lem-in.git)
