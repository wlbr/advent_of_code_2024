package main

import (
	"testing"
)

type testdata struct {
	fname         string
	expectedtask1 int
	expectedtask2 int
}

var testset []*testdata = []*testdata{{"example.txt", 143, 123}}

func TestTaskOne(t *testing.T) {

	for _, test := range testset {
		rules, updates := readdata(test.fname)
		r := task1(rules, updates)
		if r != test.expectedtask1 {
			t.Fatalf("Test '%s' failed. Got '%d' -  Wanted: '%d'", test.fname, r, test.expectedtask1)
		}
	}
}

func TestTaskTwo(t *testing.T) {
	for _, test := range testset {
		rules, updates := readdata(test.fname)
		r := task2(rules, updates)
		if r != test.expectedtask2 {
			t.Fatalf("Test '%s' failed. Got '%d' -  Wanted: '%d'", test.fname, r, test.expectedtask2)
		}
	}
}
