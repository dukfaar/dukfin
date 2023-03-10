package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/shopspring/decimal"
)

// Transaction holds the schema definition for the Transaction entity.
type SecurityPrice struct {
	ent.Schema
}

// Fields of the Transaction.
func (SecurityPrice) Fields() []ent.Field {
	return []ent.Field{
		field.String("value").
			GoType(decimal.Decimal{}).
			DefaultFunc(func() decimal.Decimal { return decimal.Decimal{} }),
		field.Time("date"),
	}
}

// Edges of the Transaction.
func (SecurityPrice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("security", Security.Type).Unique(),
	}
}
