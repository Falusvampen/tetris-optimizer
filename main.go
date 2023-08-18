package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

var board [][]string

func main() {
	if len(os.Args) < 2 {
		return
	}

	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		printError(err)
		return
	}
	defer closeFile(file)

	tetrominoes, err := readInput(file)
	if err != nil {
		printError(err)
		return
	}

	solve(tetrominoes)
	printSolution()
}

func printError(err error) {
	fmt.Printf("ERROR: %s\n", err)
}

func closeFile(file *os.File) {
	if err := file.Close(); err != nil {
		printError(err)
	}
}

func readInput(file io.Reader) ([][4][4]string, error) {
	var tetrominoes [][4][4]string
	scanner := bufio.NewScanner(file)
	var tetromino [4][4]string
	lineCount, alphaIndex := 0, 0
	alpha := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	flag := true

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if lineCount%4 != 0 || !flag {
				return nil, fmt.Errorf("bad format: unexpected empty line")
			}
			flag = false
			continue
		}
		flag = true

		if len(line) != 4 {
			return nil, fmt.Errorf("bad format: tetromino not in 4x4 grid")
		}

		for i, char := range line {
			if char == '.' {
				tetromino[lineCount%4][i] = "."
			} else if char == '#' {
				tetromino[lineCount%4][i] = string(alpha[alphaIndex])
			} else {
				return nil, fmt.Errorf("invalid character in tetromino")
			}
		}

		lineCount++
		if lineCount%4 == 0 {
			if !isValidTetromino(tetromino) {
				return nil, fmt.Errorf("invalid tetromino shape")
			}
			tetrominoes = append(tetrominoes, optimizeTetromino(tetromino))
			alphaIndex++
		}
	}

	if lineCount%4 != 0 {
		return nil, fmt.Errorf("bad format: incomplete tetromino")
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return tetrominoes, nil
}

func isValidTetromino(tetromino [4][4]string) bool {
	adjacentCount, blockCount := 0, 0

	for i, row := range tetromino {
		for j, cell := range row {
			if cell != "." {
				blockCount++
				if i+1 < 4 && tetromino[i+1][j] != "." {
					adjacentCount++
				}
				if i-1 >= 0 && tetromino[i-1][j] != "." {
					adjacentCount++
				}
				if j+1 < 4 && tetromino[i][j+1] != "." {
					adjacentCount++
				}
				if j-1 >= 0 && tetromino[i][j-1] != "." {
					adjacentCount++
				}
			}
		}
	}

	return blockCount == 4 && (adjacentCount == 6 || adjacentCount == 8)
}

func optimizeTetromino(tetromino [4][4]string) [4][4]string {
	for isEmptyRow(tetromino[0]) {
		tetromino = shiftVertical(tetromino)
	}
	for isEmptyColumn(tetromino, 0) {
		tetromino = shiftHorizontal(tetromino)
	}
	return tetromino
}

func isEmptyRow(row [4]string) bool {
	for _, cell := range row {
		if cell != "." {
			return false
		}
	}
	return true
}

func isEmptyColumn(tetromino [4][4]string, col int) bool {
	for i := 0; i < 4; i++ {
		if tetromino[i][col] != "." {
			return false
		}
	}
	return true
}

func shiftVertical(tetromino [4][4]string) [4][4]string {
	temp := tetromino[0]
	copy(tetromino[0:], tetromino[1:])
	tetromino[3] = temp
	return tetromino
}

func shiftHorizontal(tetromino [4][4]string) [4][4]string {
	tetromino = transpose(tetromino)
	tetromino = shiftVertical(tetromino)
	return transpose(tetromino)
}

func transpose(slice [4][4]string) [4][4]string {
	var result [4][4]string
	for i := range slice {
		for j := range slice[i] {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func canInsert(i, j int, tetro [4][4]string) bool {
	for a := 0; a < 4; a++ {
		for b := 0; b < 4; b++ {
			if tetro[a][b] != "." {
				if i+a >= len(board) || j+b >= len(board) || board[i+a][j+b] != "." {
					return false
				}
			}
		}
	}
	return true
}

func insert(i, j int, tetro [4][4]string) {
	for a := 0; a < 4; a++ {
		for b := 0; b < 4; b++ {
			if tetro[a][b] != "." {
				board[i+a][j+b] = tetro[a][b]
			}
		}
	}
}

func remove(i, j int, tetro [4][4]string) {
	for a := 0; a < 4; a++ {
		for b := 0; b < 4; b++ {
			if tetro[a][b] != "." {
				board[i+a][j+b] = "."
			}
		}
	}
}

func backtrackSolver(tetrominoes [][4][4]string, n int) bool {
	if n == len(tetrominoes) {
		return true
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if canInsert(i, j, tetrominoes[n]) {
				insert(i, j, tetrominoes[n])
				if backtrackSolver(tetrominoes, n+1) {
					return true
				}
				remove(i, j, tetrominoes[n])
			}
		}
	}
	return false
}

func solve(tetrominoes [][4][4]string) {
	boardSize := int(math.Ceil(math.Sqrt(float64(4 * len(tetrominoes)))))
	board = initBoard(boardSize)
	for !backtrackSolver(tetrominoes, 0) {
		boardSize++
		board = initBoard(boardSize)
	}
}

func initBoard(size int) [][]string {
	newBoard := make([][]string, size)
	for i := range newBoard {
		newBoard[i] = make([]string, size)
		for j := range newBoard[i] {
			newBoard[i][j] = "."
		}
	}
	return newBoard
}

func printSolution() {
	for _, row := range board {
		for _, cell := range row {
			fmt.Printf("%s ", cell)
		}
		fmt.Println()
	}
}
