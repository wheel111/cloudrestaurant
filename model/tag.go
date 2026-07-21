package model

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name  string `gorm:"type:varchar(50);unique;not null" json:"name"`
	Posts []Post `gorm:"many2many:post_tags;" json:"-"`
}
