package numerus

import (
	"fmt"
	"sort"
	"testing"
)

// You can generate a different list with this tool:
// https://catonmat.net/tools/generate-random-roman-numerals
var randomList = []struct {
	v uint
	s string
}{
	{1675, "MDCLXXV"},
	{860, "DCCCLX"},
	{166, "CLXVI"},
	{1639, "MDCXXXIX"},
	{849, "DCCCXLIX"},
	{139, "CXXXIX"},
	{520, "DXX"},
	{1639, "MDCXXXIX"},
	{1380, "MCCCLXXX"},
	{1602, "MDCII"},
}

func TestSort(t *testing.T) {
	var numerals Numerals
	for _, rPair := range randomList {
		numerals = append(numerals, Numeral(rPair.v))
	}
	sort.Sort(numerals)
	fmt.Println(numerals.String())
	for i, numeral := range numerals {
		fmt.Printf("%d: %s (%d)\n", i, numeral.String(), numeral.Value())
	}
}
