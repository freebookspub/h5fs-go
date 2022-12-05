package webapis

import (
	. "h5fs/handler"
	"h5fs/model"
	"h5fs/pkg/errno"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary List the users in the database
// @Description List users
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body user.ListRequest true "List users"
// @Success 200 {object} user.SwaggerListResponse "{"code":0,"message":"OK","data":{"totalCount":1,"userList":[{"id":0,"username":"admin","random":"user 'admin' get random string 'EnqntiSig'","password":"$2a$10$veGcArz47VGj7l9xN7g2iuT9TF21jLI1YGXarGzvARNdnt4inC9PG","createdAt":"2018-05-28 00:25:33","updatedAt":"2018-05-28 00:25:33"}]}}"
// @Router /user [get]
func GetIndexListsHandler(c *gin.Context) {
	var r GetIndexListsRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	page := c.Query("page")

	offset, err := strconv.Atoi(page)
	if err != nil {
		offset = 1
	}

	infos, count, err := model.GetListsNewsModel(offset)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, GetIndexListsRespones{
		TotalCount: count,
		Lists:      infos,
	})
}

func GetNewsHandler(c *gin.Context) {
	newsId := c.Param("id")
	id, err := strconv.Atoi(newsId)
	if err != nil {
		id = 1
	}

	infos, err := model.GetNewsOneModel(id)

	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, GetNewsOneRespones{
		Lists: infos,
	})
}

func GetNewsPostsHandler(c *gin.Context) {
	uuid := c.Param("uuid")

	info, err := model.GetPostsOneModel(uuid)

	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, GetPostsOneRespones{
		Lists: info,
	})
}

func GetIndexOneHandler(c *gin.Context) {
	newsId := c.Param("id")
	id, err := strconv.Atoi(newsId)
	if err != nil {
		id = 1
	}

	infos, err := model.GetIndexOneModel(id)

	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, GetIndexOneRespones{
		Lists: infos,
	})
}

func GetTagsListsHandler(c *gin.Context) {
	tag := c.Param("id")

	lists, err := model.GetTagsListsModel(tag)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, GetIndexListsRespones{
		TotalCount: 100,
		Lists:      lists,
	})

}
