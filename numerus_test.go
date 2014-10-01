// Copyright 2014, David Howden
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package numerus

import (
	"testing"
)

// TestDescNumerals is a sanity check for the internal array descNumerals
// which must be strictly decreasing and contain non-zero values
func TestDescNumerals(t *testing.T) {
	last := uint(0)
	for i, v := range descNumerals {
		if v.Value() == 0 {
			t.Errorf("invalid (zero) value in descNumerals: %v", v)
		}
		if i != 0 {
			if v.Value() >= last {
				t.Errorf("incorrect ordering in descNumerals: %v", v)
			}
		}
		last = v.Value()
	}
}

type testPair struct {
	v uint
	s string
}

var testPairs = []testPair{
	testPair{1, "I"},
	testPair{2, "II"},
	testPair{3, "III"},
	testPair{4, "IV"},
	testPair{5, "V"},
	testPair{6, "VI"},
	testPair{7, "VII"},
	testPair{8, "VIII"},
	testPair{9, "IX"},
	testPair{10, "X"},
	testPair{19, "XIX"},
	testPair{21, "XXI"},
	testPair{1944, "MCMXLIV"},
	testPair{1963, "MCMLXIII"},
	testPair{1900, "MCM"},
	testPair{2000, "MM"},
	testPair{3999, "MMMCMXCIX"},
}

func testStringMethod(n uint, s string, t *testing.T) {
	v := Numeral(n).String()
	if v != s {
		t.Errorf("input: %v, expected %v, got: %v", n, s, v)
	}
}

func TestStringMethod(t *testing.T) {
	for _, tm := range testPairs {
		testStringMethod(tm.v, tm.s, t)
	}
}

func testParseMethod(s string, u uint, t *testing.T) {
	x, err := parse(s)
	if err != nil {
		t.Errorf("input: %v, error: %v", s, err)
	}
	if x != u {
		t.Errorf("input: %v, expected: %v, got: %v", s, u, x)
	}
}

func TestParseMethod(t *testing.T) {
	for _, tm := range testPairs {
		testParseMethod(tm.s, tm.v, t)
	}
}

var invalidInputs = []string{
	"IIII",
	"VV",
	"XXXX",
	"LL",
	"CCCC",
	"DD",
	"MMMM",
	"VIV",
	"LXL",
	"DCD",
	"MCMD",
	"IM",
	"VC",
}

func TestInvalidInputs(t *testing.T) {
	for _, in := range invalidInputs {
		if _, err := Parse(in); err == nil {
			t.Errorf("expected error on invalid input: %v", in)
		}
	}
}

func TestConsistency(t *testing.T) {
	i := uint(0)
	max := Limit.Value()
	for ; i < max; i++ {
		n := Numeral(i)
		x, err := parse(n.String())
		if err != nil {
			t.Errorf("input: %v, expected: %v, error: %v", n.String(), i, err)
		}
		if x != i {
			t.Errorf("input: %v, expected: %v, error: %v", n.String(), i, x)
		}
	}
}
