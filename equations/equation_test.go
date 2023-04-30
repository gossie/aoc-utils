package equations_test

import (
	"testing"

	"github.com/gossie/aoc-utils/equations"
)

func TestOptimze(t *testing.T) {
	left := equations.Add(equations.Mul(equations.Num(4), equations.Var("r")), equations.Mul(equations.Num(0), equations.Num(7)))
	right := equations.Add(equations.Mul(equations.Num(1), equations.Var("s")), equations.Div(equations.Num(25), equations.Num(5)))
	eq := equations.NewEquation(left, right)

	if eq.Optimize().String() != "(4.000000 * r) = (s + 5.000000)" {
		t.Fatalf("expected %v to be (4.000000 * r) = (s + 5.000000)", eq)
	}
}

func TestSolveToR(t *testing.T) {
	left := equations.Add(equations.Mul(equations.Num(4), equations.Var("r")), equations.Mul(equations.Num(0), equations.Num(7)))
	right := equations.Add(equations.Mul(equations.Num(1), equations.Var("s")), equations.Div(equations.Num(25), equations.Num(5)))

	r, _ := equations.NewEquation(left, right).SolveTo("r")

	if r.String() != "((s / 4.000000) + 1.250000)" {
		t.Fatalf("expected %v to be ((s / 4.000000) + 1.250000)", r)
	}
}

func TestSolveToS(t *testing.T) {
	left := equations.Add(equations.Mul(equations.Num(4), equations.Var("r")), equations.Mul(equations.Num(0), equations.Num(7)))
	right := equations.Add(equations.Mul(equations.Num(1), equations.Var("s")), equations.Div(equations.Num(25), equations.Num(5)))

	original := equations.NewEquation(left, right)
	s, _ := original.SolveTo("s")

	if s.String() != "((4.000000 * r) - 5.000000)" {
		t.Fatalf("expected %v to be((4.000000 * r) - 5.000000)", s)
	}

	if original.String() != "((4.000000 * r) + (0.000000 * 7.000000)) = ((1.000000 * s) + (25.000000 / 5.000000))" {
		t.Fatalf("expected %v to be ((4.000000 * r) + (0.000000 * 7.000000)) = ((1.000000 * s) + (25.000000 / 5.000000))", original)
	}
}

func TestSet(t *testing.T) {
	left := equations.Add(equations.Mul(equations.Num(4), equations.Var("r")), equations.Mul(equations.Num(0), equations.Num(7)))
	right := equations.Add(equations.Mul(equations.Num(1), equations.Var("s")), equations.Div(equations.Num(25), equations.Num(5)))

	r := equations.Div(equations.Add(equations.Var("s"), equations.Num(5)), equations.Num(4))

	eq := equations.NewEquation(left, right).Set("r", r)
	if eq.String() != "((4.000000 * ((s + 5.000000) / 4.000000)) + (0.000000 * 7.000000)) = ((1.000000 * s) + (25.000000 / 5.000000))" {
		t.Fatalf("expected %v to be ((4.000000 * ((s + 5.000000) / 4.000000)) + (0.000000 * 7.000000)) = ((1.000000 * s) + (25.000000 / 5.000000))", eq)
	}

	eq = eq.Optimize()
	if eq.String() != "(s + 5.000000) = (s + 5.000000)" {
		t.Fatalf("expected %v to be (s + 5.000000) = (s + 5.000000)", eq)
	}
}

func TestOptimize_1(t *testing.T) {
	left := equations.Add(equations.Mul(equations.Num(4), equations.Var("r")), equations.Mul(equations.Num(2), equations.Var("r")))
	right := equations.Num(12.000000)

	eq := equations.NewEquation(left, right).Optimize()
	if eq.String() != "(6.000000 * r) = 12.000000" {
		t.Fatalf("expected %v to be (6.000000 * r) = 12.000000", eq)
	}
}

func TestOptimize_2(t *testing.T) {
	left := equations.Sub(equations.Mul(equations.Num(4), equations.Var("r")), equations.Mul(equations.Num(2), equations.Var("r")))
	right := equations.Num(12.000000)

	eq := equations.NewEquation(left, right).Optimize()
	if eq.String() != "(2.000000 * r) = 12.000000" {
		t.Fatalf("expected %v to be (2.000000 * r) = 12.000000", eq)
	}
}

func TestOptimize_3(t *testing.T) {
	left := equations.Mul(equations.Mul(equations.Num(4), equations.Var("r")), equations.Num(2))
	right := equations.Num(12.000000)

	eq := equations.NewEquation(left, right).Optimize()
	if eq.String() != "(8.000000 * r) = 12.000000" {
		t.Fatalf("expected %v to be (8.000000 * r) = 12.000000", eq)
	}
}

func TestOptimize_4(t *testing.T) {
	left := equations.Mul(equations.Num(2), equations.Mul(equations.Num(4), equations.Var("r")))
	right := equations.Num(12.000000)

	eq := equations.NewEquation(left, right).Optimize()
	if eq.String() != "(8.000000 * r) = 12.000000" {
		t.Fatalf("expected %v to be (8.000000 * r) = 12.000000", eq)
	}
}

func TestOptimize_5(t *testing.T) {
	left := equations.Div(equations.Mul(equations.Num(4), equations.Var("r")), equations.Num(2))
	right := equations.Num(12.000000)

	eq := equations.NewEquation(left, right).Optimize()
	if eq.String() != "(2.000000 * r) = 12.000000" {
		t.Fatalf("expected %v to be (2.000000 * r) = 12.000000", eq)
	}
}

func TestOptimize_6(t *testing.T) {
	left := equations.Mul(equations.Mul(equations.Num(4), equations.Var("r")), equations.Num(1))
	right := equations.Num(12.000000)

	eq := equations.NewEquation(left, right).Optimize()
	if eq.String() != "(4.000000 * r) = 12.000000" {
		t.Fatalf("expected %v to be (4.000000 * r) = 12.000000", eq)
	}
}

func TestOptimize_7(t *testing.T) {
	left := equations.Mul(equations.Num(1), equations.Mul(equations.Num(4), equations.Var("r")))
	right := equations.Num(12.000000)

	eq := equations.NewEquation(left, right).Optimize()
	if eq.String() != "(4.000000 * r) = 12.000000" {
		t.Fatalf("expected %v to be (4.000000 * r) = 12.000000", eq)
	}
}

func TestOptimize_8(t *testing.T) {
	left := equations.Mul(equations.Num(2), equations.Sub(equations.Mul(equations.Num(1), equations.Var("r")), equations.Num(4)))
	right := equations.Num(12.000000)

	eq := equations.NewEquation(left, right).Optimize()
	if eq.String() != "((r * 2.000000) - 8.000000) = 12.000000" {
		t.Fatalf("expected %v to be ((r * 2.000000) - 8.000000) = 12.000000", eq)
	}
}

func TestOptimize_9(t *testing.T) {
	left := equations.Mul(equations.Sub(equations.Mul(equations.Num(1), equations.Var("r")), equations.Num(4)), equations.Num(2))
	right := equations.Num(12.000000)

	eq := equations.NewEquation(left, right).Optimize()
	if eq.String() != "((r * 2.000000) - 8.000000) = 12.000000" {
		t.Fatalf("expected %v to be ((r * 2.000000) - 8.000000) = 12.000000", eq)
	}
}
