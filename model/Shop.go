package model

// 商家结构体
type Shop struct {
	Id                 int     `xorm:"pk autoincr" json:"id"`
	Name               string  `xorm:"varchar(12)" json:"name"`
	PromotionInfo      string  `xorm:"varchar(30)" json:"promotion_info"`
	Address            string  `xorm:"varchar(100)" json:"address"`
	Phone              string  `xorm:"varchar(11)" json:"phone"`
	Status             int     `xorm:"tinyint" json:"status"`
	Longitude          float64 `xorm:"" json:"longitude"`
	Latitude           float64 `xorm:"" json:"latitude"`
	ImagePath          string  `xorm:"varchar(255)" json:"image_path"`
	IsNew              bool    `xorm:"bool" json:"is_new"`
	IsPremium          bool    `xorm:"bool" json:"is_premium"`
	Rating             float32 `xorm:"float" json:"rating"`
	RatingCount        int64   `xorm:"int" json:"rating_count"`
	RecentOrderNum     int64   `xorm:"int" json:"recent_order_num"`
	MinimumOrderAmount int32   `xorm:"int" json:"minimum_order_amount"`
	DeliveryFee        int32   `xorm:"int" json:"delivery_fee"`
	OpeningHours       string  `xorm:"varchar(20)" json:"opening_hours"`
}
