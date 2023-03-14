package timeseries

type dynSeries struct {
	baseTimeSeries
}

func (ts *dynSeries) Rebuild() {}
func (ts *dynSeries) IsDynamic() bool {
	return true
}
