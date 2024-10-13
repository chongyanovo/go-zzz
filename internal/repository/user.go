package domain

import (
	"github.com/chongyanovo/go-zzz/internal/repository/dao"
	"go.uber.org/zap"
)

type UserRepository interface {
}

var _ UserRepository = (*UserRepository)(nil)

type userRepository struct {
	l   *zap.Logger
	dao dao.UserDao
}

func NewUserRepository(l *zap.Logger, dao dao.UserDao) UserRepository {
	return &userRepository{
		l:   l,
		dao: dao,
	}
}
