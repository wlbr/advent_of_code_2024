package main

import (
	"testing"
)

type testdata struct {
	fname         string
	expectedtask1 int64
	expectedtask2 int64
}

var testset []*testdata = []*testdata{{"example.txt", 3749, 11387}}

func TestTaskOne(t *testing.T) {

	for _, test := range testset {
		input := readdata(test.fname)
		r := task1(input)
		if r != test.expectedtask1 {
			t.Fatalf("Test '%s' failed. Got '%d' -  Wanted: '%d'", test.fname, r, test.expectedtask1)
		}
	}
}

func TestTaskTwo(t *testing.T) {
	for _, test := range testset {
		input := readdata(test.fname)
		r := task2(input)
		if r != test.expectedtask2 {
			t.Fatalf("Test '%s' failed. Got '%d' -  Wanted: '%d'", test.fname, r, test.expectedtask2)
		}
	}
}
