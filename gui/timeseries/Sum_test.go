package timeseries

import (
	"reflect"
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

func TestSum_GetAllValuesBetween(t *testing.T) {
	type fields struct {
		Series         []TimeSeries
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
			"Sum",
			fields{Series: []TimeSeries{
				&dynSeries{baseTimeSeries{values: []TimeSeriesValue{{Value: decimal.NewFromInt(1), Date: time.UnixMilli(5)}}}},
				&dynSeries{baseTimeSeries{values: []TimeSeriesValue{{Value: decimal.NewFromInt(2), Date: time.UnixMilli(6)}}}},
			}},
			args{from: time.UnixMilli(0), to: time.UnixMilli(10)},
			[]TimeSeriesValue{
				{Value: decimal.NewFromInt(1), Date: time.UnixMilli(5)},
				{Value: decimal.NewFromInt(3), Date: time.UnixMilli(6)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &Sum{
				Series:         tt.fields.Series,
				baseTimeSeries: tt.fields.baseTimeSeries,
			}
			if got := ts.GetAllValuesBetween(tt.args.from, tt.args.to); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sum.GetAllValuesBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
