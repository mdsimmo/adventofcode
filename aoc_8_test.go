package main

import "testing"


func TestEmpty(t *testing.T) {
	input := "\"\""
	result := count(input)
	if result != 2 {
		t.Fatalf("%s should be 2. Was %d", input, result)
	}
}


func TestBackslah(t *testing.T) {
	input := "a\\\\a"
	result := count(input)
	if result != 3 {
		t.Fatalf("%s should be 3. Was %d", input, result)
	}
}


func TestBackslashX(t *testing.T) {
	input := "a\\x90a"
	result := count(input)
	if result != 3 {
		t.Fatalf("%s should be 3. Was %d", input, result)
	}
}
