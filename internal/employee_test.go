package internal

import "testing"

func TestDeductionValueINSS(t *testing.T) {
	t.Parallel()

	// Valores relevantes para o desafio
	//
	// - Ate 1.045,00: 7,5%
	// - De 1.045,01 ate 2.089,60: 9%
	// - De 2.089,61 ate 3.134,40: 12%
	// - De 3.134,41 ate 6.101,06: 14%
	//
	// O cálculo eh feito de forma progressiva e foi tomada a
	// liberdade de arredondar para baixo os valores decimais.
	t.Run("Employee with 1,000.00 salary", func(t *testing.T) {
		e := &Employee{
			GrossWage: 100000,
		}

		got := e.deductionValueINSS()
		expected := 7500

		if got != expected {
			t.Fatalf("got=%d expected=%d", got, expected)
		}
	})

	t.Run("Employee with 2,000.00 salary", func(t *testing.T) {
		e := &Employee{
			GrossWage: 200000,
		}

		got := e.deductionValueINSS()
		expected := 16432

		if got != expected {
			t.Fatalf("got=%d expected=%d", got, expected)
		}
	})

	t.Run("Employee with 3,000.00 salary", func(t *testing.T) {
		e := &Employee{
			GrossWage: 300000,
		}

		got := e.deductionValueINSS()
		expected := 28162

		if got != expected {
			t.Fatalf("got=%d expected=%d", got, expected)
		}
	})

	t.Run("Employee with 4,000.00 salary", func(t *testing.T) {
		e := &Employee{
			GrossWage: 400000,
		}

		got := e.deductionValueINSS()
		expected := 41893

		if got != expected {
			t.Fatalf("got=%d expected=%d", got, expected)
		}
	})

	t.Run("Employee with 5,000.00 salary", func(t *testing.T) {
		e := &Employee{
			GrossWage: 500000,
		}

		got := e.deductionValueINSS()
		expected := 55893

		if got != expected {
			t.Fatalf("got=%d expected=%d", got, expected)
		}
	})

	t.Run("Employee with 6,000.00 salary", func(t *testing.T) {
		e := &Employee{
			GrossWage: 600000,
		}

		got := e.deductionValueINSS()
		expected := 69893

		if got != expected {
			t.Fatalf("got=%d expected=%d", got, expected)
		}
	})

	t.Run("Employee with 7,000.00 salary", func(t *testing.T) {
		e := &Employee{
			GrossWage: 700000,
		}

		got := e.deductionValueINSS()
		expected := 71308

		if got != expected {
			t.Fatalf("got=%d expected=%d", got, expected)
		}
	})
}

func TestDeductionValueIRPF(t *testing.T) {
	t.Parallel()

	// Valores relevantes para o desafio
	//
	// - Ate 1.903,98:
	// - De 1.903,90 até 2.826,65: 7,5% (142,80)
	// - De 2.826,66 até 3.751,05: 15% (354,80)
	// - De 3.751,06 até 4.664,68: 22,5% (636,13)
	// - Acima de 4.664,68: 27,5% (869,36)
	//
	// O cálculo eh feito de forma progressiva e foi tomada a
	// liberdade de arredondar para baixo os valores decimais.
	t.Run("Employee with 1,000.00 salary", func(t *testing.T) {
		e := &Employee{
			GrossWage: 100000,
		}

		got := e.deductionValueIRPF()
		expected := 0

		if got != expected {
			t.Fatalf("got=%d expected=%d", got, expected)
		}
	})

	t.Run("Employee with 2,000.00 salary", func(t *testing.T) {
		e := &Employee{
			GrossWage: 200000,
		}

		got := e.deductionValueIRPF()
		expected := 720

		if got != expected {
			t.Fatalf("got=%d expected=%d", got, expected)
		}
	})

	t.Run("Employee with 3,000.00 salary", func(t *testing.T) {
		e := &Employee{
			GrossWage: 300000,
		}

		got := e.deductionValueIRPF()
		expected := 9520

		if got != expected {
			t.Fatalf("got=%d expected=%d", got, expected)
		}
	})

	t.Run("Employee with 4,000.00 salary", func(t *testing.T) {
		e := &Employee{
			GrossWage: 400000,
		}

		got := e.deductionValueIRPF()
		expected := 26387

		if got != expected {
			t.Fatalf("got=%d expected=%d", got, expected)
		}
	})

	t.Run("Employee with 5,000.00 salary", func(t *testing.T) {
		e := &Employee{
			GrossWage: 500000,
		}

		got := e.deductionValueIRPF()
		expected := 50564

		if got != expected {
			t.Fatalf("got=%d expected=%d", got, expected)
		}
	})
}

