package service

import (
	"bluebell/entity"
	"bluebell/repository"
)
import (
	"bluebell/common"
	"bluebell/model"
	"bluebell/pkg/snowflake"
	"errors"
)

type IUserService interface {
	Register(p *entity.ParamRegister) error
}

var _ IUserService = (*UserService)(nil)

type UserService struct {
	userRepository repository.IUserRepository
}

func NewUserService(repository repository.IUserRepository) IUserService {
	return &UserService{userRepository: repository}
}

func (u *UserService)Register(p *entity.ParamRegister) error  {
	exist, err := u.userRepository.CheckUserExist(p.Username)
	if err != nil {
		return err
	}

	if exist {
		return errors.New("用户名已存在")
	}

	userID := snowflake.GenID()
	hashPassword, err := common.GeneratePassword(p.Password)
	if err!=nil {
		return err
	}
	user := &model.User{
		Username: p.Username,
		Password: string(hashPassword),
		UserID: userID,
	}

	_, err = u.userRepository.Insert(user)
	if err != nil {
		return err
	}
	return nil
}
