package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Greeter holds the schema definition for the Greeter entity.
type Greeter struct {
	ent.Schema
}

// Fields of the Greeter.
func (Greeter) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("name").Default("unknown"),
		field.String("description"),
		field.Time("created_at").Default(time.Now().Local),
		field.Time("updated_at").Default(time.Now().Local),
	}
}

// Edges of the Greeter.
func (Greeter) Edges() []ent.Edge {
	return nil
}
