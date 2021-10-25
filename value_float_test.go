package twostack

import "testing"

func TestFloat(t *testing.T) {
	ts := Init()
	ts.SetFloat(3.14, "pi")
	res, err := ts.GetFloat()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if res != 3.14 {
		t.Error("ts.SetFloat()/ts.GetFloat() had failed")
	}
}

func TestFloatMake(t *testing.T) {
	ts := Init()
	ts.MakeFloat("3.14", "pi")
	res, err := ts.GetFloat()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if res != 3.14 {
		t.Error("ts.SetFloat()/ts.GetFloat() had failed")
	}
}
