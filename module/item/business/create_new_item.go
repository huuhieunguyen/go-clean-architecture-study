package todobiz

import (
	todomodel "clean-architecture/module/item/model"
	"context"
)

type CreateTodoItemStorage interface {
	CreateItem(ctx context.Context, data *todomodel.ToDoItem) error
}

type createBiz struct {
	store CreateTodoItemStorage
}

// This function is a constructor that creates a new instance of the createBiz struct with the provided storage implementation.
func NewCreateToDoItemBiz(store CreateTodoItemStorage) *createBiz {
	return &createBiz{store: store}
}

func (biz *createBiz) CreateNewItem(ctx context.Context, data *todomodel.ToDoItem) error {
	if err := data.Validate(); err != nil {
		return err
	}

	// do not allow "finished" status when creating a new task
	data.Status = "Doing" // set to default

	if err := biz.store.CreateItem(ctx, data); err != nil {
		return err
	}

	return nil
}
