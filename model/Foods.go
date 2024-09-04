package model

// 食品结构体定义
type Foods struct {
	Id          int64   `xorm:"pk autoincr" json:"id"`
	Name        string  `xorm:"varchar(12)" json:"name"`
	Description string  `xorm:"varchar(32)" json:"description"`
	Icon        string  `xorm:"varchar(255)" json:"icon"`
	SellCount   int64   `xorm:"int" json:"sell_count"`
	Price       float32 `xorm:"float" json:"price"`
	OldPrice    float32 `xorm:"float" json:"old_price"`
	ShopId      int64   `xorm:"int" json:"shop_id"`
}
