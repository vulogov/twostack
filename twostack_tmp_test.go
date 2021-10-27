package twostack

import "testing"

func TestTmpSet(t *testing.T) {
	ts := Init()
	ts.setTmp("a", "b")
	if ts.getTmp("a").(string) != "b" {
		t.Error("ts.setTmp()/ts.getTmp() had failed")
	}
}

func TestTmpDel(t *testing.T) {
	ts := Init()
	ts.setTmp("a", "b")
	ts.delTmp("a")
	if ts.getTmp("a") != nil {
		t.Error("ts.delTmp() had failed")
	}
}
