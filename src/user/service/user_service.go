package service

import (
	"context"
	"golang-jwt/src/user/domain"
	"golang-jwt/src/user/repository"
)

type UserService struct {
	repo repository.UserInterface
}

func NewUserService(rp repository.UserInterface) UserServiceInterface {
	return &UserService{repo: rp}
}

func (rp *UserService) Login(username string, password string) (domain.User, error) {

	user, err := rp.repo.CheckUsername(context.Background(), username)

	if err != nil {
		return *user, err
	}

	return *user, nil
}
