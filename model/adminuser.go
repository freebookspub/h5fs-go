package model

import (
	"fmt"

	"h5fs/pkg/auth"
	"h5fs/pkg/constvar"

	validator "gopkg.in/go-playground/validator.v9"
)

// User represents a registered user.
type AdminModel struct {
	Id       int    `gorm:"id" json:"id"`             // 用户编号
	OpenId   string `gorm:"open_id" json:"open_id"`   // 微信用户唯一标识
	Username string `gorm:"username" json:"username"` // 用户名称
	Password string `gorm:"password" json:"password"` // 用户密码
	Status   int    `gorm:"status" json:"status"`     // 用户状态
	Created  string `gorm:"created" json:"created"`   // 创建时间
	Updated  string `gorm:"updated" json:"updated"`   // 更新时间
}

func (c *AdminModel) TableName() string {
	return "user"
}

// Create creates a new user account.
func (u *AdminModel) Create() error {
	return DB.Self.Create(&u).Error
}

// Compare with the plain text password. Returns true if it's the same as the encrypted one (in the `User` struct).
func (u *AdminModel) Compare(pwd string) (err error) {
	err = auth.Compare(u.Password, pwd)
	return
}

func (u *AdminModel) ComparePassword(pwd string) (err error) {
	err = nil
	return
}

// Encrypt the user password.
func (u *AdminModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

// Validate the fields.
func (u *AdminModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

// DeleteUser deletes the user by the user identifier.
func DeleteAdmin(id uint64) error {
	user := AdminModel{}
	user.Id = int(id)
	return DB.Self.Delete(&user).Error
}

// Update updates an user account information.
func (u *AdminModel) Update() error {
	return DB.Self.Save(u).Error
}

// GetUser gets an user by the user identifier.
func GetAdminUser(username string) (*AdminModel, error) {
	u := &AdminModel{}
	d := DB.Self.Where("username = ?", username).First(&u)
	return u, d.Error
}

// ListUser List all users
func ListsAdmin(username string, offset, limit int) ([]*AdminModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	users := make([]*AdminModel, 0)
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
