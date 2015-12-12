package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	run(os.Stdin)

}

func run(r io.Reader) {
	input := getArgs(r)
	fmt.Println(convertTime(input))
}

func getArgs(r io.Reader) string {
	timeArg := ""
	_, err := fmt.Fscanln(r, &timeArg)
	if err != nil {
		return "derp"
	}
	return timeArg
}

//A time in 12-hour clock format (i.e.: hh:mm:ssAM or hh:mm:ssPM), where 01≤hh≤12.
func convertTime(input string) string {
	h := input[:2]
	m := input[3:5]
	s := input[6:8]
	med := input[8:10]

	newH := ""
	newM := ""
	newS := ""

	hInt, err := strconv.Atoi(h)

	if err != nil {
		return "err"
	}

	if med == "AM" {
		newH = h
		if hInt == 12 {
			newH = "00"
		}
		newM = m
		newS = s

	} else if med == "PM" {
		newHInt, err := strconv.Atoi(h)
		if err != nil {
			return "err"
		}
		if hInt == 12 {
			newHInt = 0
		}
		newHInt = newHInt + 12
		newH = strconv.Itoa(newHInt)
		newM = m
		newS = s
	}

	return strings.Join([]string{newH, ":", newM, ":", newS}, "")
}
