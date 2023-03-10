package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/shopspring/decimal"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
		field.String("expectedAPR").
			GoType(decimal.Decimal{}).
			DefaultFunc(func() decimal.Decimal { return decimal.Decimal{} }),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("currency", Currency.Type).Unique(),
		edge.From("incomingTransactions", Transaction.Type).Ref("toAccount"),
		edge.From("outgoingTransactions", Transaction.Type).Ref("fromAccount"),
		edge.From("incomingRecurringTransactions", RecurringTransaction.Type).Ref("toAccount"),
		edge.From("outgoingRecurringTransactions", RecurringTransaction.Type).Ref("fromAccount"),
	}
}
