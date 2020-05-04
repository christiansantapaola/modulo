package modulo

import (
	"fmt"
)

type euclidExtRow struct {
	a, b, val, rem int64
}

// Init will init a row with value
func (row *euclidExtRow) Init(a, b, val, rem int64) {
	row.a = a
	row.b = b
	row.val = val
	row.rem = rem
}

// EuclidExtTable is the data to exec the Euclide extendend Algorithm
type EuclidExtTable struct {
	n1, n2 int64
	tpos   uint64
	table  []euclidExtRow
}

// ExtendedEuclideanAlgorithm will create a new euclidExtTable struct
func ExtendedEuclideanAlgorithm(n1, n2 int64) EuclidExtTable {
	res := EuclidExtTable{
		n1:    n1,
		n2:    n2,
		table: make([]euclidExtRow, 2),
	}
	res.table[0].Init(1, 0, n1, 0)
	res.table[1].Init(0, 1, n2, n1/n2)
	res.tpos = 1
	return res
}

// Next will compute the next step of the Euclid Extended Algo
func (euclid *EuclidExtTable) Next() bool {
	row1 := euclid.table[euclid.tpos-1]
	row2 := euclid.table[euclid.tpos]
	rem := euclid.table[euclid.tpos].rem
	newrow := new(euclidExtRow)
	newval := row1.val - row2.val*rem
	var newrem int64 = 0
	if newval != 0 {
		newrem = row2.val / newval
		newrow.Init(
			row1.a-row2.a*rem,
			row1.b-row2.b*rem,
			newval,
			newrem,
		)
		euclid.table = append(euclid.table, *newrow)
		euclid.tpos++
		return true
	}
	return false
}

// Gcd calculate the gcd
func (euclid *EuclidExtTable) Gcd() int64 {
	a := euclid.n1
	b := euclid.n2
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Print print the data
func (euclid *EuclidExtTable) Print() {
	fmt.Println(fmt.Sprintf("%d a + %d b = 1", euclid.n1, euclid.n2))
	fmt.Println("| i | a | b | val | rem |")
	for i, row := range euclid.table {
		fmt.Println("|", i, "|", row.a, "|", row.b, "|", row.val, "|", row.rem, "|")
	}
}

// Exec exec the algo
func (euclid *EuclidExtTable) Exec() (int64, int64, error) {
	if euclid.Gcd() != 1 {
		return 0, 0, fmt.Errorf("gcd(%s, %s) != 1", string(euclid.n1), string(euclid.n2))
	}
	for euclid.Next() == true {

	}
	return euclid.table[euclid.tpos].a, euclid.table[euclid.tpos].b, nil
}
