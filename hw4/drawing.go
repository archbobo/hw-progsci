package main

import (
  "os"
  "fmt"
  "math/rand"
  "math"
  "strconv"
)


/** The code to draw the chart ofPopsize
 *
 *
 */
func PopSize(r, x0 float64, max_t int) []float64 {
  var x float64 = x0
  var x_t []float64
  x_t = make([]float64, 0)
  for i := 0; i < max_t; i++ {
    if x < 0 {
      x = 0
    }
    x = r * x * (1 - x)
    // fmt.Println(x)
    x_t = append(x_t, x)
  }
  return x_t
}

func drawPopSize(r float64) {
  pic := CreateNewCanvas(500, 100)
  pic.SetStrokeColor(MakeColor(0, 0, 255))
  pic.SetLineWidth(1)
  var (
    x int
    y float64
  )
  var x_t []float64
  x_t = PopSize(r, 0.1, 100)
  pic.MoveTo(0, 100 - 100 * x_t[0])
  for i := 1; i < len(x_t); i++ {
    x = 5 * i
    y = 100 - 100 * x_t[i]
    pic.LineTo(float64(x), y)
  }
  pic.Stroke()
  pic.SaveToPNG("PopSize.png")
}

/** The code to draw the chart of randomwalk
 *
 *
 */

 func randDegree() float64 {
   return rand.Float64() * 360.0
 }

 func inField(coord, n float64) bool {
   return coord >= 0 && coord < n
 }

 func randStep(x, y, width, height, stepsize float64) (float64, float64) {
   var nx, ny float64 = x, y
   for (nx == x && ny == y) || !inField(nx, width) || !inField(ny, height) {
     // fmt.Println(nx, ny)
     rd := float64(randDegree())
     nx = x + stepsize * math.Cos(rd * math.Pi)
     ny = y + stepsize * math.Sin(rd * math.Pi)
   }

   return nx, ny
 }

 func drawRandomWalk(width, height, stepsize float64, steps int, seed int) {
   rand.Seed(int64(seed))
   pic := CreateNewCanvas(500, 500)
   pic.SetStrokeColor(MakeColor(0, 0, 0))
   pic.SetLineWidth(1)
   var x, y = width / 2, height / 2
   pic.MoveTo(x, y)
   for i := 0; i < steps; i++ {

     x, y = randStep(x, y, width, height, stepsize)

     pic.LineTo(x, y)
   }
   pic.Stroke()
   pic.SaveToPNG("RandomWalk.png")
 }

/** draw cellular automata
 *
 *
 */

func initialcell(width, steps int) [][]int {
  var cell [][]int = make([][]int, steps)
  for row := range cell {
    cell[row] = make([]int, width)
  }
  cell[0][int64(math.Floor(float64(width) / 2.0))] = 1
  // fmt.Println(cell[0])
  return cell
}

/*
 * the outer function used to call setNextLayer
 */
func fillCell(cell [][]int, rule int) {
  var ruletable [8]int
  // convert rule into table of array
  for i := 0; i < 8; i++ {
    ruletable[i] = rule % 10
    rule = rule / 10
  }
  for i := 0; i < len(cell) - 1; i++ {
    setNextLayer(cell[i], cell[i+1], ruletable)
  }
}

/*
 * Use the previous layer's information to set the next
 * layer
 */
func setNextLayer(currLayer, nextLayer []int, ruletable [8]int) {
  var newcurr []int
  newcurr = make([]int, 1)
  // fmt.Println(len(newcurr))
  for i := 0; i < len(currLayer); i++ {
    newcurr = append(newcurr, currLayer[i])
  }
  newcurr = append(newcurr, 0)
  // print current layer number of elements
  // fmt.Println(len(newcurr))
  for i := 0; i < len(nextLayer); i++ {
    // print reference value
    // fmt.Println(currLayer[i], currLayer[i + 1], currLayer[i + 2], i, len(nextLayer), len(currLayer))
    index_rule := newcurr[i] * 4 + newcurr[i + 1] * 2 + newcurr[i + 2]
    nextLayer[i] = ruletable[index_rule]
  }
  // test current implementation
  // fmt.Println(nextLayer)
}

/*
 * Print the final result
 */
// func drawCell() {
//   for i := 0; i < len(cell); i++ {
//     row := cell[i]
//     for j := 0; j < len(row); j++ {
//       if row[j] == 0 {
//         drawSquare
//       } else {
//         fmt.Print("#")
//       }
//     }
//     fmt.Print("\n")
//   }
// }

func drawCanvas(b Canvas, cell [][]int) {
  b.SetLineWidth(0.0)
  for i := 0; i < len(cell); i++ {
  //for i := 0; i < 2; i++ {
    row := cell[i]
    for j := 0; j < len(row); j++ {
    //for j := 0; j < 5; j++ {
      x1, y1 := float64(j * 5), float64(i * 5)
      x2, y2 := float64((j + 1) * 5), float64((i + 1) * 5)
      if row[j] == 0 {
        b.SetStrokeColor(MakeColor(255, 255, 0))
        b.SetFillColor(MakeColor(255, 255, 0))
        b.MoveTo(x1, y1)
        b.LineTo(x1, y2)
        b.LineTo(x2, y2)
        b.LineTo(x2, y1)
        b.LineTo(x1, y1)
        b.FillStroke()
      } else {
        b.SetStrokeColor(MakeColor(0, 0, 0))
        b.SetFillColor(MakeColor(0, 0, 0))
        b.MoveTo(x1, y1)
        b.LineTo(x1, y2)
        b.LineTo(x2, y2)
        b.LineTo(x2, y1)
        b.LineTo(x1, y1)
        b.FillStroke()
      }
    }
    // fmt.Print("\n")
  }
  // x1, y1 := float64(r) * h, float64(c) * w
  // x2, y2 := float64(r+5) * h, float64(c+5) * w
  // b.MoveTo(x1, y1)
  // b.LineTo(x1, y2)
  // b.LineTo(x2, y2)
  // b.LineTo(x2, y1)
  // b.LineTo(x1, y1)
  // b.FillStroke()
}

func drawCA(rule, width, steps int) {
  pic := CreateNewCanvas(500, 255)
  var cell [][]int
  cell = initialcell(width, steps + 1)
  fillCell(cell, rule)
  drawCanvas(pic, cell)
  pic.SaveToPNG("CA.png")
  // printCell(cell, )
}

func main() {
    r, err_r := strconv.ParseFloat(os.Args[1], 64)
    d, err_d := strconv.ParseFloat(os.Args[2], 64)
    rule, err_rule := strconv.Atoi(os.Args[3])
    switch {
      case err_r != nil:
        fmt.Println("Error: The type of r is not float.")
        return
      case err_d != nil:
        fmt.Println("Error: The type of d is not float.")
        return
      case err_rule != nil:
        fmt.Println("Error: The type of rule is not integer.")
        return
      case len(os.Args[3]) != 8:
        fmt.Println("Error: The length of rule is not eight.")
        return
      case r < 0:
        fmt.Println("Error: The range of r is not between 0 and 1.")
        return
      case d <= 0:
        fmt.Println("Error: d is less than 0.")
        return
    }
    drawPopSize(r)
    drawRandomWalk(500, 500, d, 1000, 12345)
    drawCA(rule, 100, 50)
}
