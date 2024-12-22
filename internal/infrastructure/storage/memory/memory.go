package memory

import (
	"context"
	"errors"
	"sync"

	"github.com/kien6034/secure-genom/internal/infrastructure/storage"
)

type MemoryStorage[T any] struct {
	store    map[string]T
	mu       sync.RWMutex
	idGetter storage.IDGetter[T]
}

func NewMemoryStorage[T any](idGetter storage.IDGetter[T]) storage.Storage[T] {
	return &MemoryStorage[T]{
		store:    make(map[string]T),
		idGetter: idGetter,
	}
}

func (m *MemoryStorage[T]) Create(ctx context.Context, data T) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	id := m.idGetter.GetID(data)
	if id == "" {
		return errors.New("invalid ID")
	}

	if _, exists := m.store[id]; exists {
		return errors.New("entity already exists")
	}

	m.store[id] = data
	return nil
}

func (m *MemoryStorage[T]) FindByID(ctx context.Context, id string) (T, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	data, exists := m.store[id]
	if !exists {
		var zero T
		return zero, errors.New("entity not found")
	}
	return data, nil
}

func (m *MemoryStorage[T]) Update(ctx context.Context, data T) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	id := m.idGetter.GetID(data)
	if _, exists := m.store[id]; !exists {
		return errors.New("entity not found")
	}

	m.store[id] = data
	return nil
}

func (m *MemoryStorage[T]) Delete(ctx context.Context, id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.store[id]; !exists {
		return errors.New("entity not found")
	}

	delete(m.store, id)
	return nil
}
