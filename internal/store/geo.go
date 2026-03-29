package store

import (
	"context"
	"github.com/ridecore/ridecore/pkg/models"
)

type GeoStore interface {
	UpdateDriverLocation(ctx context.Context, loc models.DriverLocation) error
	FindNearbyDrivers(ctx context.Context, lat, lng, radiusKm float64) ([]models.NearbyDriver, error)
	SetDriverStatus(ctx context.Context, driverID string, status models.DriverStatus) error
	GetDriver(ctx context.Context, driverID string) (models.Driver, error)
}
