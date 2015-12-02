package main

import (
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDiagonalSum(t *testing.T) {
	Convey("Difference between primary and secondary diagonal", t, func() {
		d := 3
		m := [][]int{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}}
		result := DiagonalSum(d, m)
		So(result, ShouldEqual, 15)
	})
}

func TestPrimaryDiagonal(t *testing.T) {
	Convey("Add primary diagonal", t, func() {
		d := 3
		m := [][]int{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}}
		result := primaryDiagonal(d, m)
		So(result, ShouldEqual, 4)
	})
}

func TestSecondaryDiagonal(t *testing.T) {
	Convey("Add secondary diagonal", t, func() {
		d := 3
		m := [][]int{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}}
		result := secondaryDiagonal(d, m)
		So(result, ShouldEqual, 19)
	})
}

func TestGetArgs(t *testing.T) {
	input :=
		`3
	11 2 4
	4 5 6
	10 8 -12
	`
	Convey("Get args from STDIN", t, func() {
		d, m := GetArgs(strings.NewReader(input))
		So(d, ShouldEqual, 3)
		So(m, ShouldResemble, [][]int{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}})
	})

}
