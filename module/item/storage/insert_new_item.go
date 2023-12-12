package todostorage

import (
	todomodel "clean-architecture/module/item/model"
	"context"
)

func (s *mysqlStorage) CreateItem(ctx context.Context, data *todomodel.ToDoItem) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
