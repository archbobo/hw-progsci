package main

import (
  "fmt"
  "os"
  "strconv"
)

type Board struct {
  cell [][]int
  pile int
}

// CreateBoard that returns a newboard with the right dimensions
// and the initial configuration.
func CreateBoard(size int, pile int) *Board {
  var b Board
  b.cell = make([][]int, size)
  for i := 0; i < size; i++ {
    b.cell[i] = make([]int, size)
  }
  b.pile = pile
  halfSize := size / 2
  b.cell[halfSize][halfSize] = pile
  return &b
}

func (b *Board) ComputeSteadyState() {
  for !b.IsConverged() {
    for i := 0; i < len(b.cell); i++ {
      for j := 0; j < len(b.cell); j++ {
        b.Topple(i, j)
      }
    }
  }
}

func (b *Board) DrawBoard() {
  pic := CreateNewCanvas(len(b.cell), len(b.cell))
  for i := 0; i < len(b.cell); i++ {
    for j := 0; j < len(b.cell[0]); j++ {
      b.UnitDraw(pic, i, j, b.Cell(i, j))
    }
  }
  pic.SaveToPNG("board.png")
}

func (b *Board) UnitDraw(can Canvas, r, c, number int) {
  can.SetLineWidth(0.0)
  switch number {
    case 0:
      can.SetStrokeColor(MakeColor(0, 0, 0))
      can.SetFillColor(MakeColor(0, 0, 0))
    case 1:
      can.SetStrokeColor(MakeColor(85, 85, 85))
      can.SetFillColor(MakeColor(85, 85, 85))
    case 2:
      can.SetStrokeColor(MakeColor(170, 170, 170))
      can.SetFillColor(MakeColor(170, 170, 170))
    case 3:
      can.SetStrokeColor(MakeColor(255, 255, 255))
      can.SetFillColor(MakeColor(255, 255, 255))
  }
  can.MoveTo(float64(r), float64(c))
  can.LineTo(float64(r + 1), float64(c))
  can.LineTo(float64(r + 1), float64(c + 1))
  can.LineTo(float64(r), float64(c + 1))
  can.LineTo(float64(r), float64(c))
  can.FillStroke()
}

// Topple(r, c int) that topples (r, c) until it cant be toppled any more.
// If square(r, c) has >= 4 coins on it, move 1 coin from (r, c) to each of the
// 4 neighbours of (r, c) (diagonal neighbours don't count, only north, south,
// east, and west). If square(r, c) has < 4 coins, do nothing.
func (b *Board) Topple(r, c int) {
  if b.cell[r][c] >= 4 {
    if b.Contains(r - 1, c) {
      b.Set(r - 1, c, b.cell[r - 1][c] + 1)
      b.Set(r, c, b.cell[r][c] - 1)
    } else {
      b.Set(r, c, b.cell[r][c] - 1)
    }
    if b.Contains(r, c - 1) {
      b.Set(r, c - 1, b.cell[r][c - 1] + 1)
      b.Set(r, c, b.cell[r][c] - 1)
    } else {
      b.Set(r, c, b.cell[r][c] - 1)
    }
    if b.Contains(r, c + 1) {
      b.Set(r, c + 1, b.cell[r][c + 1] + 1)
      b.Set(r, c, b.cell[r][c] - 1)
    } else {
      b.Set(r, c, b.cell[r][c] - 1)
    }
    if b.Contains(r + 1, c) {
      b.Set(r + 1, c, b.cell[r + 1][c] + 1)
      b.Set(r, c, b.cell[r][c] - 1)
    } else {
      b.Set(r, c, b.cell[r][c] - 1)
    }
  }
}

// Contains(r,c int) bool that returns true if (r, c) is within the field.
func (b *Board) Contains(r, c int) bool {
  if (r < b.NumRows()) && (c < b.NumCols()) && (r >= 0) && (c >= 0) {
    return true
  } else {
    return false
  }
}

// Set(r,c, value int) that sets the value of cell (r, c).
func (b *Board) Set(r, c, value int) {
  b.cell[r][c] = value
}

// Cell(r, c int) int that returns the value of the cell (r, c).
func (b *Board) Cell(r, c int) int {
  return b.cell[r][c]
}

// IsConverged() bool that returns true if there are no cells with  4 coins on them.
func (b *Board) IsConverged() bool {
  for r, col := range b.cell {
    for c, _ := range col {
      if b.Cell(r, c) >= 4 {
        return false
      }
    }
  }
  return true
}

// NumRows() int that returns the number of rows on the board.
func (b *Board) NumRows() int {
  return len(b.cell)
}

// NumCols() int that returns the number of columns on the board.
func (b *Board) NumCols() int {
  return len(b.cell[0])
}


func main() {
  size, errSize := strconv.Atoi(os.Args[1])
  pile, errPile := strconv.Atoi(os.Args[2])
  if (errSize != nil) || (errPile != nil) {
    fmt.Println("Error: Please input valid number for the size and pile")
    return
  }
  if (size < 0) || (pile < 0) {
    fmt.Println("Error: Please input positive integers for the size and pile")
    return
  }
  b := CreateBoard(size, pile)
  b.ComputeSteadyState()
  b.DrawBoard()
}
