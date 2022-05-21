package matrix

import (
	"bufio"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix [][]int

func New(s string) (*Matrix, error) {
	var matrix Matrix
	scanner := bufio.NewScanner(strings.NewReader(s))

	for scanner.Scan() {
		var row []int
		line := scanner.Text()
		split := strings.Split(line, " ")

		for _, s := range split {
			number, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}
			row = append(row, number)
		}
		matrix = append(matrix, row)
	}

	return &matrix, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m *Matrix) Cols() [][]int {
	var cols [][]int
	for i := range *m {
		var row []int
		for x := range *m {
			row = append(row, (*m)[x][i])
		}
		cols = append(cols, row)
	}
	return cols
}

func (m *Matrix) Rows() [][]int {
	return *m
}

func (m *Matrix) Set(row, col, val int) bool {
	if len(*m) < row || len(*m) < col {
		return false
	}

	(*m)[row][col] = val

	return true
}
