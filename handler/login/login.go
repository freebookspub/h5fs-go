package login

import (
	"fmt"
	"h5fs/handler"
	"h5fs/model"
	"h5fs/pkg/errno"
	webapijson "h5fs/pkg/redata"

	gonanoid "github.com/matoous/go-nanoid/v2"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	var u model.WebUserModel
	if err := c.ShouldBind(&u); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	fmt.Println("login:", u)
	id, err := gonanoid.New()
	if err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	fmt.Println("weblogin ", id)

	handler.SendResponse(c, nil, webapijson.WebUserResponse{})

}

func Register(c *gin.Context) {
	var u model.WebUserModel
	err := c.ShouldBind(&u)
	if err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	token, err := gonanoid.New()
	if err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	u.Token = token

	if err := u.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}

	isNull := u.FindEmail()
	if isNull.Code != 0 {
		handler.SendResponse(c, errno.ErrDatabase, isNull)
		return
	}

	ret := u.Create()
	if ret.Code != 0 {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	handler.SendResponse(c, nil, ret)

}

func UserInfo(c *gin.Context) {

}
