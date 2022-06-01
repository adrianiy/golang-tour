package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix [][]int

func splitRows(rows []string) [][]string {
	var splitted [][]string

	for _, row := range rows {
		add := strings.Split(strings.TrimSpace(row), " ")
		splitted = append(splitted, add)
	}

	return splitted
}

// Function New creates a type Matrix from a string
func New(s string) (*Matrix, error) {
	var matrix Matrix
	rows := strings.Split(strings.ReplaceAll(s, "\r\n", "\n"), "\n")
	splittedRows := splitRows(rows)
	rowLength := len(splittedRows[0])

	for _, line := range splittedRows {
		var row []int

		if rowLength != len(line) {
			return nil, errors.New("No valid length")
		}

		for _, s := range line {
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
	columns := make([][]int, 0)

	for _, row := range *m {
		for col, number := range row {
			rowCopy := []int{number}
			if len(columns) >= col+1 {
				columns[col] = append(columns[col], rowCopy...)
			} else {
				columns = append(columns, rowCopy)
			}
		}
	}

	return columns
}

func (m *Matrix) Rows() [][]int {
	rows := make([][]int, len(*m))

	for i, row := range *m {
		rowCopy := make([]int, len(row))
		copy(rowCopy, row)
		rows[i] = rowCopy
	}

	return rows
}

func (m *Matrix) Set(row, col, val int) bool {
	if row < 0 || col < 0 || len(*m) <= row || len((*m)[0]) <= col {
		return false
	}

	(*m)[row][col] = val

	return true
}
