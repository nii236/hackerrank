package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBigSum(t *testing.T) {
	Convey("Given args", t, func() {
		sumArgs := []int{}

		Convey("Should equal zero", func() {
			result := bigsum(sumArgs)
			So(result, ShouldEqual, 0)
		})
	})

	Convey("Given huge args", t, func() {
		sumArgs := []int{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}

		Convey("Should equal 5000000015", func() {
			result := bigsum(sumArgs)
			So(result, ShouldEqual, 5000000015)
		})
	})
}
