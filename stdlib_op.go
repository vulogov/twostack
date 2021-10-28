package twostack

import (
	"fmt"
	"math"
)

func NumericAdd(e1 *Elem, e2 *Elem) (*Elem, error) {
	v1 := e1.Float()
	v2 := e2.Float()
	if math.IsNaN(v1) || math.IsNaN(v2) {
		return nil, fmt.Errorf("Incompartible data for arithmetic operation (+)")
	}
	return Float(e1.Name, (v1 + v2)), nil
}
