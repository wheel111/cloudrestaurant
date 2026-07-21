package dao

import (
	"blog/model"
	"blog/tool"
	"errors"

	"gorm.io/gorm"
)

type UserDao struct {
	*tool.Gorm
}

func NewUserDao() *UserDao {
	return &UserDao{tool.DbEngine}
}

// 创建角色

func (d *UserDao) Create(user *model.User) error {
	return d.DB.Create(user).Error
}

// 根据id查找用户
func (d *UserDao) GetByID(id uint) (*model.User, error) {
	var user model.User
	err := d.DB.First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

// 根据用户名查找用户
func (d *UserDao) GetByName(name string) (*model.User, error) {
	var user model.User
	err := d.DB.Where("username=?", name).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

// 更新用户头像
func (d *UserDao) UpdateAvatar(id uint, avatar string) error {
	return d.DB.Model(&model.User{}).Where("id=?", id).Update("avatar", avatar).Error
}
