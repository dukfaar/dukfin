package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/shopspring/decimal"
)

// Security holds the schema definition for the Security entity.
type Security struct {
	ent.Schema
}

// Fields of the Security.
func (Security) Fields() []ent.Field {
	return []ent.Field{
		field.String("symbol").Unique(),
		field.String("name").Unique(),
		field.String("expectedAPR").
			GoType(decimal.Decimal{}).
			DefaultFunc(func() decimal.Decimal { return decimal.Decimal{} }),
	}
}

// Edges of the Security.
func (Security) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("currency", Currency.Type).Unique(),
		edge.From("transactions", Transaction.Type).Ref("security"),
		edge.From("prices", SecurityPrice.Type).Ref("security"),
	}
}
