package matching

import (
	"context"
	"errors"
	"github.com/ridecore/ridecore/pkg/models"
)

var (
	ErrNoDriversAvailable = errors.New("no drivers available in radius")
	ErrLockContention     = errors.New("driver lock contention, retry")
)

type MatchingEngine interface {
	Match(ctx context.Context, order models.Order) (models.Assignment, error)
}
