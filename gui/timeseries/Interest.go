package timeseries

import (
	"math"
	"time"

	"github.com/shopspring/decimal"
)

type Interest struct {
	Series TimeSeries
	APR    decimal.Decimal
	From   time.Time

	baseTimeSeries
}

func (ts *Interest) valueMultiplier(timeDiff time.Duration) decimal.Decimal {
	aprF := ts.APR.InexactFloat64()
	timeInYears := timeDiff.Hours() / (24.0 * 360.0)
	fullYears, restYear := math.Modf(timeInYears)
	aprYear := (aprF / 100.0) + 1.0
	fullValue := math.Pow(aprYear, fullYears)
	aprDay := ((aprF / 100.0) / 360.00) + 1.0
	partialValue := math.Pow(aprDay, restYear*360.0)
	return decimal.NewFromFloat(fullValue * partialValue)
}

func (ts *Interest) ValueRange(until time.Time) (decimal.Decimal, decimal.Decimal) {
	minValue, maxValue := ts.baseTimeSeries.ValueRange(until)
	if len(ts.values) == 0 {
		return minValue, maxValue
	}
	lastValue := ts.values[len(ts.values)-1]
	if lastValue.Date.Before(until) {
		futureValue := lastValue.Value.Mul(ts.valueMultiplier(until.Sub(lastValue.Date)))
		minValue = decimal.Min(minValue, futureValue)
		maxValue = decimal.Max(maxValue, futureValue)
	}
	return minValue, maxValue
}

func (ts *Interest) GetValue(at time.Time) decimal.Decimal {
	lastValue := ts.values[len(ts.values)-1]
	if lastValue.Date.Equal(at) {
		return lastValue.Value
	}
	if lastValue.Date.After(at) {
		valueIndex := ts.baseTimeSeries.GetValueIndex(at)
		if valueIndex == -1 {
			return decimal.Decimal{}
		}
		v := ts.values[valueIndex]
		return v.Value.Mul(ts.valueMultiplier(at.Sub(v.Date)))
	}
	for lastValue.Date.Before(at) {
		baseValue := ts.Series.GetValue(lastValue.Date)
		nextDate := ts.Series.GetNext(lastValue.Date).Date
		if nextDate.Equal(lastValue.Date) {
			nextDate = at
		}
		lastValue.Value = baseValue.Mul(ts.valueMultiplier(nextDate.Sub(lastValue.Date))).Sub(baseValue)
		lastValue.Date = nextDate
		ts.values = append(ts.values, lastValue)
	}
	return lastValue.Value
}

func (ts *Interest) Rebuild() {
	ts.Series.Rebuild()
	ts.values = make([]TimeSeriesValue, 0)
	ts.values = append(ts.values, TimeSeriesValue{Value: decimal.Decimal{}, Date: time.Now()})
}

func (ts *Interest) IsDynamic() bool {
	return true
}
