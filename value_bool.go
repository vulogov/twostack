package twostack

import (
	"fmt"

	"github.com/cstockton/go-conv"
	"github.com/google/uuid"
)

func Bool(name string, v bool, labels ...string) *Elem {
	e := NewElem(name, Bool_t, v, labels...)
	e.FromString = BoolFromString
	e.ToString = BoolToString
	return e
}

func MakeBool(name string, p string, labels ...string) *Elem {
	val := BoolFromString(p)
	if val == nil {
		return nil
	}
	return Bool(name, val.(bool), labels...)
}

func BoolFromString(d string) interface{} {
	res, err := conv.Bool(d)
	if err != nil {
		return nil
	}
	return res
}

func BoolToString(e *Elem) string {
	if e.Type == Bool_t {
		return fmt.Sprintf("%v", e.V.(bool))
	} else {
		return "UNKNOWN"
	}
}

func (ts *TwoStack) SetBool(v bool, labels ...string) {
	var name string
	if len(labels) > 0 {
		name = labels[0]
	} else {
		name = uuid.New().String()
	}
	e := Bool(name, v, labels...)
	ts.Set(e)
}

func (ts *TwoStack) MakeBool(p string, labels ...string) {
	var name string
	if len(labels) > 0 {
		name = labels[0]
	} else {
		name = uuid.New().String()
	}
	e := MakeBool(name, p, labels...)
	if e != nil {
		ts.Set(e)
	}
}

func (ts *TwoStack) GetBool() (bool, error) {
	e, err := ts.GetElem()
	if err != nil {
		return false, err
	}
	if e.Type != Bool_t {
		return false, fmt.Errorf("There is no boolean value on top of the stack: %v", e.Type)
	}
	return e.V.(bool), nil
}

func (ts *TwoStack) TakeBool() (bool, error) {
	e, err := ts.TakeElem()
	if err != nil {
		return false, err
	}
	if e.Type != Bool_t {
		return false, fmt.Errorf("There is no boolean value on top of the stack: %v", e.Type)
	}
	e, err = ts.TakeElem()
	if err != nil {
		return false, err
	}
	return e.V.(bool), nil
}
