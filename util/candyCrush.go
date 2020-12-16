// https://play.golang.org/p/Ffts4qrX_U
// implement with workers

package main

import (
	"fmt"
	"sync"
	"time"
)

func checkRow(row []string) bool {
	for i := 3; i <= len(row); i++ {
		sub := row[i-3 : i]

		if sub[0] == sub[1] && sub[1] == sub[2] {
			return true
		}
	}

	return false
}

func checkRows(board [][]string) bool {
	if len(board[0][:]) < 3 {
		return false
	}
	for row := 0; row < len(board); row++ {
		if checkRow(board[row]) == true {
			return true
		}

	}

	return false
}

func checkColumns(board [][]string) bool {
	rowCount, colCount := len(board), len(board[0][:])

	if rowCount < 3 {
		return false
	}

	for col := 0; col < colCount; col++ {
		r := make([]string, rowCount)
		for row := 0; row < rowCount; row++ {
			r[row] = board[row][col]
		}
		// fmt.Println(r)
		if checkRow(r) == true {
			return true
		}
	}

	return false
}

func play2(board [][]string) bool {
	var (
		workers sync.WaitGroup
		inrow   bool
		incol   bool
	)

	workers.Add(2)

	go func() {
		inrow = checkRows(board)
		workers.Done()
	}()

	go func() {
		incol = checkColumns(board)
		workers.Done()
	}()

	workers.Wait()
	return (inrow || incol)
}

func play1(board [][]string) bool {
	if checkRows(board) {
		return true
	}
	return checkColumns(board)
}

type Board struct {
	id     int
	grid   [][]string
	expect bool
}

func main() {

	boards := []Board{
		{id: 1, expect: true, grid: [][]string{{"R", "B", "G"}, {"R", "Y", "B"}, {"R", "O", "Y"}}},
		{id: 2, expect: true, grid: [][]string{{"R", "B", "G", "G", "G", "R", "B", "O", "P", "Y"}}},
		{id: 3, expect: false, grid: [][]string{{"R", "Y", "R", "Y"}, {"Y", "R", "Y", "R"}, {"R", "Y", "R", "Y"}, {"Y", "R", "Y", "R"}}},
		{id: 4, expect: true, grid: [][]string{{"R"}, {"B"}, {"P"}, {"Y"}, {"G"}, {"G"}, {"G"}, {"R"}, {"B"}}},
	}

	t0 := time.Now().UnixNano()
	for _, board := range boards {
		win := play1(board.grid)
		fmt.Printf("id %d) win? %v expect: %v\n", board.id, win, board.expect)
	}
	t1 := time.Now().UnixNano()
	fmt.Printf("elapsed: %f micro seconds\n", float64(t1-t0)/1e3)
}
