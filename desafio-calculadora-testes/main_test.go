package main

import "testing"

func TestSum1(t *testing.T) {
	test := sum(2, 2)
	expected := 4
	if test != expected {
		t.Error("Valor esperado é: ", expected, " Valor retornado: ", test)
	}
}

func TestSum2(t *testing.T) {
	test := sum(2, 2, 2)
	expected := 6
	if test != expected {
		t.Error("Valor esperado é: ", expected, " Valor retornado: ", test)
	}
}

func TestSub1(t *testing.T) {
	test := subtract(6, 5)
	expected := 1

	if test != expected {
		t.Error("Valor esperado:", expected, "Valor retornado:", test)
	}
}

func TestSub2(t *testing.T) {
	test := subtract(6, 5, 1)
	expected := 0

	if test != expected {
		t.Error("Valor esperado:", expected, "Valor retornado:", test)
	}
}

func TestMult1(t *testing.T) {
	test := multiply(10, 10)
	expected := 100
	if test != expected {
		t.Error("Valor esperado:", expected, "Valor retornado:", test)
	}
}

func TestMult2(t *testing.T) {
	test := multiply(10, 10, 2)
	expected := 200
	if test != expected {
		t.Error("Valor esperado:", expected, "Valor retornado:", test)
	}
}

func TestDiv1(t *testing.T) {
	test := divide(12, 2)
	expected := 6
	if test != expected {
		t.Error("Valor esperado:", expected, "Valor retornado:", test)
	}
}
func TestDiv2(t *testing.T) {
	test := divide(12, 2, 2)
	expected := 3
	if test != expected {
		t.Error("Valor esperado:", expected, "Valor retornado:", test)
	}
}
