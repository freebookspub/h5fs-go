package webapis

import (
	. "h5fs/handler"
	"h5fs/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCategoryLists(c *gin.Context) {
	parentID := c.Query("id")
	id, err := strconv.Atoi(parentID)
	if err != nil {
		id = 1
	}
	var count uint64

	infos, count, err := model.GetCategoryListsModel(id)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, GetCategoryListsRespones{
		TotalCount: count,
		Lists:      infos,
	})

}
