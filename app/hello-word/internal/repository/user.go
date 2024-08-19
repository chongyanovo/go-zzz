package repository

import "github.com/chongyanovo/go-zzz/app/hello-word/internal/repository/dao"

type UserRepository interface {
	FindUserById(id int) (*dao.User, error)
}

type UserRepositoryImpl struct {
	userDao *dao.UserDao
}

func NewUserRepository(userDao *dao.UserDao) UserRepository {
	return &UserRepositoryImpl{userDao: userDao}
}

func (u UserRepositoryImpl) FindUserById(id int) (*dao.User, error) {
	//TODO implement me
	panic("implement me")
}
