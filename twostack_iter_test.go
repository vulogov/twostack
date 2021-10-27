package twostack

import (
	"fmt"
	"testing"
)

func TestCZip(t *testing.T) {
	ts := Init()
	ts.Put(1, "")
	ts.Put(2, "")
	ts.Put(3, "")
	ts.Add()
	ts.Put(4, "")
	ts.Put(5, "")
	ts.Put(6, "")
	ts.Add()
	ts.Put(7, "")
	ts.Put(8, "")
	ts.Put(9, "")
	ts.Add()
	ts.Put(10, "")
	ts.Put(11, "")
	ts.Put(12, "")
	ts.CZip()
	r := ts.String()
	fmt.Println(r)
	v, err := ts.G()
	if err != nil {
		t.Error("#1 ts.CZip() had failed")
	}
	if v != int64(12) {
		t.Error("#2 ts.CZip() had failed")
	}
}
