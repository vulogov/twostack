package twostack

import "testing"

func TestInt(t *testing.T) {
	ts := Init()
	ts.SetInt(42, "answer")
	res, err := ts.GetInt()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if res != 42 {
		t.Error("ts.SetInt()/ts.GetInt() had failed")
	}
}

func TestIntMake(t *testing.T) {
	ts := Init()
	ts.MakeInt("42", "answer")
	res, err := ts.GetInt()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if res != 42 {
		t.Error("ts.SetInt()/ts.GetInt() had failed")
	}
}
