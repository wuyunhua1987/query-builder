package builder

import (
	"fmt"
	"strings"
)

type AndBuilder struct {
	ops []QueryBuilder
}

type OrBuilder struct {
	ops []QueryBuilder
}

type SingleBuilder struct {
	opt QueryBuilder
}

func (*AndBuilder) Operate()    {}
func (*OrBuilder) Operate()     {}
func (*SingleBuilder) Operate() {}

func (b *AndBuilder) Parse() (string, interface{}) {
	cols := []string{}
	vals := []interface{}{}
	for _, v := range b.ops {
		switch t := v.(type) {
		case *AndBuilder:
			col, val := t.Parse()
			cols = append(cols, col)
			if valslice, ok := val.([]interface{}); ok {
				vals = append(vals, valslice...)
			}
		case *OrBuilder:
			col, val := t.Parse()
			cols = append(cols, col)
			if valslice, ok := val.([]interface{}); ok {
				vals = append(vals, valslice...)
			}
		default:
			col, val := t.Parse()
			cols = append(cols, col)
			if val != nil {
				if valslice, ok := val.([]interface{}); ok {
					vals = append(vals, valslice...)
				} else {
					vals = append(vals, val)
				}
			}
		}
	}
	and := strings.Join(cols, " AND ")
	return fmt.Sprintf("(%s)", and), vals
}

func (b *OrBuilder) Parse() (string, interface{}) {
	cols := []string{}
	vals := []interface{}{}
	for _, v := range b.ops {
		switch t := v.(type) {
		case *OrBuilder:
			col, val := t.Parse()
			cols = append(cols, col)
			if valslice, ok := val.([]interface{}); ok {
				vals = append(vals, valslice...)
			}
		case *AndBuilder:
			col, val := t.Parse()
			cols = append(cols, col)
			if valslice, ok := val.([]interface{}); ok {
				vals = append(vals, valslice...)
			}
		default:
			col, val := t.Parse()
			cols = append(cols, col)
			if val != nil {
				if valslice, ok := val.([]interface{}); ok {
					vals = append(vals, valslice...)
				} else {
					vals = append(vals, val)
				}
			}
		}
	}
	or := strings.Join(cols, " OR ")
	return fmt.Sprintf("(%s)", or), vals
}

func (b *SingleBuilder) Parse() (string, interface{}) {
	switch t := b.opt.(type) {
	case *OrBuilder:
		q, arg := t.Parse()
		return q[1 : len(q)-1], arg.([]interface{})
	case *AndBuilder:
		q, arg := t.Parse()
		return q[1 : len(q)-1], arg.([]interface{})
	case *SingleBuilder:
		return t.Parse()
	default:
		vals := []interface{}{}
		col, val := t.Parse()
		if val != nil {
			if valslice, ok := val.([]interface{}); ok {
				vals = append(vals, valslice...)
			} else {
				vals = append(vals, val)
			}
		}
		return col, vals
	}
}
