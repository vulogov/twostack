package twostack

import "fmt"

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
