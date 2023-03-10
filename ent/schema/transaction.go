package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/shopspring/decimal"
)

// Transaction holds the schema definition for the Transaction entity.
type Transaction struct {
	ent.Schema
}

// Fields of the Transaction.
func (Transaction) Fields() []ent.Field {
	return []ent.Field{
		field.String("securityAmount").
			GoType(decimal.Decimal{}).
			DefaultFunc(func() decimal.Decimal { return decimal.Decimal{} }),
		field.String("currencyValue").
			GoType(decimal.Decimal{}).
			DefaultFunc(func() decimal.Decimal { return decimal.Decimal{} }),
		field.Time("date"),
	}
}

// Edges of the Transaction.
func (Transaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("toPortfolio", Portfolio.Type).Unique(),
		edge.To("fromPortfolio", Portfolio.Type).Unique(),
		edge.To("toAccount", Account.Type).Unique(),
		edge.To("fromAccount", Account.Type).Unique(),
		edge.To("security", Security.Type).Unique(),
	}
}
