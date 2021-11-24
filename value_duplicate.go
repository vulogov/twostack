package twostack

func (e *Elem) Duplicate() *Elem {
	var res *Elem
	switch e.Type {
	case Bool_t:
		res = Bool(e.Name, e.V.(bool))
	case Int_t:
		res = Int64(e.Name, e.V.(int64))
	case Float_t:
		res = Float(e.Name, e.V.(float64))
	case String_t:
		res = String(e.Name, e.V.(string)[0:len(e.V.(string))])
	case None_t:
		res = None(e.Name)
	}
	for v := range e.Labels.Iterator().C {
		res.Labels.Add(v)
	}
	return res
}
