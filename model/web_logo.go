package model

import (
	"fmt"

	"h5fs/pkg/auth"
	"h5fs/pkg/constvar"
	webapijson "h5fs/pkg/redata"

	validator "gopkg.in/go-playground/validator.v9"
)

type WebUserModel struct {
	Id       int    `gorm:"id" json:"id"`
	Username string `form:"username" json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=100"`
	Password string `form:"password" json:"password" gorm:"column:passwrod;not null" binding:"required" validate:"min=1,max=256"`
	Email    string `form:"email" json:"email" gorm:"column:email;not null" binding:"required" validate:"min=1,max=256"`
	Token    string `gorm:"token" json:"token"`
}

type RegisterModel struct {
	Username string `gorm:"username" json:"username"`
	Password string `gorm:"password" json:"password"`
	Email    string `gorm:"email" json:"email"`
	Token    string `gorm:"token" json:"token"`
}

func (c *WebUserModel) TableName() string {
	return "user"
}

// Create creates a new user account.
func (u *WebUserModel) Create() webapijson.Register {
	ret := webapijson.Register{
		Code:    0,
		Message: "OK",
	}
	register := RegisterModel{
		Username: u.Username,
		Password: u.Password,
		Email:    u.Email,
		Token:    u.Token,
	}
	if err := DB.Self.Table("user").Create(&register).Error; err != nil {
		ret.Code = 400
		ret.Message = "DataBase Error in Create user."
	}
	return ret
}

func (u *WebUserModel) FindEmail() webapijson.Register {
	ret := webapijson.Register{
		Code:    0,
		Message: "OK",
	}
	d := DB.Self.Where("email = ?", u.Email).First(&u)
	if d.Error != nil {
		ret.Code = 400
		ret.Message = "DataBase Error in Find email."
	}

	if d.RowsAffected != 0 {
		ret.Code = 500
		ret.Message = "You are already registered."
	}

	return ret
}

// Compare with the plain text password. Returns true if it's the same as the encrypted one (in the `User` struct).
func (u *WebUserModel) Compare(pwd string) (err error) {
	err = auth.Compare(u.Password, pwd)
	return
}

func (u *WebUserModel) ComparePassword(pwd string) (err error) {
	err = nil
	return
}

// Encrypt the user password.
func (u *WebUserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

// Validate the fields.
func (u *WebUserModel) Validate() error {
	fmt.Println("aaaaabbbbb", u)
	validate := validator.New()
	return validate.Struct(u)
}

// DeleteUser deletes the user by the user identifier.
func DeleteWebUser(id uint64) error {
	user := WebUserModel{}
	user.Id = int(id)
	return DB.Self.Delete(&user).Error
}

// Update updates an user account information.
func (u *WebUserModel) Update() error {
	return DB.Self.Save(u).Error
}

// GetUser gets an user by the user identifier.
func GetWebuser(username string) (*WebUserModel, error) {
	u := &WebUserModel{}
	d := DB.Self.Where("username = ?", username).First(&u)
	return u, d.Error
}

// ListUser List all users
func ListsWebUser(username string, offset, limit int) ([]*WebUserModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	users := make([]*WebUserModel, 0)
	var count uint64

	where := fmt.Sprintf("username like '%%%s%%'", username)
	if err := DB.Self.Model(&UserModel{}).Where(where).Count(&count).Error; err != nil {
		return users, count, err
	}

	if err := DB.Self.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err
	}

	return users, count, nil
}
