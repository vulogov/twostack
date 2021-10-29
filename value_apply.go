package twostack

import (
	"fmt"

	mapset "github.com/deckarep/golang-set"
	"github.com/lrita/cmap"
)

func (ts *TwoStack) ApplyGen(f GenFun) (*Elem, error) {
	err := GenericGenFun(ts, f)
	if err == nil {
		res, _ := ts.GetElem()
		return res, err
	}
	return nil, err
}

func (ts *TwoStack) ApplyFun(f FunFun) (*Elem, error) {
	if ts.Len() < 1 {
		return nil, fmt.Errorf("Stack is to shallow for executing .ApplyFun")
	}
	e, err := ts.TakeElem()
	if err != err {
		return nil, err
	}
	err = e.F(ts, e, f)
	if err == nil {
		res, _ := ts.GetElem()
		return res, err
	}
	return nil, err
}

func (ts *TwoStack) ApplyFunAll(f FunFun) (*Elem, error) {
	for i := 0; i < ts.Q().Len(); i++ {
		_, err := ts.ApplyFun(f)
		if err != nil {
			return nil, err
		}
		ts.CLeft()
	}
	return nil, nil
}

func (ts *TwoStack) ApplyOp(f OpFun) (*Elem, error) {
	if ts.Len() < 2 {
		return nil, fmt.Errorf("Stack is to shallow for executing .ApplyOp")
	}
	e1, err := ts.TakeElem()
	if err != err {
		return nil, err
	}
	e2, err := ts.TakeElem()
	if err != err {
		return nil, err
	}
	err = e1.Op(ts, e1, e2, f)
	if err == nil {
		res, _ := ts.GetElem()
		return res, err
	}
	return nil, err
}

func (ts *TwoStack) ApplyOpAll(f OpFun) (*Elem, error) {
	for ts.Len() > 1 {
		_, err := ts.ApplyOp(f)
		if err != nil {
			return nil, err
		}
	}
	res, _ := ts.GetElem()
	return res, nil
}

func (ts *TwoStack) Eval(f EvalFun) (*Elem, error) {
	p := mapset.NewSet()
	var kv cmap.Cmap

	q := ts.Q()
	if q.Len() < 1 {
		return nil, fmt.Errorf("Cwell stack is too shallow for .Eval()")
	}
	if ts.Mode == true {
		call := q.At(q.Len() - 1).(*Elem)
		kv.Store("__call__", call)
		if call.Type != Call_t {
			return nil, fmt.Errorf("#1 .Eval() expects Call_t element")
		}
		for i := 0; i < q.Len(); i++ {
			v := q.At(i).(*Elem)
			if v.Name == "#" {
				p.Add(v)
			}
			kv.Store(v.Name, v)
		}
	} else {
		call := q.At(0).(*Elem)
		kv.Store("__call__", call)
		if call.Type != Call_t {
			return nil, fmt.Errorf("#2 .Eval() expects Call_t element")
		}
		for i := 1; i < q.Len(); i++ {
			v := q.At(i).(*Elem)
			if v.Name == "#" {
				p.Add(v)
			}
			kv.Store(v.Name, v)
		}
	}
	res, err := f(&kv, &p)
	if err != nil {
		return nil, err
	}
	q.Clear()
	q.PushBack(res)
	return res, nil
}
