package twostack

import "testing"

func TestApplyGen(t *testing.T) {
	ts := Init()
	_, err := ts.ApplyGen(Answer)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	res, _ := ts.G()
	if res != int64(42) {
		t.Error("ts.ApplyGen() had failed")
	}
}

func TestApplyFun(t *testing.T) {
	ts := Init()
	ts.Put(42)
	_, err := ts.ApplyFun(NumericIncrease)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	res, _ := ts.G()
	if res != 43.0 {
		t.Errorf("#1 ts.ApplyFun() had failed: %v", res)
	}
}

func TestApplyFunFail(t *testing.T) {
	ts := Init()
	_, err := ts.ApplyFun(NumericIncrease)
	if err == nil {
		t.Errorf("#2 ts.ApplyFun() had failed")
	}
}

func TestApplyOp(t *testing.T) {
	ts := Init()
	ts.Put(42)
	ts.Put(42)
	_, err := ts.ApplyOp(NumericAdd)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	res, _ := ts.G()
	if res != 84.0 {
		t.Errorf("#1 ts.ApplyOp() had failed: %v", res)
	}
}

func TestApplyOpAll(t *testing.T) {
	ts := Init()
	ts.Put(1)
	ts.Put(2)
	ts.Put(3)
	_, err := ts.ApplyOpAll(NumericAdd)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	res, _ := ts.G()
	if res != 6.0 {
		t.Errorf("#1 ts.ApplyOpAll() had failed: %v", res)
	}
}
