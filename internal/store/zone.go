package store

import (
	"context"
	"github.com/ridecore/ridecore/pkg/models"
)

type ZoneStore interface {
	UpsertZoneStats(ctx context.Context, stats models.ZoneStats) error
	GetZoneStats(ctx context.Context, h3Cell string) (models.ZoneStats, error)
	GetTopDemandZones(ctx context.Context, limit int) ([]models.ZoneStats, error)
}
