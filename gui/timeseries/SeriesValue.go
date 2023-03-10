package timeseries

import (
	"time"

	"github.com/shopspring/decimal"
)

type TimeSeriesValue struct {
	Value decimal.Decimal
	Date  time.Time
}
