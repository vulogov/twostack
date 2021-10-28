package twostack

import (
	"fmt"
	"math"

	"github.com/cstockton/go-conv"
	"github.com/google/uuid"
)

func Float(name string, v float64, labels ...string) *Elem {
	e := NewElem(name, Float_t, v, labels...)
	e.FromString = FloatFromString
	e.ToString = FloatToString
	return e
}

func MakeFloat(name string, p string, labels ...string) *Elem {
	val := FloatFromString(p)
	if val == nil {
		return nil
	}
	return Float(name, val.(float64), labels...)
}

func FloatFromString(d string) interface{} {
	res, err := conv.Float64(d)
	if err != nil {
		return nil
	}
	return res
}

func FloatToString(e *Elem) string {
	if e.Type == Float_t {
		return fmt.Sprintf("%v", e.V.(float64))
	} else {
		return "UNKNOWN"
	}
}

func (ts *TwoStack) SetFloat(v float64, labels ...string) {
	var name string
	if len(labels) > 0 {
		name = labels[0]
	} else {
		name = uuid.New().String()
	}
	e := Float(name, v, labels...)
	ts.Set(e)
}

func (ts *TwoStack) MakeFloat(p string, labels ...string) {
	var name string
	if len(labels) > 0 {
		name = labels[0]
	} else {
		name = uuid.New().String()
	}
	e := MakeFloat(name, p, labels...)
	if e != nil {
		ts.Set(e)
	}
}

func (ts *TwoStack) GetFloat() (float64, error) {
	e, err := ts.GetElem()
	if err != nil {
		return float64(0), err
	}
	if e.Type != Float_t {
		return float64(0), fmt.Errorf("There is no float value on top of the stack: %v", e.Type)
	}
	return e.V.(float64), nil
}

func (ts *TwoStack) Float() float64 {
	e, err := ts.GetElem()
	if err != nil {
		return math.NaN()
	}
	switch e.Type {
	case Float_t:
		return e.V.(float64)
	case Int_t:
		return float64(e.V.(int64))
	case String_t:
		v := FloatFromString(e.V.(string))
		if v != nil {
			return v.(float64)
		}
	}
	return math.NaN()
}

func (ts *TwoStack) TakeFloat() (float64, error) {
	e, err := ts.TakeElem()
	if err != nil {
		return float64(0), err
	}
	if e.Type != Float_t {
		return float64(0), fmt.Errorf("There is no float value on top of the stack: %v", e.Type)
	}
	e, err = ts.TakeElem()
	if err != nil {
		return float64(0), err
	}
	return e.V.(float64), nil
}
