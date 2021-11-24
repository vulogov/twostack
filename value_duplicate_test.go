package twostack

import "testing"

func TestDupBool(t *testing.T) {
	e := Bool("boolean", true).Duplicate()
	if e.V.(bool) != true {
		t.Errorf("Can not duplicate bool element")
	}
}

func TestDupInt(t *testing.T) {
	e := Int64("answer", int64(42)).Duplicate()
	if e.V.(int64) != 42 {
		t.Errorf("Can not duplicate int64 element")
	}
}

func TestDupFloat(t *testing.T) {
	e := Float("pi", float64(3.14)).Duplicate()
	if e.V.(float64) != 3.14 {
		t.Errorf("Can not duplicate float64 element")
	}
}

func TestDupString(t *testing.T) {
	e := String("answer", "42").Duplicate()
	if e.V.(string) != "42" {
		t.Errorf("Can not duplicate string element")
	}
}
