package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*===============================================================
 * Functions to manipulate a "field" of cells --- the main data
 * that must be managed by this program.
 *==============================================================*/

// The data stored in a single cell of a field
type Cell struct {
	kind  string
	
	// This variable used to store the kind value from the previous field
	preKind string
	score float64
}

// createField should create a new field of the ysize rows and xsize columns,
// so that field[r][c] gives the Cell at position (r,c).
func createField(rsize, csize int) [][]Cell {
	f := make([][]Cell, rsize)
	for i := range f {
		f[i] = make([]Cell, csize)
	}
	return f
}

// inField returns true iff (row,col) is a valid cell in the field
// This method has not been used in this program
func inField(field [][]Cell, row, col int) bool {
	return row >= 0 && row < len(field) && col >= 0 && col < len(field[0])
}

// readFieldFromFile should open the given file and read the initial
// values for the field. The first line of the file will contain
// two space-separated integers saying how many rows and columns
// the field should have:
//    10 15
// each subsequent line will consist of a string of Cs and Ds, which
// are the initial strategies for the cells:
//    CCCCCCDDDCCCCCC
//
// If there is ever an error reading, this function should cause the
// program to quit immediately.
// read the field
// field := readFieldFromFile(fieldFile)
// fmt.Println("Field dimensions are:", len(field), "by", len(field[0]))

func readFieldFromFile(filename string) [][]Cell {
  // WRITE YOUR CODE HERE
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error: something went wrong opening the file.")
		fmt.Println("Probably you gave the wrong filename.")
	}
	scanner := bufio.NewScanner(file)
	// move the index to point the first line
	scanner.Scan()
	var items []string = strings.Split(scanner.Text(), " ")
	rsize, _ := strconv.Atoi(items[0])
	csize, _ := strconv.Atoi(items[1])
	// use the given number of rows and number of cols to create the field
	field := createField(rsize, csize)
	// copy the corresponding value into the field
	var i int = 0
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		for j := 0; j < len(line); j++{
			//fmt.Println(line[j])
			field[i][j].kind = string(line[j])
		}
		i++
	}
	file.Close()
	return field
	// return nil // This is included only so this template will compile
}

// drawUnitField is used by drawField to draw single rectangle shape
func drawUnitField (b Canvas, rindex, cindex int, kind string) {
	if strings.EqualFold(kind, "D") {
		b.SetStrokeColor(MakeColor(255, 0, 0))
		b.SetFillColor(MakeColor(255, 0, 0))
	} else {
		b.SetStrokeColor(MakeColor(0, 0, 255))
		b.SetFillColor(MakeColor(0, 0, 255))
	}
	row := float64(rindex)
	col := float64(cindex)
	b.MoveTo(row * 5.0, col * 5.0)
	b.LineTo(row * 5.0 + 5.0, col * 5.0)
	b.LineTo(row * 5.0 + 5.0, col * 5.0 + 5.0)
	b.LineTo(row * 5.0, col * 5.0 + 5.0)
	b.LineTo(row * 5.0, col * 5.0)
	b.FillStroke()
}

// drawField should draw a representation of the field on a canvas and save the
// canvas to a PNG file with a name given by the parameter filename.  Each cell
// in the field should be a 5-by-5 square, and cells of the "D" kind should be
// drawn red and cells of the "C" kind should be drawn blue.
func drawField(field [][]Cell, filename string) {
  // WRITE YOUR CODE HERE
	pic := CreateNewCanvas(500, 500)
	// Visit every unit in the field and draw the corresponding brick
	for i := 0; i < len(field); i++ {
		line := field[i]
		for j := 0; j < len(line); j++ {
			drawUnitField(pic, i, j, line[j].kind)
		}
	}
	pic.SaveToPNG(filename)
}

