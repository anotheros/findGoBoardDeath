package normal

import (
	"testing"
)

func hasV(vs [][2]int, v [2]int) bool {
	for _, n := range vs {
		if n == v {
			return true
		}
	}
	return false
}


func TestFindGoBoardDeath(t *testing.T) {
	board := [][]int{
		{0, 1, 1, 1},
		{0, 1, 1, 2},
		{2, 2, 1, 2},
		{1, 2, 1, 2},
	}
	vs := FindGoBoardDeath(board)
	if !hasV(vs, [2]int{1, 3}) || !hasV(vs, [2]int{2, 3}) || !hasV(vs, [2]int{3, 3}) || !hasV(vs, [2]int{3, 0}) {
		t.Error("FindGoBoardDeath test failed!")
	}

	board = [][]int{
		{0, 2, 1, 1, 1},
		{0, 2, 1, 2, 1},
		{0, 2, 1, 2, 1},
		{0, 0, 1, 1, 1},
	}

	vs = FindGoBoardDeath(board)
	if !hasV(vs, [2]int{1, 3}) || !hasV(vs, [2]int{2, 3}) {
		t.Error("FindGoBoardDeath test failed!")
	}
}

/**
func TestFindGoBoardDeath2(t *testing.T) {
	board := [][]int{
		{0, 1, 1, 1},
		{0, 1, 1, 2},
		{2, 2, 1, 2},
		{1, 2, 1, 2},
	}
	for index,value := range board {
		fmt.Print(index,":")
		for i,_ := range value {
			fmt.Print(board[index][i])
		}
		fmt.Println()
	}
	fmt.Println(len(board))
}

func TestFindGoBoardDeath2(t *testing.T) {
	board := []int{
		0, 1, 1, 1,
	}

	fmt.Println(board[:0])
}**/