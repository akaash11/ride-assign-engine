package events

type OrderRequestEvent struct {
	OrderID     string  `json:"order_id"`
	PickupLat   float64 `json:"pickup_lat"`
	PickupLng   float64 `json:"pickup_lng"`
	RadiusKm    float64 `json:"radius_km"`
	MaxWaitMs   int64   `json:"max_wait_ms"`
	PublishedAt int64   `json:"published_at"`
}

type AssignmentEvent struct {
	OrderID    string `json:"order_id"`
	DriverID   string `json:"driver_id"`
	MatchedMs  int64  `json:"matched_ms"`
	AssignedAt int64  `json:"assigned_at"`
}
