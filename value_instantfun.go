package twostack

func (ts *TwoStack) SetIF(f FunFun) {
	ts.InstF = f
	ts.IsIF = true
}

func (ts *TwoStack) StopIF() {
	ts.InstF = Passthrough
	ts.IsIF = false
}
