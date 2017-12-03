/*
* Shi Ruitao  2017-12-3
*/

package main

import (
	"fmt"
	// "strings"
)

func main() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s \n", board[i])
	}
	s := []int{1, 2}
	fmt.Printf("%d\n", s)
	s = append(s, 3, 4)
	fmt.Printf("s=%d, len=%d, cap=%d, \n", s, len(s), cap(s))
	for i, r := range s{
		fmt.Printf("index: %d, value: %d \n", i, r)
	}
}
