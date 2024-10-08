package model

// 食品种类结构体
type FoodCategory struct {
	//类别id
	Id int64 `xorm:"pk autoincr" json:"id"`
	//食品类别标题
	Title string `xorm:"varchar(20)" json:"title"`
	//食品描述
	Description string `xorm:"varchar(30)" json:"description"`
	//食品种类图片
	ImageUrl string `xorm:"varchar(255)" json:"image_url"`
	//食品类别链接
	LinkUrl string `xorm:"varchar(255)" json:"link_url"`
	//该类别服务是否在线
	IsInServing bool `json:"is_in_serving"`
}
