package todobiz

import (
	todomodel "clean-architecture/module/item/model"
	pagination "clean-architecture/utils"
	"context"
)

type ListTodoItemStorage interface {
	ListItem(
		ctx context.Context,
		condition map[string]interface{},
		paging *pagination.DataPaging,
	) ([]todomodel.ToDoItem, error)
}

type listBiz struct {
	store ListTodoItemStorage
}

func NewListToDoItemBiz(store ListTodoItemStorage) *listBiz {
	return &listBiz{store: store}
}

func (biz *listBiz) ListItems(ctx context.Context,
	condition map[string]interface{},
	paging *pagination.DataPaging,
) ([]todomodel.ToDoItem, error) {
	result, err := biz.store.ListItem(ctx, condition, paging)

	if err != nil {
		return nil, err
	}

	return result, err
}
