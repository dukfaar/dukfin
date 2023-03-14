package timeseries

import (
	"sort"
	"time"

	"github.com/shopspring/decimal"
)

type Multiply struct {
	Series1 TimeSeries
	Series2 TimeSeries

	baseTimeSeries
}

func (ts *Multiply) calculateValue(at time.Time) decimal.Decimal {
	return ts.Series1.GetValue(at).Mul(ts.Series2.GetValue(at))
}

func (ts *Multiply) Get(at time.Time) TimeSeriesValue {
	if ts.IsDynamic() {
		return TimeSeriesValue{
			Value: ts.calculateValue(at),
			Date:  at,
		}
	}
	return ts.baseTimeSeries.Get(at)
}

func (ts *Multiply) GetValue(at time.Time) decimal.Decimal {
	return ts.Get(at).Value
}

func (ts *Multiply) ValueRange(until time.Time) (decimal.Decimal, decimal.Decimal) {
	if ts.IsDynamic() {
		minValue := ts.GetValue(until)
		maxValue := ts.GetValue(until)
		return minValue, maxValue
	}
	return ts.baseTimeSeries.ValueRange(until)
}

func (ts *Multiply) Rebuild() {
	ts.Series1.Rebuild()
	ts.Series2.Rebuild()
	if !ts.IsDynamic() {
		values1 := ts.Series1.GetAllValues()
		values2 := ts.Series2.GetAllValues()
		ts.values = make([]TimeSeriesValue, 0, len(values1)+len(values2))
		for _, v := range values1 {
			ts.values = append(ts.values, TimeSeriesValue{Value: ts.calculateValue(v.Date), Date: v.Date})
		}
		for _, v := range values2 {
			ts.values = append(ts.values, TimeSeriesValue{Value: ts.calculateValue(v.Date), Date: v.Date})
		}
		sort.Slice(ts.values, func(i, j int) bool {
			return ts.values[i].Date.Before(ts.values[j].Date)
		})
	}
}

func (ts *Multiply) GetAllValuesBetween(from time.Time, to time.Time) []TimeSeriesValue {
	if ts.IsDynamic() {
		values1 := ts.Series1.GetAllValues()
		values2 := ts.Series2.GetAllValues()
		ts.values = make([]TimeSeriesValue, 0, len(values1)+len(values2))
		for _, v := range values1 {
			ts.values = append(ts.values, TimeSeriesValue{Value: ts.calculateValue(v.Date), Date: v.Date})
		}
		for _, v := range values2 {
			ts.values = append(ts.values, TimeSeriesValue{Value: ts.calculateValue(v.Date), Date: v.Date})
		}
		sort.Slice(ts.values, func(i, j int) bool {
			return ts.values[i].Date.Before(ts.values[j].Date)
		})
	}
	return ts.baseTimeSeries.GetAllValuesBetween(from, to)
}

func (ts *Multiply) IsDynamic() bool {
	return ts.Series1.IsDynamic() || ts.Series2.IsDynamic()
}
