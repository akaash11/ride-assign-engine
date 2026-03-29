package simulator

import "context"

type Simulator interface {
	Run(ctx context.Context, driverCount int) error
}
