package events

type LocationEvent struct {
	DriverID string  `json:"driver_id"`
	Lat      float64 `json:"lat"`
	Lng      float64 `json:"lng"`
	Status   string  `json:"status"`
	TsUnixMs int64   `json:"ts_unix_ms"`
	SimBatch string  `json:"sim_batch"`
}
