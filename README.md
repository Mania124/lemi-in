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

## Usage

```bash
go run cmd/lem-in/main.go <filename>
```

### Input File Format

The input file should follow this format:
```
<number_of_ants>
<room_definitions>
##start
<start_room> <x_coord> <y_coord>
##end
<end_room> <x_coord> <y_coord>
<room_connections>
```

Example:
```
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

### Output Format

The program outputs the input data followed by the ant movements:
```
L<ant_number>-<room_number>
```

Example output:
```
L1-2 L2-3
L1-1 L2-1 L3-2
L3-1
```

## Implementation Details

- Uses Breadth-First Search for pathfinding
- Implements node-disjoint path finding
- Optimizes ant distribution based on path lengths
- Handles concurrent ant movements efficiently
- Supports both simple and complex graph configurations

## Project Structure

```
lem-in/
├── api/
│   ├── colony.go       # Colony and room structures
│   ├── pathfinder.go   # Path finding and ant movement logic
│   └── reader.go       # Input file parsing
├── cmd/
│   └── lem-in/
│       └── main.go     # Main program entry
└── README.md
```

## Error Handling

The program includes robust error handling for:
- Invalid file formats
- Missing start/end rooms
- Invalid room connections
- Invalid number of ants
- Unreachable end room

## Contributing

Feel free to submit issues and enhancement requests.

## Contributors

- mmoffat
- hshikuku
- oragwelr

## Repository

[lem-in Repository](https://learn.zone01kisumu.ke/git/hshikuku/lem-in.git)

