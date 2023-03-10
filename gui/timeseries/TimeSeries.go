package timeseries

import (
	"math"
	"sort"
	"time"

	"github.com/shopspring/decimal"
)

type TimeSeries interface {
	Get(at time.Time) TimeSeriesValue
	GetNext(at time.Time) TimeSeriesValue
	GetValue(at time.Time) decimal.Decimal
	GetValueIndex(at time.Time) int
	GetAllValues() []TimeSeriesValue
	GetAllValuesBetween(from time.Time, to time.Time) []TimeSeriesValue
	MinDate() time.Time
	MaxDate() time.Time
	ValueRange(until time.Time) (decimal.Decimal, decimal.Decimal)
	Rebuild()
	IsDynamic() bool
}

type baseTimeSeries struct {
	values []TimeSeriesValue
}

func (ts *baseTimeSeries) MinDate() time.Time {
	if len(ts.values) == 0 {
		return time.Now()
	}
	return ts.values[0].Date
}

func (ts *baseTimeSeries) MaxDate() time.Time {
	return ts.values[len(ts.values)-1].Date
}

func (ts *baseTimeSeries) GetAllValues() []TimeSeriesValue {
	return ts.values
}

func (ts *baseTimeSeries) GetAllValuesBetween(from time.Time, to time.Time) []TimeSeriesValue {
	fromIndex := ts.GetValueIndex(from)
	if fromIndex < 1 {
		fromIndex = 1
	}
	toIndex := ts.GetValueIndex(to)
	return ts.values[fromIndex-1 : toIndex+1]
}

func (ts *baseTimeSeries) GetAccountValueOrder(until time.Time) []TimeSeriesValue {
	lastIndex := sort.Search(len(ts.values), func(i int) bool { return ts.values[i].Date.After(until) })
	return ts.values[0:lastIndex]
}

func (ts *baseTimeSeries) GetValueIndex(at time.Time) int {
	lastIndex := sort.Search(len(ts.values), func(i int) bool { return ts.values[i].Date.After(at) })
	if lastIndex == 0 {
		return -1
	}
	return lastIndex - 1
}

func (ts *baseTimeSeries) Get(at time.Time) TimeSeriesValue {
	index := ts.GetValueIndex(at)
	if index < 0 {
		return TimeSeriesValue{Date: at, Value: decimal.Decimal{}}
	}
	return ts.values[index]
}

func (ts *baseTimeSeries) GetValue(at time.Time) decimal.Decimal {
	return ts.Get(at).Value
}

func (ts *baseTimeSeries) GetNext(at time.Time) TimeSeriesValue {
	index := ts.GetValueIndex(at)
	if index < 0 {
		return TimeSeriesValue{Date: at, Value: decimal.Decimal{}}
	}
	if index > len(ts.values) {
		index -= 1
	}
	return ts.values[index]
}

func (ts *baseTimeSeries) ValueRange(until time.Time) (decimal.Decimal, decimal.Decimal) {
	minValue := decimal.NewFromInt(math.MaxInt64)
	maxValue := decimal.NewFromInt(math.MinInt64)
	for _, s := range ts.values {
		minValue = decimal.Min(minValue, s.Value)
		maxValue = decimal.Max(maxValue, s.Value)
	}
	return minValue, maxValue
}
