package twostack

import "testing"

func TestBool(t *testing.T) {
	ts := Init()
	ts.SetBool(true, "boolean")
	res, err := ts.GetBool()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if res != true {
		t.Error("ts.SetBool()/ts.GetBool() had failed")
	}
}

func TestBoolMake(t *testing.T) {
	ts := Init()
	ts.MakeBool("True", "boolean")
	res, err := ts.GetBool()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if res != true {
		t.Error("ts.SetBool()/ts.GetBool() had failed")
	}
}
