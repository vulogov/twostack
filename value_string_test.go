package twostack

import "testing"

func TestString(t *testing.T) {
	ts := Init()
	ts.SetString("42", "answer")
	res, err := ts.GetString()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if res != "42" {
		t.Error("ts.SetString()/ts.GetString() had failed")
	}
}

func TestStringMake(t *testing.T) {
	ts := Init()
	ts.MakeString("Привет", "string")
	res, err := ts.GetString()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if res != "Привет" {
		t.Error("ts.SetString()/ts.GetString() had failed")
	}
}

func TestStringMakeWithEmptyName(t *testing.T) {
	ts := Init()
	ts.MakeString("Привет")
	res, err := ts.GetString()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if res != "Привет" {
		t.Error("ts.SetString()/ts.GetString() had failed")
	}
}
