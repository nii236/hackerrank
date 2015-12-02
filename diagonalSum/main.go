package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	run(os.Stdin)
}

func run(r io.Reader) {
	d, m := GetArgs(os.Stdin)
	fmt.Println(DiagonalSum(d, m))
}

// GetArgs reads from STDIN and returns stuff in the correct format
func GetArgs(r io.Reader) (int, [][]int) {
	d := 0

	fmt.Fscanln(r, &d)

	m := [][]int{}
	var val int

	for i := 0; i < d; i++ {
		row := []int{}
		for j := 0; j < d; j++ {
			fmt.Fscan(r, &val)
			row = append(row, val)
		}
		m = append(m, row)
	}

	return d, m
}

// DiagonalSum sums diagonal stuff
func DiagonalSum(d int, m [][]int) int {
	a := primaryDiagonal(d, m)
	b := secondaryDiagonal(d, m)
	return intAbs(a - b)
}

func primaryDiagonal(d int, m [][]int) int {
	res := 0
	for i := 0; i < d; i++ {
		res += m[i][i]
	}
	return res
}
func secondaryDiagonal(d int, m [][]int) int {
	res := 0
	for i := 0; i < d; i++ {
		res += m[d-i-1][i]
	}
	return res
}

func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
