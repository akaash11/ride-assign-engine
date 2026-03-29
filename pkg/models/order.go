package models

import "time"

type OrderStatus string

const (
	OrderPending   OrderStatus = "pending"
	OrderMatched   OrderStatus = "matched"
	OrderUnmatched OrderStatus = "unmatched"
)

type Order struct {
	ID         string      `json:"order_id"`
	PickupLat  float64     `json:"pickup_lat"`
	PickupLng  float64     `json:"pickup_lng"`
	Status     OrderStatus `json:"status"`
	AssignedTo string      `json:"driver_id,omitempty"`
	CreatedAt  time.Time   `json:"created_at"`
	MatchedAt  *time.Time  `json:"matched_at,omitempty"`
}

type Assignment struct {
	OrderID   string    `json:"order_id"`
	DriverID  string    `json:"driver_id"`
	CreatedAt time.Time `json:"created_at"`
	MatchedMs int64     `json:"matched_ms"`
}
