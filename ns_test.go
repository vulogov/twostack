package twostack

import (
	"testing"
)

func TestNS(t *testing.T) {
	ns, err := NewNS("abc")
	if err != nil {
		t.Error("#1 .NewNS() had failed")
	}
	ns1, err := ns.NewNS("cde")
	if err != nil {
		t.Error("#2 .NewNS() had failed")
	}
	if ns1.RootNS.ID != ns.ID {
		t.Error("#3 .NewNS() had failed")
	}
}
