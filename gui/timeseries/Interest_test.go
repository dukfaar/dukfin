package timeseries

import (
	"reflect"
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

func TestInterest_calculateNextValue(t *testing.T) {
	type fields struct {
		Series         TimeSeries
		APR            decimal.Decimal
		From           time.Time
		Resolution     time.Duration
		baseTimeSeries baseTimeSeries
	}
	tests := []struct {
		name   string
		fields fields
		want   TimeSeriesValue
	}{
		{"Interest",
			fields{
				Series: &dynSeries{baseTimeSeries{values: []TimeSeriesValue{
					TimeSeriesValue{Value: decimal.NewFromInt(100), Date: time.UnixMilli(0)},
				}}},
				APR:            decimal.NewFromInt(10),
				Resolution:     time.Hour * 24 * 360,
				baseTimeSeries: baseTimeSeries{values: []TimeSeriesValue{TimeSeriesValue{Value: decimal.NewFromInt(0), Date: time.UnixMilli(0)}}},
			},
			TimeSeriesValue{Value: decimal.NewFromInt(10), Date: time.UnixMilli(0).Add(time.Hour * 24 * 360)},
		},
		{"Interest with startValue",
			fields{
				Series: &dynSeries{baseTimeSeries{values: []TimeSeriesValue{
					TimeSeriesValue{Value: decimal.NewFromInt(100), Date: time.UnixMilli(0)},
				}}},
				APR:            decimal.NewFromInt(10),
				Resolution:     time.Hour * 24 * 360,
				baseTimeSeries: baseTimeSeries{values: []TimeSeriesValue{TimeSeriesValue{Value: decimal.NewFromInt(10), Date: time.UnixMilli(0)}}},
			},
			TimeSeriesValue{Value: decimal.NewFromInt(21), Date: time.UnixMilli(0).Add(time.Hour * 24 * 360)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &Interest{
				Series:         tt.fields.Series,
				APR:            tt.fields.APR,
				From:           tt.fields.From,
				Resolution:     tt.fields.Resolution,
				baseTimeSeries: tt.fields.baseTimeSeries,
			}
			ts.calculateNextValue()
			if got := ts.getLastValue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculateNextValue; new last value = %v, want %v", got, tt.want)
			}
		})
	}
}
