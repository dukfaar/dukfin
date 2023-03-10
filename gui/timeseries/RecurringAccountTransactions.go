package timeseries

import (
	"math"
	"time"

	"github.com/dukfaar/dukfin/ent"
	"github.com/shopspring/decimal"
)

type RecurringAccountTransactions struct {
	addRecurringTransactions ent.RecurringTransactions
	subRecurringTransactions ent.RecurringTransactions

	_addRecurringTransactions ent.RecurringTransactions
	_subRecurringTransactions ent.RecurringTransactions

	baseTimeSeries
}

func NewRecurringAccountTransactions(addRecurringTransactions ent.RecurringTransactions, subRecurringTransactions ent.RecurringTransactions) *RecurringAccountTransactions {
	result := &RecurringAccountTransactions{
		addRecurringTransactions: addRecurringTransactions,
		subRecurringTransactions: subRecurringTransactions,
	}
	return result
}

func (ts *RecurringAccountTransactions) calculateNextValue() {
	lastValue := ts.values[len(ts.values)-1]
	nextValue := TimeSeriesValue{
		Value: lastValue.Value,
		Date:  time.UnixMilli(math.MaxInt64),
	}
	var appliedRecurringTransaction *ent.RecurringTransaction
	var isAdd bool
	for _, tx := range ts._addRecurringTransactions {
		if tx.NextDate.Before(nextValue.Date) {
			nextValue.Date = tx.NextDate
			appliedRecurringTransaction = tx
			isAdd = true
		}
	}
	for _, tx := range ts._subRecurringTransactions {
		if tx.NextDate.Before(nextValue.Date) {
			nextValue.Date = tx.NextDate
			appliedRecurringTransaction = tx
			isAdd = false
		}
	}
	if appliedRecurringTransaction != nil {
		if isAdd {
			nextValue.Value = lastValue.Value.Add(appliedRecurringTransaction.CurrencyValue)
		} else {
			nextValue.Value = lastValue.Value.Sub(appliedRecurringTransaction.CurrencyValue)
		}
		nextValue.Date = appliedRecurringTransaction.NextDate
		appliedRecurringTransaction.NextDate = appliedRecurringTransaction.NextDate.AddDate(appliedRecurringTransaction.IntervalYears, appliedRecurringTransaction.IntervalMonths, appliedRecurringTransaction.IntervalDays)
	}
	ts.values = append(ts.values, nextValue)
}

func (ts *RecurringAccountTransactions) GetValue(at time.Time) decimal.Decimal {
	lastValue := ts.values[len(ts.values)-1]
	if lastValue.Date.Equal(at) {
		return lastValue.Value
	}
	if lastValue.Date.After(at) {
		return ts.baseTimeSeries.GetValue(at)
	}
	for ; lastValue.Date.Before(at); lastValue = ts.values[len(ts.values)-1] {
		ts.calculateNextValue()
	}
	return lastValue.Value
}

func (ts *RecurringAccountTransactions) GetNext(at time.Time) TimeSeriesValue {
	index := ts.GetValueIndex(at)
	if index < 0 {
		return TimeSeriesValue{Date: at, Value: decimal.Decimal{}}
	}
	if index+1 >= len(ts.values) {
		ts.calculateNextValue()
	}
	return ts.values[index+1]
}

func (ts *RecurringAccountTransactions) Rebuild() {
	ts.values = make([]TimeSeriesValue, 0)
	ts.values = append(ts.values, TimeSeriesValue{Value: decimal.Decimal{}, Date: time.Now()})
	lastValue := ts.values[len(ts.values)-1]
	ts._addRecurringTransactions = make(ent.RecurringTransactions, len(ts.addRecurringTransactions))
	for i, tx := range ts.addRecurringTransactions {
		ts._addRecurringTransactions[i] = &ent.RecurringTransaction{
			CurrencyValue:  tx.CurrencyValue,
			NextDate:       tx.NextDate,
			IntervalYears:  tx.IntervalYears,
			IntervalMonths: tx.IntervalMonths,
			IntervalDays:   tx.IntervalDays,
		}
		for ts._addRecurringTransactions[i].NextDate.Before(lastValue.Date) {
			ts._addRecurringTransactions[i].NextDate = ts._addRecurringTransactions[i].NextDate.AddDate(ts._addRecurringTransactions[i].IntervalYears, ts._addRecurringTransactions[i].IntervalMonths, ts._addRecurringTransactions[i].IntervalDays)
		}
	}
	ts._subRecurringTransactions = make(ent.RecurringTransactions, len(ts.subRecurringTransactions))
	for i, tx := range ts.subRecurringTransactions {
		ts._subRecurringTransactions[i] = &ent.RecurringTransaction{
			CurrencyValue:  tx.CurrencyValue,
			NextDate:       tx.NextDate,
			IntervalYears:  tx.IntervalYears,
			IntervalMonths: tx.IntervalMonths,
			IntervalDays:   tx.IntervalDays,
		}
		for ts._subRecurringTransactions[i].NextDate.Before(lastValue.Date) {
			ts._subRecurringTransactions[i].NextDate = ts._subRecurringTransactions[i].NextDate.AddDate(ts._subRecurringTransactions[i].IntervalYears, ts._subRecurringTransactions[i].IntervalMonths, ts._subRecurringTransactions[i].IntervalDays)
		}
	}
}

func (ts *RecurringAccountTransactions) IsDynamic() bool {
	return true
}
