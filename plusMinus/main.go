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
	input := getArgs(r)
	fmt.Println(getPositives(input))
	fmt.Println(getNegatives(input))
	fmt.Println(getZeroes(input))
}

func getArgs(r io.Reader) []int {
	numArgs := 0

	fmt.Fscanln(r, &numArgs)

	sumArgs := []int{}
	var val int
	for i := 0; i < numArgs; i++ {
		fmt.Fscan(r, &val)
		sumArgs = append(sumArgs, val)
	}

	return sumArgs
}

func getNegatives(input []int) float64 {
	num := 0
	for _, v := range input {
		if v < 0 {
			num++
		}
	}
	den := count(input)
	fnum := float64(num)
	fden := float64(den)
	return fnum / fden
}

func getPositives(input []int) float64 {
	num := 0
	for _, v := range input {
		if v > 0 {
			num++
		}
	}
	den := count(input)
	fnum := float64(num)
	fden := float64(den)
	return fnum / fden
}

func getZeroes(input []int) float64 {
	num := 0
	for _, v := range input {
		if v == 0 {
			num++
		}
	}
	den := count(input)
	fnum := float64(num)
	fden := float64(den)
	return fnum / fden
}

func count(input []int) int {
	count := 0
	for range input {
		count++
	}
	return count
}
