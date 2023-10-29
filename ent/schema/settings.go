package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Settings holds the schema definition for the Settings entity.
type Settings struct {
	ent.Schema
}

// Fields of the Settings.
func (Settings) Fields() []ent.Field {
	return []ent.Field{
		field.Int("category_id").
			Positive(),
		field.String("name").
			Default("unknown"),
		field.String("value").
			Default("unknown"),
		field.Int("uid").
			Positive(),
		field.Int("created").
			Default(0),
		field.Int("actuality").
			Default(0),
	}
}

// Edges of the Settings.
func (Settings) Edges() []ent.Edge {
	return nil
}
