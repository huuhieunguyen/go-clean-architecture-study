package todostorage

import (
	todomodel "clean-architecture/module/item/model"
	"context"
)

func (s *mysqlStorage) UpdateItem(
	ctx context.Context,
	condition map[string]interface{},
	dataUpdate *todomodel.ToDoItem,
) error {
	if err := s.db.Where(condition).Updates(dataUpdate).Error; err != nil {
		return err
	}

	return nil
}
