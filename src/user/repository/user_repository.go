package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-jwt/src/user/domain"
)

type UserRepository struct {
	repo *sql.DB
}

func NewUserRepo(rp *sql.DB) UserInterface {
	return &UserRepository{repo: rp}
}

func (userRepo *UserRepository) CheckUsername(ctx context.Context, userName string) (*domain.User, error) {
	script := "Select * from users where username=? limit 1"
	result, err := userRepo.repo.QueryContext(ctx, script, userName)
	user := domain.User{}
	if err != nil {
		return &user, err
	}
	defer result.Close()

	if result.Next() {
		result.Scan(&user.Username)
		return &user, nil
	} else {
		return &user, errors.New("Id: " + userName + " tidak ditemukan")
	}
}

func (userRepo *UserRepository) Register(user *domain.User) error {
	return nil
}
