package equations

import (
	"fmt"
	"testing"
)

func TestOptimze(t *testing.T) {
	left := Add(Var(4, "r"), Mul(Num(0), Num(7)))
	right := Add(Var(1, "s"), Div(Num(25), Num(5)))
	eq := NewEquation(left, right)

	if eq.optimize().String() != "4.000000r = (1.000000s + 5.000000)" {
		t.Fatalf("expected %v to be 4.000000r = (1.000000s + 5.000000)", eq)
	}
}

func TestSolveToR(t *testing.T) {
	left := Add(Var(4, "r"), Mul(Num(0), Num(7)))
	right := Add(Var(1, "s"), Div(Num(25), Num(5)))

	r, _ := NewEquation(left, right).SolveTo("r")

	if r.String() != "(0.250000s + 1.250000)" {
		t.Fatalf("expected %v to be (0.250000s + 1.250000)", r)
	}
}

func TestSolveToS(t *testing.T) {
	left := Add(Var(4, "r"), Mul(Num(0), Num(7)))
	right := Add(Var(1, "s"), Div(Num(25), Num(5)))

	original := NewEquation(left, right)
	s, _ := original.SolveTo("s")

	if s.String() != "(4.000000r + -5.000000)" {
		t.Fatalf("expected %v to be (4.000000r + -5.000000)", s)
	}

	if original.String() != "(4.000000r + (0.000000 * 7.000000)) = (1.000000s + (25.000000 / 5.000000))" {
		t.Fatalf("expected %v to be (4.000000r + (0.000000 * 7.000000)) = (1.000000s + (25.000000 / 5.000000))", original)
	}
}

func TestSolveTo_variableOnBothSides(t *testing.T) {
	left := Add(Var(4, "x"), Mul(Num(2), Num(7)))
	right := Add(Var(2, "x"), Div(Num(25), Num(5)))

	original := NewEquation(left, right)
	s, _ := original.SolveTo("x")

	if s.String() != fmt.Sprintf("%f", -9.0/2.0) {
		t.Fatalf("expected %v to be 4.5", s)
	}
}

func TestSet(t *testing.T) {
	left := Add(Var(4, "r"), Mul(Num(0), Num(7)))
	right := Add(Var(1, "s"), Div(Num(25), Num(5)))

	r := Div(Add(Var(1, "s"), Num(5)), Num(4))

	eq := NewEquation(left, right).Set("r", r)
	if eq.String() != "((4.000000 * ((1.000000s + 5.000000) / 4.000000)) + (0.000000 * 7.000000)) = (1.000000s + (25.000000 / 5.000000))" {
		t.Fatalf("expected %v to be ((4.000000 * ((1.000000s + 5.000000) / 4.000000)) + (0.000000 * 7.000000)) = (1.000000s + (25.000000 / 5.000000))", eq)
	}

	eq = eq.optimize()
	if eq.String() != "(1.000000s + 5.000000) = (1.000000s + 5.000000)" {
		t.Fatalf("expected %v to be (1.000000s + 5.000000) = (1.000000s + 5.000000)", eq)
	}
}

func TestOptimize_1(t *testing.T) {
	left := Add(Var(4, "r"), Var(2, "r"))
	right := Num(12.000000)

	eq := NewEquation(left, right).optimize()
	if eq.String() != "6.000000r = 12.000000" {
		t.Fatalf("expected %v to be 6.000000r = 12.000000", eq)
	}
}

func TestOptimize_2(t *testing.T) {
	left := Sub(Var(4, "r"), Var(2, "r"))
	right := Num(12.000000)

	eq := NewEquation(left, right).optimize()
	if eq.String() != "2.000000r = 12.000000" {
		t.Fatalf("expected %v to be 2.000000r = 12.000000", eq)
	}
}

func TestOptimize_3(t *testing.T) {
	left := Mul(Var(4, "r"), Num(2))
	right := Num(12.000000)

	eq := NewEquation(left, right).optimize()
	if eq.String() != "8.000000r = 12.000000" {
		t.Fatalf("expected %v to be 8.000000r = 12.000000", eq)
	}
}

func TestOptimize_4(t *testing.T) {
	left := Mul(Num(2), Var(4, "r"))
	right := Num(12.000000)

	eq := NewEquation(left, right).optimize()
	if eq.String() != "8.000000r = 12.000000" {
		t.Fatalf("expected %v to be 8.000000r = 12.000000", eq)
	}
}

func TestOptimize_5(t *testing.T) {
	left := Div(Var(4, "r"), Num(2))
	right := Num(12.000000)

	eq := NewEquation(left, right).optimize()
	if eq.String() != "2.000000r = 12.000000" {
		t.Fatalf("expected %v to be 2.000000r = 12.000000", eq)
	}
}

func TestOptimize_6(t *testing.T) {
	left := Mul(Var(4, "r"), Num(1))
	right := Num(12.000000)

	eq := NewEquation(left, right).optimize()
	if eq.String() != "4.000000r = 12.000000" {
		t.Fatalf("expected %v to be 4.000000r = 12.000000", eq)
	}
}

func TestOptimize_7(t *testing.T) {
	left := Mul(Num(1), Var(4, "r"))
	right := Num(12.000000)

	eq := NewEquation(left, right).optimize()
	if eq.String() != "4.000000r = 12.000000" {
		t.Fatalf("expected %v to be 4.000000r = 12.000000", eq)
	}
}

func TestOptimize_8(t *testing.T) {
	left := Mul(Num(2), Sub(Var(1, "r"), Num(4)))
	right := Num(12.000000)

	eq := NewEquation(left, right).optimize()
	if eq.String() != "(2.000000r + -8.000000) = 12.000000" {
		t.Fatalf("expected %v to be (2.000000r + -8.000000) = 12.000000", eq)
	}
}

func TestOptimize_9(t *testing.T) {
	left := Mul(Sub(Var(1, "r"), Num(4)), Num(2))
	right := Num(12.000000)

	eq := NewEquation(left, right).optimize()
	if eq.String() != "(2.000000r + -8.000000) = 12.000000" {
		t.Fatalf("expected %v to be (2.000000r + -8.000000) = 12.000000", eq)
	}
}
