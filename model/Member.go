package model

type Member struct {
	Id           int64   `xorm:"pk autoincr " json:"id"`
	Username     string  `xorm:"varchar(20) " json:"username"`
	Mobile       string  `xorm:"varchar(11) " json:"mobile"`
	Password     string  `xorm:"varchar(255) " json:"password"`
	RegisterTime int64   `xorm:"bigint " json:"register_time"`
	Avatar       string  `xorm:"varchar(255) " json:"avatar"`
	Balance      float64 `xorm:"double " json:"balance"`
	IsActive     int8    `xorm:"tinyint " json:"is_active"`
	City         string  `xorm:"varchar(10) " json:"city"`
}
