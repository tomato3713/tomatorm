package tomatorm

func newSelectBuilder() *SelectBuilder {
	sb := &SelectBuilder{}
	return sb
}

func NewSelectBuilder() *SelectBuilder {
	builder := newSelectBuilder()
	return builder
}
