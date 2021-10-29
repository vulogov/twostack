package twostack

import (
	"fmt"

	mapset "github.com/deckarep/golang-set"
	"github.com/lrita/cmap"
)

func GiveAnswer(kv *cmap.Cmap, p *mapset.Set) (*Elem, error) {
	if r, ok := kv.Load("answer"); ok {
		return r.(*Elem), nil
	}
	return nil, fmt.Errorf("No answer found")
}
