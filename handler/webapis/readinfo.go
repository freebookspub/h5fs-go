package webapis

import (
	. "h5fs/handler"
	"h5fs/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetReadHandler(c *gin.Context) {
	newsId := c.Param("id")

	id, err := strconv.Atoi(newsId)

	if err != nil {
		SendResponse(c, err, GetReadListsRespones{
			Lists: nil,
		})
	}

	infos, err := model.GetReadListsModel(id)

	SendResponse(c, err, GetReadListsRespones{
		Lists: infos,
	})

}
