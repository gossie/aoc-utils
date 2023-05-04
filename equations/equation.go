package equations

import (
	"errors"
	"fmt"
	"reflect"
)

type BinaryOp func(value, value) value

var operators = map[string]BinaryOp{
	"+": Add,
	"*": Mul,
	"-": Sub,
	"/": Div,
}

var rightComplements = map[string]string{
	"+": "-",
	"*": "/",
	"-": "+",
	"/": "*",
}

var leftComplements = map[string]string{
	"+": "-",
	"*": "/",
	"-": "-",
	"/": "/",
}

type opValuePair struct {
	op   string
	val  value
	swap bool
}

type path []*opValuePair

func findValue(val *value, name string) (*value, path, path, error) {
	if variable(name)(val) {
		return val, make(path, 0), append(make(path, 0), &opValuePair{"/", Num(val.number), false}), nil
	}

	if val.left != nil || val.right != nil {
		if val.left != nil {
			found, pathToValue, complementaryPath, err := findValue(val.left, name)
			if err == nil {
				op := rightComplements[val.op]
				pathToValue = append(pathToValue, &opValuePair{op, *val.left, val.op == "-" || val.op == "/"})
				complementaryPath = append(complementaryPath, &opValuePair{op, *val.right, false})
				return found, pathToValue, complementaryPath, nil
			}
		}

		if val.right != nil {
			found, pathToValue, complementaryPath, err := findValue(val.right, name)
			if err == nil {
				op := leftComplements[val.op]
				pathToValue = append(pathToValue, &opValuePair{op, *val.right, false})
				complementaryPath = append(complementaryPath, &opValuePair{op, *val.left, val.op == "-" || val.op == "/"})
				return found, pathToValue, complementaryPath, nil
			}
		}
	}

	return nil, nil, nil, errors.New("variable " + name + " not found")
}

type equation struct {
	left, right value
}

func NewEquation(left, right value) equation {
	return equation{left: left, right: right}
}

func (e equation) Optimize() equation {
	l := e.left.execute()
	r := e.right.execute()
	return NewEquation(l, r)
}

func (e equation) SolveTo(varName string) (*value, error) {
	left, _, leftComplementaryPath, errLeft := findValue(&e.left, varName)
	right, rightPath, rightComplementaryPath, errRight := findValue(&e.right, varName)

	fmt.Println(e)

	if left != nil && right != nil {
		eq := NewEquation(processPathElement(rightPath[len(rightPath)-1], e.left), processPathElement(rightPath[len(rightPath)-1], e.right))
		eq = eq.Optimize()
		fmt.Println(eq)
		return eq.SolveTo(varName)
	}

	if errLeft != nil && errRight != nil {
		return nil, errors.New(varName + " could not be found")
	}

	if left != nil {
		return processPath(e.right, leftComplementaryPath)
	} else {
		return processPath(e.left, rightComplementaryPath)
	}
}

func (e equation) Set(varName string, val value) equation {
	newLeft := insert(e.left, varName, val)
	newRight := insert(e.right, varName, val)
	return NewEquation(newLeft, newRight)
}

func insert(current value, varName string, val value) value {
	if variable(varName)(&current) {
		return Mul(Num(current.number), val)
	}

	if op, present := operators[current.op]; present {
		return op(insert(*current.left, varName, val), insert(*current.right, varName, val))
	}

	return current
}

func processPath(val value, p path) (*value, error) {
	current := val
	for i := len(p) - 1; i >= 0; i-- {
		current = processPathElement(p[i], current)
	}
	result := current.execute()
	return &result, nil
}

func processPathElement(v *opValuePair, current value) value {
	switch v.op {
	default:
		panic("unkown operator " + v.op)
	case "+":
		if v.swap {
			return Add(v.val, current)
		} else {
			return Add(current, v.val)
		}
	case "*":
		if v.swap {
			return Mul(v.val, current)
		} else {
			return Mul(current, v.val)
		}
	case "-":
		if v.swap {
			return Sub(v.val, current)
		} else {
			return Sub(current, v.val)
		}
	case "/":
		if v.swap {
			return Div(v.val, current)
		} else {
			return Div(current, v.val)
		}
	}
}

func (e equation) String() string {
	return fmt.Sprintf("%v = %v", e.left, e.right)
}

type value struct {
	left, right *value
	op          string
	number      float64
	name        string
}

