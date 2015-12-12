package main

import (
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCheckPalindrome(t *testing.T) {
	Convey("Shouldn't be a palindrome", t, func() {
		input := "bcbc"
		output := checkPalindrome(input)
		expected := false
		So(output, ShouldResemble, expected)
	})
	Convey("Should be a palindrome", t, func() {
		input := "abba"
		output := checkPalindrome(input)
		expected := true
		So(output, ShouldResemble, expected)
	})
}

func TestRunIndexer(t *testing.T) {
	Convey("Check if strings can become a palindrome", t, func() {
		input1 := "aaab"
		output1 := runIndexer(input1)
		expected1 := 3

		input2 := "baa"
		output2 := runIndexer(input2)
		expected2 := 0

		input3 := "aaa"
		output3 := runIndexer(input3)
		expected3 := -1

		input4 := "aaabbs"
		output4 := runIndexer(input4)
		expected4 := -2

		input5 := "hgygsvlfcwnswtuhmyaljkqlqjjqlqkjlaymhutwsnwcwflvsgygh"
		output5 := runIndexer(input5)
		expected5 := 44
		So(output1, ShouldEqual, expected1)
		So(output2, ShouldEqual, expected2)
		So(output3, ShouldEqual, expected3)
		So(output4, ShouldEqual, expected4)
		So(output5, ShouldEqual, expected5)
	})
}

func TestGetArgs(t *testing.T) {
	input := `3
  aaab
  baa
  aaa
  `
	Convey("Get args from STDIN", t, func() {
		args := getArgs(strings.NewReader(input))
		So(args, ShouldResemble, []string{"aaab", "baa", "aaa"})
	})
}
