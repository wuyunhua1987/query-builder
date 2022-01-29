package builder

type OP interface {
	Operate()
}

type QueryBuilder interface {
	Parse() (string, interface{})
}

type Builder struct{}

func New() *Builder {
	return &Builder{}
}

func (b *Builder) And(query ...QueryBuilder) QueryBuilder {
	return &AndBuilder{query}
}

func (b *Builder) Or(query ...QueryBuilder) QueryBuilder {
	return &OrBuilder{query}
}
