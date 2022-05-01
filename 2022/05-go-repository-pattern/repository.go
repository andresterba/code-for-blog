package main

import "github.com/google/uuid"

type storableConstraint interface {
	ID() uuid.UUID
}

type GenericRepository[T storableConstraint] struct {
	Store map[uuid.UUID]T
}

func NewGenericRepository[T storableConstraint]() GenericRepository[T] {
	return GenericRepository[T]{
		Store: make(map[uuid.UUID]T),
	}
}

func (t *GenericRepository[T]) Add(item T) error {
	_, ok := t.Store[item.ID()]
	if ok {
		return nil
	}

	t.Store[item.ID()] = item

	return nil
}

func (t *GenericRepository[T]) Update(item T) error {
	_, ok := t.Store[item.ID()]
	if ok {
		return nil
	}

	t.Store[item.ID()] = item

	return nil
}

func (t *GenericRepository[T]) Delete(item T) error {
	delete(t.Store, item.ID())

	return nil
}

func (t *GenericRepository[T]) Get(id uuid.UUID) (T, error) {
	item, ok := t.Store[id]
	if !ok {
		return *new(T), nil
	}

	return item, nil
}

func (t *GenericRepository[T]) GetAll() ([]T, error) {
	var allItems []T

	for _, item := range t.Store {
		allItems = append(allItems, item)
	}

	return allItems, nil
}
