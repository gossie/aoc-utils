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

	if r.String() != "((s + 5.000000) / 4.000000)" {
		t.Fatalf("expected %v to be ((s + 5.000000) / 4.000000)", r)
	}
}

func TestSolveToS(t *testing.T) {
	left := equations.Add(equations.Mul(equations.Num(4), equations.Var("r")), equations.Mul(equations.Num(0), equations.Num(7)))
	right := equations.Add(equations.Mul(equations.Num(1), equations.Var("s")), equations.Div(equations.Num(25), equations.Num(5)))

	s, _ := equations.NewEquation(left, right).SolveTo("s")

	if s.String() != "((4.000000 * r) - 5.000000)" {
		t.Fatalf("expected %v to be((4.000000 * r) - 5.000000)", s)
	}
}
