package model

type Readinfo struct {
	Id     int `gorm:"id" json:"id"`           // ID
	NewsId int `gorm:"news_id" json:"news_id"` // 新闻ID
	Read   int `gorm:"read" json:"read"`       // 文章阅读数量
}

func (*Readinfo) TableName() string {
	return "readinfo"
}

func GetReadListsModel(id int) (*Readinfo, error) {
	u := &Readinfo{}
	d := DB.Self.Model(&Readinfo{}).Where("news_id = ?", id).First(&u)

	return u, d.Error
}
