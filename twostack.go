package twostack

import (
	"fmt"

	"github.com/gammazero/deque"
	"github.com/google/uuid"
	"github.com/lrita/cmap"
)

type TwoStack struct {
	R      deque.Deque
	C      cmap.Cmap
	tmp    cmap.Cmap
	ID     string
	Status bool
	Mode   bool
}

const minCap = 1024

//
// Initialize new 2-dimentional stack structure
//
func Init() *TwoStack {
	ts := &TwoStack{
		Status: true,
		Mode:   true,
		ID:     uuid.New().String(),
	}
	ts.R.PushBack(deque.New(0, minCap))
	if ts.R.Len() != 1 {
		ts.Status = false
	}
	return ts
}

func (ts *TwoStack) Left() {
	ts.R.Rotate(-1)
}

func (ts *TwoStack) CLeft() {
	q := ts.Q()
	q.Rotate(-1)
}

func (ts *TwoStack) Right() {
	ts.R.Rotate(1)
}

func (ts *TwoStack) CRight() {
	q := ts.Q()
	q.Rotate(1)
}

func (ts *TwoStack) Normal() {
	ts.Mode = true
}

func (ts *TwoStack) Reverse() {
	ts.Mode = false
}

func (ts *TwoStack) Set(data interface{}) {
	if ts.R.Len() == 0 {
		ts.R.PushBack(deque.New(0, minCap))
	}
	if ts.Mode == true {
		ts.R.Back().(*deque.Deque).PushBack(data)
	} else {
		ts.R.Front().(*deque.Deque).PushFront(data)
	}
}

func (ts *TwoStack) Get() (interface{}, error) {
	if ts.R.Len() == 0 {
		return nil, fmt.Errorf("Stack is to shallow for .Get() operation")
	}
	if ts.Mode == true {
		if ts.R.Back().(*deque.Deque).Len() == 0 {
			return nil, fmt.Errorf("Stack(Cell) is to shallow for .Take() operation")
		}
		return ts.R.Back().(*deque.Deque).Back(), nil
	} else {
		if ts.R.Front().(*deque.Deque).Len() == 0 {
			return nil, fmt.Errorf("Stack(Cell) is to shallow for .Take() operation")
		}
		return ts.R.Front().(*deque.Deque).Front(), nil
	}
}

func (ts *TwoStack) Take() (interface{}, error) {
	if ts.R.Len() == 0 {
		return nil, fmt.Errorf("Stack is to shallow for .Take() operation")
	}
	if ts.Mode == true {
		if ts.R.Back().(*deque.Deque).Len() == 0 {
			return nil, fmt.Errorf("Stack(Cell) is to shallow for .Take() operation")
		}
		return ts.R.Back().(*deque.Deque).PopBack(), nil
	} else {
		if ts.R.Front().(*deque.Deque).Len() == 0 {
			return nil, fmt.Errorf("Stack(Cell) is to shallow for .Take() operation")
		}
		return ts.R.Front().(*deque.Deque).PopFront(), nil
	}
}

func (ts *TwoStack) Add() {
	if ts.Mode == true {
		ts.R.PushBack(deque.New(0, minCap))
	} else {
		ts.R.PushFront(deque.New(0, minCap))
	}
}

func (ts *TwoStack) addQ(q *deque.Deque) {
	if ts.Mode == true {
		ts.R.PushBack(q)
	} else {
		ts.R.PushFront(q)
	}
}

func (ts *TwoStack) Del() {
	if ts.R.Len() > 0 {
		if ts.Mode == true {
			ts.R.PopBack()
		} else {
			ts.R.PopFront()
		}
	}
}

func (ts *TwoStack) GLen() int {
	return ts.R.Len()
}

func (ts *TwoStack) Len() int {
	if ts.R.Len() > 0 {
		if ts.Mode == true {
			return ts.R.Back().(*deque.Deque).Len()
		} else {
			return ts.R.Front().(*deque.Deque).Len()
		}
	}
	return 0
}

func (ts *TwoStack) Q() *deque.Deque {
	if ts.R.Len() > 0 {
		if ts.Mode == true {
			return ts.R.Back().(*deque.Deque)
		} else {
			return ts.R.Front().(*deque.Deque)
		}
	}
	return nil
}
