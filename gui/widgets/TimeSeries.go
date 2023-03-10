package widgets

import (
	"math"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/dukfaar/dukfin/gui/renderer"
	"github.com/shopspring/decimal"
)

type TimeSeries struct {
	widget.BaseWidget
	Series map[string]renderer.Series

	endDate *time.Time
}

func NewTimeSeries() *TimeSeries {
	result := &TimeSeries{
		Series: make(map[string]renderer.Series),
	}
	result.ExtendBaseWidget(result)
	return result
}

func (s *TimeSeries) SetEndDate(endDate time.Time) {
	s.endDate = &endDate
}

func (s *TimeSeries) CreateRenderer() fyne.WidgetRenderer {
	var minDate time.Time = time.Now()
	var maxDate time.Time = time.Now()
	if s.endDate != nil {
		maxDate = *s.endDate
	}
	minValue := decimal.NewFromInt(math.MaxInt)
	maxValue := decimal.NewFromInt(math.MinInt)
	for _, series := range s.Series {
		minV, maxV := series.Series.ValueRange(maxDate)
		minD := series.Series.MinDate()
		if minD.Before(minDate) {
			minDate = minD
		}
		minValue = decimal.Min(minValue, minV)
		maxValue = decimal.Max(maxValue, maxV)
	}
	minValue = decimal.Min(minValue, decimal.Decimal{})
	valueRange := maxValue.Sub(minValue)
	valueRangeSkip := valueRange.Div(decimal.NewFromInt(10))
	result := &renderer.TimeSeriesRenderer{}
	result.SetDateRange(minDate, maxDate)
	result.SetValueRange(minValue.Sub(valueRangeSkip), maxValue.Add(valueRangeSkip))
	for name, series := range s.Series {
		result.SetSeries(name, series)
	}
	return result
}
