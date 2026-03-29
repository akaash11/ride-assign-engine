package models

import "time"

type DriverStatus string

const (
	StatusAvailable DriverStatus = "available"
	StatusBusy      DriverStatus = "busy"
	StatusOffline   DriverStatus = "offline"
)

type Driver struct {
	ID        string       `json:"driver_id"`
	Lat       float64      `json:"lat"`
	Lng       float64      `json:"lng"`
	H3Cell    string       `json:"h3_cell"`
	Status    DriverStatus `json:"status"`
	UpdatedAt time.Time    `json:"updated_at"`
}

type DriverLocation struct {
	DriverID  string       `json:"driver_id"`
	Lat       float64      `json:"lat"`
	Lng       float64      `json:"lng"`
	Status    DriverStatus `json:"status"`
	Timestamp time.Time    `json:"ts"`
}

type NearbyDriver struct {
	Driver
	DistanceKm float64 `json:"distance_km"`
}
