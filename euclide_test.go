package modulo

import (
	"fmt"
	"testing"
)

func TestEuclid(t *testing.T) {
	eu := ExtendedEuclideanAlgorithm(2152151, 51251521)
	a, b, _ := eu.Exec()
	eu.Print()
	fmt.Printf("(%d) %d + (%d) %d = %d\n", a, eu.n1, b, eu.n2, 1)
}
