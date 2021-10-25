package twostack

import (
	"fmt"

	"github.com/cstockton/go-conv"
)

func Int(name string, v int, labels ...string) *Elem {
	return Int64(name, int64(v), labels...)
}

func Int64(name string, v int64, labels ...string) *Elem {
	e := NewElem(name, Int_t, v, labels...)
	e.FromString = IntFromString
	e.ToString = IntToString
	return e
}

func MakeInt(name string, p string, labels ...string) *Elem {
	val := IntFromString(p)
	if val == nil {
		return nil
	}
	return Int64(name, val.(int64), labels...)
}

func IntFromString(d string) interface{} {
	res, err := conv.Int64(d)
	if err != nil {
		return nil
	}
	return res
}

func IntToString(e *Elem) string {
	if e.Type == Int_t {
		return fmt.Sprintf("%v", e.V.(int64))
	} else {
		return "UNKNOWN"
	}
}

func (ts *TwoStack) SetInt(v int, name string, labels ...string) {
	ts.SetInt64(int64(v), name, labels...)
}

func (ts *TwoStack) SetInt64(v int64, name string, labels ...string) {
	e := Int64(name, v, labels...)
	ts.Set(e)
}

func (ts *TwoStack) MakeInt(p string, name string, labels ...string) {
	e := MakeInt(name, p, labels...)
	if e != nil {
		ts.Set(e)
	}
}

func (ts *TwoStack) GetInt() (int64, error) {
	e, err := ts.GetElem()
	if err != nil {
		return int64(0), err
	}
	if e.Type != Int_t {
		return int64(0), fmt.Errorf("There is no int value on top of the stack: %v", e.Type)
	}
	return e.V.(int64), nil
}

func (ts *TwoStack) TakeInt() (int64, error) {
	e, err := ts.TakeElem()
	if err != nil {
		return int64(0), err
	}
	if e.Type != Int_t {
		return int64(0), fmt.Errorf("There is no int value on top of the stack: %v", e.Type)
	}
	e, err = ts.TakeElem()
	if err != nil {
		return int64(0), err
	}
	return e.V.(int64), nil
}
