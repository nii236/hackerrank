package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSimpleArraySum(t *testing.T) {
	Convey("Given args", t, func() {
		sumArgs := []int{}

		Convey("Should equal zero", func() {
			result := simpleArraySum(sumArgs)
			So(result, ShouldEqual, 0)
		})
	})

	Convey("Given normal args", t, func() {
		sumArgs := []int{1, 2, 3, 4}

		Convey("Should equal 5000000015", func() {
			result := simpleArraySum(sumArgs)
			So(result, ShouldEqual, 10)
		})
	})
}
