package main

import (
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStaircase(t *testing.T) {
	input := 6
	Convey("Print staircase", t, func() {
		output := staircase(input)
		expected := []string{"     #", "    ##", "   ###", "  ####", " #####", "######"}
		So(output, ShouldResemble, expected)
	})
}

func TestGetArgs(t *testing.T) {
	input := `6 `
	Convey("Get args from STDIN", t, func() {
		args := getArgs(strings.NewReader(input))
		So(args, ShouldEqual, 6)
	})
}
