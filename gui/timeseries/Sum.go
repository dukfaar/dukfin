package timeseries

import (
	"sort"
	"time"

	"github.com/shopspring/decimal"
)

type Sum struct {
	Series []TimeSeries

	baseTimeSeries
}

func (ts *Sum) calculateValue(at time.Time) decimal.Decimal {
	result := decimal.Decimal{}
	for _, series := range ts.Series {
		result = result.Add(series.GetValue(at))
	}
	return result
}

func (ts *Sum) Get(at time.Time) TimeSeriesValue {
	if ts.IsDynamic() {
		result := TimeSeriesValue{
			Value: ts.calculateValue(at),
			Date:  at,
		}
		return result
	}
	return ts.baseTimeSeries.Get(at)
}

func (ts *Sum) GetValue(at time.Time) decimal.Decimal {
	return ts.Get(at).Value
}

func (ts *Sum) ValueRange(until time.Time) (decimal.Decimal, decimal.Decimal) {
	if ts.IsDynamic() {
		minValue := ts.GetValue(until)
		maxValue := ts.GetValue(until)
		return minValue, maxValue
	}
	return ts.baseTimeSeries.ValueRange(until)
}

func (ts *Sum) Rebuild() {
	for _, s := range ts.Series {
		s.Rebuild()
	}
	if !ts.IsDynamic() {
		ts.values = make([]TimeSeriesValue, 0)
		for _, s := range ts.Series {
			for _, v := range s.GetAllValues() {
				ts.values = append(ts.values, TimeSeriesValue{Value: ts.calculateValue(v.Date), Date: v.Date})
			}
		}
		sort.Slice(ts.values, func(i, j int) bool {
			return ts.values[i].Date.Before(ts.values[j].Date)
		})
	}
}

func (ts *Sum) IsDynamic() bool {
	for _, s := range ts.Series {
		if s.IsDynamic() {
			return true
		}
	}
	return false
}
