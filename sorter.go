package numerus

import (
	"strings"
)

// Numerals defines a new type as an array of Numeral structures.
// The type is defined because it's going to implement the Sort
// interface. This interface requires three methods:
//
//  1. Len (the number of array elements)
//  2. Less (true if the first element is less than the second)
//  3. Swap (exchanges two element in the array)
//
// To invoke the sort on a Numerals object `numerals`, call
// one of these two methods, depending the desired order:
//
//	sort.Sort(numerals) // Ascending order
//	sort.Sort(sort.Reverse(numerals)) // Descending order
//
// This will sort the array in place.
type Numerals []Numeral

// Len returns the number of Numeral structures in the array.
func (ns Numerals) Len() int {
	return len(ns)
}

// Less compares two Numeral structures and returns true if the first
// one is less than the second one.
//
// In the case of Numeral, this is easy because the underlying data
// structure is simply an unsigned integer.
func (ns Numerals) Less(i int, j int) bool {
	return ns[i] < ns[j]
}

// Exchanges two Numeral objects in the array.
func (ns Numerals) Swap(i int, j int) {
	ns[i], ns[j] = ns[j], ns[i]
}

// String returns a string representation of the array
func (ns Numerals) String() string {
	numeralStrings := make([]string, 0, len(ns))
	for _, numeral := range ns {
		numeralStrings = append(numeralStrings, numeral.String())
	}
	sb := "["
	sb += strings.Join(numeralStrings, " ")
	sb += "]"
	return sb
}
