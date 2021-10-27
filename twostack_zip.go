package twostack

import (
	"fmt"
)

func (ts *TwoStack) CZip() error {
	if (ts.GLen() - 1) != ts.Len() {
		return fmt.Errorf("Global size doesnt match a Cell stack size=%v", ts.Len())
	}
	q := ts.Q()
	if czippre(ts) {
		TSIterator(ts, q, czippre, czip, czippost)
		czippost(ts)
	}
	return nil
}

func czip(ts *TwoStack, f *Elem) {
	ts.Left()
	ts.SetElem(f)
}

func czippre(ts *TwoStack) bool {
	return true
}

func czippost(ts *TwoStack) bool {
	ts.Left()
	ts.Del()
	return true
}
