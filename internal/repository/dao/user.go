package dao

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserDao interface {
}

var _ UserDao = (*userDao)(nil)

type userDao struct {
	l  *zap.Logger
	db *gorm.DB
}

func NewUserDao(l *zap.Logger, db *gorm.DB) UserDao {
	return &userDao{
		l:  l,
		db: db,
	}
}
