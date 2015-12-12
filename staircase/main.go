package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	run(os.Stdin)

}

func run(r io.Reader) {
	input := getArgs(r)
	result := staircase(input)

	for _, v := range result {
		fmt.Println(v)
	}

}

func getArgs(r io.Reader) int {
	numArgs := 0
	fmt.Fscanln(r, &numArgs)
	return numArgs
}

func staircase(num int) []string {
	var result []string

	for i := 0; i < num; i++ {
		spaces := strings.Repeat(" ", num-i-1)
		hashes := strings.Repeat("#", i+1)
		row := strings.Join([]string{spaces, hashes}, "")
		result = append(result, row)
	}

	return result
}
