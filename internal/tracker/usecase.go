package tracker

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type Input interface {
	Get() string
}

type Output interface {
	Out(string)
}

type Store interface {
	Create(ctx context.Context, item Item) error
	List(ctx context.Context) ([]Item, error)
	Get(ctx context.Context, id string) (Item, error)
}

type Usecase interface {
	Done(ctx context.Context, in Input, out Output, store Store) error
}

type AddUsecase struct{}

func (u AddUsecase) Done(ctx context.Context, in Input, out Output, store Store) error {
	out.Out("enter name:")
	name := in.Get()
	id := uuid.New().String()

	if err := store.Create(ctx, Item{ID: id, Name: name}); err != nil {
		return fmt.Errorf("failed to create item: %w", err)
	}
	return nil
}

type GetUsecase struct{}

func (u GetUsecase) Done(
	ctx context.Context,
	in Input,
	out Output,
	store Store,
) error {
	items, err := store.List(ctx)
	if err != nil {
		return fmt.Errorf("failed to get items: %w", err)
	}
	for _, item := range items {
		out.Out(item.ID + " " + item.Name)
	}
	return nil
}
