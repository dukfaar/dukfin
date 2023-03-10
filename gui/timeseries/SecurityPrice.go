package timeseries

import (
	"context"

	"github.com/dukfaar/dukfin/ent"
)

type SecurityPrice struct {
	Security *ent.Security

	baseTimeSeries
}

func NewSecurityPrice(security *ent.Security) *SecurityPrice {
	result := &SecurityPrice{Security: security}
	return result
}

func (ts *SecurityPrice) Rebuild() {
	allPrices, _ := ts.Security.QueryPrices().All(context.Background())
	ts.values = make([]TimeSeriesValue, 0, len(allPrices))
	for _, price := range allPrices {
		ts.values = append(ts.values, TimeSeriesValue{Value: price.Value, Date: price.Date})
	}
}

func (ts *SecurityPrice) IsDynamic() bool {
	return false
}
