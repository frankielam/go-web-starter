package db

import (
	"context"
	"gowebstarter/model/entity"
)

func GetSysUserById(ctx context.Context, id int64) (*entity.SysUser, error) {
	var user entity.SysUser
	err := db.Where("id = ?", id).First(&user).Error
	return &user, err
}

func GetSysUserByName(ctx context.Context, name string) (*entity.SysUser, error) {
	var user entity.SysUser
	err := db.Where("name = ?", name).First(&user).Error
	return &user, err
}

func CreateUser(ctx context.Context, user *entity.SysUser) error {
	return db.Create(user).Error
}

func UpdateUser(ctx context.Context, user *entity.SysUser) error {
	return db.Save(user).Error
}

func DeleteUser(ctx context.Context, id int64) error {
	return db.Delete(&entity.SysUser{}, id).Error
}
