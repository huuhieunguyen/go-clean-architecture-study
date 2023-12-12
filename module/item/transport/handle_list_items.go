package todotrpt

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	todobiz "clean-architecture/module/item/business"
	todostorage "clean-architecture/module/item/storage"
	pagination "clean-architecture/utils"
)

func HandleListItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging pagination.DataPaging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		paging.Process()

		storage := todostorage.NewMySQLStorage(db)
		biz := todobiz.NewListToDoItemBiz(storage)

		result, err := biz.ListItems(c.Request.Context(), nil, &paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": result, "paging": paging})
	}
}
