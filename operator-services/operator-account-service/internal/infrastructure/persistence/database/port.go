package database

import "context"

type Port interface {
	Connect(ctx context.Context) error
	Close() error
}
