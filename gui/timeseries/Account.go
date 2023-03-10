package timeseries

import (
	"context"
	"sort"
	"time"

	"github.com/dukfaar/dukfin/ent"
	"github.com/dukfaar/dukfin/ent/transaction"
	"github.com/shopspring/decimal"
)

type Account struct {
	Account *ent.Account

	baseTimeSeries
}

func NewAccount(account *ent.Account) *Account {
	result := &Account{Account: account}
	return result
}

func (ts *Account) GetAllAccountTransactionsInOrder(until time.Time) []*ent.Transaction {
	incomingTxs, err := ts.Account.QueryIncomingTransactions().Where(transaction.DateLTE(until)).All(context.Background())
	if err != nil {
		return nil
	}
	outgoingTxs, err := ts.Account.QueryOutgoingTransactions().Where(transaction.DateLTE(until)).All(context.Background())
	if err != nil {
		return nil
	}
	allTxs := append(incomingTxs, outgoingTxs...)
	sort.Slice(allTxs, func(i, j int) bool {
		return allTxs[i].Date.Before(allTxs[j].Date)
	})
	return allTxs
}

func (ts *Account) Rebuild() {
	allTxs := ts.GetAllAccountTransactionsInOrder(time.Now())
	ts.values = make([]TimeSeriesValue, 0, len(allTxs))
	currentValue := decimal.Decimal{}
	for _, tx := range allTxs {
		from, _ := tx.QueryFromAccount().First(context.Background())
		to, _ := tx.QueryToAccount().First(context.Background())
		if to != nil && to.ID == ts.Account.ID {
			currentValue = currentValue.Add(tx.CurrencyValue)
		} else if from != nil && from.ID == ts.Account.ID {
			currentValue = currentValue.Sub(tx.CurrencyValue)
		}
		ts.values = append(ts.values, TimeSeriesValue{Value: currentValue, Date: tx.Date})
	}
}

func (ts *Account) GetAccountValueOrder(until time.Time) []TimeSeriesValue {
	lastIndex := sort.Search(len(ts.values), func(i int) bool { return ts.values[i].Date.After(until) })
	return ts.values[0:lastIndex]
}

func (ts *Account) IsDynamic() bool {
	return true
}
