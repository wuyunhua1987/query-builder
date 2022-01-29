package builder

import (
	"bytes"
	"fmt"
)

type EqualBuilder struct {
	col string
	val interface{}
}

type LikeBuilder struct {
	col    string
	val    string
	prefix bool
	suffix bool
}

type NotEqualBuilder struct {
	col string
	val interface{}
}

type InBuilder struct {
	col string
	val interface{}
}

type NotInBuilder struct {
	col string
	val interface{}
}

type NULLBuilder struct {
	col string
}

type NotNULLBuilder struct {
	col string
}

func (b *EqualBuilder) Parse() (string, interface{}) {
	return fmt.Sprintf("`%s` = ?", b.col), b.val
}

func (b *LikeBuilder) Parse() (string, interface{}) {
	buf := new(bytes.Buffer)
	if b.prefix {
		buf.WriteString("%")
	}
	buf.WriteString(b.val)
	if b.suffix {
		buf.WriteString("%")
	}
	return fmt.Sprintf("`%s` like ?", b.col), buf.String()
}

func (b *NotEqualBuilder) Parse() (string, interface{}) {
	return fmt.Sprintf("`%s` <> ?", b.col), b.val
}

func (b *InBuilder) Parse() (string, interface{}) {
	return fmt.Sprintf("`%s` IN (?)", b.col), b.val
}

func (b *NotInBuilder) Parse() (string, interface{}) {
	return fmt.Sprintf("`%s` NOT IN (?)", b.col), b.val
}

func (b *NULLBuilder) Parse() (string, interface{}) {
	return fmt.Sprintf("`%s` IS NULL", b.col), nil
}

func (b *NotNULLBuilder) Parse() (string, interface{}) {
	return fmt.Sprintf("`%s` IS NOT NULL", b.col), nil
}

func (*Builder) Equal(col string, val interface{}) QueryBuilder {
	return &EqualBuilder{col, val}
}

func (*Builder) Like(col string, val string, prefix, suffix bool) QueryBuilder {
	return &LikeBuilder{col, val, prefix, suffix}
}

func (*Builder) NotEqual(col string, val interface{}) QueryBuilder {
	return &NotEqualBuilder{col, val}
}

func (*Builder) In(col string, val interface{}) QueryBuilder {
	return &InBuilder{col, val}
}

func (*Builder) NotIn(col string, val interface{}) QueryBuilder {
	return &NotEqualBuilder{col, val}
}

func (*Builder) NULL(col string) QueryBuilder {
	return &NULLBuilder{col}
}

func (*Builder) NotNULL(col string) QueryBuilder {
	return &NotNULLBuilder{col}
}
