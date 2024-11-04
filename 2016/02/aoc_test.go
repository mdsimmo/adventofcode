package main

import (
	"testing"
)

func TestKeymap(t *testing.T) {
	input := map[complex128]int {
		complex(0, 0): 5, 
		complex(1, 0): 6,
		complex(-1, 0): 4, 
		complex(-1, 1): 1, 
		complex(1, -1): 9, 
	}
	fail := false
	for in, expect := range(input) {
		output := Keymap(in)
		if output != expect {
			t.Logf("In: %+v, Expect: %d, Got: %d", in, expect, output)
			fail = true
		}
	}
	if fail {
		t.FailNow()
	}

}

func TestLine(t *testing.T) {
	input := []struct {
		start complex128
		line string
		expect complex128
	} {
		{
			start: complex(0, 0),
			line: "",
			expect: complex(0, 0),
		},
		{
			start: complex(0, 0),
			line: "UR",
			expect: complex(1, 1),
		},
		{
			start: complex(-1, -1),
			line: "UURRDDLL",
			expect: complex(-1, -1),
		},
		{
			start: complex(-1, -1),
			line: "LLUUUUU",
			expect: complex(-1, 1),
		},
	}
	fail := false
	for _, in := range(input) {
		output := LinePosition(in.line, in.start, AddPosition1)
		if output != in.expect {
			t.Logf("In: %+v, Got: %+v", in, output)
			fail = true
		}
	}
	if fail {
		t.FailNow()
	}

}

func TestGiven(t *testing.T) {
	result := process("in2.txt", AddPosition1, Keymap)
	expect := 0x1985
	if result != expect {
		t.Fatalf("Failed example. Expect: %x, Result %x", expect, result)
	}

}

func TestGiven2(t *testing.T) {
	result := process("in2.txt", AddPosition2, Keymap2)
	expect := 0x5db3
	if result != expect {
		t.Fatalf("Failed example. Expect: %x, Result %x", expect, result)
	}

}
