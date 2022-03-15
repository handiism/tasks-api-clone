package service

import (
	"tasks/repo"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Verify(email, password string) error
}

type authService struct {
	userRepo repo.UserRepo
}

func NewAuthService(userRepo repo.UserRepo) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

func (a *authService) Verify(email string, password string) error {
	user, err := a.userRepo.FindByEmail(email)

	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password), []byte(password),
	); err != nil {
		return err
	}

	return nil
}