// drawUnitFieldColor is used to draw the single rectangle shape of color scheme
func drawUnitFieldColor (b Canvas, rindex, cindex int, kind string, preKind string) {
	if strings.EqualFold(kind, "C") && strings.EqualFold(preKind, "C") {
		b.SetStrokeColor(MakeColor(0, 0, 255))
		b.SetFillColor(MakeColor(0, 0, 255))
	} else if strings.EqualFold(kind, "D") && strings.EqualFold(preKind, "D"){
		b.SetStrokeColor(MakeColor(255, 0, 0))
		b.SetFillColor(MakeColor(255, 0, 0))
	} else if strings.EqualFold(kind, "D") && strings.EqualFold(preKind, "C"){
		b.SetStrokeColor(MakeColor(255, 255, 0))
		b.SetFillColor(MakeColor(255, 255, 0))
	} else if strings.EqualFold(kind, "C") && strings.EqualFold(preKind, "D"){
		b.SetStrokeColor(MakeColor(0, 255, 0))
		b.SetFillColor(MakeColor(0, 255, 0))
	}
	row := float64(rindex)
	col := float64(cindex)
	b.MoveTo(row * 5.0, col * 5.0)
	b.LineTo(row * 5.0 + 5.0, col * 5.0)
	b.LineTo(row * 5.0 + 5.0, col * 5.0 + 5.0)
	b.LineTo(row * 5.0, col * 5.0 + 5.0)
	b.LineTo(row * 5.0, col * 5.0)
	b.FillStroke()
}

// drawField should draw a color scheme representation of the field on a canvas and save the
// canvas to a PNG file with a name given by the parameter filename.  Each cell
// in the field should be a 5-by-5 square, and cells of the "D" kind should be
// drawn red and cells of the "C" kind should be drawn blue.
func drawFieldColor (field [][]Cell, filename string) {
		// WRITE YOUR CODE HERE
	pic := CreateNewCanvas(500, 500)
	for i := 0; i < len(field); i++ {
		line := field[i]
		for j := 0; j < len(line); j++ {
			drawUnitFieldColor(pic, i, j, line[j].kind, line[j].preKind)
		}
	}
	pic.SaveToPNG(filename)
}

/*===============================================================
 * Functions to simulate the spatial games
 *==============================================================*/

// play a game between a cell of type "me" and a cell of type "them" (both me
// and them should be either "C" or "D"). This returns the reward that "me"
// gets when playing against them.
func gameBetween(me, them string, b float64) float64 {
	if me == "C" && them == "C" {
		return 1
	} else if me == "C" && them == "D" {
		return 0
	} else if me == "D" && them == "C" {
		return b
	} else if me == "D" && them == "D" {
		return 0
	} else {
		fmt.Println("type ==", me, them)
		panic("This shouldn't happen")
	}
}

