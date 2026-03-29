package consumer

import "context"

type LocationConsumer interface {
	Run(ctx context.Context) error
}
