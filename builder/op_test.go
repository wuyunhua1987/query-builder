package builder

import (
	"reflect"
	"testing"
)

func TestAndBuilder_Parse(t *testing.T) {
	type fields struct {
		ops []QueryBuilder
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		want1  interface{}
	}{
		// TODO: Add test cases.
		{
			"single",
			fields{
				[]QueryBuilder{
					&EqualBuilder{"id", 1},
				},
			},
			"(`id` = ?)",
			[]interface{}{1},
		},
		{
			"multi",
			fields{
				[]QueryBuilder{
					&EqualBuilder{"id", 1},
					&EqualBuilder{"name", "a"},
				},
			},
			"(`id` = ? AND `name` = ?)",
			[]interface{}{1, "a"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &AndBuilder{
				ops: tt.fields.ops,
			}
			got, got1 := b.Parse()
			if got != tt.want {
				t.Errorf("AndBuilder.Parse() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("AndBuilder.Parse() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOrBuilder_Parse(t *testing.T) {
	type fields struct {
		ops []QueryBuilder
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		want1  interface{}
	}{
		// TODO: Add test cases.
		{
			"single",
			fields{
				[]QueryBuilder{
					&EqualBuilder{"id", 1},
				},
			},
			"(`id` = ?)",
			[]interface{}{1},
		},
		{
			"multi",
			fields{
				[]QueryBuilder{
					&EqualBuilder{"id", 1},
					&EqualBuilder{"name", "a"},
				},
			},
			"(`id` = ? OR `name` = ?)",
			[]interface{}{1, "a"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &OrBuilder{
				ops: tt.fields.ops,
			}
			got, got1 := b.Parse()
			if got != tt.want {
				t.Errorf("OrBuilder.Parse() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrBuilder.Parse() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
