// Copyright 2014, David Howden
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package numerus provides methods for parsing and creating Roman Numerals.  See
// http://en.wikipedia.org/wiki/Roman_numerals for more details.
package numerus

import (
	"errors"
	"fmt"
	"strings"
)

type numeral interface {
	Value() uint
	String() string
}

type sym struct {
	s string
	v uint
}

func (s sym) Value() uint {
	return s.v
}

func (s sym) String() string {
	return s.s
}

type comb struct {
	a, b sym // NB: a < b
}

func (c comb) Value() uint {
	return c.b.Value() - c.a.Value()
}

func (c comb) String() string {
	return c.a.String() + c.b.String()
}

var (
	_I = sym{"I", 1}
	_V = sym{"V", 5}
	_X = sym{"X", 10}
	_L = sym{"L", 50}
	_C = sym{"C", 100}
	_D = sym{"D", 500}
	_M = sym{"M", 1000}

	_IV = comb{_I, _V}
	_IX = comb{_I, _X}
	_XL = comb{_X, _L}
	_XC = comb{_X, _C}
	_CD = comb{_C, _D}
	_CM = comb{_C, _M}
)

var descNumerals = []numeral{_M, _CM, _D, _CD, _C, _XC, _L, _XL, _X, _IX, _V, _IV, _I}

// Limit is the upper bound of possible numerals allowed by this package
// (this limit is set by the rule which prohibits more than three consecutive Ms).
const Limit = Numeral(3999)

// Numeral represents a Roman Numeral value.
type Numeral uint

// String returns a string representing the underlying Numeral in standard
// Roman Numeral notation.
func (n Numeral) String() string {
	result := ""
	i := uint(n)

	for _, v := range descNumerals {
		for {
			if i < v.Value() {
				break
			}
			result += v.String()
			i -= v.Value()
		}
	}
	return result
}

// Value returns the underlying value of the numeral as a uint.
func (n Numeral) Value() uint {
	return uint(n)
}

// parse takes a string representation of a Roman Numeral (in standard form)
// and returns a uint and error value, which is set if the given input is not
// in the standard representation.
func parse(s string) (uint, error) {
	// As overflowing doesn't catch this, we test for it first
	if strings.Contains(s, "MMMM") {
		return 0, errors.New("invalid numeral near MMMM")
	}

	// Check the running totals so that we don't accept invalid input
	// i.e. MCMD should be MMCD
	check := make([]uint, len(descNumerals))

	orig := s
	n := uint(0)
	for i, v := range descNumerals {
		vs := v.String()
		for {
			if !strings.HasPrefix(s, vs) {
				break
			}

			for j := 0; j < i; j++ {
				check[j] += v.Value()
				if check[j] >= descNumerals[j].Value() {
					return 0, fmt.Errorf("invalid numeral %#v after %#v", s, orig[:len(orig)-len(s)])
				}
			}
			n += v.Value()
			s = s[len(vs):]
		}
	}

	if len(s) > 0 {
		return 0, fmt.Errorf("invalid numeral %#v after %#v", s, orig[:len(orig)-len(s)])
	}
	return n, nil
}

// Parse takes a string in standard Roman Numeral notation and returns a Numeral.
// If the given representation is invalid an error is returned.
func Parse(s string) (Numeral, error) {
	n, err := parse(s)
	if err != nil {
		return Numeral(0), err
	}
	return Numeral(n), err
}
