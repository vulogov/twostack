package twostack

import (
	mapset "github.com/deckarep/golang-set"
	"github.com/google/uuid"
	"github.com/lrita/cmap"
)

const (
	Bool_t     = 0
	Int_t      = 1
	Float_t    = 2
	String_t   = 3
	Duration_t = 4
	Call_t     = 5
	None_t     = 99
)

const (
	GT  = 10
	GTE = 11
	LS  = 20
	LSE = 21
	EQ  = 0
	NEQ = 1
	IDK = 99
)

type OpFun func(e1 *Elem, e2 *Elem) (*Elem, error)
type FunFun func(e1 *Elem) (*Elem, error)
type GenFun func() (*Elem, error)
type EvalFun func(kv *cmap.Cmap, p *mapset.Set) (*Elem, error)
type FromStringFun func(d string) interface{}
type ToStringFun func(e *Elem) string
type ApplyOpFun func(ts *TwoStack, e1 *Elem, e2 *Elem, f OpFun) error
type ApplyFunFun func(ts *TwoStack, e1 *Elem, f FunFun) error
type ApplyGenFun func(ts *TwoStack, f GenFun) error

type Elem struct {
	Type       int
	V          interface{}
	Name       string
	Labels     mapset.Set
	FromString FromStringFun
	ToString   ToStringFun
	Op         ApplyOpFun
	F          ApplyFunFun
	G          ApplyGenFun
}

func NewElem(name string, t int, val interface{}, labels ...string) *Elem {
	e := Elem{Type: t, V: val, Labels: mapset.NewSet()}
	e.Labels.Add(name)
	e.Name = name
	for _, v := range labels {
		e.Labels.Add(v)
	}
	e.FromString = NoneFromString
	e.ToString = NoneToString
	e.F = GenericFunFun
	e.Op = GenericOpFun
	e.G = GenericGenFun
	return &e
}

func (ts *TwoStack) GetElem() (*Elem, error) {
	e, err := ts.Get()
	return e.(*Elem), err
}

func (ts *TwoStack) TakeElem() (*Elem, error) {
	e, err := ts.Take()
	return e.(*Elem), err
}

func (ts *TwoStack) SetElem(e *Elem) {
	ts.Set(e)
}

func (e *Elem) String() string {
	return e.ToString(e)
}

func (ts *TwoStack) G() (interface{}, error) {
	_v, err := ts.Get()
	if err != nil {
		return nil, err
	}
	switch v := _v.(type) {
	case *Elem:
		if err != nil {
			return nil, err
		}
		switch v.Type {
		case Bool_t:
			return v.V.(bool), nil
		case Int_t:
			return v.V.(int64), nil
		case String_t:
			return v.V.(string), nil
		case Float_t:
			return v.V.(float64), nil
		case Call_t:
			return v.V.(string), nil
		case None_t:
			return nil, nil
		}
	}
	return nil, nil
}

func (ts *TwoStack) T() (interface{}, error) {
	_v, err := ts.Take()
	if err != nil {
		return nil, err
	}
	switch v := _v.(type) {
	case *Elem:
		if err != nil {
			return nil, err
		}
		switch v.Type {
		case Bool_t:
			return v.V.(bool), nil
		case Int_t:
			return v.V.(int64), nil
		case String_t:
			return v.V.(string), nil
		case Call_t:
			return v.V.(string), nil
		case Float_t:
			return v.V.(float64), nil
		case None_t:
			return nil, nil
		}
	}
	return nil, nil
}

func (ts *TwoStack) Put(d interface{}, labels ...string) bool {
	var name string
	var e *Elem
	if len(labels) > 0 {
		name = labels[0]
	} else {
		name = uuid.New().String()
	}
	switch d.(type) {
	case int:
		e = Int(name, d.(int), labels...)
	case int64:
		e = Int64(name, d.(int64), labels...)
	case float32:
		e = Float(name, float64(d.(float32)), labels...)
	case float64:
		e = Float(name, d.(float64), labels...)
	case bool:
		e = Bool(name, d.(bool), labels...)
	case string:
		e = String(name, d.(string), labels...)
	}
	if e != nil {
		if ts.IsIF {
			e1, err := ts.InstF(e)
			if err != nil {
				return false
			}
			ts.Set(e1)
			return true
		}
		ts.Set(e)
		return true
	} else {
		return false
	}
}

func NoneFromString(d string) interface{} {
	return nil
}

func NoneToString(e *Elem) string {
	return "None"
}

func GenericGenFun(ts *TwoStack, f GenFun) error {
	res, err := f()
	ts.SetElem(res)
	return err
}

func GenericFunFun(ts *TwoStack, e1 *Elem, f FunFun) error {
	res, err := f(e1)
	ts.SetElem(res)
	return err
}

func GenericOpFun(ts *TwoStack, e1 *Elem, e2 *Elem, f OpFun) error {
	res, err := f(e1, e2)
	ts.SetElem(res)
	return err
}
