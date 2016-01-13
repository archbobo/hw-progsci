package main

import (
  "fmt"
  "os"
  "math"
  "strconv"
)

/*
 * Print the final result
 */
func printCell(cell [][]int) {
  for i := 0; i < len(cell); i++ {
    row := cell[i]
    for j := 0; j < len(row); j++ {
      if row[j] == 0 {
        fmt.Print(" ")
      } else {
        fmt.Print("#")
      }
    }
    fmt.Print("\n")
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
 * the function used to initialize cell diagram
 * , and make the middle element in the first row labeled
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

func main() {
  if len(os.Args) != 4 {
    fmt.Println("Error: informative error message here")
    return
  }
  if len(os.Args[1]) != 8 {
    fmt.Println("Error: informative error message here")
    return
  }
  rule, err_rule := strconv.Atoi(os.Args[1])
  width, err_width := strconv.Atoi(os.Args[2])
  steps, err_steps := strconv.Atoi(os.Args[3])
  if (err_rule != nil) || (err_width != nil) ||
  (err_steps != nil) || (width < 0) || (steps < 0){
    fmt.Println("Error: informative error message here")
    return
  }
  var cell [][]int
  cell = initialcell(width, steps + 1)
  // Print cell
  // fmt.Println(cell, rule)
  fillCell(cell, rule)
  printCell(cell)
  // fmt.Println(steps, len(cell))
}
