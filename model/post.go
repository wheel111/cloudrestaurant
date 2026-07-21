// Package model 定义了数据模型
package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title   string `gorm:"type:varchar(255);not null" json:"title"` // 标题
	Content string `gorm:"type:longtext;not null" json:"content"`   // 论坛内容
	Summary string `gorm:"type:varchar(500)" json:"summary"`        // 摘要
	Status  int    `gorm:"default:1;not null" json:"status"`
	// 外键关联
	AuthorID   uint     `gorm:"index;not null" json:"author_id"`       // 发帖人id
	Author     User     `gorm:"foreignKey:AuthorID" json:"author"`     // 关联用户信息
	CategoryID uint     `gorm:"index;not null" json:"category_id"`     // 所属板块id
	Category   Category `gorm:"foreignKey:CategoryID" json:"category"` // 关联板块信息
	Tags       []Tag    `gorm:"many2many:post_tags;" json:"tags"`      // 标签
	ViewCount  int64    `gorm:"default:0" json:"view_count"`           // 浏览量
	LikeCount  int64    `gorm:"default:0" json:"like_count"`           // 点击量
}
