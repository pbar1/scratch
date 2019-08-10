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

func make2D(x, y int) [][]uint {
	m := make([][]uint, y)
	for i := range m {
		m[i] = make([]uint, x)
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

/******************************************************************************
Notes on traversing an Ulam spiral
	Moving right: (x+1, y)
	Moving left:  (x-1, y)
	Moving up:    (x, y-1) <- because of top-down indexing
	Moving down:  (x, y+1) <- "


******************************************************************************/

// ulam creates the Ulam Spiral matrix of unsigned integers up to a given number
func ulam(n uint) [][]uint {
	side := int(math.Ceil(math.Sqrt(float64(n))))
	part := make2D(side, side)

	originX := int(math.Ceil(float64(side)/2 - 1))
	originY := int(math.Floor(float64(side) / 2))

	i := 0
	for x := 0; x < side; x++ {
		for y := 0; y < side; y++ {
			part[x][y] = uint(i)
			// TODO
		}
	}
	part[originX][originY] = 1 // TODO: marking origin, remove

	return part
}

// stringMatrix builds the string form of a given matrix
func stringMatrix(matrix [][]uint) string {
	var ys []string
	for y := range matrix {
		var xs []string
		for x := range matrix[y] {
			xs = append(xs, fmt.Sprintf("%d", matrix[x][y]))
		}
		ys = append(ys, strings.Join(xs, "\t"))
	}
	return strings.Join(ys, "\n")
}
