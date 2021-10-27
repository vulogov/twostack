package twostack

func (ts *TwoStack) setTmp(name string, v interface{}) {
	ts.tmp.Store(name, v)
}

func (ts *TwoStack) getTmp(name string) interface{} {
	if v, ok := ts.tmp.Load(name); ok {
		return v
	}
	return nil
}

func (ts *TwoStack) delTmp(name string) {
	ts.tmp.Delete(name)
}
