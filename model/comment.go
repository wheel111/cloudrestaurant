package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `gorm:"type:text;not null" json:"content"` // 评论内容

	PostID uint `gorm:"index;not null" json:"post_id"` // 对应文章的评论
	Post   Post `gorm:"foreignKey:PostID" json:"-"`    // 避免循环引用

	UserID uint `gorm:"index;not null" json:"user_id"` // 发表评论的用户
	User   User `gorm:"foreignKey:UserID" json:"user"` // 关联

	// 支持楼中楼回复 (ParentID 为 0 表示顶级评论)
	ParentID uint `gorm:"index;default:0" json:"parent_id"` // 嵌套评论基础
}
