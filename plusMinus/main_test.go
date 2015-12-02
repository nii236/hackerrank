package main

import (
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPlusMinus(t *testing.T) {
	input := []int{-1, 1, 0, 1}
	Convey("Get fraction of negative numbers", t, func() {
		result := getNegatives(input)
		So(result, ShouldEqual, 0.25)
	})

	Convey("Get fraction of positive numbers", t, func() {
		result := getPositives(input)
		So(result, ShouldEqual, 0.5)
	})

	Convey("Get fraction of zeroes", t, func() {
		result := getZeroes(input)
		So(result, ShouldEqual, 0.25)
	})

	Convey("Get count of numbers", t, func() {
		result := count(input)
		So(result, ShouldEqual, 4)
	})

}

func TestGetArgs(t *testing.T) {
	input := `6
  -4 3 -9 0 4 1
  `
	Convey("Get args from STDIN", t, func() {
		args := getArgs(strings.NewReader(input))
		So(args, ShouldResemble, []int{-4, 3, -9, 0, 4, 1})
	})
}
