package twostack

import (
	"fmt"

	"github.com/gammazero/deque"
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

func (ts *TwoStack) GZip() error {
	if !TSGlobalTaker(ts, gzippre, gzip, gzippost) {
		return fmt.Errorf(".Gzip() operation had failed")
	}
	return nil
}

func gzip(ts *TwoStack, f *Elem) {
	q := ts.getTmp("_gzipq").(*deque.Deque)
	if q != nil {
		if ts.Mode == true {
			q.PushBack(f)
		} else {
			q.PushFront(f)
		}
	}
}

func gzippre(ts *TwoStack) bool {
	ts.setTmp("_gzipq", deque.New(0, minCap))
	return true
}

func gzippost(ts *TwoStack) bool {
	q := ts.getTmp("_gzipq").(*deque.Deque)
	if q != nil {
		ts.addQ(q)
		ts.delTmp("_gzipq")
		return true
	}
	return false
}

func (ts *TwoStack) GMerge() error {
	if ts.GLen() < 2 {
		return fmt.Errorf("Global Stack too shallow for .GMerge()")
	}
	q1 := ts.Q()
	ts.Left()
	for q1.Len() > 0 {
		if ts.Mode == true {
			v := q1.PopBack().(*Elem)
			ts.SetElem(v)
		} else {
			v := q1.PopFront().(*Elem)
			ts.SetElem(v)
		}
	}
	ts.Right()
	ts.Del()
	return nil
}
