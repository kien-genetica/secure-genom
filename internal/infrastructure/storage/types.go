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

func getID[T any](data T) string {
	if entity, ok := any(data).(IDGetter[T]); ok {
		return entity.GetID(data)
	}
	return ""
}

//
// type DataIDGetter struct{}

// func (g DataIDGetter) GetID(data *entity.Data) string {
//     return data.ID
// }

// // Usage
// data := &entity.Data{ID: "123"}
// id := getID(data) // Returns "123" if Data implements IDGetter
