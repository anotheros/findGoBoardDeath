package normal

import (
	"fmt"
	"gopkg.in/fatih/set.v0"
)

//FindGoBoardDeath1.
// 实现找出棋盘所有死子的算法。0：空白，1：黑子，2：白子。
// 两个同样颜色的子临近（在上下左右）计算成为一片棋子
// 一个或者一片棋子上下左右没有空白了就算是死子
// 以下棋盘矩阵中，B[0][2]是死子。
// B = [[0,1,2],
// [0,2,1],
// [0,0,0]]。
//返回所有死子坐标
func FindGoBoardDeath(board [][]int) [][2]int {
	captureSet := set.New(set.ThreadSafe)
	abstractBoard, _ := NewBoard(board)
	for index,value := range board {
		for i,_ := range value {
			if board[index][i]!= 0 {
				if(!captureSet.Has(Position{uint8(index),uint8(i)})){
					captures  :=abstractBoard.getNoLibertyStones(uint8(index),uint8(i))
					for _,cap:= range captures {
						captureSet.Add(cap)
					}
				}
			}
		}
	}
	var result [][2]int;
	for _,val := range captureSet.List() {
		aa := val.(Position);
		result =append(result, [2]int{int(aa.X),int(aa.Y)})
	}
	return result
}

type Position struct {
	X uint8
	Y uint8
}

type AbstractBoard struct {
	BoardSize uint8
	data      [][]int
}







func NewBoard(board [][]int) (*AbstractBoard, error) {
	boardSize := uint8(len(board))
	if boardSize < 1 {
		return nil, fmt.Errorf("Boardsize can not be less than 1!")
	}

	return &AbstractBoard{
		boardSize,
		board ,
	}, nil
}
func (board *AbstractBoard) getNeighbours(x uint8, y uint8) (neighbourIndexes []Position) {
	neighbourIndexes = []Position{}

	// Check for board borders
	if x > 0 {
		neighbourIndexes = append(neighbourIndexes, Position{(x - 1), y})
	}
	if x < board.BoardSize-1 {
		neighbourIndexes = append(neighbourIndexes, Position{(x + 1), y})
	}
	if y > 0 {
		neighbourIndexes = append(neighbourIndexes, Position{x, y - 1})
	}
	if y < board.BoardSize-1 {
		neighbourIndexes = append(neighbourIndexes, Position{x, y + 1})
	}

	return
}

func (board *AbstractBoard) getStatus(x uint8, y uint8) int {
	return board.data[x][y]
}

func (a *Position) isSamePosition(b Position) bool {
	return a.X == b.X && a.Y == b.Y
}


// Get all stones with no liberties left on given position
func (board *AbstractBoard) getNoLibertyStones(x uint8, y uint8) (noLibertyStones []Position) {
	//log.Printf("Get no liberty stones for (%d, %d)", x, y)

	noLibertyStones = []Position{}
	newlyFoundStones := []Position{Position{x, y}}
	foundNew := true
	var groupStones []Position = nil

	// Search until no new stones are found
	for foundNew == true {
		foundNew = false
		groupStones = []Position{}

		for _, newlyFoundStone := range newlyFoundStones {
			neighbours := board.getNeighbours(newlyFoundStone.X, newlyFoundStone.Y)

			// Check liberties of stone newlyFoundStone.X, newlyFoundStone.Y by checking the neighbours
			for _, neighbour := range neighbours {
				nbX := neighbour.X
				nbY := neighbour.Y

				// Has newlyFoundStone a free liberty?
				if board.getStatus(nbX, nbY) == 0 {
					// Neighbour is empty and not origPosition so newlyFoundStone has at least one liberty
					return noLibertyStones[:0]
				} else {
					// Is the neighbour of newlyFoundStone.X, newlyFoundStone.Y the same color? Then we have a group here
					if board.getStatus(newlyFoundStone.X, newlyFoundStone.Y) == board.getStatus(nbX, nbY) {
						foundNewHere := true
						nbGroupStone := Position{nbX, nbY}



						// Check if found stone is already in our group list
						for _, groupStone := range groupStones {
							if groupStone.isSamePosition(nbGroupStone) {
								foundNewHere = false
								break
							}
						}

						// Check if found stone is already in result set list
						if foundNewHere {
							for _, noLibertyStone := range noLibertyStones {
								if noLibertyStone.isSamePosition(nbGroupStone) {
									foundNewHere = false
									break
								}
							}
						}

						// If groupStone is not known yet, add it
						if foundNewHere {
							groupStones = append(groupStones, nbGroupStone)
							foundNew = true
						}
					}
				}
			}
		}

		// Add newly found stones to the resultset
		noLibertyStones = append(noLibertyStones, newlyFoundStones...)

		// Now check the found group stones
		newlyFoundStones = groupStones
	}

	return
}
