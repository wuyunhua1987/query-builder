package builder

type OP interface {
	Operate()
}

type QueryBuilder interface {
	Parse() (string, interface{})
}

type Builder struct {
	op OP
}

func New() *Builder {
	return &Builder{}
}

func (b *Builder) And(query ...QueryBuilder) QueryBuilder {
	and := &AndBuilder{query}
	b.op = and
	return and
}

func (b *Builder) Or(query ...QueryBuilder) QueryBuilder {
	or := &OrBuilder{query}
	b.op = or
	return or
}

func (b *Builder) Parse() (string, []interface{}) {
	switch t := b.op.(type) {
	case *AndBuilder:
		q, arg := t.Parse()
		return q[1 : len(q)-1], arg.([]interface{})
	case *OrBuilder:
		q, arg := t.Parse()
		return q[1 : len(q)-1], arg.([]interface{})
	default:
		return "", nil
	}
}
