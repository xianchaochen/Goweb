package service

import (
	"bluebell/entity/user"
	"bluebell/pkg/jwt"
	"bluebell/repository"
)
import (
	"bluebell/common"
	"bluebell/model"
	"bluebell/pkg/snowflake"
	"errors"
)

type IUserService interface {
	Register(p *user.ParamRegister) error
	Login(p *user.ParamLogin) (string,string,error)
}

var _ IUserService = (*UserService)(nil)

type UserService struct {
	userRepository repository.IUserRepository
}

func NewUserService(repository repository.IUserRepository) IUserService {
	return &UserService{userRepository: repository}
}

func (u *UserService) Register(p *user.ParamRegister) error {
	if exist := u.userRepository.CheckUserExist(p.Username); exist {
		return errors.New("用户名已存在")
	}

	userID := snowflake.GenID()
	hashPassword, err := common.GeneratePassword(p.Password)
	if err != nil {
		return err
	}
	user := &model.User{
		Username: p.Username,
		Password: string(hashPassword),
		UserID:   userID,
	}
	_, err = u.userRepository.Insert(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) Login(p *user.ParamLogin) (string,string,error) {
	user := u.userRepository.FindUserByUsername(p.Username)
	if user == nil {
		return "","",errors.New("无效用户")
	}

	isOk, err := common.ValidatePassword(p.Password, user.Password)
	if !isOk {
		return "","",err
	}

	aToken, rToken, err := jwt.Generate(p.Username, user.UserID)
	if err != nil {
		return "","",err
	}
	return aToken,rToken,nil
}
