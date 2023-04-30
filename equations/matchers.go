package equations

type matcher func(*value) bool

func num(n float64) matcher {
	return func(v *value) bool {
		return v.op == "num" && v.number == n
	}
}

func anyNum() matcher {
	return func(v *value) bool {
		return v.op == "num"
	}
}

func variable(name string) matcher {
	return func(val *value) bool {
		return val.op == "var" && val.name == name
	}
}

func any() matcher {
	return func(v *value) bool {
		return v != nil
	}
}

func bin(leftOperand matcher, operation string, rightOperand matcher) matcher {
	return func(val *value) bool {
		return val.op == operation && leftOperand(val.left) && rightOperand(val.right)
	}
}
