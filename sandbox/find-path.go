// https://codefights.com/challenge/n927Ns3HD2q3EMeXB

package main

import "fmt"

// define data here...
type Cell struct {
    row int
    col int
    val int
}

func findOne(m [][]int) (*Cell, error) {
    for row := range m {
        for col := range m[row] {
            if m[row][col] == 1 {
                return &Cell{ row, col, 1 }, nil
            }
        }
    }

    return nil, fmt.Errorf("one not found in matrix...")
}

func findNext(m [][]int, cell *Cell) (*Cell, error) {
    y := cell.val + 1
    rowCount := len(m)
    colCount := len(m[0])

    // up
    if cell.row > 0 && m[cell.row - 1][cell.col] == y {
        return &Cell{cell.row - 1, cell.col, y}, nil
    }

    // down
    if cell.row < (rowCount - 1) && m[cell.row + 1][cell.col] == y {
        return &Cell{cell.row + 1, cell.col, y}, nil
    }

    // left
    if cell.col > 0 && m[cell.row][cell.col - 1] == y {
        return &Cell{cell.row, cell.col - 1, y}, nil
    }

    // right
    if cell.col < (colCount - 1) && m[cell.row][cell.col + 1] == y {
        return &Cell{cell.row, cell.col + 1, y}, nil
    }

    // not found...

    return nil, fmt.Errorf("not found...")
}

func findPath(matrix [][]int) bool {
    count := len(matrix) * len(matrix[0])

    cell, err := findOne(matrix)
    if err != nil {
        return false
    }

    fmt.Printf("cell %d:%d = %d\n", cell.row, cell.col, cell.val)

    for i := 1; i < count; i++ {
        cell, err = findNext(matrix, cell)
        if err != nil {
            return false
        }

        fmt.Printf("cell %d:%d = %d\n", cell.row, cell.col, cell.val)
    }

    return true
}

func main() {
    matrix := [][]int{{1, 4, 5}, {2, 3, 6}}
    findPath(matrix)
}

