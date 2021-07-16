package service

import "golang-jwt/src/user/domain"

type UserServiceInterface interface {
	Login(username string, password string) (domain.User, error)
}
