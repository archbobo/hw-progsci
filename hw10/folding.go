package main

import (
  "fmt"
  "os"
  "math/rand"
  "strings"
  "math"
  // "strconv"
  // "code.google.com/p/draw2d/draw2d"
)

const RADIUS float64 = 10.0
const EDGE_LENGTH float64 = 50.0
const CANVAS_SIZE int = 10000

func main() {
  if len(os.Args) != 2 {
    fmt.Println("Error: Please input the appropriate arguments!")
    return
  }
  // Test for parsing the arguments from the command line
  prot_seq := os.Args[1]
  // fmt.Println("HP protein sequence is " + prot_seq)
  // Test for drawing fold
  drawFold(prot_seq, randomFold(os.Args[1]))
  // Test for drawH
  // pic := CreateNewCanvas(500, 500)
  // x := 25.0
  // y := 250.0
  // drawP(pic, x, y)
  // pic.SaveToPNG("fold.png")
}

// Generate random fold
func randomFold(seq string) []string {
  direct := []string{"r", "l", "f"}
  fold := make([]string, 0)
  for _ = range seq {
    fold = append(fold, direct[rand.Intn(3)])
  }
  // Test for random fold
  // fmt.Println("Random Fold can be " + "<" + strings.Join(fold, ",") + ">")
  return fold
}

func drawFold(seq string, fold []string) {
  pic := CreateNewCanvas(CANVAS_SIZE, CANVAS_SIZE)
  degree := 0
  x := float64(CANVAS_SIZE) / 2.0
  y := float64(CANVAS_SIZE) / 2.0
  pic.SetLineWidth(1.0)
  pic.MoveTo(x, y)
  seq_arr := strings.Split(seq, "")
  for i := 0; i < len(seq_arr); i++ {
    switch seq_arr[i] {
      case "H" :
        drawH(pic, x, y)
      case "P" :
        drawP(pic, x, y)
    }
    drawEdge(pic, &x, &y, fold[i], &degree)
    // fmt.Println("@@@@x, y" + strconv.Itoa(int(x)) + "," + strconv.Itoa(int(y)))
  }
  pic.SaveToPNG("fold.png")
}

// Use 0, 1, 2, 3 to denote the four directions for the current direction.
// 0 means UP. 1 means right. 2 means down. 3 means left.
func drawEdge(b Canvas, x, y *float64, dir string, degree *int) {
  b.SetLineWidth(1)
  b.SetStrokeColor(MakeColor(0, 0, 0))
  b.MoveTo(*x, *y)
  switch dir {
    case "r" : {
      switch *degree {
        case 0 : {
          *x += EDGE_LENGTH
          b.LineTo(*x, *y)
          *degree = 1
        }
        case 1 : {
          *y -= EDGE_LENGTH
          b.LineTo(*x, *y)
          *degree = 2
        }
        case 2 : {
          *x -= EDGE_LENGTH
          b.LineTo(*x, *y)
          *degree = 3
        }
        case 3 : {
          *y += EDGE_LENGTH
          b.LineTo(*x, *y)
          *degree = 0
        }
      }
      b.FillStroke()
    }
    case "l" : {
      switch *degree {
        case 0 : {
          *x -= EDGE_LENGTH
          b.LineTo(*x, *y)
          *degree = 3
        }
        case 1 : {
          *y += EDGE_LENGTH
          b.LineTo(*x, *y)
          *degree = 0
        }
        case 2 : {
          *x += EDGE_LENGTH
          b.LineTo(*x, *y)
          *degree = 1
        }
        case 3 : {
          *y -= EDGE_LENGTH
          b.LineTo(*x, *y)
          *degree = 2
        }
      }
      b.FillStroke()
    }
    case "f" : {
      switch *degree {
        case 0 : {
          *y += EDGE_LENGTH
          b.LineTo(*x, *y)
          // degree = 3
        }
        case 1 : {
          *x += EDGE_LENGTH
          b.LineTo(*x, *y)
          // degree = 0
        }
        case 2 : {
          *y -= EDGE_LENGTH
          b.LineTo(*x, *y)
          // sdegree = 1
        }
        case 3 : {
          *x -= EDGE_LENGTH
          b.LineTo(*x, *y)
          // degree = 2
        }
      }
      b.FillStroke()
    }
  }
}

// H residue
func drawH(b Canvas, x, y float64) {
  b.SetLineWidth(1)
  b.SetStrokeColor(MakeColor(0, 0, 0))
  b.SetFillColor(MakeColor(0, 0, 0))
  b.MoveTo(x, y)
  b.gc.ArcTo(x, y, RADIUS, RADIUS, 0, math.Pi * 2)
  b.FillStroke()
}

// P residue
func drawP(b Canvas, x, y float64) {
  b.SetLineWidth(1)
  b.SetStrokeColor(MakeColor(0, 0, 0))
  b.SetFillColor(MakeColor(255, 255, 255))
  b.MoveTo(x, y)
  b.gc.ArcTo(x, y, RADIUS, RADIUS, 0, math.Pi * 2)
  b.FillStroke()
}
