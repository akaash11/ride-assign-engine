package store

import (
	"context"
	"errors"
	"time"
)

var ErrLockNotAcquired = errors.New("lock not acquired")

type LockStore interface {
	Acquire(ctx context.Context, resource string, ttl time.Duration) (token string, err error)
	Release(ctx context.Context, resource string, token string) error
}