func TestDeductionValueHealthcare(t *testing.T) {
	t.Parallel()

	t.Run("Employee with 1,000.00 salary and with active healthcare", func(t *testing.T) {
		e := &Employee{
			GrossWage: 100000,
			HasHealthcare: true,
		}

		got := e.deductionValueHealthcare()
		expected := 1000

		if got != expected {
			t.Fatalf("got=%d expected=%d", got, expected)
		}
	})

	t.Run("Employee with 1,000.00 salary and without active healthcare", func(t *testing.T) {
		e := &Employee{
			GrossWage: 100000,
			HasHealthcare: false,
		}

		got := e.deductionValueHealthcare()
		expected := 0

		if got != expected {
			t.Fatalf("got=%d expected=%d", got, expected)
		}
	})
}

func TestDeductionValueDentalcare(t *testing.T) {
	t.Parallel()

	t.Run("Employee with 1,000.00 salary and with active dentalcare", func(t *testing.T) {
		e := &Employee{
			GrossWage: 100000,
			HasDentalcare: true,
		}

		got := e.deductionValueDentalcare()
		expected := 500

		if got != expected {
			t.Fatalf("got=%d expected=%d", got, expected)
		}
	})

	t.Run("Employee with 1,000.00 salary and without active dentalcare", func(t *testing.T) {
		e := &Employee{
			GrossWage: 100000,
			HasDentalcare: false,
		}

		got := e.deductionValueDentalcare()
		expected := 0

		if got != expected {
			t.Fatalf("got=%d expected=%d", got, expected)
		}
	})
}

func TestDeductionValueTransportationAllowance(t *testing.T) {
	t.Parallel()

	t.Run("Employee with 1,000.00 salary and with active transportation allowance", func(t *testing.T) {
		e := &Employee{
			GrossWage: 100000,
			HasTransportationAllowance: true,
		}

		got := e.deductionValueTransportationAllowance()
		expected := 0

		if got != expected {
			t.Fatalf("got=%d expected=%d", got, expected)
		}
	})

	t.Run("Employee with 1,800.00 salary and with active transportation allowance", func(t *testing.T) {
		e := &Employee{
			GrossWage: 180000,
			HasTransportationAllowance: true,
		}

		got := e.deductionValueTransportationAllowance()
		expected := 10800

		if got != expected {
			t.Fatalf("got=%d expected=%d", got, expected)
		}
	})

	t.Run("Employee with 1,800.00 salary and without active transportation allowance", func(t *testing.T) {
		e := &Employee{
			GrossWage: 180000,
			HasTransportationAllowance: false,
		}

		got := e.deductionValueTransportationAllowance()
		expected := 0

		if got != expected {
			t.Fatalf("got=%d expected=%d", got, expected)
		}
	})
}

func TestDeductionValueFGTS(t *testing.T) {
	t.Parallel()

	t.Run("Employee with 1,000.00 salary", func(t *testing.T) {
		e := &Employee{
			GrossWage: 100000,
		}

		got := e.deductionValueFGTS()
		expected := 8000

		if got != expected {
			t.Fatalf("got=%d expected=%d", got, expected)
		}
	})
}
