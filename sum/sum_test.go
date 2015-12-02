package main

import (
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSumArray(t *testing.T) {
	Convey("Given args", t, func() {
		sumArgs := []int{}

		Convey("Should equal 0", func() {
			result := SumArray(sumArgs)
			So(result, ShouldEqual, 0)
		})

		Convey("Should equal 10", func() {
			sumArgs = []int{0, 1, 2, 3, 4}
			result := SumArray(sumArgs)
			So(result, ShouldEqual, 10)
		})
	})
}

func TestGetArgs(t *testing.T) {
	Convey("Given STDIN", t, func() {
		input := `5
		0 1 2 3 4
		`
		expectedNumArgs := 5
		expectedSumArgs := []int{0, 1, 2, 3, 4}
		r := strings.NewReader(input)
		numArgs, sumArgs := GetArgs(r)
		Convey("Get the number of arguments", func() {
			So(numArgs, ShouldEqual, expectedNumArgs)
		})

		Convey("Get the summable arguments", func() {
			So(sumArgs, ShouldResemble, expectedSumArgs)
		})
	})
}
