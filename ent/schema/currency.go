package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Currency holds the schema definition for the Currency entity.
type Currency struct {
	ent.Schema
}

// Fields of the Currency.
func (Currency) Fields() []ent.Field {
	return []ent.Field{
		field.String("symbol").Unique(),
		field.String("name").Unique(),
	}
}

// Edges of the Currency.
func (Currency) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("accounts", Account.Type),
		edge.To("portfolios", Portfolio.Type),
	}
}

func (Currency) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("symbol"),
	}
}
