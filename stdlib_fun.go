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
