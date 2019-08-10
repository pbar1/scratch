package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	m := ulam(100)
	s := stringMatrix(m)
	fmt.Println(s)
}

func make2D(x, y int) [][]int {
	m := make([][]int, y)
	for i := range m {
		m[i] = make([]int, x)
	}
	return m
}

//func make3D(x, y, z int) [][][]uint {
//	m := make([][][]uint, z)
//	for i := range m {
//		m[i] = make2D(x, y)
//	}
//	return m
//}

// triangular finds the Nth triangular number by evaluating the expression n(n+1)/2
func triangular(n int) int {
	return n * (n + 1) >> 1
}

// pentagonal finds the Nth pentagonal number by evaluating the expression (3n^2-n)/2
func pentagonal(n int) int {
	return (3*n*n - n) >> 1
}

// ulam creates the partial Ulam Spiral of unsigned integers up to a given number
func ulam(n int) [][]int {
	side := int(math.Ceil(math.Sqrt(float64(n))))
	part := make2D(side, side)

	originX := int(math.Ceil(float64(side)/2 - 1))
	originY := int(math.Floor(float64(side) / 2))

	counter := 1
	posX := originX
	posY := originY
	facing := 0 // initially facing south

	for {
		part[posX][posY] = counter
		counter++
		if counter > n {
			break
		}

		_, peekLeftX, peekLeftY := moveLeft(facing, posX, posY)
		if peekLeftX < 0 || peekLeftX >= side || peekLeftY < 0 || peekLeftY >= side {
			panic(fmt.Sprintf("n: %d, side: %d, counter: %d, posX: %d, posY: %d, facing: %d", n, side, counter, posX, posY, facing))
		} else if part[peekLeftX][peekLeftY] == 0 {
			facing, posX, posY = moveLeft(facing, posX, posY)
		} else {
			_, posX, posY = moveForward(facing, posX, posY)
		}
	}

	return part
}

/******************************************************************************
Notes on traversing an Ulam spiral
	Moving right: (x+1, y)
	Moving left:  (x-1, y)
	Moving up:    (x, y-1) <- because of top-down indexing
	Moving down:  (x, y+1) <- "

Directions
	0: south
	1: east
	2: north
	3: west

Interesting, the first move algorithmically looks like a turn left...
******************************************************************************/
func moveLeft(facing, posX, posY int) (int, int, int) {
	switch facing {
	case 0:
		// while facing south, a left turn is a move right
		return 1, posX + 1, posY
	case 1:
		// while facing east, a left turn is a move up
		return 2, posX, posY - 1
	case 2:
		// while facing north, a left turn is a move left
		return 3, posX - 1, posY
	case 3:
		// while facing west, a left turn is a move down
		return 0, posX, posY + 1
	default:
		panic(fmt.Sprintf("facing: %d, posX: %d, posY: %d", facing, posX, posY))
	}
}

func moveForward(facing, posX, posY int) (int, int, int) {
	switch facing {
	case 0:
		// while facing south, a move forward is a move down
		return facing, posX, posY + 1
	case 1:
		// while facing east, a move forward is a move right
		return facing, posX + 1, posY
	case 2:
		// while facing north, a move forward is a move up
		return facing, posX, posY - 1
	case 3:
		// while facing west, a move forward is a move left
		return facing, posX - 1, posY
	default:
		panic(fmt.Sprintf("facing: %d, posX: %d, posY: %d", facing, posX, posY))
	}
}

// stringMatrix builds the string form of a given matrix
func stringMatrix(matrix [][]int) string {
	var ys []string
	for y := range matrix {
		var xs []string
		for x := range matrix[y] {
			var s string
			// removing zeros for ease of viewing
			if matrix[x][y] != 0 {
				s = fmt.Sprintf("%d", matrix[x][y])
			}
			xs = append(xs, s)
		}
		ys = append(ys, strings.Join(xs, "\t"))
	}
	return strings.Join(ys, "\n")
}
