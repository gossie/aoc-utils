package equations

import (
	"errors"
	"fmt"
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

func findValue(val *value, name string) (*value, path, error) {
	if variable(name)(val) {
		return val, make(path, 0), nil
	}

	if val.left != nil || val.right != nil {
		if val.left != nil {
			found, p, err := findValue(val.left, name)
			if err == nil {
				op := rightComplements[val.op]
				p = append(p, &opValuePair{op, *val.right, false})
				return found, p, nil
			}
		}

		if val.right != nil {
			found, p, err := findValue(val.right, name)
			if err == nil {
				op := leftComplements[val.op]
				p = append(p, &opValuePair{op, *val.left, val.op == "-" || val.op == "/"})
				return found, p, nil
			}
		}
	}

	return nil, nil, errors.New("variable " + name + " not found")
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
	foundLeft, pLeft, errLeft := findValue(&e.left, varName)
	foundRight, pRight, errRight := findValue(&e.right, varName)

	if foundLeft != nil && foundRight != nil {
		return nil, errors.New(varName + " found on both sides, currently that cannot be handle")
	}

	if errLeft != nil && errRight != nil {
		return nil, errors.New(varName + " could not be found")
	}

	if foundLeft != nil {
		return processPath(e.right, pLeft)
	} else {
		return processPath(e.left, pRight)
	}
}

func (e equation) Set(varName string, val value) equation {
	return NewEquation(insert(e.left, varName, val), insert(e.right, varName, val))
}

func insert(current value, varName string, val value) value {
	if variable(varName)(&current) {
		return val
	}

	if op, present := operators[current.op]; present {
		return op(insert(*current.left, varName, val), insert(*current.right, varName, val))
	}

	return current
}

func processPath(val value, p path) (*value, error) {
	current := val
	for i := len(p) - 1; i >= 0; i-- {
		v := p[i]
		switch v.op {
		case "+":
			if v.swap {
				current = Add(v.val, current)
			} else {
				current = Add(current, v.val)
			}
		case "*":
			if v.swap {
				current = Mul(v.val, current)
			} else {
				current = Mul(current, v.val)
			}
		case "-":
			if v.swap {
				current = Sub(v.val, current)
			} else {
				current = Sub(current, v.val)
			}
		case "/":
			if v.swap {
				current = Div(v.val, current)
			} else {
				current = Div(current, v.val)
			}
		}
	}
	result := current.execute()
	return &result, nil
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

	var number1, number2 float64
	var varName1, varName2 string

	switch {
	case bin(anyNum(&number1), "+", anyNum(&number2))(&v):
		v.setNumber(v.left.number + v.right.number)
	case bin(anyNum(&number1), "*", anyNum(&number2))(&v):
		v.setNumber(v.left.number * v.right.number)
	case bin(anyNum(&number1), "-", anyNum(&number2))(&v):
		v.setNumber(v.left.number - v.right.number)
	case bin(anyNum(&number1), "/", anyNum(&number2))(&v):
		v.setNumber(v.left.number / v.right.number)
	case bin(any(), "*", num(0))(&v) || bin(num(0), "*", any())(&v):
		v.setNumber(0)
	case bin(any(), "*", num(1))(&v) || bin(any(), "/", num(1))(&v):
		v = *v.left
	case bin(num(1), "*", any())(&v):
		v = *v.right
	case bin(any(), "+", num(0))(&v) || bin(any(), "-", num(0))(&v):
		v = *v.left
	case bin(num(0), "+", any())(&v):
		v = *v.right
	case bin(bin(anyNum(&number1), "*", anyVariable(&varName1)), "+", bin(anyNum(&number2), "*", anyVariable(&varName2)))(&v) && varName1 == varName2:
		v.setVariable(number1+number2, varName1)
	case bin(bin(anyNum(&number1), "*", anyVariable(&varName1)), "-", bin(anyNum(&number2), "*", anyVariable(&varName2)))(&v) && varName1 == varName2:
		v.setVariable(number1-number2, varName1)
	case bin(bin(anyNum(&number1), "*", anyVariable(&varName1)), "*", anyNum(&number2))(&v):
		v.setVariable(number1*number2, varName1)
	case bin(anyNum(&number1), "*", bin(anyNum(&number2), "*", anyVariable(&varName1)))(&v):
		v.setVariable(number1*number2, varName1)
	case bin(bin(anyNum(&number1), "*", anyVariable(&varName1)), "/", anyNum(&number2))(&v):
		v.setVariable(number1/number2, varName1)
	case bin(anyNum(&number1), "*", bin(any(), "/", anyNum(&number2)))(&v) && number1 == number2:
		v = *v.right.left
	}
	return v
}

func (v *value) setNumber(number float64) {
	v.op = "num"
	v.number = number
	v.left = nil
	v.right = nil
}

func (v *value) setVariable(number float64, varName string) {
	factor := Num(number)
	variable := Var(varName)

	v.op = "*"
	v.left = &factor
	v.right = &variable
}

func (v value) String() string {
	switch v.op {
	default:
		panic("unknown operator: " + v.op)
	case "num":
		return fmt.Sprintf("%f", v.number)
	case "var":
		return v.name
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

func Var(name string) value {
	return value{name: name, op: "var"}
}
