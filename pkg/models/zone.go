package models

import "time"

const H3Resolution = 8

type ZoneStats struct {
	H3Cell         string    `json:"h3_cell"`
	Resolution     int       `json:"resolution"`
	DriverCount    int       `json:"driver_count"`
	AvailableCount int       `json:"available_count"`
	PendingOrders  int       `json:"pending_orders"`
	DemandRatio    float64   `json:"demand_ratio"`
	ComputedAt     time.Time `json:"computed_at"`
}
