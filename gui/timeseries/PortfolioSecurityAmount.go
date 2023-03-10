package timeseries

import (
	"context"
	"sort"
	"time"

	"github.com/dukfaar/dukfin/ent"
	"github.com/dukfaar/dukfin/ent/security"
	"github.com/dukfaar/dukfin/ent/transaction"
	"github.com/shopspring/decimal"
)

type PortfolioSecurityAmount struct {
	Portfolio *ent.Portfolio
	Security  *ent.Security

	baseTimeSeries
}

func NewPortfolioSecurityAmount(portfolio *ent.Portfolio, security *ent.Security) *PortfolioSecurityAmount {
	result := &PortfolioSecurityAmount{Portfolio: portfolio, Security: security}
	return result
}

func (ts *PortfolioSecurityAmount) GetAllTransactionsInOrder(until time.Time) []*ent.Transaction {
	incomingTxs, err := ts.Portfolio.QueryIncomingTransactions().Where(transaction.And(transaction.DateLTE(until), transaction.HasSecurityWith(security.ID(ts.Security.ID)))).All(context.Background())
	if err != nil {
		return nil
	}
	outgoingTxs, err := ts.Portfolio.QueryOutgoingTransactions().Where(transaction.And(transaction.DateLTE(until), transaction.HasSecurityWith(security.ID(ts.Security.ID)))).All(context.Background())
	if err != nil {
		return nil
	}
	allTxs := append(incomingTxs, outgoingTxs...)
	sort.Slice(allTxs, func(i, j int) bool {
		return allTxs[i].Date.Before(allTxs[j].Date)
	})
	return allTxs
}

func (ts *PortfolioSecurityAmount) Rebuild() {
	allTxs := ts.GetAllTransactionsInOrder(time.Now())
	ts.values = make([]TimeSeriesValue, 0, len(allTxs))
	currentValue := decimal.Decimal{}
	for _, tx := range allTxs {
		from, _ := tx.QueryFromPortfolio().First(context.Background())
		to, _ := tx.QueryToPortfolio().First(context.Background())
		if to != nil && to.ID == ts.Portfolio.ID {
			currentValue = currentValue.Add(tx.SecurityAmount)
		} else if from != nil && from.ID == ts.Portfolio.ID {
			currentValue = currentValue.Sub(tx.SecurityAmount)
		}
		ts.values = append(ts.values, TimeSeriesValue{Value: currentValue, Date: tx.Date})
	}
}

func (ts *PortfolioSecurityAmount) IsDynamic() bool {
	return false
}
