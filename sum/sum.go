package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	_, sumArgs := GetArgs(os.Stdin)
	SumArray(sumArgs)
}

func run(r io.Reader) {
	_, sumArgs := GetArgs(r)
	fmt.Println(SumArray(sumArgs))
}

// GetArgs reads from STDIN and returns stuff in the correct format
func GetArgs(r io.Reader) (int, []int) {
	var numArgs int
	_, err := fmt.Fscanln(r, &numArgs)
	if err != nil {
		fmt.Println(err)
	}

	sumArgs := make([]int, numArgs)
	for i := range sumArgs {
		_, err = fmt.Fscan(r, &sumArgs[i])
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println(sumArgs)

	return numArgs, sumArgs
}

// SumArray sums arrays
func SumArray(sumArgs []int) int {
	var s int
	for _, v := range sumArgs {
		s += v
	}
	return s
}
