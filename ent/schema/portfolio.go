package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Portfolio holds the schema definition for the Portfolio entity.
type Portfolio struct {
	ent.Schema
}

// Fields of the Portfolio.
func (Portfolio) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
	}
}

// Edges of the Portfolio.
func (Portfolio) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("currency", Currency.Type).Unique(),
		edge.To("referenceAccount", Account.Type).Unique(),
		edge.From("incomingTransactions", Transaction.Type).Ref("toPortfolio"),
		edge.From("outgoingTransactions", Transaction.Type).Ref("fromPortfolio"),
		edge.From("incomingRecurringTransactions", RecurringTransaction.Type).Ref("toPortfolio"),
		edge.From("outgoingRecurringTransactions", RecurringTransaction.Type).Ref("fromPortfolio"),
	}
}
