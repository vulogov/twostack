package twostack

import (
	"github.com/gammazero/deque"
)

func (ts *TwoStack) String() string {
	out := ""
	for i := 0; i < ts.R.Len(); i++ {
		q := ts.R.At(i).(*deque.Deque)
		if q.Len() > 0 {
			for j := 0; j < q.Len(); j++ {
				switch q.At(j).(type) {
				case *Elem:
					out += q.At(j).(*Elem).String()
					out += " "
				}
			}
			out += "\n"
		}
	}
	return out
}
