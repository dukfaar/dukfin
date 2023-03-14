package timeseries

import (
	"reflect"
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

func TestMultiply_GetAllValuesBetween(t *testing.T) {
	type fields struct {
		Series1        TimeSeries
		Series2        TimeSeries
		baseTimeSeries baseTimeSeries
	}
	type args struct {
		from time.Time
		to   time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []TimeSeriesValue
	}{
		{
			"Multiply",
			fields{
				Series1: &dynSeries{baseTimeSeries{values: []TimeSeriesValue{{Value: decimal.NewFromInt(2), Date: time.UnixMilli(5)}}}},
				Series2: &dynSeries{baseTimeSeries{values: []TimeSeriesValue{{Value: decimal.NewFromInt(5), Date: time.UnixMilli(6)}}}},
			},
			args{from: time.UnixMilli(0), to: time.UnixMilli(10)},
			[]TimeSeriesValue{
				{Value: decimal.NewFromInt(0), Date: time.UnixMilli(5)},
				{Value: decimal.NewFromInt(10), Date: time.UnixMilli(6)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &Multiply{
				Series1:        tt.fields.Series1,
				Series2:        tt.fields.Series2,
				baseTimeSeries: tt.fields.baseTimeSeries,
			}
			if got := ts.GetAllValuesBetween(tt.args.from, tt.args.to); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Multiply.GetAllValuesBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
