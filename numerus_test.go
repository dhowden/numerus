// Copyright 2014, David Howden
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package numerus

import "testing"

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

var testPairs = []struct {
	v uint
	s string
}{
	{0, ""},
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{6, "VI"},
	{7, "VII"},
	{8, "VIII"},
	{9, "IX"},
	{10, "X"},
	{19, "XIX"},
	{21, "XXI"},
	{1944, "MCMXLIV"},
	{1963, "MCMLXIII"},
	{1900, "MCM"},
	{2000, "MM"},
	{3999, "MMMCMXCIX"},
}

func TestStringMethod(t *testing.T) {
	for _, tm := range testPairs {
		got := Numeral(tm.v).String()
		if got != tm.s {
			t.Errorf("Numeral(%d).String() = %v, expected %v", tm.v, got, tm.s)
		}
	}
}

func TestParseMethod(t *testing.T) {
	for _, tm := range testPairs {
		t.Run(tm.s, func(t *testing.T) {
			got, err := Parse(tm.s)
			if err != nil {
				t.Fatalf("parse(%v) gave error error: %v", tm.s, err)
			}
			if got.Value() != tm.v {
				t.Errorf("parse(%v) = %d, expected: %d", tm.s, got, tm.v)
			}
		})
	}
}

func BenchmarkParseMethod(b *testing.B) {
	for _, tm := range testPairs {
		b.Run(tm.s, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Parse(tm.s)
			}
		})
	}
}

func BenchmarkStringMethod(b *testing.B) {
	for _, tm := range testPairs {
		b.Run(tm.s, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Numeral(tm.v).String()
			}
		})
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
			t.Errorf("Parse(%v) should have returned an error", in)
		}
	}
}

func TestConsistency(t *testing.T) {
	i := uint(0)
	max := Limit.Value()
	for ; i < max; i++ {
		s := Numeral(i).String()
		x, err := parse(s)
		if err != nil {
			t.Errorf("Numeral(%d).String() = %v, parse(%v) gave error: %v", i, s, s, err)
		}
		if x != i {
			t.Errorf("Numeral(%d).String() = %v, parse(%v) = %d, expected: %d", i, s, s, x, i)
		}
	}
}