// updateScores goes through every cell, and plays the Prisoner's dilema game
// with each of it's in-field nieghbors (including itself). It updates the
// score of each cell to be the sum of that cell's winnings from the game.
func updateScores(field [][]Cell, b float64) {
    // WRITE YOUR CODE HERE
		var sum float64 = 0.0
		for i := 0; i < len(field); i++ {
			for j := 0; j < len(field[i]); j++ {
				switch {
					// when unit_field is at the upper left corner
					case i == 0 && j == 0 : {
						sum = (gameBetween(field[i][j].kind, field[i + 1][j].kind, b) +
						gameBetween(field[i][j].kind, field[i + 1][j + 1].kind, b) +
						gameBetween(field[i][j].kind, field[i][j + 1].kind, b) +
						gameBetween(field[i][j].kind, field[i][j].kind, b))
					}

					// when unit_field is the top level
					case (i == 0) && (j != (len(field[i]) - 1)) : {
						sum = (gameBetween(field[i][j].kind, field[i][j].kind, b) +
						gameBetween(field[i][j].kind, field[i][j + 1].kind, b) +
						gameBetween(field[i][j].kind, field[i][j - 1].kind, b) +
						gameBetween(field[i][j].kind, field[i + 1][j].kind, b) +
						gameBetween(field[i][j].kind, field[i + 1][j + 1].kind, b) +
						gameBetween(field[i][j].kind, field[i + 1][j - 1].kind, b))
					}

					// left most level
					case (j == 0) && (i != (len(field) - 1)) : {
						sum = (gameBetween(field[i][j].kind, field[i][j].kind, b) +
						gameBetween(field[i][j].kind, field[i + 1][j].kind, b) +
						gameBetween(field[i][j].kind, field[i - 1][j].kind, b) +
						gameBetween(field[i][j].kind, field[i][j + 1].kind, b) +
						gameBetween(field[i][j].kind, field[i + 1][j + 1].kind, b) +
						gameBetween(field[i][j].kind, field[i - 1][j + 1].kind, b))
					}

					// right lower corner
					case (i == (len(field) - 1)) && (j == (len(field[i]) - 1)) : {
						sum = (gameBetween(field[i][j].kind, field[i][j - 1].kind, b) +
						gameBetween(field[i][j].kind, field[i - 1][j].kind, b) +
						gameBetween(field[i][j].kind, field[i - 1][j - 1].kind, b) +
						gameBetween(field[i][j].kind, field[i][j].kind, b))
					}

					// right most level
					case (i == (len(field) - 1)) && (j != 0): {
						sum = (gameBetween(field[i][j].kind, field[i][j].kind, b) +
						gameBetween(field[i][j].kind, field[i][j - 1].kind, b) +
						gameBetween(field[i][j].kind, field[i - 1][j].kind, b) +
						gameBetween(field[i][j].kind, field[i][j + 1].kind, b) +
						gameBetween(field[i][j].kind, field[i - 1][j - 1].kind, b) +
						gameBetween(field[i][j].kind, field[i - 1][j + 1].kind, b))
					}

					//down most level
					case (j == (len(field[i]) - 1)) && (i != 0): {
						sum = (gameBetween(field[i][j].kind, field[i][j].kind, b) +
						gameBetween(field[i][j].kind, field[i][j - 1].kind, b) +
						gameBetween(field[i][j].kind, field[i - 1][j].kind, b) +
						gameBetween(field[i][j].kind, field[i + 1][j].kind, b) +
						gameBetween(field[i][j].kind, field[i + 1][j - 1].kind, b) +
						gameBetween(field[i][j].kind, field[i - 1][j - 1].kind, b))
					}

					// right upper corner
					case (i == 0) && (j == (len(field[i]) - 1)) : {
						sum = (gameBetween(field[i][j].kind, field[i][j - 1].kind, b) +
						gameBetween(field[i][j].kind, field[i + 1][j].kind, b) +
						gameBetween(field[i][j].kind, field[i + 1][j - 1].kind, b) +
						gameBetween(field[i][j].kind, field[i][j].kind, b))
					}

					// left lower corner
					case (i == len(field) - 1) && (j == 0) : {
						sum = (gameBetween(field[i][j].kind, field[i][j + 1].kind, b) +
						gameBetween(field[i][j].kind, field[i - 1][j + 1].kind, b) +
						gameBetween(field[i][j].kind, field[i - 1][j].kind, b) +
						gameBetween(field[i][j].kind, field[i][j].kind, b))
					}

					default : {
						sum = (gameBetween(field[i][j].kind, field[i][j].kind, b) +
						gameBetween(field[i][j].kind, field[i][j - 1].kind, b) +
						gameBetween(field[i][j].kind, field[i - 1][j].kind, b) +
						gameBetween(field[i][j].kind, field[i + 1][j].kind, b) +
						gameBetween(field[i][j].kind, field[i + 1][j - 1].kind, b) +
						gameBetween(field[i][j].kind, field[i - 1][j - 1].kind, b) +
						gameBetween(field[i][j].kind, field[i - 1][j + 1].kind, b) +
						gameBetween(field[i][j].kind, field[i][j + 1].kind, b) +
						gameBetween(field[i][j].kind, field[i + 1][j + 1].kind, b))
					}
				}
				field[i][j].score = sum
			}
	 }
}

