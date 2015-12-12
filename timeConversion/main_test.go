package main

import (
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStaircase(t *testing.T) {
	Convey("Convert time for PM", t, func() {
		input := "07:05:45PM"
		output := convertTime(input)
		expected := "19:05:45"
		So(output, ShouldEqual, expected)
	})

	Convey("Convert time for AM", t, func() {
		input := "07:05:45AM"
		output := convertTime(input)
		expected := "07:05:45"
		So(output, ShouldEqual, expected)
	})
}

func TestGetArgs(t *testing.T) {
	input := "07:05:45PM"
	Convey("Get args from STDIN", t, func() {
		args := getArgs(strings.NewReader(input))
		So(args, ShouldEqual, "07:05:45PM")
	})
}
