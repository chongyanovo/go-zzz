package dao

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type User struct {
	ID   int64  `gorm:"primaryKey autoIncrement column:id"`
	Name string `gorm:"column:name"`
}

type UserDao interface {
}

type UserDaoImpl struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewUserDao(db *gorm.DB) UserDao {
	return &UserDaoImpl{db: db}
}
