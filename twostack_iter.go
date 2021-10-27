package twostack

import (
	"github.com/gammazero/deque"
)

type TSIterp_f func(ts *TwoStack) bool
type TSIter_f func(ts *TwoStack, f *Elem)

func TSIterator(ts *TwoStack, q *deque.Deque, fpre TSIterp_f, f TSIter_f, fpost TSIterp_f) {
	if ts.Mode == true {
		for i := (q.Len() - 1); i >= 0; i-- {
			f(ts, q.At(i).(*Elem))
		}
	} else {
		for i := 0; i < q.Len(); i++ {
			f(ts, q.At(i).(*Elem))
		}
	}
}

func TSGlobalIterator(ts *TwoStack, fpre TSIterp_f, f TSIter_f, fpost TSIterp_f) bool {
	if fpre(ts) {
		if ts.Mode == true {
			for i := (ts.R.Len() - 1); i >= 0; i-- {
				f(ts, ts.R.At(i).(*deque.Deque).Back().(*Elem))
			}
		} else {
			for i := 0; i < ts.R.Len(); i++ {
				f(ts, ts.R.At(i).(*deque.Deque).Front().(*Elem))
			}
		}
		return fpost(ts)
	}
	return false
}

func TSGlobalTaker(ts *TwoStack, fpre TSIterp_f, f TSIter_f, fpost TSIterp_f) bool {
	if fpre(ts) {
		if ts.Mode == true {
			for i := (ts.R.Len() - 1); i >= 0; i-- {
				if ts.R.At(i).(*deque.Deque).Len() > 0 {
					f(ts, ts.R.At(i).(*deque.Deque).PopBack().(*Elem))
				} else {
					f(ts, None(""))
				}
			}
		} else {
			for i := 0; i < ts.R.Len(); i++ {
				if ts.R.At(i).(*deque.Deque).Len() > 0 {
					f(ts, ts.R.At(i).(*deque.Deque).PopFront().(*Elem))
				} else {
					f(ts, None(""))
				}
			}
		}
		return fpost(ts)
	}
	return false
}

func TSCellIterator(ts *TwoStack, fpre TSIterp_f, f TSIter_f, fpost TSIterp_f) bool {
	var q *deque.Deque
	if ts.Mode == true {
		q = ts.R.Back().(*deque.Deque)
	} else {
		q = ts.R.Front().(*deque.Deque)
	}
	if fpre(ts) {
		TSIterator(ts, q, fpre, f, fpost)
		return fpost(ts)
	}
	return false
}

func IterPre(ts *TwoStack) bool {
	return true
}

func IterPost(ts *TwoStack) bool {
	return true
}
