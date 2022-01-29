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

func (*AndBuilder) Operate() {}
func (*OrBuilder) Operate()  {}

func (b *AndBuilder) Parse() (string, interface{}) {
	cols := []string{}
	vals := []interface{}{}
	for _, v := range b.ops {
		switch t := v.(type) {
		case *AndBuilder:
			col, val := t.Parse()
			cols = append(cols, col)
			valslice := val.([]interface{})
			vals = append(vals, valslice...)
		case *OrBuilder:
			col, val := t.Parse()
			cols = append(cols, col)
			valslice := val.([]interface{})
			vals = append(vals, valslice...)
		default:
			col, val := t.Parse()
			cols = append(cols, col)
			if val != nil {
				vals = append(vals, val)
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
			valslice := val.([]interface{})
			vals = append(vals, valslice...)
		case *AndBuilder:
			col, val := t.Parse()
			cols = append(cols, col)
			valslice := val.([]interface{})
			vals = append(vals, valslice...)
		default:
			col, val := t.Parse()
			cols = append(cols, col)
			if val != nil {
				vals = append(vals, val)
			}
		}
	}
	and := strings.Join(cols, " OR ")
	return fmt.Sprintf("(%s)", and), vals
}
