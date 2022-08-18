package main

import (
	"fmt"
	"testing"
)

type CaseTest struct {
	data     string
	expected string
}

var testCases = []CaseTest{
	{
		"q3e4t1p1p2b4",
		"qqqeeeetpppbbbb",
	},
	{
		"cvbn",
		"cvbn",
	},
	{
		`q\\5we\\5`,
		`q\\\\\we\\\\\`,
	},
	{
		`a\4b\5`,
		`a4b5`,
	},
	{
		`bn\55`,
		`bn55555`,
	},
	{
		"",
		"",
	},
	{
		"    ",
		"    ",
	},
	{
		"-5",
		"-----",
	},
	{
		"-5",
		"-----",
	},
	{
		"\\45",
		"-----",
	},
}

var negativeTestCases = []CaseTest{
	{
		"554",
		"incorrect string, first element is number",
	},
	{
		`4\\5`,
		"incorrect string, first element is number",
	},
	{
		`\\\\4\\5`,
		"incorrect string, first element is number",
	},
}

func TestUnpack(t *testing.T) {
	t.Run("Positive Tests Result:", func(t *testing.T) {
		for idx, cs := range testCases {
			res, _ := Unpack(cs.data)
			if *res != cs.expected {
				t.Errorf("error on #%d test: unexpected result %v (%v)", idx, cs.expected, res)
			}
		}
	})

	t.Run("Negative Tests Result:", func(t *testing.T) {
		for idx, cs := range negativeTestCases {
			res, err := Unpack(cs.data)
			if fmt.Sprint(err) != cs.expected {
				t.Errorf("error on #%d test: unexpected result %v (%v)", idx, cs.expected, res)
			}
		}
	})
}
