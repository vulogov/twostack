package twostack

func (ts *TwoStack) SetIF(f FunFun) {
	ts.InstF = f
	ts.IsIF = true
}

func (ts *TwoStack) StopIF() {
	ts.InstF = Passthrough
	ts.IsIF = false
}

func (ts *TwoStack) SetIFilter(f FilterFun) {
	ts.InstFF = f
	ts.IsFF = true
}

func (ts *TwoStack) StopIFilter() {
	ts.InstFF = PassthroughFilter
	ts.IsFF = false
}
