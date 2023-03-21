package renderer

import (
	"bytes"
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"github.com/dukfaar/dukfin/gui/timeseries"
	"github.com/shopspring/decimal"
	"github.com/wcharczuk/go-chart"
)

type series struct {
	Series timeseries.TimeSeries
	name   string
	index  int
}

type TimeSeriesRenderer struct {
	fyne.WidgetRenderer

	series map[string]*series

	image          *canvas.Image
	redrawRequired bool

	chart chart.Chart

	fromDate         time.Time
	toDate           time.Time
	dateRangeSeconds float64

	fromValue  decimal.Decimal
	toValue    decimal.Decimal
	valueRange decimal.Decimal
}

func (r *TimeSeriesRenderer) buildTS(s *series) chart.TimeSeries {
	ts := chart.TimeSeries{
		Name:    s.name,
		XValues: []time.Time{},
		YValues: []float64{},
		YAxis:   chart.YAxisPrimary,
		Style: chart.Style{
			Show:        true,
			StrokeColor: chart.GetDefaultColor(s.index).WithAlpha(128),
			FillColor:   chart.GetDefaultColor(s.index).WithAlpha(0),
		},
	}
	for _, value := range s.Series.GetAllValuesBetween(r.fromDate, r.toDate) {
		ts.XValues = append(ts.XValues, value.Date)
		ts.YValues = append(ts.YValues, value.Value.InexactFloat64())
	}
	return ts
}

func (r *TimeSeriesRenderer) SetSeries(name string, s timeseries.TimeSeries) {
	if r.series == nil {
		r.series = make(map[string]*series)
		r.chart.Series = []chart.Series{}
	}
	ts, ok := r.series[name]
	if !ok {
		ts = &series{
			Series: s,
			name:   name,
			index:  len(r.series),
		}
		r.chart.Series = append(r.chart.Series, r.buildTS(ts))
		r.series[name] = ts
		return
	}
	r.chart.Series[ts.index] = r.buildTS(ts)
}

func (r *TimeSeriesRenderer) SetDateRange(from time.Time, to time.Time) {
	r.fromDate = from
	r.toDate = to
	r.dateRangeSeconds = r.toDate.Sub(r.fromDate).Seconds()
	r.redrawRequired = true
}

func (r *TimeSeriesRenderer) SetValueRange(from decimal.Decimal, to decimal.Decimal) {
	r.fromValue = from
	r.toValue = to
	r.valueRange = r.toValue.Sub(r.fromValue)
	r.redrawRequired = true
}

func (r *TimeSeriesRenderer) MinSize() fyne.Size {
	if r.image == nil {
		return fyne.Size{Width: 0, Height: 0}
	}
	return r.image.MinSize()
}

func (r *TimeSeriesRenderer) Destroy() {
}

func (r *TimeSeriesRenderer) Objects() []fyne.CanvasObject {
	if r.image == nil {
		return []fyne.CanvasObject{}
	}
	return []fyne.CanvasObject{r.image}
}

func (r *TimeSeriesRenderer) Refresh() {
	if r.image != nil {
		r.image.Refresh()
	}
}

func (r *TimeSeriesRenderer) rerender(s fyne.Size) {
	imageReadwriter := &bytes.Buffer{}
	r.chart.XAxis = chart.XAxis{
		Name:           "X",
		ValueFormatter: chart.TimeDateValueFormatter,
		Style: chart.Style{
			Show: true,
		},
	}
	r.chart.YAxis = chart.YAxis{
		Name: "Y",
		Style: chart.Style{
			Show: true,
		},
	}
	r.chart.Width = int(s.Width)
	r.chart.Height = int(s.Height)
	err := r.chart.Render(chart.PNG, imageReadwriter)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	r.image = canvas.NewImageFromReader(imageReadwriter, "chart.png")
	r.image.Resize(s)
}

func (r *TimeSeriesRenderer) Layout(s fyne.Size) {
	if r.redrawRequired || (r.image != nil && !s.Subtract(r.image.Size()).IsZero()) {
		r.rerender(s)
	}
	r.redrawRequired = false
}
