package main

import (
	"os"
	"fmt"
	"math/rand"
	"math"
	"strconv"
)

func randDegree() float64 {
	return (float64(rand.Int() % 361) / 360.0)
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
	  // test
		// fmt.Println(nx, ny, x, y, width, height)
		// fmt.Println(!inField(nx, width), !inField(ny, width))
	}

	return nx, ny
}

// width = 10 height = 100.1 stepsize = 0.2 steps = 1000
func randomWalk(width, height, stepsize float64, steps int) {
	var x, y = width / 2, height / 2
	fmt.Println(x, y)
	for i := 0; i < steps; i++ {
		x, y = randStep(x, y, width, height, stepsize)
		fmt.Println(x, y)
	}
	dist := math.Sqrt((x - width / 2) * (x - width / 2) + (y - height / 2) * (y - height / 2))
	fmt.Println("Distance =", dist)
}

/* Read the parameters
 *
 *
 */
func main() {
	if len(os.Args) != 6 {
		fmt.Println("Error: informative error message here")
		return
	}
	width, err_wid := strconv.ParseFloat(os.Args[1], 64)
	height, err_hei := strconv.ParseFloat(os.Args[2], 64)
	stepsize, err_ste := strconv.ParseFloat(os.Args[3], 64)
	number_of_steps, err_num:= strconv.Atoi(os.Args[4])
	seed, err_see := strconv.Atoi(os.Args[5])
	rand.Seed(int64(seed))
	if (err_wid != nil) || (err_hei != nil) ||
		(err_ste != nil) || (err_num != nil) || (err_see != nil) {
		fmt.Println("Error: informative error message here")
		return
	}
	if (width <= 0) || (height <= 0) ||
		(stepsize <= 0) || (number_of_steps <= 0) || (seed <= 0) {
		fmt.Println("Error: informative error message here")
		return	
	}
 	randomWalk(width, height, stepsize, number_of_steps)
	return
}
