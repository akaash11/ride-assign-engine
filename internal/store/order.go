package store

import (
	"context"
	"github.com/ridecore/ridecore/pkg/models"
)

type OrderStore interface {
	SaveOrder(ctx context.Context, order models.Order) error
	SaveAssignment(ctx context.Context, assignment models.Assignment) error
	GetOrdersByStatus(ctx context.Context, status models.OrderStatus, limit int) ([]models.Order, error)
}