// updateStrategies create a new field by going through every cell (r,c), and
// looking at each of the cells in its neighborhood (including itself) and the
// setting the kind of cell (r,c) in the new field to be the kind of the
// neighbor with the largest score
func updateStrategies(field [][]Cell) [][]Cell {
  // WRITE YOUR CODE HERE
	newField := createField(len(field), len(field[0]))
	var temp Cell
	temp.score = 0
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			curr_score := field[i][j].score
			temp.kind = field[i][j].kind
			newField[i][j].preKind = field[i][j].kind
			switch {
				// left upper corner
				case i == 0 && j == 0 : {
					if field[i + 1][j].score > curr_score {
						temp.kind = string(field[i + 1][j].kind)
						curr_score = field[i + 1][j].score
					}
					if field[i + 1][j + 1].score > curr_score {
						temp.kind = string(field[i + 1][j + 1].kind)
						curr_score = field[i + 1][j + 1].score
					}
					if field[i][j + 1].score > curr_score {
						temp.kind = string(field[i][j + 1].kind)
						curr_score = field[i][j + 1].score
					}
				}

				// upper most line
				case (i == 0) && (j != (len(field[i]) - 1)): {
					if field[i][j + 1].score > curr_score {
						temp.kind = string(field[i][j + 1].kind)
						curr_score = field[i][j + 1].score
					}
					if field[i + 1][j].score > curr_score {
						temp.kind = string(field[i + 1][j].kind)
						curr_score = field[i + 1][j].score
					}
					if field[i][j - 1].score > curr_score {
						temp.kind = string(field[i][j - 1].kind)
						curr_score = field[i][j - 1].score
					}
					if field[i + 1][j + 1].score > curr_score {
						temp.kind = string(field[i + 1][j + 1].kind)
						curr_score = field[i + 1][j + 1].score
					}
					if field[i + 1][j - 1].score > curr_score {
						temp.kind = string(field[i + 1][j - 1].kind)
						curr_score = field[i + 1][j - 1].score
					}
				}

				// left most line
				case (j == 0) && (i != (len(field) - 1)): {
					if field[i - 1][j].score > curr_score {
						temp.kind = string(field[i - 1][j].kind)
						curr_score = field[i - 1][j].score
					}
					if field[i + 1][j].score > curr_score {
						temp.kind = string(field[i + 1][j].kind)
						curr_score = field[i + 1][j].score
					}
					if field[i][j + 1].score > curr_score {
						temp.kind = string(field[i][j + 1].kind)
						curr_score = field[i][j + 1].score
					}
					if field[i + 1][j + 1].score > curr_score {
						temp.kind = string(field[i + 1][j + 1].kind)
						curr_score = field[i + 1][j + 1].score
					}
					if field[i - 1][j + 1].score > curr_score {
						temp.kind = string(field[i - 1][j + 1].kind)
						curr_score = field[i - 1][j + 1].score
					}
				}

				// right lower corner
				case i == (len(field) - 1) && j == (len(field[i]) - 1) : {
					if field[i][j - 1].score > curr_score {
						temp.kind = string(field[i][j - 1].kind)
						curr_score = field[i][j - 1].score
					}
					if field[i - 1][j].score > curr_score {
						temp.kind = string(field[i - 1][j].kind)
						curr_score = field[i - 1][j].score
					}
					if field[i - 1][j - 1].score > curr_score {
						temp.kind = string(field[i - 1][j - 1].kind)
						curr_score = field[i - 1][j - 1].score
					}
				}

				// right most line
				case (i == (len(field) - 1)) && (j != 0): {
					if field[i][j - 1].score > curr_score {
						temp.kind = string(field[i][j - 1].kind)
						curr_score = field[i][j - 1].score
					}
					if field[i][j + 1].score > curr_score {
						temp.kind = string(field[i][j + 1].kind)
						curr_score = field[i][j + 1].score
					}
					if field[i - 1][j].score > curr_score {
						temp.kind = string(field[i - 1][j].kind)
						curr_score = field[i - 1][j].score
					}
					if field[i - 1][j - 1].score > curr_score {
						temp.kind = string(field[i - 1][j - 1].kind)
						curr_score = field[i - 1][j - 1].score
					}
					if field[i - 1][j + 1].score > curr_score {
						temp.kind = string(field[i - 1][j + 1].kind)
						curr_score = field[i - 1][j + 1].score
					}
				}

				// lower most line
				case (j == (len(field[i]) - 1)) && (i != 0): {
					if field[i][j - 1].score > curr_score {
						temp.kind = string(field[i][j - 1].kind)
						curr_score = field[i][j - 1].score
					}
					if field[i + 1][j].score > curr_score {
						temp.kind = string(field[i + 1][j].kind)
						curr_score = field[i + 1][j].score
					}
					if field[i - 1][j].score > curr_score {
						temp.kind = string(field[i - 1][j].kind)
						curr_score = field[i - 1][j].score
					}
					if field[i - 1][j - 1].score > curr_score {
						temp.kind = string(field[i - 1][j - 1].kind)
						curr_score = field[i - 1][j - 1].score
					}
					if field[i + 1][j - 1].score > curr_score {
						temp.kind = string(field[i + 1][j - 1].kind)
						curr_score = field[i + 1][j - 1].score
					}
				}

				// right upper corner
				case (i == 0) && (j == (len(field[i]) - 1)) : {
					if field[i][j - 1].score > curr_score {
						temp.kind = string(field[i][j - 1].kind)
						curr_score = field[i][j - 1].score
					}
					if field[i + 1][j].score > curr_score {
						temp.kind = string(field[i + 1][j].kind)
						curr_score = field[i + 1][j].score
					}
					if field[i + 1][j - 1].score > curr_score {
						temp.kind = string(field[i + 1][j - 1].kind)
						curr_score = field[i + 1][j - 1].score
					}
				}

				// left lower corner
				case (i == len(field) - 1) && (j == 0) : {
					if field[i][j + 1].score > curr_score {
						temp.kind = string(field[i][j + 1].kind)
						curr_score = field[i][j + 1].score
					}
					if field[i - 1][j + 1].score > curr_score {
						temp.kind = string(field[i - 1][j + 1].kind)
						curr_score = field[i - 1][j + 1].score
					}
					if field[i - 1][j].score > curr_score {
						temp.kind = string(field[i - 1][j].kind)
						curr_score = field[i - 1][j].score
					}
				}

				default : {
					if field[i][j - 1].score > curr_score {
						temp.kind = string(field[i][j - 1].kind)
						curr_score = field[i][j - 1].score
					}
					if field[i + 1][j].score > curr_score {
						temp.kind = string(field[i + 1][j].kind)
						curr_score = field[i + 1][j].score
					}
					if field[i - 1][j].score > curr_score {
						temp.kind = string(field[i - 1][j].kind)
						curr_score = field[i - 1][j].score
					}
					if field[i - 1][j - 1].score > curr_score {
						temp.kind = string(field[i - 1][j - 1].kind)
						curr_score = field[i - 1][j - 1].score
					}
					if field[i + 1][j - 1].score > curr_score {
						temp.kind = string(field[i + 1][j - 1].kind)
						curr_score = field[i + 1][j - 1].score
					}
					if field[i - 1][j + 1].score > curr_score {
						temp.kind = string(field[i - 1][j + 1].kind)
						curr_score = field[i - 1][j + 1].score
					}
					if field[i][j + 1].score > curr_score {
						temp.kind = string(field[i][j + 1].kind)
						curr_score = field[i][j + 1].score
					}
					if field[i + 1][j + 1].score > curr_score {
						temp.kind = string(field[i + 1][j + 1].kind)
						curr_score = field[i + 1][j + 1].score
					}
				}
			}
			newField[i][j].score = 0
			newField[i][j].kind = temp.kind
		}
	}
	return newField
	//return nil // This is included only so this template will compile
}

