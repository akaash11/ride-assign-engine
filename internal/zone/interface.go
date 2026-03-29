package zone

import (
	"context"
	"github.com/ridecore/ridecore/pkg/models"
)

type ZoneService interface {
	ComputeZoneStats(ctx context.Context) error
	GetZoneStats(ctx context.Context, h3Cell string) (models.ZoneStats, error)
}
