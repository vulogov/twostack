package twostack

import "testing"

func TestCall(t *testing.T) {
	ts := Init()
	ts.SetCall("Hello/World")
	res, err := ts.GetCall()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if res != "Hello/World" {
		t.Error("ts.SetCall()/ts.GetCall() had failed")
	}
}
