package dao

import (
	"blog/model"
	"blog/tool"
)

type CommentDao struct {
	*tool.Gorm
}

func NewCommentDao() *CommentDao {
	return &CommentDao{tool.DbEngine}
}

// 创建评论
func (d *CommentDao) Create(comment *model.Comment) error {
	return d.DB.Create(comment).Error
}

// 删除评论
func (d *CommentDao) Delete(id uint) error {
	return d.DB.Delete(&model.Comment{}, id).Error
}

// 检索所有评论
func (d *CommentDao) GetByPostID(postID uint) ([]model.Comment, error) {
	var comments []model.Comment
	err := d.DB.Preload("User").Where("post_id = ?", postID).Find(&comments).Error
	return comments, err
}
