package builder

import "testing"

func TestBuilder(t *testing.T) {
	b := New()
	b.And(
		b.Equal("id", 1),
		b.Like("name", "a", true, true),
		b.In("status", 1, 2),
		b.NULL("delete"),
		b.Or(
			b.NotEqual("email", "a@b.com"),
			b.NotIn("state", 2, 3),
			b.And(
				b.NotNULL("phone"),
				b.Equal("tel", "01"),
			),
			b.Equal("del", 0),
		),
	)
	cond, val := b.Parse()
	t.Log(cond)
	t.Log(val)
}

func TestSingleBuilder(t *testing.T) {
	b := New()
	b.Single(
		b.Equal("id", 1),
	)
	cond, val := b.Parse()
	t.Log(cond)
	t.Log(val)
}
