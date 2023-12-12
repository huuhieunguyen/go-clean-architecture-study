package todostorage

import (
	"context"

	todomodel "clean-architecture/module/item/model"
	pagination "clean-architecture/utils"
)

func (s *mysqlStorage) ListItem(
	ctx context.Context,
	condition map[string]interface{},
	paging *pagination.DataPaging,
) ([]todomodel.ToDoItem, error) {
	offset := (paging.Page - 1) * paging.Limit

	var result []todomodel.ToDoItem

	if err := s.db.Table(todomodel.ToDoItem{}.TableName()).
		Where(condition).
		Count(&paging.Total).
		Offset(offset).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
