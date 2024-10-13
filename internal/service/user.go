package domain

import (
	repository "github.com/chongyanovo/go-zzz/internal/repository"
	"go.uber.org/zap"
)

type UserService interface {
}

var _ UserService = (*userService)(nil)

type userService struct {
	l    *zap.Logger
	repo repository.UserRepository
}

func NewUserService(l *zap.Logger, repo repository.UserRepository) UserService {
	return &userService{
		l:    l,
		repo: repo,
	}
}
