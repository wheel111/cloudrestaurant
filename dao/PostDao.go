package dao

import (
	"blog/model"
	"blog/tool"
	"errors"

	"gorm.io/gorm"
)

type PostDao struct {
	*tool.Gorm
}

func NewPostDao() *PostDao {
	return &PostDao{tool.DbEngine}
}

// 创建文章
func (d *PostDao) Create(post *model.Post) error {
	return d.DB.Create(post).Error
}

// 根据id 查询单篇文章（带作者、分类、标签）
func (d *PostDao) GetByID(id uint) (*model.Post, error) {
	var post model.Post
	err := d.DB.Preload("Author").Preload("Category").Preload("Tags").First(&post, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &post, err
}

// 分页查询文章列表
func (d *PostDao) List(offset, limit int) ([]model.Post, int64, error) {
	var posts []model.Post
	var total int64
	d.DB.Model(&model.Post{}).Count(&total)
	err := d.DB.Preload("Author").Preload("Category").Preload("Tags").Offset(offset).Limit(limit).Order("created_at DESC").Find(&posts).Error
	return posts, total, err
}

// 删除文章
func (d *PostDao) Delete(id uint) error {
	return d.DB.Delete(&model.Post{}, id).Error
}

// 更新文章
func (d *PostDao) Update(id int, field map[string]interface{}) error {
	return d.DB.Model(&model.Post{}).Where("id = ?", id).Updates(field).Error
}

// 标签关联
func (d *PostDao) CreateWithTags(post *model.Post, tagIDs []uint) error {
	return d.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(post).Error; err != nil {
			return err
		}
		if len(tagIDs) > 0 {
			var tags []model.Tag
			tx.Find(&tags, tagIDs)
			return tx.Model(post).Association("Tags").Replace(tags)
		}
		return nil
	})
}