// evolve takes an intial field and evolves it for nsteps according to the game
// rule. At each step, it should call "updateScores()" and the updateStrategies
func evolve(field [][]Cell, nsteps int, b float64) [][]Cell {
	for i := 0; i < nsteps; i++ {
		updateScores(field, b)
		field = updateStrategies(field)
	}
	return field
}

// Implements a Spatial Games version of prisoner's dilemma. The command-line
// usage is:
//     ./spatial field_file b nsteps
// where 'field_file' is the file continaing the initial arrangment of cells, b
// is the reward for defecting against a cooperator, and nsteps is the number
// of rounds to update stategies.
//
func main() {
	// parse the command line
	if len(os.Args) != 4 {
		fmt.Println("Error: should spatial field_file b nsteps")
		return
	}

	fieldFile := os.Args[1]

	b, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil || b <= 0 {
		fmt.Println("Error: bad b parameter.")
		return
	}

	nsteps, err := strconv.Atoi(os.Args[3])
	if err != nil || nsteps < 0 {
		fmt.Println("Error: bad number of steps.")
		return
	}

  // read the field
	field := readFieldFromFile(fieldFile)
    fmt.Println("Field dimensions are:", len(field), "by", len(field[0]))

  // evolve the field for nsteps and write it as a PNG
	field = evolve(field, nsteps, b)
	drawField(field, "Prisoners.png")
	drawFieldColor(field, "ColorPrisoners.png")
}
