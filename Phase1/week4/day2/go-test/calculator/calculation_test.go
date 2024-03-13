package calculator

import "testing"

func TestPenjumlahan(t *testing.T) {
	a := 10
	b := 5

	hasil := Penjumlahan(a, b)

	if hasil != 15 {
		t.Errorf("Expected 15 but got %d", hasil)
	}
}

func TestPengurangan(t *testing.T) {
	a := 10
	b := 5

	hasil := Pengurangan(a, b)

	if hasil != 5 {
		t.Errorf("Expected 5 but got %d", hasil)
	}
}
