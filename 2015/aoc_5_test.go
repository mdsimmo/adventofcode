package main

import "testing"

func TestGood(t *testing.T) {
	input := "aabbaa dgd"
	result := test2(input)
	if !result {
		t.Fatalf("%s is good", input)
	}
}


func TestNoDouble(t *testing.T) {
	input := "aabba dgd"
	result := test2(input)
	if result {
		t.Fatalf("%s is bad", input)
	}
}


func TestNoRepeat(t *testing.T) {
	input := "aabbaa dggd"
	result := test2(input)
	if result {
		t.Fatalf("%s is bad", input)
	}
}

func TestNoDoubleContinued(t *testing.T) {
	input := "aaa dgd"
	result := test2(input)
	if result {
		t.Fatalf("%s is bad", input)
	}
}


func TestAllAtStart(t *testing.T) {
	input := "aaaa"
	result := test2(input)
	if !result {
		t.Fatalf("%s is good", input)
	}
}
