package timeseries

import (
	"fmt"
	"math"
	"time"

	"github.com/shopspring/decimal"
)

type Interest struct {
	Series     TimeSeries
	APR        decimal.Decimal
	From       time.Time
	Resolution time.Duration

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

func (ts *Interest) getLastValue() TimeSeriesValue {
	return ts.values[len(ts.values)-1]
}

func (ts *Interest) calculateNextValue() {
	lastValue := ts.getLastValue()
	baseValue := ts.Series.GetValue(lastValue.Date)
	nextDate := ts.Series.GetNext(lastValue.Date).Date
	nextResolutionDate := lastValue.Date.Add(ts.Resolution)
	if nextResolutionDate.Before(nextDate) || lastValue.Date == nextDate {
		nextDate = nextResolutionDate
	}
	fmt.Println(lastValue.Value.InexactFloat64())
	lastValue.Value = baseValue.Add(lastValue.Value).Mul(ts.valueMultiplier(nextDate.Sub(lastValue.Date))).Sub(lastValue.Value)
	lastValue.Date = nextDate
	fmt.Println(lastValue.Value.InexactFloat64())
	ts.values = append(ts.values, lastValue)
}

func (ts *Interest) GetValue(at time.Time) decimal.Decimal {
	lastValue := ts.values[len(ts.values)-1]
	if ts.getLastValue().Date.Equal(at) {
		return lastValue.Value
	}
	for ts.getLastValue().Date.Before(at) {
		ts.calculateNextValue()
	}
	valueIndex := ts.baseTimeSeries.GetValueIndex(at)
	if valueIndex == -1 {
		return decimal.Decimal{}
	}
	v := ts.values[valueIndex]
	return v.Value.Mul(ts.valueMultiplier(at.Sub(v.Date)))
}

func (ts *Interest) GetAllValuesBetween(from time.Time, to time.Time) []TimeSeriesValue {
	for ts.getLastValue().Date.Before(to) {
		ts.calculateNextValue()
	}
	return ts.baseTimeSeries.GetAllValuesBetween(from, to)
}

func (ts *Interest) Rebuild() {
	ts.Series.Rebuild()
	ts.values = make([]TimeSeriesValue, 0)
	ts.values = append(ts.values, TimeSeriesValue{Value: decimal.Decimal{}, Date: time.Now()})
}

func (ts *Interest) IsDynamic() bool {
	return true
}
