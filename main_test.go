package main

import (
	"encoding/json"
	"testing"
)

func (p *EightPuzzle) pathToJson() []byte {
	path := make([]Board, 0)
	node := &p.closed[len(p.closed)-1]
	for node != nil {
		path = append(path, node.board)
		node = node.parent
	}

	v, _ := json.Marshal(path)

	return v
}

func TestSolvingBoard1(t *testing.T) {
	solution := Board{
		{1, 2, 3},
		{8, 0, 4},
		{7, 6, 5},
	}

	var puzzle EightPuzzle
	puzzle.Init(Board{
		{0, 3, 2},
		{1, 4, 6},
		{8, 7, 5},
	}, solution)

	for !puzzle.isSolved() {
		currentNode := puzzle.open[0]
		movements := currentNode.Movements()

		for _, movement := range movements {
			if !puzzle.IsClosed(movement) {
				movement.cost = puzzle.f(movement)
				puzzle.open = append(puzzle.open, movement)
			}
		}

		puzzle.Close(currentNode)
	}

	expected := "[[[1,2,3],[8,0,4],[7,6,5]],[[1,0,3],[8,2,4],[7,6,5]],[[1,3,0],[8,2,4],[7,6,5]],[[1,3,4],[8,2,0],[7,6,5]],[[1,3,4],[8,0,2],[7,6,5]],[[1,3,4],[8,6,2],[7,0,5]],[[1,3,4],[8,6,2],[0,7,5]],[[1,3,4],[0,6,2],[8,7,5]],[[0,3,4],[1,6,2],[8,7,5]],[[3,0,4],[1,6,2],[8,7,5]],[[3,4,0],[1,6,2],[8,7,5]],[[3,4,2],[1,6,0],[8,7,5]],[[3,4,2],[1,0,6],[8,7,5]],[[3,0,2],[1,4,6],[8,7,5]],[[0,3,2],[1,4,6],[8,7,5]]]"

	if string(puzzle.pathToJson()) != expected {
		t.Error("board 1 failed")
	}
}

func TestSolvingBoard2(t *testing.T) {
	solution := Board{
		{1, 2, 3},
		{8, 0, 4},
		{7, 6, 5},
	}

	var puzzle EightPuzzle
	puzzle.Init(Board{
		{0, 3, 2},
		{1, 4, 6},
		{8, 7, 5},
	}, solution)

	for !puzzle.isSolved() {
		currentNode := puzzle.open[0]
		movements := currentNode.Movements()

		for _, movement := range movements {
			if !puzzle.IsClosed(movement) {
				movement.cost = puzzle.f(movement)
				puzzle.open = append(puzzle.open, movement)
			}
		}

		puzzle.Close(currentNode)
	}

	expected := "[[[1,2,3],[8,0,4],[7,6,5]],[[1,0,3],[8,2,4],[7,6,5]],[[1,3,0],[8,2,4],[7,6,5]],[[1,3,4],[8,2,0],[7,6,5]],[[1,3,4],[8,0,2],[7,6,5]],[[1,3,4],[8,6,2],[7,0,5]],[[1,3,4],[8,6,2],[0,7,5]],[[1,3,4],[0,6,2],[8,7,5]],[[0,3,4],[1,6,2],[8,7,5]],[[3,0,4],[1,6,2],[8,7,5]],[[3,4,0],[1,6,2],[8,7,5]],[[3,4,2],[1,6,0],[8,7,5]],[[3,4,2],[1,0,6],[8,7,5]],[[3,0,2],[1,4,6],[8,7,5]],[[0,3,2],[1,4,6],[8,7,5]]]"

	if string(puzzle.pathToJson()) != expected {
		t.Error("board 2 failed")
	}
}

func TestSolvingBoard3(t *testing.T) {
	solution := Board{
		{1, 2, 3},
		{8, 0, 4},
		{7, 6, 5},
	}

	var puzzle EightPuzzle
	puzzle.Init(Board{
		{8, 3, 5},
		{4, 1, 6},
		{2, 7, 0},
	}, solution)

	for !puzzle.isSolved() {
		currentNode := puzzle.open[0]
		movements := currentNode.Movements()

		for _, movement := range movements {
			if !puzzle.IsClosed(movement) {
				movement.cost = puzzle.f(movement)
				puzzle.open = append(puzzle.open, movement)
			}
		}

		puzzle.Close(currentNode)
	}

	expected := "[[[1,2,3],[8,0,4],[7,6,5]],[[1,0,3],[8,2,4],[7,6,5]],[[0,1,3],[8,2,4],[7,6,5]],[[8,1,3],[0,2,4],[7,6,5]],[[8,1,3],[2,0,4],[7,6,5]],[[8,1,3],[2,4,0],[7,6,5]],[[8,1,3],[2,4,5],[7,6,0]],[[8,1,3],[2,4,5],[7,0,6]],[[8,1,3],[2,4,5],[0,7,6]],[[8,1,3],[0,4,5],[2,7,6]],[[8,1,3],[4,0,5],[2,7,6]],[[8,0,3],[4,1,5],[2,7,6]],[[8,3,0],[4,1,5],[2,7,6]],[[8,3,5],[4,1,0],[2,7,6]],[[8,3,5],[4,1,6],[2,7,0]]]"

	if string(puzzle.pathToJson()) != expected {
		t.Error("board 3 failed")
	}
}

func TestSolvingBoard4(t *testing.T) {
	solution := Board{
		{1, 2, 3},
		{8, 0, 4},
		{7, 6, 5},
	}

	var puzzle EightPuzzle
	puzzle.Init(Board{
		{0, 2, 3},
		{1, 4, 5},
		{8, 7, 6},
	}, solution)

	for !puzzle.isSolved() {
		currentNode := puzzle.open[0]
		movements := currentNode.Movements()

		for _, movement := range movements {
			if !puzzle.IsClosed(movement) {
				movement.cost = puzzle.f(movement)
				puzzle.open = append(puzzle.open, movement)
			}
		}

		puzzle.Close(currentNode)
	}

	expected := "[[[1,2,3],[8,0,4],[7,6,5]],[[1,2,3],[8,4,0],[7,6,5]],[[1,2,3],[8,4,5],[7,6,0]],[[1,2,3],[8,4,5],[7,0,6]],[[1,2,3],[8,4,5],[0,7,6]],[[1,2,3],[0,4,5],[8,7,6]],[[0,2,3],[1,4,5],[8,7,6]]]"

	if string(puzzle.pathToJson()) != expected {
		t.Error("board 4 failed")
	}
}
