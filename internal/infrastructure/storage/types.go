package storage

import (
	"context"
)

// Generic type to represent a storage
type Storage[T any] interface {
	Create(ctx context.Context, data T) error
	FindByID(ctx context.Context, id string) (T, error)
	Update(ctx context.Context, data T) error
	Delete(ctx context.Context, id string) error
}

type IDGetter[T any] interface {
	GetID(data T) string
}
