package twostack

import (
	"math"
	"testing"
)

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

func TestFloat1(t *testing.T) {
	ts := Init()
	ts.MakeFloat("3.14", "pi")
	res := ts.Float()
	if res != 3.14 {
		t.Error("#1 ts.Float() had failed")
	}
}

func TestFloat2(t *testing.T) {
	ts := Init()
	ts.MakeInt("42", "answer")
	res := ts.Float()
	if res != 42.0 {
		t.Error("#2 ts.Float() had failed")
	}
}

func TestFloat3(t *testing.T) {
	ts := Init()
	ts.MakeString("42", "answer")
	res := ts.Float()
	if res != 42.0 {
		t.Error("#3 ts.Float() had failed")
	}
}

func TestFloat4(t *testing.T) {
	ts := Init()
	ts.MakeBool("true", "answer")
	res := ts.Float()
	if !math.IsNaN(res) {
		t.Error("#4 ts.Float() had failed")
	}
}
