package repository

import (
	"context"
	"golang-jwt/src/user/domain"
)

type UserInterface interface {
	CheckUsername(ctx context.Context, name string) (*domain.User, error)
	Register(user *domain.User) error
}
