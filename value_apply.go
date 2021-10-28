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
