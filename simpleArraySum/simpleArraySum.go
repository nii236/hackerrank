package main

import "fmt"

func main() {
	run()
}

func run() {
	numArgs := 0
	fmt.Scanln(&numArgs)

	sumArgs := make([]int, numArgs)
	for i := range sumArgs {
		fmt.Scan(&sumArgs[i])
	}

	result := simpleArraySum(sumArgs)
	fmt.Println(result)
}

func simpleArraySum(sumArgs []int) int {
	sum := 0
	for _, v := range sumArgs {
		sum += v
	}
	return sum
}
