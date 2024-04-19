package internal

import "testing"

func TestDeductINSS(t *testing.T) {
	// Valores relevantes para o desafio
	//
	// - ate 1.045,00: 7,5%
	// - de 1.045,01 ate 2.089,60: 9%
	// - de 2.089,61 ate 3.134,40: 12%
	// - de 3.134,41 ate 6.101,06: 14%
	//
	// O cálculo eh feito de forma progressiva e foi tomada a
	// liberdade de arredondar para baixo os valores decimais.
	t.Run("Employee with 1,000.00 salary", func(t *testing.T) {
		e := &Employee{
			GrossWage: 100000,
		}

		got := e.deductINSS()
		expected := 7500

		if got != expected {
			t.Fatalf("got=%d expected=%d", got, expected)
		}
	})

	t.Run("Employee with 2,000.00 salary", func(t *testing.T) {
		e := &Employee{
			GrossWage: 200000,
		}

		got := e.deductINSS()
		expected := 7837 + 8594

		if got != expected {
			t.Fatalf("got=%d expected=%d", got, expected)
		}
	})

	t.Run("Employee with 3,000.00 salary", func(t *testing.T) {
		e := &Employee{
			GrossWage: 300000,
		}

		got := e.deductINSS()
		expected := 7837 + 9401 + 10924

		if got != expected {
			t.Fatalf("got=%d expected=%d", got, expected)
		}
	})

	t.Run("Employee with 4,000.00 salary", func(t *testing.T) {
		e := &Employee{
			GrossWage: 400000,
		}

		got := e.deductINSS()
		expected := 7837 + 9401 + 12537 + 12118

		if got != expected {
			t.Fatalf("got=%d expected=%d", got, expected)
		}
	})

	t.Run("Employee with 5,000.00 salary", func(t *testing.T) {
		e := &Employee{
			GrossWage: 500000,
		}

		got := e.deductINSS()
		expected := 7837 + 9401 + 12537 + 26118

		if got != expected {
			t.Fatalf("got=%d expected=%d", got, expected)
		}
	})

	t.Run("Employee with 6,000.00 salary", func(t *testing.T) {
		e := &Employee{
			GrossWage: 600000,
		}

		got := e.deductINSS()
		expected := 7837 + 9401 + 12537 + 40118

		if got != expected {
			t.Fatalf("got=%d expected=%d", got, expected)
		}
	})

	t.Run("Employee with 7,000.00 salary", func(t *testing.T) {
		e := &Employee{
			GrossWage: 700000,
		}

		got := e.deductINSS()
		expected := 7837 + 9401 + 12537 + 41533

		if got != expected {
			t.Fatalf("got=%d expected=%d", got, expected)
		}
	})
}
