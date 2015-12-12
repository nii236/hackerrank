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
	for _, v := range input {
		fmt.Println(runIndexer(v))
	}
}

func getArgs(r io.Reader) []string {
	numArgs := 0
	_, err := fmt.Fscanln(r, &numArgs)

	timeArgs := []string{}
	for {
		timeArg := ""
		_, err = fmt.Fscanln(r, &timeArg)
		if err == io.EOF {
			break
		}
		timeArgs = append(timeArgs, timeArg)
	}
	return timeArgs
}

func runIndexer(input string) int {
	if checkPalindrome(input) {
		return -1
	}
	for i := 0; i < len(input); i++ {
		newString := strings.Join([]string{input[:i], input[i+1:]}, "")
		if checkPalindrome(newString) {
			return i
		}
	}
	return -2
}

func checkPalindrome(input string) bool {
	for i := 0; i < len(input); i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}
