package equations

type matcher func(*value) bool

func num(n float64) matcher {
	return func(v *value) bool {
		return v.op == "num" && v.number == n
	}
}

func anyNum(n *float64) matcher {
	return func(v *value) bool {
		if v.op == "num" {
			*n = v.number
			return true
		}
		return false
	}
}

func variable(name string) matcher {
	return func(val *value) bool {
		return val.op == "var" && val.name == name
	}
}

func anyVariable(name *string) matcher {
	return func(v *value) bool {
		if v.op == "var" {
			*name = v.name
			return true
		}
		return false
	}
}

func any(val *value) matcher {
	return func(v *value) bool {
		*val = *v
		return v != nil
	}
}

func bin(leftOperand matcher, operation string, rightOperand matcher) matcher {
	return func(val *value) bool {
		return val.op == operation && leftOperand(val.left) && rightOperand(val.right)
	}
}
