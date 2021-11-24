package twostack

import (
	"fmt"

	"github.com/gammazero/deque"
	"github.com/google/uuid"
	"github.com/lrita/cmap"
)

type NS struct {
	Name    string
	ID      string
	IsTemp  bool
	IsRoot  bool
	TS      *TwoStack
	RootNS  *NS
	NSstack deque.Deque
	NScat   cmap.Cmap
	EFun    cmap.Cmap
	GFun    cmap.Cmap
	FFun    cmap.Cmap
	OFun    cmap.Cmap
	IFun    cmap.Cmap
	IFFun   cmap.Cmap
}

func NewNS(name string) (*NS, error) {
	ns := NS{
		Name:    name,
		ID:      uuid.New().String(),
		TS:      Init(),
		IsTemp:  false,
		IsRoot:  true,
		NSstack: *deque.New(0, minCap),
	}
	ns.RootNS = &ns
	return &ns, nil
}

func (ns *NS) NewNS(name string) (*NS, error) {
	nns := NS{
		Name:    name,
		ID:      uuid.New().String(),
		TS:      Init(),
		IsRoot:  false,
		NSstack: *deque.New(0, minCap),
	}
	nns.RootNS = ns.RootNS
	ns.RootNS.NScat.Store(nns.ID, &nns)
	ns.RootNS.NSstack.PushBack(&nns)
	return ns, nil
}

func (ns *NS) NewTempNS() (*NS, error) {
	nns, err := ns.NewNS("")
	if err != nil {
		return nil, err
	}
	nns.Name = nns.ID
	nns.IsTemp = true
	return nns, nil
}

func (ns *NS) Current() (*NS, error) {
	if ns.RootNS.NSstack.Len() == 0 {
		return nil, fmt.Errorf("NS stack is empty")
	}
	return ns.RootNS.NSstack.Back().(*NS), nil
}

func (ns *NS) EndNS() error {
	if ns.RootNS.NSstack.Len() == 0 {
		return fmt.Errorf("NS stack is empty")
	}
	nns := ns.RootNS.NSstack.Back().(*NS)
	if nns.ID != ns.ID {
		return fmt.Errorf("You are trying to end NS, which is not current")
	}
	if nns.IsTemp == true {
		ns.RootNS.NScat.Delete(nns.ID)
	}
	ns.RootNS.NSstack.PopBack()
	return nil
}

func (ns *NS) Put(v interface{}, labels ...string) error {
	if !ns.TS.Put(v, labels...) {
		return fmt.Errorf(".Put() operation is failed for NS: %v", ns.Name)
	}
	return nil
}

func (ns *NS) Get() (interface{}, error) {
	res, err := ns.TS.G()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (ns *NS) Take() (interface{}, error) {
	res, err := ns.TS.T()
	if err != nil {
		return nil, err
	}
	return res, nil
}
