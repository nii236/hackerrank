package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	input := `5
	1 2 3 4 5`
	r := strings.NewReader(input)
	fmt.Println(GetArgs(r))
}

// GetArgs reads from STDIN and returns stuff in the correct format
func GetArgs(r io.Reader) (int, []int) {
	var numArgs = 0
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

	return numArgs, sumArgs
}
