package twostack

import "math"

func NumericIncrease(e1 *Elem) (*Elem, error) {
	v := e1.Float()
	if math.IsNaN(v) {
		return Float(e1.Name, v), nil
	}
	v += 1.0
	return Float(e1.Name, v), nil
}

func NumericIncreaseIfNum(e1 *Elem) (*Elem, error) {
	v := e1.Float()
	if math.IsNaN(v) {
		return e1, nil
	}
	v += 1.0
	return Float(e1.Name, v), nil
}

func OnlyIf42(e1 *Elem) bool {
	v := e1.Float()
	if math.IsNaN(v) {
		return false
	}
	if v == 42.0 {
		return true
	}
	return false
}

func Passthrough(e1 *Elem) (*Elem, error) {
	return e1, nil
}

func PassthroughFilter(e1 *Elem) bool {
	return true
}
