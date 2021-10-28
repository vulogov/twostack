package twostack

import (
	"fmt"

	"github.com/google/uuid"
)

func Call(name string, v string, labels ...string) *Elem {
	e := NewElem(name, Call_t, v, labels...)
	e.FromString = CallFromString
	e.ToString = CallToString
	return e
}

func MakeCall(name string, p string, labels ...string) *Elem {
	val := CallFromString(p)
	return Call(name, val.(string), labels...)
}

func CallFromString(d string) interface{} {
	return d
}

func CallToString(e *Elem) string {
	if e.Type == Call_t {
		return e.V.(string)
	} else {
		return "UNKNOWN"
	}
}

func (ts *TwoStack) SetCall(v string, labels ...string) {
	var name string
	if len(labels) > 0 {
		name = labels[0]
	} else {
		name = uuid.New().String()
	}
	e := Call(name, v, labels...)
	ts.Set(e)
}

func (ts *TwoStack) MakeCall(p string, labels ...string) {
	var name string
	if len(labels) > 0 {
		name = labels[0]
	} else {
		name = uuid.New().String()
	}
	e := MakeCall(name, p, labels...)
	ts.Set(e)
}

func (ts *TwoStack) GetCall() (string, error) {
	e, err := ts.GetElem()
	if err != nil {
		return "", err
	}
	if e.Type != Call_t {
		return "", fmt.Errorf("There is no call value on top of the stack: %v", e.Type)
	}
	return e.V.(string), nil
}

func (ts *TwoStack) TakeCall() (string, error) {
	e, err := ts.TakeElem()
	if err != nil {
		return "", err
	}
	if e.Type != Call_t {
		return "", fmt.Errorf("There is no call value on top of the stack: %v", e.Type)
	}
	e, err = ts.TakeElem()
	if err != nil {
		return "", err
	}
	return e.V.(string), nil
}
