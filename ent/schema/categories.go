package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Categories holds the schema definition for the Categories entity.
type Categories struct {
	ent.Schema
}

// Fields of the Categories.
func (Categories) Fields() []ent.Field {
	return []ent.Field{
		field.Int("parent_id").
			Default(0),
		field.String("name").
			Default("unknown"),
		field.String("system_name").
			Default("unknown"),
		field.String("description").
			Default("unknown"),
	}
}

// Edges of the Categories.
func (Categories) Edges() []ent.Edge {
	return nil
}
