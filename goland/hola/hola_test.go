package main

import "testing"

func TestHolaMundo(t *testing.T) {

	t.Run("Caso 01.-", func(t *testing.T) {
		expected := "Hola, Andres"
		got := HolaMundo("Andres")
		isMensajeCorrecto(got, expected, t)
	})

	t.Run("Caso 02.-", func(t *testing.T) {
		expected := "Hola, Parrita"
		got := HolaMundo("")
		isMensajeCorrecto(got, expected, t)
	})
}	

func isMensajeCorrecto(got string, expected string, t *testing.T) {
	t.Helper()
	if expected != got {
		t.Errorf("got: '%s' expected: '%s'", got, expected)
	}
}
