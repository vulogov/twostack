package twostack

import (
	"fmt"

	"github.com/google/uuid"
)

func String(name string, v string, labels ...string) *Elem {
	e := NewElem(name, String_t, v, labels...)
	e.FromString = StringFromString
	e.ToString = StringToString
	return e
}

func MakeString(name string, p string, labels ...string) *Elem {
	val := StringFromString(p)
	return String(name, val.(string), labels...)
}

func StringFromString(d string) interface{} {
	return d
}

func StringToString(e *Elem) string {
	if e.Type == String_t {
		return e.V.(string)
	} else {
		return "UNKNOWN"
	}
}

func (ts *TwoStack) SetString(v string, labels ...string) {
	var name string
	if len(labels) > 0 {
		name = labels[0]
	} else {
		name = uuid.New().String()
	}
	e := String(name, v, labels...)
	ts.Set(e)
}

func (ts *TwoStack) MakeString(p string, labels ...string) {
	var name string
	if len(labels) > 0 {
		name = labels[0]
	} else {
		name = uuid.New().String()
	}
	e := MakeString(name, p, labels...)
	ts.Set(e)
}

func (ts *TwoStack) GetString() (string, error) {
	e, err := ts.GetElem()
	if err != nil {
		return "", err
	}
	if e.Type != String_t {
		return "", fmt.Errorf("There is no string value on top of the stack: %v", e.Type)
	}
	return e.V.(string), nil
}

func (ts *TwoStack) TakeString() (string, error) {
	e, err := ts.TakeElem()
	if err != nil {
		return "", err
	}
	if e.Type != String_t {
		return "", fmt.Errorf("There is no string value on top of the stack: %v", e.Type)
	}
	e, err = ts.TakeElem()
	if err != nil {
		return "", err
	}
	return e.V.(string), nil
}
