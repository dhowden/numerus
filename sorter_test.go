package numerus

import (
	"fmt"
	"sort"
	"testing"
)

func TestSortAscending(t *testing.T) {
	numerals := loadRandomList()

	// Sort the list in ascending order
	sort.Sort(numerals)

	// Verify that the list is sorted in ascending order
	verifyList(numerals, ascending, t)

	// Print the list if desired by setting showList to true.
	showList := false
	if showList {
		printList(numerals)
	}
}

func TestSortDescending(t *testing.T) {
	numerals := loadRandomList()

	// Sort the list in descending order
	sort.Sort(sort.Reverse(numerals))

	// Verify that the list is sorted in descending order
	verifyList(numerals, descending, t)

	// Print the list if desired by setting showList to true.
	showList := false
	if showList {
		printList(numerals)
	}
}

type valueAndString struct {
	v uint
	s string
}

// You can generate a different list with this tool:
// https://catonmat.net/tools/generate-random-roman-numerals
var randomList = []valueAndString{
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

func (vands valueAndString) String() string {
	return fmt.Sprintf("{%d, %q}", vands.v, vands.s)
}

type order int

const (
	ascending  = 1
	descending = -1
)

func loadRandomList() Numerals {
	var numerals Numerals
	for _, vands := range randomList {
		numerals = append(numerals, Numeral(vands.v))
	}
	return numerals
}

func verifyList(numerals Numerals, order order, t *testing.T) {
	for i := 1; i < len(numerals); i++ {

		// Get the two values to be compared
		thisValue := numerals[i]
		prevValue := numerals[i-1]

		// Determine what condition is to be tested
		switch order {
		case ascending:
			if thisValue < prevValue {
				t.Errorf("list is out of order at index %d", i)
			}
		case descending:
			if thisValue > prevValue {
				t.Errorf("list is out of order at index %d", i)
			}
		}
	}
}

func printList(numerals Numerals) {
	for _, numeral := range numerals {
		v := numeral.Value()
		s := numeral.String()
		vands := valueAndString{v, s}
		fmt.Println(vands.String())
	}
}
