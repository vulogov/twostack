package twostack

func None(name string, labels ...string) *Elem {
	return NewElem(name, None_t, nil, labels...)
}

func MakeNone(name string, p string, labels ...string) *Elem {
	return NewElem(name, None_t, nil, labels...)
}
