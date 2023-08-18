package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tetromino struct {
	shape [][]bool
	label rune
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <path_to_file>")
		return
	}

	tetrominoes, err := readTetrominoes(os.Args[1])
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	size := 4 // starting grid size
	label := 'A'
	for i := range tetrominoes {
		tetrominoes[i].label = label
		label++
	}

	for {
		grid := make([][]rune, size)
		for i := range grid {
			grid[i] = make([]rune, size)
			for j := range grid[i] {
				grid[i][j] = '.'
			}
		}

		if fitTetrominoes(tetrominoes, grid, size) {
			printGrid(grid)
			break
		}
		size++
	}
}

func readTetrominoes(path string) ([]Tetromino, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tetrominoes []Tetromino
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		if len(lines) == 4 {
			tetromino, err := linesToTetromino(lines)
			if err != nil {
				return nil, err
			}
			tetrominoes = append(tetrominoes, tetromino)
			lines = nil
		}
	}

	return tetrominoes, nil
}

func linesToTetromino(lines []string) (Tetromino, error) {
	shape := make([][]bool, 4)
	for i := range shape {
		shape[i] = make([]bool, 4)
		for j, ch := range lines[i] {
			if ch == '#' {
				shape[i][j] = true
			} else if ch != '.' {
				return Tetromino{}, fmt.Errorf("invalid character in tetromino")
			}
		}
	}
	return Tetromino{shape: shape}, nil
}

func fitTetrominoes(tetrominoes []Tetromino, grid [][]rune, size int) bool {
	if len(tetrominoes) == 0 {
		return true
	}

	t := tetrominoes[0]
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if canPlace(t, grid, i, j, size) {
				place(t, grid, i, j)
				if fitTetrominoes(tetrominoes[1:], grid, size) {
					return true
				}
				remove(t, grid, i, j)
			}
		}
	}
	return false
}

func canPlace(t Tetromino, grid [][]rune, x, y, size int) bool {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if t.shape[i][j] {
				if x+i >= size || y+j >= size || grid[x+i][y+j] != '.' {
					return false
				}
			}
		}
	}
	return true
}

func place(t Tetromino, grid [][]rune, x, y int) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if t.shape[i][j] {
				grid[x+i][y+j] = t.label
			}
		}
	}
}

func remove(t Tetromino, grid [][]rune, x, y int) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if t.shape[i][j] {
				grid[x+i][y+j] = '.'
			}
		}
	}
}

func printGrid(grid [][]rune) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}
