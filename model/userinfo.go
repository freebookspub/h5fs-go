package model

import (
	"h5fs/pkg/auth"

	validator "gopkg.in/go-playground/validator.v9"
)

type Userinfo struct {
	Id          int    `gorm:"id" json:"id"`                     // 用户ID
	Nickname    string `gorm:"nickname" json:"nickname"`         // 昵称
	Username    string `gorm:"username" json:"username"`         // 用户名
	Phone       string `gorm:"phone" json:"phone"`               // 手机号
	Password    string `gorm:"password" json:"password"`         // MD5密码
	Salt        string `gorm:"salt" json:"salt"`                 // 盐值
	Status      int    `gorm:"status" json:"status"`             // 状态，1正常，2停用
	Avatar      string `gorm:"avatar" json:"avatar"`             // 用户头像
	Balance     int    `gorm:"balance" json:"balance"`           // 用户积分
	IsAdmin     int    `gorm:"is_admin" json:"is_admin"`         // 是否管理员
	IsDel       int    `gorm:"is_del" json:"is_del"`             // 是否删除 0 为未删除、1 为已删除
	DeletedTime int    `gorm:"deleted_time" json:"deleted_time"` // 删除时间
	CreateTime  string `gorm:"create_time" json:"create_time"`   // 创建时间
	UpdateTime  string `gorm:"update_time" json:"update_time"`   // 更新时间
	Emal        string `gorm:"emal" json:"emal"`                 // 邮箱
	LoginTime   string `gorm:"login_time" json:"login_time"`     // 最后login时间
	From        string `gorm:"from" json:"from"`                 // 用户来源,facebook, google etc.
	Birthday    string `gorm:"birthday" json:"birthday"`         // 用户生日
}

func (*Userinfo) TableName() string {
	return "userinfo"
}

// Create creates a new user account.
func (u *Userinfo) CreateUserInfo() error {
	return DB.Self.Create(&u).Error
}

// DeleteUser deletes the user by the user identifier.
func (u *Userinfo) Delete(id uint64) error {
	u.IsDel = 1
	return DB.Self.Delete(u).Error
}

// Update updates an user account information.
func (u *Userinfo) Update() error {
	return DB.Self.Save(u).Error
}

// Compare with the plain text password. Returns true if it's the same as the encrypted one (in the `User` struct).
func (u *Userinfo) Compare(pwd string) (err error) {
	err = auth.Compare(u.Password, pwd)
	return
}

// Encrypt the user password.
func (u *Userinfo) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

// Validate the fields.
func (u *Userinfo) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

// GetUser gets an user by the user identifier.
func LoginUser(username string, _ string) (*Userinfo, error) {
	u := &Userinfo{}
	d := DB.Self.Where("username = ?", username).First(&u)
	return u, d.Error
}

// ListUser List all users
// func ListUser(username string, offset, limit int) ([]*UserModel, uint64, error) {
// 	if limit == 0 {
// 		limit = constvar.DefaultLimit
// 	}

// 	users := make([]*UserModel, 0)
// 	var count uint64

// 	where := fmt.Sprintf("username like '%%%s%%'", username)
// 	if err := DB.Self.Model(&UserModel{}).Where(where).Count(&count).Error; err != nil {
// 		return users, count, err
// 	}

// 	if err := DB.Self.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
// 		return users, count, err
// 	}

// 	return users, count, nil
// }
