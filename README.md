# Tetris Optimizer

## Overview

The provided code is a Tetromino Solver written in Go. It takes a file containing tetromino shapes as input and attempts to fit them into the smallest possible square grid. Each tetromino shape is represented in a 4x4 grid format, where `#` represents a block and `.` represents an empty space. The program will then optimize the placement of these tetrominoes in a square grid and print the solution.

## Prerequisites

- Go (latest version recommended)

## How to Use

1. **Compile the Code** :

```bash

go build -o solver
```

2. **Run the Program** :

```bash

./solver <path_to_input_file>
```

Replace `<path_to_input_file>` with the path to your file containing tetromino shapes. 3. **Test the Code** :
If you want to test the code, you can use the provided `test.sh` script:

```bash

./test.sh
```

## Input File Format

The input file should contain tetromino shapes, each represented in a 4x4 grid format. Here's an example of how tetromino shapes should be formatted:

```bash

#...
#...
#...
#...

..#.
..#.
..#.
..#.
```

- `#` represents a block.
- `.` represents an empty space.
- Each tetromino shape should be separated by an empty line.

## Code Structure

- **Main Function** : The entry point of the program. It reads the input file and initiates the solving process.
- **Utility Functions** : Functions like `printError`, `closeFile`, and `readInput` help in reading the input file and handling errors.
- **Tetromino Validation** : Functions like `isValidTetromino` and `optimizeTetromino` validate and optimize the tetromino shapes.
- **Solver Logic** : The `solve` function uses a backtracking approach to fit the tetrominoes into the smallest possible square grid.
- **Board Manipulation** : Functions like `canInsert`, `insert`, `remove`, and `initBoard` help in manipulating the board during the solving process.
- **Output** : The `printSolution` function prints the optimized placement of tetrominoes in the grid.

## Limitations

- The program assumes that the input file is correctly formatted. Incorrectly formatted files may lead to unexpected behavior.
- The program uses a backtracking approach, which may not be the most efficient for very large input sizes.

## Contributing

If you find any bugs or have suggestions for improvements, please open an issue or submit a pull request.

## License

This code is provided under the MIT License. Please refer to the LICENSE file for more details.---

Thank you for using the Tetromino Solver! If you have any questions or feedback, please don't hesitate to reach out.
