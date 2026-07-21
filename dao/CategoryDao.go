package dao

import (
	"blog/model"
	"blog/tool"
	"errors"

	"gorm.io/gorm"
)

type CategoryDao struct {
	*tool.Gorm
}

func NewCategoryDao() *CategoryDao {
	return &CategoryDao{tool.DbEngine}
}

// 创建板块
func (d *CategoryDao) Create(category *model.Category) error {
	return d.DB.Create(category).Error
}

// 搜索所有板块
func (d *CategoryDao) GetAll() ([]model.Category, error) {
	var category []model.Category
	err := d.DB.Find(&category).Error
	return category, err
}

// 根据id搜索特定的板块
func (d *CategoryDao) GetByID(id uint) (*model.Category, error) {
	var category model.Category
	err := d.DB.First(&category, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &category, err
}

// 更新板块信息
func (d *CategoryDao) Update(id uint, field map[string]interface{}) error {
	return d.DB.Model(&model.Category{}).Where("id= ?", id).Updates(field).Error
}

// 删除板块信息
func (d *CategoryDao) Delete(id uint) error {
	return d.DB.Delete(&model.Category{}, id).Error
}
