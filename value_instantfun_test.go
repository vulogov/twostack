package twostack

import "testing"

func TestIF(t *testing.T) {
	ts := Init()
	ts.Put(42)
	res, _ := ts.G()
	if res != int64(42) {
		t.Errorf("#1 ts.SetIF() had failed: %v", res)
	}
	ts.SetIF(NumericIncreaseIfNum)
	ts.Put(42)
	res, _ = ts.G()
	if res != 43.0 {
		t.Errorf("#2 ts.SetIF() had failed: %v", res)
	}
	ts.StopIF()
	ts.Put(42)
	res, _ = ts.G()
	if res != int64(42) {
		t.Errorf("#3 ts.SetIF() had failed: %v", res)
	}
}
