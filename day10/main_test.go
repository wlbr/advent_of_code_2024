package main

import (
	"testing"
)

type testdata struct {
	fname         string
	expectedtask1 int
	expectedtask2 int
}

var testset []*testdata = []*testdata{
	{"example1.txt", 2, 2},
	{"example2.txt", 1, 16},
	{"example3.txt", 4, 13},
	{"example4.txt", 3, 3},
	{"example5.txt", 36, 81},
}

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
