# Lem-in: Ant Farm Pathfinding
Student project

## Objectives

This project involves creating a digital version of an ant farm. 
The goal is to develop a program called `lem-in` that reads from a file describing ants and their colony, finds the quickest path for the ants, and displays their movements.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Features](#features)
- [Input Format](#input-format)
- [Examples](#examples)
- [Error Handling](#error-handling)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## Installation

To run the project, you need to have [Go](https://golang.org/) installed.

Clone the repository and navigate to the project directory:

```bash
git clone https://github.com/ClemNTTS/lem-in.git
cd lem-in
```

## Usage

The program reads from a file that describes the ant colony and its paths, finds the quickest way to move all ants from the `##start` room to the `##end` room, and displays the results.

```bash
go run . example.txt
```

## Features

- **Pathfinding Algorithm**: Implements an efficient algorithm to find the quickest path(s) for the ants.
- **Error Handling**: Detects and reports various input errors.
- **Visualization**: Displays the ants' movements in a step-by-step format.

## Input Format

The input file should follow these rules:

- The first line contains the number of ants.
- Rooms are defined by `name coord_x coord_y`.
- Links between rooms are defined by `name1-name2`.
- Comments start with `#`.
- Special rooms are marked with `##start` and `##end`.

### Example Input

```plaintext
3
##start
1 23 3
2 16 7
##end
0 9 5
0-4
0-6
1-3
4-3
5-2
3-5
4-2
2-1
7-6
7-2
7-4
6-5
```

## Examples

```bash
$ go run . test0.txt
```

example.txt
```plaintext
3
##start
1 23 3
2 16 7
3 16 3
4 16 5
5 9 3
6 1 5
7 4 8
##end
0 9 5
0-4
0-6
1-3
4-3
5-2
3-5
4-2
2-1
7-6
7-2
7-4
6-5

L1-3 L2-2
L1-4 L2-5 L3-3
L1-0 L2-6 L3-4
L2-0 L3-0
```


## Error Handling

The program will handle various errors related to input format:

- Missing `##start` or `##end` room.
- Invalid room coordinates.
- Invalid number of ants.
- Circular paths or rooms linking to themselves.
- Duplicated rooms or invalid links.

Example of error message:

```plaintext
ERROR: invalid data format, no start room found
```

## Contributing

To contribute:

1. Fork the repository.
2. Create a new branch.
3. Make your changes.
4. Submit a pull request.

## License

This project is licensed under the MIT License.

## Contact

For any questions or suggestions, feel free to reach out:

- Email: clement.nutt@gmail.com
- GitHub: [ClemNTTS](https://github.com/ClemNTTS)
