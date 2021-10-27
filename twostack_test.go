package twostack

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	ts := Init()
	if ts.Status != true {
		t.Fatalf("TwoStack is failed to initialize")
	}
}

func TestSetGet(t *testing.T) {
	ts := Init()
	ts.Set("42")
	r, _ := ts.Get()
	if r != "42" {
		t.Error("ts.Set()/ts.Get() had failed")
	}
}

func TestTake(t *testing.T) {
	ts := Init()
	ts.Set("42")
	r, _ := ts.Take()
	if r != "42" {
		t.Error("#1 ts.Take() had failed")
	}
	if ts.Len() != 0 {
		t.Error("#2 ts.Take() had failed")
	}
}

func TestMode(t *testing.T) {
	ts := Init()
	ts.Set("42")
	ts.Set("41")
	ts.Reverse()
	r, _ := ts.Get()
	if r != "42" {
		t.Error("ts.Mode() had failed")
	}
	ts.Normal()
	r, _ = ts.Get()
	if r != "41" {
		t.Error("ts.Mode() had failed")
	}
}

func TestLeftRight(t *testing.T) {
	ts := Init()
	ts.Set("42")
	ts.Add()
	ts.Set("41")
	ts.Left()
	r, _ := ts.Get()
	if r != "42" {
		t.Errorf("#1 ts.Mode() had failed, as got %s", r)
	}
	ts.Right()
	r, _ = ts.Get()
	if r != "41" {
		t.Errorf("#2 ts.Mode() had failed, as got %s", r)
	}
}

func TestGLen(t *testing.T) {
	ts := Init()
	if ts.GLen() != 1 {
		t.Error("ts.GLen() had failed")
	}
}

func TestLen(t *testing.T) {
	ts := Init()
	ts.Set("42")
	if ts.Len() != 1 {
		t.Error("ts.Len() had failed")
	}
}

func TestQ(t *testing.T) {
	ts := Init()
	ts.Set("42")
	q := ts.Q()
	if q.Len() != 1 {
		t.Error("ts.Q() had failed")
	}
}

func TestDel(t *testing.T) {
	ts := Init()
	ts.Del()
	if ts.GLen() != 0 {
		t.Error("ts.Del() had failed")
	}
}

func TestPut1(t *testing.T) {
	ts := Init()
	if !ts.Put(42, "answer") {
		t.Error("#1 ts.Put(int) had failed")
	}
	r, _ := ts.G()
	if r != int64(42) {
		t.Errorf("#2 ts.Put(int) had failed: %v", r)
	}
}

func TestPut2(t *testing.T) {
	ts := Init()
	if !ts.Put("42", "answer") {
		t.Error("#1 ts.Put(string) had failed")
	}
	r, _ := ts.T()
	if r != "42" {
		t.Errorf("#2 ts.Put(string) had failed: %v", r)
	}
}

func TestPut3(t *testing.T) {
	ts := Init()
	if !ts.Put(true, "boolean") {
		t.Error("#1 ts.Put(bool) had failed")
	}
	r, _ := ts.G()
	if r != true {
		t.Errorf("#2 ts.Put(bool) had failed: %v", r)
	}
}

func TestPut4(t *testing.T) {
	ts := Init()
	if !ts.Put(3.14, "pi") {
		t.Error("#1 ts.Put(float) had failed")
	}
	r, _ := ts.G()
	if r != 3.14 {
		t.Errorf("#2 ts.Put(float) had failed: %v", r)
	}
}

func TestToString(t *testing.T) {
	ts := Init()
	ts.Put(1, "")
	ts.Put(2, "")
	ts.Add()
	ts.Put(3, "")
	ts.Put(4, "")
	r := ts.String()
	fmt.Println(r)
	if len(r) == 0 {
		t.Errorf("#2 ts.String() had failed")
	}
}
