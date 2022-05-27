package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Board [][]int

type EightPuzzle struct {
	solution    Board
	open        []Node
	closed      []Node
	openedNodes int
	closedNodes int
}

type Node struct {
	parent *Node
	board  Board
	g      int
	cost   float64
}

type Position struct {
	i, j int
}

func (ep *EightPuzzle) Init(board, solution Board) {
	node := Node{
		parent: nil,
		board:  board,
		g:      0,
	}
	ep.solution = solution
	node.cost = ep.f(node)
	ep.open = append(ep.open, node)
}

func GetInitialBoard() [][]int {
	scanner := bufio.NewScanner(os.Stdin)
	initialBoard := make(Board, 3)

	for i := 0; i < 3; i++ {
		initialBoard[i] = make([]int, 3)
	}

	for i := 0; i < 3; i++ {
		fmt.Printf("Digite a %da linha do tabuleiro com os valores separados por espaço. Exemplo: 8 0 3 (use 0 para o espaço vazio)\n", i+1)
		fmt.Printf(">> ")
		scanner.Scan()
		userInput := scanner.Text()
		for j, input := range strings.Fields(userInput) {
			p, err := strconv.Atoi(input)

			if err == nil {
				initialBoard[i][j] = p
			}
		}

	}
	return initialBoard
}

func (ep *EightPuzzle) f(node Node) float64 {
	return float64(node.g) + ep.h(node)
}

func (ep *EightPuzzle) GetTargetPosition(id int) Position {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if ep.solution[i][j] == id {
				return Position{i, j}
			}
		}
	}

	panic("the target board has missing values")
}

func (ep *EightPuzzle) h(node Node) float64 {
	h := 0.0

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if node.board[i][j] == 0 {
				continue
			}
			t := ep.GetTargetPosition(node.board[i][j])
			dx := math.Abs(float64(t.j - j))
			dy := math.Abs(float64(t.i - i))

			h += math.Sqrt(dx*dx + dy*dy)
		}
	}

	return h
}

func (ep *EightPuzzle) isSolved() bool {
	if len(ep.closed) == 0 {
		return false
	}

	b := ep.closed[len(ep.closed)-1].board

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b[i][j] != ep.solution[i][j] {
				return false
			}
		}
	}

	return true
}

func (ep *EightPuzzle) Close(node Node) {
	ep.closed = append(ep.closed, node)
	ep.open = ep.open[1:]
	ep.closedNodes++

	sort.SliceStable(ep.open, func(i, j int) bool {
		return ep.open[i].cost < ep.open[j].cost
	})
}

func (n *Node) GetEmptyTile() Position {
	return n.GetTile(0)
}

func (n *Node) GetTile(id int) Position {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if n.board[i][j] == id {
				return Position{i, j}
			}
		}
	}

	panic("the board has missing values")
}

func (n *Node) Movements() []Node {
	b0 := n.GetEmptyTile()
	movements := make([]Node, 0)

	if b0.i > 0 {
		movements = append(movements, n.Up())
	}

	if b0.i < 2 {
		movements = append(movements, n.Bottom())
	}

	if b0.j < 2 {
		movements = append(movements, n.Right())
	}

	if b0.j > 0 {
		movements = append(movements, n.Left())
	}

	return movements
}

func (ep *EightPuzzle) IsClosed(node Node) bool {
	for _, closed := range ep.closed {
		if closed.Equals(node) {
			return true
		}
	}
	return false
}

func (n *Node) Equals(other Node) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if n.board[i][j] != other.board[i][j] {
				return false
			}
		}
	}
	return true
}

func (n *Node) Up() Node {
	b0 := n.GetEmptyTile()
	board := n.CopyBoard()
	board[b0.i][b0.j] = board[b0.i-1][b0.j]
	board[b0.i-1][b0.j] = 0

	return Node{
		parent: n,
		board:  board,
		g:      n.g + 1,
	}
}

func (n *Node) Bottom() Node {
	b0 := n.GetEmptyTile()
	board := n.CopyBoard()
	board[b0.i][b0.j] = board[b0.i+1][b0.j]
	board[b0.i+1][b0.j] = 0

	return Node{
		parent: n,
		board:  board,
		g:      n.g + 1,
	}
}

func (n *Node) Right() Node {
	b0 := n.GetEmptyTile()
	board := n.CopyBoard()
	board[b0.i][b0.j] = board[b0.i][b0.j+1]
	board[b0.i][b0.j+1] = 0

	return Node{
		parent: n,
		board:  board,
		g:      n.g + 1,
	}
}

func (n *Node) Left() Node {
	b0 := n.GetEmptyTile()
	board := n.CopyBoard()
	board[b0.i][b0.j] = board[b0.i][b0.j-1]
	board[b0.i][b0.j-1] = 0

	return Node{
		parent: n,
		board:  board,
		g:      n.g + 1,
	}
}

func (n *Node) CopyBoard() Board {
	board := make(Board, 3)
	for i := 0; i < 3; i++ {
		board[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			board[i][j] = n.board[i][j]
		}
	}

	return board
}

func (ep *EightPuzzle) PrintPath(n Node) {
	if n.parent != nil {
		ep.PrintPath(*n.parent)
	}

	fmt.Print("\n\n\n")
	for _, row := range n.board {
		fmt.Println(row)
	}
	fmt.Println()
	fmt.Printf("g: %d\n", n.g)
	fmt.Printf("h: %f\n", n.cost-float64(n.g))
	fmt.Printf("f: %f", n.cost)
}

func main() {
	solution := Board{
		{1, 2, 3},
		{8, 0, 4},
		{7, 6, 5},
	}

	var puzzle EightPuzzle
	puzzle.Init(GetInitialBoard(), solution)

	for !puzzle.isSolved() {
		currentNode := puzzle.open[0]
		movements := currentNode.Movements()

		for _, movement := range movements {
			if !puzzle.IsClosed(movement) {
				movement.cost = puzzle.f(movement)
				puzzle.open = append(puzzle.open, movement)
				puzzle.openedNodes++
			}
		}

		puzzle.Close(currentNode)
	}

	fmt.Print("\n\n")
	puzzle.PrintPath(puzzle.closed[len(puzzle.closed)-1])
	fmt.Print("\n\n")
	fmt.Printf("Total de vértices abertos: %d\n", puzzle.openedNodes)
	fmt.Printf("Total de vértices fechados: %d\n", puzzle.closedNodes)
}