func (v value) execute() value {
	if v.left != nil && v.right != nil {
		l := v.left.execute()
		r := v.right.execute()
		v.left = &l
		v.right = &r
	}

	var val1, val2, val3 value
	var number1, number2, number3 float64
	var varName1, varName2 string

	switch {
	// Grundrechenarten
	case bin(anyNum(&number1), "+", anyNum(&number2))(&v):
		v = Num(v.left.number + v.right.number)
	case bin(anyNum(&number1), "*", anyNum(&number2))(&v):
		v = Num(v.left.number * v.right.number)
	case bin(anyNum(&number1), "-", anyNum(&number2))(&v):
		v = Num(v.left.number - v.right.number)
	case bin(anyNum(&number1), "/", anyNum(&number2))(&v):
		v = Num(v.left.number / v.right.number)
	// Multiplikation mit 0 (kommutativ)
	case bin(any(&val1), "*", num(0))(&v) || bin(num(0), "*", any(&val2))(&v):
		v = Num(0)
	// Multiplikation mit 1 (kommutativ), Division durch 1, Addition mit 0 (kommutativ) oder Subtraktion von 0
	case bin(any(&val1), "*", num(1))(&v) || bin(num(1), "*", any(&val1))(&v) || bin(any(&val1), "/", num(1))(&v) || bin(any(&val1), "+", num(0))(&v) || bin(num(0), "+", any(&val1))(&v) || bin(any(&val1), "-", num(0))(&v):
		v = val1
	case bin(bin(any(&val1), "+", any(&val2)), "-", any(&val3))(&v) && reflect.DeepEqual(val1, val3):
		v = val2
	case bin(bin(any(&val1), "+", any(&val2)), "-", any(&val3))(&v) && reflect.DeepEqual(val2, val3):
		v = val1
	// Rechnen mit Variable
	case bin(anyVariable(&number1, &varName1), "/", anyNum(&number2))(&v):
		v = Var(number1/number2, varName1)
	case bin(anyVariable(&number1, &varName1), "*", anyNum(&number2))(&v) || bin(anyNum(&number1), "*", anyVariable(&number2, &varName1))(&v):
		v = Var(number1*number2, varName1)
	case bin(anyVariable(&number1, &varName1), "+", anyVariable(&number2, &varName2))(&v) && varName1 == varName2:
		v = Var(number1+number2, varName1)
	case bin(anyVariable(&number1, &varName1), "-", anyVariable(&number2, &varName2))(&v) && varName1 == varName2:
		v = Var(number1-number2, varName1)
	case bin(anyVariable(&number1, &varName1), "*", anyNum(&number2))(&v) || bin(anyNum(&number1), "*", anyVariable(&number2, &varName1))(&v):
		v = Var(number1*number2, varName1)
	case bin(anyVariable(&number1, &varName1), "/", anyNum(&number2))(&v):
		v = Var(number1/number2, varName1)
	// Distributivgesetze
	case bin(bin(any(&val1), "+", any(&val2)), "*", anyNum(&number1))(&v) || bin(anyNum(&number1), "*", bin(any(&val1), "+", any(&val2)))(&v):
		v = Add(Mul(Num(number1), val1), Mul(Num(number1), val2)).execute()
	case bin(bin(any(&val1), "-", any(&val2)), "*", anyNum(&number1))(&v) || bin(anyNum(&number1), "*", bin(any(&val1), "-", any(&val2)))(&v):
		v = Sub(Mul(Num(number1), val1), Mul(Num(number1), val2)).execute()
	case bin(bin(any(&val1), "+", any(&val2)), "/", anyNum(&number1))(&v):
		v = Add(Div(val1, Num(number1)), Div(val2, Num(number1))).execute()
	case bin(bin(any(&val1), "-", any(&val2)), "/", anyNum(&number1))(&v):
		v = Sub(Div(val1, Num(number1)), Div(val2, Num(number1))).execute()
	// Assoziativgesetze? (varible wird bearbeitet)
	case bin(bin(anyVariable(&number1, &varName1), "+", anyNum(&number2)), "-", anyVariable(&number3, &varName2))(&v):
		v = Add(Var(number1-number3, varName1), Num(number2))
	}
	return v
}

func (v value) String() string {
	switch v.op {
	default:
		panic("unknown operator: " + v.op)
	case "num":
		return fmt.Sprintf("%f", v.number)
	case "var":
		return fmt.Sprintf("%f%v", v.number, v.name)
	case "+":
		return fmt.Sprintf("(%v + %v)", v.left, v.right)
	case "*":
		return fmt.Sprintf("(%v * %v)", v.left, v.right)
	case "-":
		return fmt.Sprintf("(%v - %v)", v.left, v.right)
	case "/":
		return fmt.Sprintf("(%v / %v)", v.left, v.right)
	}
}

func Num(number float64) value {
	return value{number: number, op: "num"}
}

func Add(left, right value) value {
	return value{left: &left, right: &right, op: "+"}
}

func Sub(left, right value) value {
	return value{left: &left, right: &right, op: "-"}
}

func Mul(left, right value) value {
	return value{left: &left, right: &right, op: "*"}
}

func Div(left, right value) value {
	return value{left: &left, right: &right, op: "/"}
}

func Var(factor float64, name string) value {
	return value{number: factor, name: name, op: "var"}
}
