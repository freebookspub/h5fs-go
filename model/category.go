package model

import (
	"fmt"
)

type Category struct {
	Id       int    `gorm:"id" json:"id"`               // 类目编号
	Name     string `gorm:"name" json:"name"`           // 类目名称
	ParentId int    `gorm:"parent_id" json:"parent_id"` // 父级类目编号
	Level    int    `gorm:"level" json:"level"`         // 类目级别
	Sort     int    `gorm:"sort" json:"sort"`           // 类目排序
	Created  string `gorm:"created" json:"created"`     // 创建时间
	Updated  string `gorm:"updated" json:"updated"`     // 更新时间
}

func (*Category) TableName() string {
	return "category"
}

func GetCategoryListsModel(id int) ([]*Category, uint64, error) {

	lists := make([]*Category, 0)
	var total uint64

	str := fmt.Sprintf("parent_id = %d", id)
	err := DB.Self.Model(&Category{}).Where(str).Count(&total).Error
	if err != nil {
		return lists, total, err
	}

	err2 := DB.Self.Where(str).Find(&lists).Error
	if err2 != nil {
		return lists, total, err
	}

	return lists, total, nil
}
