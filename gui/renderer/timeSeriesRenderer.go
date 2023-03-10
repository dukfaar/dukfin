package renderer

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"github.com/dukfaar/dukfin/gui/timeseries"
	"github.com/shopspring/decimal"
	"github.com/wcharczuk/go-chart"
)

type Series struct {
	Series timeseries.TimeSeries
	Color  color.RGBA
}

type TimeSeriesRenderer struct {
	fyne.WidgetRenderer

	series map[string]Series

	image          *canvas.Image
	drawImage      *image.RGBA
	redrawRequired bool

	chart chart.Chart

	decHeight decimal.Decimal

	fromDate         time.Time
	toDate           time.Time
	dateRangeSeconds float64
	pixelDuration    time.Duration

	fromValue  decimal.Decimal
	toValue    decimal.Decimal
	valueRange decimal.Decimal
}

func (r *TimeSeriesRenderer) SetSeries(name string, s Series) {
	if r.series == nil {
		r.series = make(map[string]Series)
	}
	r.series[name] = s
	r.chart.Series = []chart.Series{}
	for name, series := range r.series {
		ts := chart.TimeSeries{
			Name:    name,
			XValues: []time.Time{},
			YValues: []float64{},
			YAxis:   chart.YAxisPrimary,
			Style: chart.Style{
				Show:        true,
				StrokeColor: chart.GetDefaultColor(0).WithAlpha(128),
				FillColor:   chart.GetDefaultColor(0).WithAlpha(0),
			},
		}
		for _, value := range series.Series.GetAllValuesBetween(r.fromDate, r.toDate) {
			ts.XValues = append(ts.XValues, value.Date)
			ts.YValues = append(ts.YValues, value.Value.InexactFloat64())
		}
		r.chart.Series = append(r.chart.Series, ts)
	}
	r.redrawRequired = true
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
