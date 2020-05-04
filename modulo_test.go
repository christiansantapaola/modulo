package modulo

import (
	"fmt"
	"testing"
)

func TestPowerMod(t *testing.T) {
	c, _ := PowerMod(2, 29, 247)
	t.Error("debug")
	fmt.Println("c: ", c)
	m, _ := PowerMod(c, 149, 247)
	fmt.Println("m: ", m)
	i, _ := PowerMod(29, -2, 216)
	fmt.Println(i)
	j, _ := PowerMod(149, 2, 216)
	fmt.Println(j)
}

func TestInverseMod(t *testing.T) {
	i, err := InverseMod(29, 216)
	if err != nil {
		t.Errorf("err != nil")
	}
	if i != 149 {
		t.Errorf("error")
	}
	t.Error("debug")
	fmt.Println(i)
}
